APPLICATION_NAME=llama-server
LOCAL_GOOS=${shell go env GOOS}
LOCAL_GOARCH=${shell go env GOARCH}
VERSION=${shell cat version}

.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt -C src

vet: fmt
	go vet -C src

build: fmt vet
	GOOS=windows GOARCH=amd64 go build -C src -o ../bin/${APPLICATION_NAME}_${VERSION}_windows_amd64
	GOOS=windows GOARCH=arm64 go build -C src -o ../bin/${APPLICATION_NAME}_${VERSION}_windows_arm64
	GOOS=linux GOARCH=amd64 go build -C src -o ../bin/${APPLICATION_NAME}_${VERSION}_linux_arm64
	GOOS=linux GOARCH=arm64 go build -C src -o ../bin/${APPLICATION_NAME}_${VERSION}_linux_amd64
	GOOS=darwin GOARCH=arm64 go build -C src -o ../bin/${APPLICATION_NAME}_${VERSION}_darwin_arm64
	GOOS=darwin GOARCH=amd64 go build -C src -o ../bin/${APPLICATION_NAME}_${VERSION}_darwin_amd64

clean:
	rm -rf ./bin/${APPLICATION_NAME}_*
	docker rm -f ${APPLICATION_NAME}-redis

run: clean build
	docker pull redis
	docker run --name ${APPLICATION_NAME}-redis -d redis
	./bin/${APPLICATION_NAME}_${VERSION}_${LOCAL_GOOS}_${LOCAL_GOARCH}
