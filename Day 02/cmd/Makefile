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

run: run_$(00) run_$(01) run_$(02)

run_$(00):
	./$(00) -sl /home/Nebomoe/Downloads
	./$(00) -d /home/Nebomoe/Downloads
	./$(00) -f -ext 'txt' /home/Nebomoe/Downloads

run_$(01):
	./$(01) -w /home/Nebomoe/Downloads/text1.txt /home/Nebomoe/Downloads/text2.txt
	./$(01) -l /home/Nebomoe/Downloads/text1.txt /home/Nebomoe/Downloads/text2.txt
	./$(01) -m /home/Nebomoe/Downloads/text1.txt /home/Nebomoe/Downloads/text2.txt

run_$(02):
	echo /home/Nebomoe/Downloads | ./$(02) ls -la
	./$(00) -f -ext 'txt' /home/Nebomoe/Downloads | ./$(02) ./$(01) -l

clean:
	rm -f $(00) $(01) $(02)