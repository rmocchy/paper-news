FROM golang:1.22.2-alpine3.18

WORKDIR /app

COPY ./ ./
RUN go mod download

# バイナリファイルにビルド
RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o server /app/cmd/main.go

EXPOSE 8080

# バイナリファイルを実行
CMD ./server