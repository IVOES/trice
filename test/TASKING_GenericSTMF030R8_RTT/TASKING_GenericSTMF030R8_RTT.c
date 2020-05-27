/*****************************************************
 *
 * TASKING_GenericSTMF030R8_RTT.c
 *
 * Description : Hello World in C, ANSI-style
 *
 */

#include "SEGGER_RTT.h"
#include "trice.h"

int main(void)
{
	SEGGER_RTT_printf( 0, "Hello world\n" );
	*(uint32_t*)0xE000E010 = 0;        // SysTick->CTRL = 0; // Disable SysTick
	*(uint32_t*)0xE000E014 = 0xFFFFFF; // SysTick->LOAD = 999; // Count down from 999 to 0
	*(uint32_t*)0xE000E018 = 0;        // SysTick->VAL = 0; // Clear current value to 0
	*(uint32_t*)0xE000E010 = 0x5;      // SysTick->CTRL = 0x7; // Enable SysTick, and use processor clock
	// exception and use processor clock
    TRICE_RTTD_SYNC;
    TRICE16_1( Id(45373), "tim:timing      message, SysTick is %6d\n", SYSTICKVAL16 );
    TRICE0( Id(35421), "att:TASKING_GenericSTMF030R8_RTT\n" );
    TRICE16_1( Id(45373), "tim:timing      message, SysTick is %6d\n", SYSTICKVAL16 );
    TRICE16_1( Id(45373), "tim:timing      message, SysTick is %6d\n", SYSTICKVAL16 );
    TRICE16_1( Id(45373), "tim:timing      message, SysTick is %6d\n", SYSTICKVAL16 );
    TRICE16_1( Id(45373), "tim:timing      message, SysTick is %6d\n", SYSTICKVAL16 );
    for(;;){
    	triceServe();
    }
}
