test:
	go test ./...

test.race:
	go test -race ./...

fmt:
	gofumpt -w -l .
