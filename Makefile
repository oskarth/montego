pprof:
	go build montego.go
	./montego
	go tool pprof montego montego.pprof
