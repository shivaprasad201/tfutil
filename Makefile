clean:
	rm -rf bin/

build-mac:
	env GOARCH=amd64 GOOS=darwin go build -o bin/mac/servicebackend

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/servicebackend
	
clean-ws:
	rm -rf temp-ws/