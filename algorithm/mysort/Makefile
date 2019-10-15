default:
	./mergeSort -cpuprofile=cpu.pprof -memprofile=mem.pprof
cpu.pprof:
	./mergeSort -cpuprofile=cpu.pprof
mem.pprof:
	./mergeSort -memprofile=mem.pprof
build:
	go build /data/github.com/tkstorm/go-example/algorithm/mysort/mergeSort.go

pprof-mem:
	go tool pprof ./mergeSort mem.pprof
pprof-cpu:
	go tool pprof ./mergeSort cpu.pprof


.PHONY: build cpu.pprof mem.pprof default
