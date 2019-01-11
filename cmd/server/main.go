package main

import "demo"

const (
	port = ":50051"
)

func main() {
	demo.Listen(port)
}
