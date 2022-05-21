
.PYHON: test


test:
	go test -v -json -gcflags=all=-l ./...