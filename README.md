# Paper-News Backend

## これは何？
論文情報を収集して保存してくれるサーバーです。

## 開発者向け情報
以下の機能が有効になっています
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