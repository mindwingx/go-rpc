### Efficient Go RPC Usage

This guide outlines a practical approach to leverage Go's built-in `RPC (Remote Procedure Call)` mechanism, which is
sometimes confused with the more powerful `gRPC`. It focuses on the `master` and `worker` dummy microservices, employing RPC
over REST API or other alternatives.

##### Scenario Overview

In our scenario, we have two modules: master and worker, serving as mock microservices. We're using Go's RPC for
communication.

- Role of the master Module

The `master` module acts as the server. It offers a method named `Multiply`, which returns results to the caller, i.e.,
the `worker` module. It utilizes the server method to register the RpcServer struct, enabling it to listen on
port `9090` over `TCP`. The listener is continuously accepting incoming requests, passing them to the `ServeConn` method
via `listener.Accept()`.

The `Multiply` method is accessible via the `RpcServer` struct as well. Launch the master service with the following
command to enable `RPC` on port `9090`:

```go
go run master/main.go
```

- Role of the `worker` Module

The worker module functions as the client/caller. It sends requests to the `master` module, which is actively listening
for `RPC` calls. Additionally, the module has a `Dial` method for instantiating `RPC`. It prepares the `Call` method to
invoke the served `RPC` method (`Multiply`) on the same network and port. This call triggers the `Multiply` method on
the master module concurrently.

Run the following command to initiate the worker service:

```go
go run worker/main.go
```

To send a custom count of requests, use:

```go
go run worker/main.go --count = 2000
```

Additional Notes:

- The caller module's `rpc.Call` function takes three parameters: the registered struct's name concatenated with the
  method name (e.g., `RpcServer.Multiply`), an argument struct, and a response-compatible struct.
- Ensure the caller module's struct matches the one used in the server methods (`Multiply`).
- This example demonstrates efficient `RPC` use in a simplified scenario. Remember to adjust and expand this pattern according to your actual use case and project requirements.