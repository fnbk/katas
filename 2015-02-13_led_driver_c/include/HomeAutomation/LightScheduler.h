#ifndef D_LightScheduler_H
#define D_LightScheduler_H

enum Day {
	EVERYDAY = 10
};

enum {
	UNUSED = -1
};

void LightScheduler_Create(void);
void LightScheduler_Destroy(void);
void LightScheduler_ScheduleTurnOn(int id, int day, int minuteOfDay);
void LightScheduler_Wakeup(void);

#endif /* D_LightScheduler_H */
