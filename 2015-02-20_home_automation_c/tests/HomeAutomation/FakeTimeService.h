#ifndef D_FakeTimeService_H
#define D_FakeTimeService_H

#include "TimeService.h"

void FakeTimeService_SetMinute(int);
void FakeTimeService_SetDay(int);
void FakeTimeService_Create(void);
void FakeTimeService_Destroy(void);

#endif /* D_FakeTimeService_H */
