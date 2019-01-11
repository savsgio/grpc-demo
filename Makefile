build_server:
	go build -o bin/server  cmd/server/main.go

build_client:
	go build -o bin/client  cmd/client/main.go

clean:
	rm -f ./bin/*
