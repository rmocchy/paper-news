FROM golang:1.23.3

WORKDIR /app
COPY . .
RUN go mod download
# for development
RUN go install github.com/air-verse/air@latest 

# RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -v -o /app/server /app/cmd/main.go

# EXPOSE 8080