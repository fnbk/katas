.PHONY: all clean

all: fizzbuzz

fizzbuzz:
	gcc -o AllTests tests/AllTests.cpp tests/RingBufferTest.cpp -lCppUTest -lCppUTestExt -lstdc++

check:
	./AllTests

clean:
	rm AllTests
