.PHONY: all clean

all: stdio

stdio:
	gcc -o AllTests tests/AllTests.cpp tests/SprintfTest.cpp -lCppUTest -lCppUTestExt -lstdc++

check:
	./AllTests

clean:
	rm AllTests
