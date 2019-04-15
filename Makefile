image:
	GOOS=linux GOARCH=amd64 go build -o ./docker/<YOUR-APPLICATION> ./cmd/<YOUR-APPLICATION>
	cd docker && docker build -t <YOUR-APPLICATION-AT-REG>:<VERSION> .  && cd -
	rm ./docker/<YOUR-APPLICATION>

test:
	go test -v $(shell go list ./... | grep -v /vendor/)

