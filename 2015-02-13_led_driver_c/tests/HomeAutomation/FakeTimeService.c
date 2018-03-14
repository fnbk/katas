#include "FakeTimeService.h"

/*
 * FakeTimeService
 */

struct Time _time;

void FakeTimeService_Destroy()
{
    _time.minuteOfDay = TIME_UNKNOWN;
    _time.dayOfWeek   = TIME_UNKNOWN;
}

void FakeTimeService_Create()
{
    _time.minuteOfDay = TIME_UNKNOWN;
    _time.dayOfWeek   = TIME_UNKNOWN;
}

void FakeTimeService_SetMinute(int minute)
{
    _time.minuteOfDay = minute;
}

void FakeTimeService_SetDay(int day)
{
    _time.dayOfWeek = day;
}

/*
 * TimeService
 */

void TimeService_GetTime(struct Time * time)
{
    time->minuteOfDay = _time.minuteOfDay;
    time->dayOfWeek   = _time.dayOfWeek;
}
