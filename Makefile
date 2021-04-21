#!/usr/bin/make -f

rosetta: go.sum
	go build -mod=readonly ./cmd/rosetta

install:
	go install ./cmd/rosetta

test:
	go test -mod=readonly -race github.com/CortexFoundation/rosetta-cortex/...

clean:
	rm -f rosetta coverage.txt

.PHONY: rosetta test clean
