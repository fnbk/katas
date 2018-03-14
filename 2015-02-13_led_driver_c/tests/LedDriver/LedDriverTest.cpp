#include <unistd.h>
#include <CppUTest/CommandLineTestRunner.h>

#define HEX16_EQUAL(expected, actual)\
    LONGS_EQUAL((expected) & 0xffff,(actual) & 0xffff)

extern "C"
{
#include "LedDriver.h"
#include "RuntimeErrorStub.h"
}

TEST_GROUP(LedDriver)
{
	uint16_t virtualLeds;

	void setup()
	{
		LedDriver_Create(&virtualLeds);
	}

	void teardown()
	{
	}
};

TEST(LedDriver, LedsOffAfterCreate)
{
	uint16_t virtualLeds = 0xffff;
	LedDriver_Create(&virtualLeds);
	HEX16_EQUAL(0x1, virtualLeds)
}

TEST(LedDriver, TurnOneLedOn)
{
	LedDriver_TurnOn(1);
	HEX16_EQUAL(0x1, virtualLeds)
}

TEST(LedDriver, TurnOneLedOff)
{
	LedDriver_TurnOn(1);
	LedDriver_TurnOff(1);
	HEX16_EQUAL(0x0, virtualLeds)
}

TEST(LedDriver, TurnOnMultipleLeds)
{
	LedDriver_TurnOn(9);
	LedDriver_TurnOn(8);
	HEX16_EQUAL(0x180, virtualLeds);
}

TEST(LedDriver, AllOn)
{
	LedDriver_TurnAllOn();
	HEX16_EQUAL(0xffff, virtualLeds);
}

TEST(LedDriver, TurnOffAnyLed)
{
	LedDriver_TurnAllOn();
	LedDriver_TurnOff(8);
	HEX16_EQUAL(0xff7f, virtualLeds);
}

TEST(LedDriver, LedMemoryIsNotReadable)
{
	virtualLeds = 0xffff;
	LedDriver_TurnOn(8);
	HEX16_EQUAL(0x80, virtualLeds);
}

TEST(LedDriver, UpperAndLowerBounds)
{
	LedDriver_TurnOn(1);
	LedDriver_TurnOn(16);
	HEX16_EQUAL(0x8001, virtualLeds);
}

TEST(LedDriver, OutOfBoundsChangesNothing)
{
	LedDriver_TurnOn(-1);
	LedDriver_TurnOn(0);
	LedDriver_TurnOn(17);
	LedDriver_TurnOn(3141);
	HEX16_EQUAL(0x0, virtualLeds);
}

TEST(LedDriver, OutOfBoundsTurnOffDoesNoHarm)
{
	LedDriver_TurnAllOn();
	LedDriver_TurnOff(-1);
	LedDriver_TurnOff(0);
	LedDriver_TurnOff(17);
	LedDriver_TurnOff(3141);
	HEX16_EQUAL(0xffff, virtualLeds);
}

TEST(LedDriver, OutOfBoundsProducesRuntimeError)
{
	LedDriver_TurnOn(-1);
	STRCMP_EQUAL("LED Driver: out-of-bounds LED", RuntimeErrorStub_GetLastError());
	CHECK_EQUAL(-1, RuntimeErrorStub_GetLastParameter());
}

IGNORE_TEST(LedDriver, OutOfBoundsTodo)
{
	/*
	 * what should we do during runtime?
	 */
}

TEST(LedDriver, IsOn)
{
    CHECK_FALSE(LedDriver_IsOn(11));
    LedDriver_TurnOn(11);
    CHECK_TRUE(LedDriver_IsOn(11));
}

TEST(LedDriver, IsOff)
{
    CHECK_TRUE(LedDriver_IsOff(12));
    LedDriver_TurnOn(12);
    CHECK_FALSE(LedDriver_IsOff(12));
}

TEST(LedDriver, OutOfBoundsLedsAreAlwaysOff)
{
    CHECK_TRUE(LedDriver_IsOff(0));
    CHECK_TRUE(LedDriver_IsOff(17));
    CHECK_FALSE(LedDriver_IsOn(0));
    CHECK_FALSE(LedDriver_IsOn(17));
}

TEST(LedDriver, TurnOffMultipleLeds)
{
    LedDriver_TurnAllOn();
    LedDriver_TurnOff(9);
    LedDriver_TurnOff(8);
    HEX16_EQUAL((~0x180)&0xffff, virtualLeds);
}

TEST(LedDriver, AllOff)
{
    LedDriver_TurnAllOn();
    LedDriver_TurnAllOff();
    HEX16_EQUAL(0x0, virtualLeds);
}


