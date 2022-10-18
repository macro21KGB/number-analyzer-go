converter: main.go
	go build -o converter main.go 

run: converter
	./converter 10 15
