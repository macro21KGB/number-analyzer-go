converter: src/main.go
	go build -o converter src/main.go 

run: converter
	./converter 10 15
	@echo "#########"
	./converter 1 3
