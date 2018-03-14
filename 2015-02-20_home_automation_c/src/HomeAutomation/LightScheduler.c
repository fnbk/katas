#include "LightController.h"
#include "LightScheduler.h"
#include "TimeService.h"

struct ScheduledLightEvent
{
	int id;
	int minuteOfDay;
};

static struct ScheduledLightEvent scheduledEvent;

void LightScheduler_Create(void)
{
	scheduledEvent.id = UNUSED;
}

void LightScheduler_Destroy(void)
{
}

void LightScheduler_ScheduleTurnOn(int id, int day, int minuteOfDay)
{
	scheduledEvent.id = id;
	scheduledEvent.minuteOfDay = minuteOfDay;
}

void LightScheduler_Wakeup(void)
{
	struct Time time;
	TimeService_GetTime(&time);

	if(scheduledEvent.id == UNUSED)
		return;
	if(time.minuteOfDay != scheduledEvent.minuteOfDay)
		return;

	LightController_On(scheduledEvent.id);
}


