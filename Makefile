build:
	docker-compose build

run:
	docker-compose up

compile:
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-freebsd-386 main.go