build:
	go build -o journal main.go 

clean:
	go clean 
	rm -rf journal journal.yaml