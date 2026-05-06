all: read-st

read-st:
	go build

clean:
	- go clean
	- rm -f read-st-*

cross:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o read-st-linux-amd64
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o read-st-windows-amd64.exe
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o read-st-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o read-st-darwin-arm64
