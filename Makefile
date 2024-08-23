BINARY_NAME=brp

build:
	go build -o build/${BINARY_NAME} main.go

full: build
	GOARCH=amd64 GOOS=darwin  go build -o build/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux   go build -o build/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o build/${BINARY_NAME}-windows.exe main.go
	GOARCH=arm64 GOOS=darwin  go build -o build/${BINARY_NAME}-darwin-arm main.go
	GOARCH=arm64 GOOS=linux   go build -o build/${BINARY_NAME}-linux-arm main.go
	GOARCH=arm64 GOOS=windows go build -o build/${BINARY_NAME}-windows-arm.exe main.go

test:
	# rm -r test 2> /dev/null
	mkdir -p test/one test/two test/three
	touch test/one/abc.txt test/one/bar.txt test/two/car.md test/two/door.md test/three/"hello world.md" test/three/"good bye.md"

	# rm -r auto 2> /dev/null
	mkdir auto
	echo "123 oo\n321 a" > auto/ap.txt
	echo "789 oo\n987 e" > auto/pre.txt
	echo "re a /-\nre oo __\nre c 1\nre b 2\nre oo __\ncase title kebab\ninsert '  ' 2 -c\ncase trim" > auto/auto.txt


clean:
	go clean
	rm -r test auto build/* 2> /dev/null


.PHONY: all build full test clean
