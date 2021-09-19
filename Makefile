init-project:
	go mod init golang-gin-poc

build:
	go build -race -o app . 

run: build 
	./app 

	