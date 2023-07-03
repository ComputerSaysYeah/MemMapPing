
all: clean binary

clean:
	rm -f mmping
	go clean -testcache -testcache

binary:
	go build -o mmping mmping.go

