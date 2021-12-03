.PHONY: build
build:
	go build -o cli-sample .
.PHONY: test
test:
	go test . -shuffle on