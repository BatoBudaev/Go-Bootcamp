.PHONY: all clean

00 = myFind
01 = myWc
02 = myXargs

all: build

build: $(00) $(01) $(02)

$(00): ex00/$(00).go
	go build -o $@ $^

$(01): ex01/$(01).go
	go build -o $@ $^

$(02): ex02/$(02).go
	go build -o $@ $^

clean:
	rm -f $(00) $(01) $(02)