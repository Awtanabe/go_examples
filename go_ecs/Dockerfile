FROM golang:1.23.4-alpine

WORKDIR /app

# アプリケーションのソースコードをコピー
COPY . .

# モジュール依存関係をダウンロード
RUN go mod download

# 正しいアーキテクチャ (linux/amd64) でバイナリをビルド
RUN GOOS=linux GOARCH=amd64 go build -o main main.go

# ポート8080を公開
EXPOSE 8080

# ビルド済みのバイナリを実行
CMD ["./main"]