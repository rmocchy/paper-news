# Dockerfile
FROM golang:1.23.3

# 必要なツールをインストール
RUN apt-get update && apt-get install -y unzip

# プロトコルバッファのコンパイラを取得
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v24.4/protoc-24.4-linux-x86_64.zip && \
    unzip protoc-24.4-linux-x86_64.zip -d /usr/local && \
    rm protoc-24.4-linux-x86_64.zip

# Goのprotobufコード生成プラグインをインストール
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# PATHの設定
ENV PATH=$PATH:/go/bin

# ワークディレクトリ
WORKDIR /app

# ソースコードをコンテナにコピー
COPY . .
