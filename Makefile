dev:
	fresh

build:
	go build -o ./bin/app .

audit:
	go run golang.org/x/vuln/cmd/govulncheck@latest -show verbose ./...
