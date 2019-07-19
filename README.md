# nats-streaming-ui

Simple Dashboard UI for inspecting a NATS Streaming Cluster

* [x] List Channles and Subscriptions
* [ ] Channel activity graphs

## Development

You need a running NATS Streaming cluster in order run the UI.

```
kubectl port-forward stan-cluster-nats-streaming-1 8222
```

Fetch all dependencies:

```

Run the server using `air` to automatically reload on changes:

```
go get github.com/cosmtrek/air/cmd/air
 ~/go/bin/air -c .air.conf
```
