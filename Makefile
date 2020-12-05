deps:
	dep ensure -v

run: deps
	go run *.go