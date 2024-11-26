# 雛形 of Gin-server (by mocchy)

## これは何？
Golang(Gin)サーバーの雛形です。
以下の機能が使える状態になっていますので、すぐにコード実装が行えます！
- Airによるホットリロード
- Protobufによるコード生成

## 使い方
### サーバーの起動
1. 以下コードを実行すればOK
```
docker compose up
```

### Protobufの生成
1. gen/apiフォルダの用意
2. protoフォルダの用意
3. 以下コードの実行
```
make protoc_docker
```
gen/api傘下にファイルたちが生成されます