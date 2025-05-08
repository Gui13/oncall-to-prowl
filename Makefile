

.PHONY: all clean amd64-linux arm64-linux arm64-mac

all: arm64-mac

amd64-linux:
	GOOS=linux GOARCH=amd64 go build
arm64-linux:
	GOOS=linux GOARCH=arm64 go build

arm64-mac:
	GOOS=darwin GOARCH=arm64 go build

clean:
	rm oncall-to-prowl

docker: amd64-linux
	docker build -t ghcr.io/gui13/oncall-to-prowl .

pushdocker:
	docker push ghcr.io/gui13/oncall-to-prowl

test:
	go test
