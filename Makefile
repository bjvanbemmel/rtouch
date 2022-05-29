build:
	go build -o ./bin/rtouch ./cmd
install:
	go build -o ~/.local/bin/rtouch ./cmd