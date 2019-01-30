SRCS=$(wildcard *.go)

9cc: $(SRCS)
	go build -o 9cc $(SRCS)

test: 9cc
	./test.sh


clean:
	rm -rf 9cc *.o *~ tmp*
