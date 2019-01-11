package main

import "demo/client"

const (
	address = "localhost:50051"
)

func main() {
	client.Main(address)
}
