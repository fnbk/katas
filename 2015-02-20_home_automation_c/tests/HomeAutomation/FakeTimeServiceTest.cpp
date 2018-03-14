#include <CppUTest/CommandLineTestRunner.h>

extern "C"
{
#include "FakeTimeService.h"
}

TEST_GROUP(FakeTimeService)
{
    void setup()
    {
        FakeTimeService_Create();
    }

    void teardown()
    {
        FakeTimeService_Destroy();
    }
};

TEST(FakeTimeService, Set)
{
    struct Time time;
    FakeTimeService_SetMinute(42);
    FakeTimeService_SetDay(SATURDAY);
    TimeService_GetTime(&time);
    LONGS_EQUAL(42, time.minuteOfDay);
    LONGS_EQUAL(SATURDAY, time.dayOfWeek);
}


TEST(FakeTimeService, Create)
{
    struct Time time;
    TimeService_GetTime(&time);
    LONGS_EQUAL(TIME_UNKNOWN, time.minuteOfDay);
    LONGS_EQUAL(TIME_UNKNOWN, time.dayOfWeek);
}
