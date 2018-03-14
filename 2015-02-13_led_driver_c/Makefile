.PHONY: all clean

all: homeautomation

stdio:
	gcc -o AllTests tests/AllTests.cpp tests/stdio/* -lCppUTest -lCppUTestExt -lstdc++

leddriver:
	gcc -o AllTests tests/AllTests.cpp tests/LedDriver/*.cpp src/LedDriver/* mocks/*.c -lCppUTest -lCppUTestExt -lstdc++ -Iinclude/LedDriver -Iinclude/util -Imocks

homeautomation:
	gcc -c src/HomeAutomation/*.c -Iinclude/HomeAutomation
	ar crs libhomeautomation.a *.o
	rm *.o
	gcc -o AllTests tests/AllTests.cpp tests/HomeAutomation/*.cpp tests/HomeAutomation/*.c -Iinclude/HomeAutomation -lCppUTest -lCppUTestExt -lstdc++ -L. -lhomeautomation

check:
	./AllTests
