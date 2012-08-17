pprof:
	go build montego.go
	./montego -cpuprofile=montego.pprof
	go tool pprof montego montego.pprof
