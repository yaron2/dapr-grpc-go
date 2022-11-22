# gRPC proxy example with Dapr and Go

This example shows how to use Dapr gRPC proxying with Go using native gRPC clients and servers.

## Getting started

Prerequisites:

* [Go](https://go.dev/dl/)
* [Dapr CLI](https://docs.dapr.io/getting-started/install-dapr-cli/)

### Install Dapr

If you haven't already run `dapr init`, do so via your terminal:

```bash
$ dapr init
```

### Run the gRPC Go server

```bash
cd greeter_server
go mod tidy

dapr run --app-id server --app-port 50051 -- go run main.go
```

### Run the gRPC Go client

```bash
cd greeter_client
go mod tidy

dapr run --app-id client --dapr-grpc-port 3501 -- go run main.go
```

### Verify it works

If everything is set up correctly, you should see an output similar to this:

```bash
== APP == 2022/11/21 16:12:07 Greeting: Hello 0, 20221121161207
== APP == 2022/11/21 16:12:08 Greeting: Hello 1, 20221121161208
== APP == 2022/11/21 16:12:09 Greeting: Hello 2, 20221121161209
== APP == 2022/11/21 16:12:10 Greeting: Hello 3, 20221121161210
== APP == 2022/11/21 16:12:11 Greeting: Hello 4, 20221121161211
== APP == 2022/11/21 16:12:12 Greeting: Hello 5, 20221121161212
== APP == 2022/11/21 16:12:13 Greeting: Hello 6, 20221121161213
== APP == 2022/11/21 16:12:14 Greeting: Hello 7, 20221121161214
== APP == 2022/11/21 16:12:15 Greeting: Hello 8, 20221121161215
== APP == 2022/11/21 16:12:16 Greeting: Hello 9, 20221121161216
```
