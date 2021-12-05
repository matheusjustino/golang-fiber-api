build:
	go build -o server main.go

start:
	go run main.go

dev:
	nodemon --exec go run main.go

prod:
	./server