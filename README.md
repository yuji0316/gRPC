# Goで初めてのgRPC

## 環境構築

1. GoのgRPCライブラリを取得

    go get -u google.golang.org/grpc

1. Protocol Buffersをインストール(Macの場合)

    brew install protobuf

1. Protocol Buffersが標準ではGoのコード生成に対応していないので、そのためのプラグインを取得

    go get -u github.com/golang/protobuf/protoc-gen-go

1. 上記のライブラリがデフォルトで`$GOPATH/bin`にインストールされるのでパスを通す

    export PATH=$PATH:$GOPATH/bin

## Protocol BuffersからGoのgRPC用コードを生成

1. Protocol BuffersでgRPCのイターフェースを定義  
[ソースコード](./hellovoicy/hellovoicy.proto)

1. Protocol BuffersからGoのソースコードを生成  

    protoc -I hellovoicy/ hellovoicy/hellovoicy.proto --go_out=plugins=grpc:hellovoicy  
    [生成されたコード](./hellovoicy/hellovoicy.pb.go)

## サーバーサイドの構築

[ソースコード](./server/main.go)

実行

    go run server/main.go

## クライアントサイドの構築

[ソースコード](./client/main.go)

実行

    go run client/main.go

## その他

[こちら](https://github.com/grpc/grpc-go)のコードを参考にいたしました。

[This software includes the work that is distributed in the Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0)