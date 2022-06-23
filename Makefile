build:
	go build -o ./bin/rtouch ./cmd
install:
	go build -o ~/.local/bin/rtouch ./cmd
release:
	GOOS=windows GOARCH=amd64 go build -o ./bin/rtouch_windows_amd64.exe ./cmd
	GOOS=windows GOARCH=arm64 go build -o ./bin/rtouch_windows_arm64.exe ./cmd
	GOOS=darwin GOARCH=amd64 go build -o ./bin/rtouch_mac_amd64 ./cmd
	GOOS=darwin GOARCH=arm64 go build -o ./bin/rtouch_mac_arm64 ./cmd
	GOOS=linux GOARCH=amd64 go build -o ./bin/rtouch_linux_amd64 ./cmd
	GOOS=linux GOARCH=arm64 go build -o ./bin/rtouch_linux_arm64 ./cmd