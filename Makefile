clean:
	rm -rf bin/

build-mac:
	env GOARCH=amd64 GOOS=darwin go build -o bin/mac/tfutil -v

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/tfutil
	
clean-ws:
	rm -rf temp-ws/