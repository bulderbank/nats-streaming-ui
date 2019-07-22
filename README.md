# nats-streaming-ui

Simple Dashboard UI for inspecting a NATS Streaming Cluster

* [x] List Channles and Subscriptions
* [ ] Channel activity graphs

### Tech Stack

* Go 1.12
* [Gin](https://github.com/Masterminds/sprig) web frmework
* [Sprig](https://github.com/Masterminds/sprig) template functions
* [Semantic UI](https://semantic-ui.com/)

## Development

You need a running NATS Streaming cluster in order run the UI.

```
kubectl port-forward stan-cluster-nats-streaming-1 8222
```

Run the server using `air` to automatically reload on changes:

```
go get github.com/cosmtrek/air/cmd/air
 ~/go/bin/air -c .air.conf
```
