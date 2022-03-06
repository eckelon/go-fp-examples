.PHONY: goderive run clean test
default:
	make clean
	make goderive
	make run
goderive:
	go install github.com/awalterschulze/goderive
	goderive --autoname=true --dedup=true --prefix="" ./...
test:
	go test -json -v ./... 2>&1 | tee /tmp/gotest.log | gotestfmt
run:
	go run .
clean:
	go clean
	rm derived.gen.go || true
