test:
	go test ./...

test.race:
	go test -race ./...

dep.update:
	go get -u ./...

fmt:
	gofumpt -w -l .
