#include <CppUTest/CommandLineTestRunner.h>


/*
 *
 */

typedef enum {
	BFALSE = 0,
	BTRUE = !0
} bool_t;

typedef enum {
	u8GlobalZero  = (uint8_t)0U,	/**!  8bit zero value */
	u16GlobalZero = (uint16_t)0U,	/**! 16bit zero value */
	u32GlobalZero = (uint32_t)0U,	/**! 32bit zero value */
	u8GlobalMax   = 0xFFU,		/**!  8bit unsigned maximum value */
	u16GlobalMax  = 0xFFFFU		/**! 16bit unsigned maximum value */
} Globals_LimitOfTypes;


void sample_out(void *target, size_t target_size, uint8_t sample)
{
	uint8_t *tgt = (uint8_t*)target;
	for (int i=0; i < target_size; i++)
	{
		tgt[i] = sample;
	}
}

bool_t is_sampled_out(void *target, size_t target_size, uint8_t sample)
{
	bool_t ret = BTRUE;
	uint8_t *tgt = (uint8_t*)target;
	for (int i=0; i < target_size; i++)
	{
		if (tgt[i] != sample)
		{
			ret = BFALSE;
			break;
		}
	}
	return ret;
}

TEST_GROUP(sample_out)
{
	void setup()
	{
	}

	void teardown()
	{
	}
};

TEST(sample_out, AllElementsZero)
{
	uint8_t len = 10;
	uint8_t buff[len];
	sample_out((void*)&buff, len, 0);
	bool_t res = is_sampled_out((void*)&buff, len, 0);
	CHECK_EQUAL(BTRUE, res);
}

TEST(sample_out, AllElementsOne)
{
	uint8_t len = 10;
	uint8_t buff[len];
	sample_out((void*)&buff, len, 1);
	bool_t res = is_sampled_out((void*)&buff, len, 1);
	CHECK_EQUAL(BTRUE, res);
}

TEST(sample_out, AllElementsZeroExceptOne)
{
	uint8_t buff[2];
	sample_out((void*)&buff, 2, 1);
	buff[0] = 0;
	CHECK_EQUAL(0, buff[0]);
	CHECK_EQUAL(1, buff[1]);
}

//
// TEST_GROUP(fizzbuzz)
// {
// 	char output[100];
// 	const char *expected;
//
// 	void setup()
// 	{
// 		memset(output, 0xaa, sizeof(output));
// 		expected = "";
// 	}
//
// 	void teardown()
// 	{
// 	}
//
// 	void expect(const char *s)
// 	{
// 		expected = s;
// 	}
//
// 	void given(int charsWritten)
// 	{
// 		LONGS_EQUAL(strlen(expected), charsWritten);
// 		STRCMP_EQUAL(expected, output);
// 		BYTES_EQUAL(0xaa, output[strlen(expected) + 1]);
// 	}
// };
//
// TEST(fizzbuzz, FizzForMultiplesOfThree)
// {
// 	//char output[100][10] = {0};
// 	//sprintf(output[2], "fizz");
// 	//STRCMP_EQUAL(output[2], "fizz");
//
// 	char output[100][10] = {0};
// 	strcpy(output[2],"fizz");
// 	STRCMP_EQUAL(output[2], "fizz");
// 	
// 	//fizzbuzz(&output);
// }
