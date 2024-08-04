package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ControlMessage struct {
	Target string
	Count  int
}

func main() {
	controlChannel := make(chan ControlMessage)
	workerCompleteChan := make(chan bool)
	statusPollChannel := make(chan chan bool)
	workerActive := false

	go admin(controlChannel, statusPollChannel)

	for {
		select {
		case respChan := <-statusPollChannel:
			respChan <- workerActive
			log.Println("Status requested, responded with:", workerActive)
		case msg := <-controlChannel:
			workerActive = true
			log.Println("Received control message:", msg)
			// ワーカーが処理を開始する
			go func() {
				time.Sleep(time.Duration(msg.Count) * time.Second)
				workerCompleteChan <- false
				log.Println("Work complete for target:", msg.Target)
			}()
		case status := <-workerCompleteChan:
			workerActive = status
			log.Println("Worker status updated to:", workerActive)
		}
	}
}

func admin(cc chan ControlMessage, statusPollChannel chan chan bool) {
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		hostTokens := strings.Split(r.Host, ":")
		log.Println("Host tokens:", hostTokens)

		r.ParseForm()
		count, err := strconv.ParseInt(r.FormValue("count"), 10, 32)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		cc <- ControlMessage{
			Target: r.FormValue("target"),
			Count:  int(count),
		}

		fmt.Fprintf(w, "Control message issued for Target %s with Count %d", html.EscapeString(r.FormValue("target")), int(count))
		log.Println("Control message issued for Target", r.FormValue("target"), "with Count", int(count))
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		reqChan := make(chan bool)
		statusPollChannel <- reqChan
		timeout := time.After(1 * time.Second)

		select {
		case result := <-reqChan:
			if result {
				fmt.Fprint(w, "ACTIVE")
			} else {
				fmt.Fprint(w, "INACTIVE")
			}
		case <-timeout:
			fmt.Fprint(w, "INACTIVE")
		}
	})

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
