# Go HTTP Server with Control and Status Endpoints on a UNIQLO T-shirt

## Overview

This repository contains a Go-based HTTP server designed to handle control messages and provide status updates. The server demonstrates basic usage of goroutines and channels to manage concurrent tasks efficiently. This code is featured on a UNIQLO T-shirt, showcasing practical Go programming in a stylish way.

### Build and Run with Docker:
```
docker build -t go-http-server .
docker run -p 8080:8080 go-http-server
```

## Example Usage
### Send a control message:
```
curl -X POST "http://localhost:8080/admin" -d "target=mytarget&count=10"
```

### Check the server status:
```
curl "http://localhost:8080/status"
```