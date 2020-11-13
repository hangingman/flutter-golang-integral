# flutter-golang-integral

## サーバのビルド

```sh
$ make
```

## gRPCのサーバを動かす

- 日付と使用しているポートが出力されれば成功

```sh
$ make run
...
2020/11/13 18:52:27 start server on port:50051
```

## gRPCと通信テスト

- `grpcurl` を使う

```sh
$ go get github.com/fullstorydev/grpcurl/...
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

- インストールが出来たら、サーバを動かした状態で, サービス一覧を出してみる

```sh
$ grpcurl -plaintext localhost:50051 list
grpc.reflection.v1alpha.ServerReflection
pb.EchoService
```
