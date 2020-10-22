build:
	go build -o ./bin/cryptonotify ./cmd/cryptonotify

run:
	go run ./cmd/cryptonotify

test:
	go test ./internal/rules
	go test ./pkg/coinlore
	go test ./cmd/cryptonotify

compile:
	# Linux 64-bit
	GOOS=linux GOARCH=amd64 go build -o ./bin/cryptonotify-linux-amd64 ./cmd/cryptonotify
	# Windows 64-bit
	GOOS=windows GOARCH=amd64 go build -o ./bin/cryptonotify-windows-amd64.exe ./cmd/cryptonotify
