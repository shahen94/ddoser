default: build
build:
	echo "Building DDoSer..."
	go build -o ddoser cmd/ddoser/main.go
	echo "Done!"
	echo "Run ./ddoser to start the program"
clean:
	rm -f ddoser
run:
	go run cmd/ddoser/main.go