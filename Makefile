${GOPATH}/bin/broker: $(shell find cmd pkg)
	go install ./cmd/broker

run: ${GOPATH}/bin/broker vendor
	@${GOPATH}/bin/broker

clean:
	@rm -f ${GOPATH}/bin/broker

vendor:
	@glide install

test: vendor
	go test ./pkg/...

.PHONY: run broker clean
