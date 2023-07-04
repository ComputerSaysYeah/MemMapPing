
all: clean binary

clean:
	rm -f mmping mmping.file
	go clean -testcache -testcache

binary:
	go build -o mmping mmping.go

