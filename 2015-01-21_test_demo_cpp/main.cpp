#include <CppUTest/CommandLineTestRunner.h>

int main(int ac, char** av)
{
	return CommandLineTestRunner::RunAllTests(ac, av);
}

TEST_GROUP(FirstTestGroup)
{
};

TEST(FirstTestGroup, StringCompare)
{
	STRCMP_EQUAL("hello", "hello2world");
}

TEST(FirstTestGroup, LongEqual)
{
   	LONGS_EQUAL(1, 2);
}

TEST(FirstTestGroup, Check)
{
	CHECK(false);
}

TEST(FirstTestGroup, Fail)
{
	FAIL("Fail me!");
}

