all:
	go build

benchmark:
	go test -test.bench .

cpuprofile-run:
	go test -test.cpuprofile cpu.out
	
cpuprofile-open:
	go tool pprof calculordssolver cpu.out

cpuprofile-run-open: cpuprofile-run cpuprofile-open

memprofile-run:
	go test -test.memprofile mem.out
	
memprofile-open:
	go test -test.memprofile mem.out

memprofile-run-open: memprofile-run memprofile-open