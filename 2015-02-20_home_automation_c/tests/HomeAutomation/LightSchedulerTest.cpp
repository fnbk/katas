#include <CppUTest/CommandLineTestRunner.h>

extern "C"
{
#include "LightControllerSpy.h"
#include "LightScheduler.h"
#include "FakeTimeService.h"
}

TEST_GROUP(LightScheduler)
{
    void setup()
    {
        LightController_Create();
        LightScheduler_Create();
    }

    void teardown()
    {
        LightController_Destroy();
        LightScheduler_Destroy();
    }
};

TEST(LightScheduler, ScheduleOnEverydayNotTimeYet)
{
    LightScheduler_ScheduleTurnOn(3, EVERYDAY, 1200);
    FakeTimeService_SetDay(MONDAY);
    FakeTimeService_SetMinute(1199);
    LightScheduler_Wakeup();

    LONGS_EQUAL(LIGHT_ID_UNKNOWN, LightControllerSpy_GetLastId());
    LONGS_EQUAL(LIGHT_STATE_UNKNOWN, LightControllerSpy_GetLastState());
}

// IGNORE_TEST(LightScheduler, NoScheduleNothingHappens)
// {
//     FakeTimeService_SetDay(MONDAY);
//     FakeTimeService_SetMinute(100);
//     LightScheduler_Wakeup();
//     LONGS_EQUAL(LIGHT_ID_UNKNOWN, LightControllerSpy_GetLastId());
//     LONGS_EQUAL(LIGHT_STATE_UNKNOWN, LightControllerSpy_GetLastState());
// }
//

TEST(LightScheduler, ScheduleOnEverydayItsTime)
{
	LightScheduler_ScheduleTurnOn(3, EVERYDAY, 1200);
	FakeTimeService_SetDay(MONDAY);
	FakeTimeService_SetMinute(1200);

	LightScheduler_Wakeup();

	LONGS_EQUAL(3, LightControllerSpy_GetLastId());
	LONGS_EQUAL(LIGHT_ON, LightControllerSpy_GetLastState());
}

