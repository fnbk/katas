#ifndef D_TimeService_H
#define D_TimeService_H

struct Time {
	int minuteOfDay;
	int dayOfWeek;
};

enum {
	TIME_UNKNOWN = -1,
	MONDAY = 1,
	TUESDAY = 2,
	WEDNESDAY = 3,
	THURSDAY = 4,
	FRIDAY = 5,
	SATURDAY = 6,
	SUNDAY = 7
};

void TimeService_GetTime(struct Time *);

#endif /* D_TimeService_H */
