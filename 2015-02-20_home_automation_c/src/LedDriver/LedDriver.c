#include <unistd.h>
#include "LedDriver.h"
#include "RuntimeError.h"

enum {
    FIRST_LED = 1,
    LAST_LED = 16
};

static uint16_t *ledsAddress;
static uint16_t ledsImage;

static uint16_t convertLedNumberToBit(int ledNumber)
{
	return 1 << (ledNumber -1);
}

static void updateHardware(void)
{
	*ledsAddress = ledsImage;
}

static int isLedOutOfBounds(int ledNumber)
{
    return (ledNumber < FIRST_LED) || (ledNumber > LAST_LED);
}

static void setLedImageBit(int ledNumber)
{
	ledsImage |= convertLedNumberToBit(ledNumber);
}

static void clearLedImageBit(int ledNumber)
{
	ledsImage &= ~(convertLedNumberToBit(ledNumber));
}

void LedDriver_Create(uint16_t * address)
{
	ledsAddress = address;
	ledsImage = 0x0;
	*ledsAddress = ledsImage;
}

void LedDriver_Destroy(void)
{
}

void LedDriver_TurnOn(int ledNumber)
{
    if (isLedOutOfBounds(ledNumber))
	{
		RUNTIME_ERROR("LED Driver: out-of-bounds LED", -1);
		return;
	}

    setLedImageBit(ledNumber);
	updateHardware();
}

void LedDriver_TurnOff(int ledNumber)
{
    if (isLedOutOfBounds(ledNumber))
            return;

    clearLedImageBit(ledNumber);
	updateHardware();
}

void LedDriver_TurnAllOn(void)
{
	ledsImage = 0xffff;
	updateHardware();
}

void LedDriver_TurnAllOff(void)
{
	ledsImage = 0x0;
	updateHardware();
}

int LedDriver_IsOn(int ledNumber)
{
    if (isLedOutOfBounds(ledNumber))
        return 0;

    return ledsImage & (convertLedNumberToBit(ledNumber));
}

int LedDriver_IsOff(int ledNumber)
{
    return !LedDriver_IsOn(ledNumber);
}

