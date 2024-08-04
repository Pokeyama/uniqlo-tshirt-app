# ベースイメージとして公式のGoイメージを使用
FROM golang:1.20

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールのキャッシュを利用するため、go.modとgo.sumをコピー
COPY go.mod ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o main .

# ポート8080を開放
EXPOSE 8080

# コンテナ起動時に実行されるコマンド
CMD ["./main"]
