// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package decoder_test

var (
	// til is the trace id list content for tests
	til = `[
		{
			"id": 18,
			"fmtType": "TRICE_S",
			"fmtStrg": "%s",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 1603,
			"fmtType": "TRICE0",
			"fmtStrg": "rd_:triceBareFifo.c",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 1716,
			"fmtType": "TRICE0",
			"fmtStrg": "rd_:triceBareFifoToBytesBuffer.c",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 44155,
			"fmtType": "TRICE0",
			"fmtStrg": "rd_:triceBareFifoToEscFifo.c",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 59283,
			"fmtType": "TRICE0",
			"fmtStrg": "rd_:triceCheck.c",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 24481,
			"fmtType": "TRICE0",
			"fmtStrg": "--------------------------------------------------\\r\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 13428,
			"fmtType": "TRICE0",
			"fmtStrg": "--------------------------------------------------\\r\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 18133,
			"fmtType": "TRICE16_1",
			"fmtStrg": "dbg:12345 as 16bit is %#016b\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 3791,
			"fmtType": "TRICE0",
			"fmtStrg": "--------------------------------------------------\\r\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 17125,
			"fmtType": "TRICE0",
			"fmtStrg": "sig:This ASSERT error is just a demo and no real error:\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 26797,
			"fmtType": "TRICE0",
			"fmtStrg": "--------------------------------------------------\\r\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 18113,
			"fmtType": "TRICE16_1",
			"fmtStrg": "ERR:error       message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 28289,
			"fmtType": "TRICE16_1",
			"fmtStrg": "WRN:warning     message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 53560,
			"fmtType": "TRICE16_1",
			"fmtStrg": "ATT:attention   message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 16672,
			"fmtType": "TRICE16_1",
			"fmtStrg": "DIA:diagnostics message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 42206,
			"fmtType": "TRICE16_1",
			"fmtStrg": "TIM:timing      message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 23973,
			"fmtType": "TRICE16_1",
			"fmtStrg": "DBG:debug       message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 13730,
			"fmtType": "TRICE16_1",
			"fmtStrg": "SIG:signal      message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 24302,
			"fmtType": "TRICE16_1",
			"fmtStrg": "RD:read         message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 14138,
			"fmtType": "TRICE16_1",
			"fmtStrg": "WR:write        message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 55445,
			"fmtType": "TRICE16_1",
			"fmtStrg": "ISR:interrupt   message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 41495,
			"fmtType": "TRICE16_1",
			"fmtStrg": "TST:test        message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 45697,
			"fmtType": "TRICE16_1",
			"fmtStrg": "MSG:normal      message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 24589,
			"fmtType": "TRICE16_1",
			"fmtStrg": "INFO:informal   message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 2527,
			"fmtType": "TRICE8_4",
			"fmtStrg": "tst:TRICE8  %%03x -\u003e  %03x  %03x  %03x  %03x\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 63423,
			"fmtType": "TRICE8_4",
			"fmtStrg": "tst:TRICE8   %%4d -\u003e %4d %4d %4d %4d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 40063,
			"fmtType": "TRICE8_4",
			"fmtStrg": "tst:TRICE8   %%4o -\u003e %4o %4o %4o %4o\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 8674,
			"fmtType": "TRICE16_4",
			"fmtStrg": "tst:TRICE16  %%05x -\u003e   %05x   %05x   %05x   %05x\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 59869,
			"fmtType": "TRICE16_4",
			"fmtStrg": "tst:TRICE16   %%6d -\u003e  %6d  %6d  %6d  %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 24781,
			"fmtType": "TRICE16_4",
			"fmtStrg": "tst:TRICE16   %%7o -\u003e %7o %7o %7o %7o\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 16811,
			"fmtType": "TRICE32_4",
			"fmtStrg": "tst:TRICE32_4 %%09x -\u003e      %09x      %09x       %09x     %09x\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 28944,
			"fmtType": "TRICE32_4",
			"fmtStrg": "tst:TRICE32_4 %%10d -\u003e     %10d     %10d     %10d    %10x\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 19049,
			"fmtType": "TRICE64_1",
			"fmtStrg": "att:64bit %#b\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 41267,
			"fmtType": "TRICE8_1",
			"fmtStrg": "tst:TRICE8_1 %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 28873,
			"fmtType": "TRICE8_2",
			"fmtStrg": "tst:TRICE8_2 %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 18501,
			"fmtType": "TRICE8_3",
			"fmtStrg": "tst:TRICE8_3 %d %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 34892,
			"fmtType": "TRICE8_4",
			"fmtStrg": "tst:TRICE8_4 %d %d %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 2043,
			"fmtType": "TRICE8_5",
			"fmtStrg": "tst:TRICE8_5 %d %d %d %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 36242,
			"fmtType": "TRICE8_6",
			"fmtStrg": "tst:TRICE8_6 %d %d %d %d %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 23187,
			"fmtType": "TRICE8_7",
			"fmtStrg": "tst:TRICE8_7 %d %d %d %d %d %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 5279,
			"fmtType": "TRICE8_8",
			"fmtStrg": "tst:TRICE8_8 %d %d %d %d %d %d %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 50977,
			"fmtType": "TRICE16_1",
			"fmtStrg": "tst:TRICE16_1 %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 65136,
			"fmtType": "TRICE16_2",
			"fmtStrg": "tst:TRICE16_2 %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 58915,
			"fmtType": "TRICE16_3",
			"fmtStrg": "tst:TRICE16_3 %d %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 45117,
			"fmtType": "TRICE16_4",
			"fmtStrg": "tst:TRICE16_4 %d %d %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 6384,
			"fmtType": "TRICE32_1",
			"fmtStrg": "tst:TRICE32_1 %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 11505,
			"fmtType": "TRICE32_2",
			"fmtStrg": "tst:TRICE32_2 %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 56189,
			"fmtType": "TRICE32_3",
			"fmtStrg": "tst:TRICE32_3 %d %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 36573,
			"fmtType": "TRICE32_4",
			"fmtStrg": "tst:TRICE32_4 %d %d %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 43171,
			"fmtType": "TRICE64_1",
			"fmtStrg": "tst:TRICE64_1 %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 8402,
			"fmtType": "TRICE64_2",
			"fmtStrg": "tst:TRICE64_2 %d %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 44690,
			"fmtType": "TRICE0",
			"fmtStrg": "e:A",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 1471,
			"fmtType": "TRICE0",
			"fmtStrg": "w:B",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 52563,
			"fmtType": "TRICE0",
			"fmtStrg": "a:c",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 61604,
			"fmtType": "TRICE0",
			"fmtStrg": "wr:d",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 39587,
			"fmtType": "TRICE0",
			"fmtStrg": "rd:e\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 53436,
			"fmtType": "TRICE0",
			"fmtStrg": "diag:f",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 32678,
			"fmtType": "TRICE0",
			"fmtStrg": "d:G",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 13587,
			"fmtType": "TRICE0",
			"fmtStrg": "t:H",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 8059,
			"fmtType": "TRICE0",
			"fmtStrg": "time:i",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 26357,
			"fmtType": "TRICE0",
			"fmtStrg": "message:J",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 2812,
			"fmtType": "TRICE0",
			"fmtStrg": "dbg:k\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 43782,
			"fmtType": "TRICE0",
			"fmtStrg": "1",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 31945,
			"fmtType": "TRICE0",
			"fmtStrg": "2",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 57633,
			"fmtType": "TRICE0",
			"fmtStrg": "3",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 32117,
			"fmtType": "TRICE0",
			"fmtStrg": "4",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 3996,
			"fmtType": "TRICE0",
			"fmtStrg": "e:7",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 6559,
			"fmtType": "TRICE0",
			"fmtStrg": "m:12",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 30479,
			"fmtType": "TRICE0",
			"fmtStrg": "m:123\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 41878,
			"fmtType": "TRICE16_1",
			"fmtStrg": "err:error       message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 3242,
			"fmtType": "TRICE16_1",
			"fmtStrg": "wrn:warning     message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 10588,
			"fmtType": "TRICE16_1",
			"fmtStrg": "att:attention   message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 15045,
			"fmtType": "TRICE16_1",
			"fmtStrg": "dia:diagnostics message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 20195,
			"fmtType": "TRICE16_1",
			"fmtStrg": "tim:timing      message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 56298,
			"fmtType": "TRICE16_1",
			"fmtStrg": "dbg:debug       message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 62837,
			"fmtType": "TRICE16_1",
			"fmtStrg": "sig:signal      message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 35517,
			"fmtType": "TRICE16_1",
			"fmtStrg": "rd_:read        message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 47873,
			"fmtType": "TRICE16_1",
			"fmtStrg": "wr:write        message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 64995,
			"fmtType": "TRICE16_1",
			"fmtStrg": "isr:interrupt   message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 42962,
			"fmtType": "TRICE16_1",
			"fmtStrg": "tst:test        message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 59159,
			"fmtType": "TRICE16_1",
			"fmtStrg": "msg:normal      message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 48090,
			"fmtType": "TRICE16_1",
			"fmtStrg": "info:informal   message, SysTick is %6d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 55237,
			"fmtType": "TRICE0",
			"fmtStrg": "e:A",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 38195,
			"fmtType": "TRICE0",
			"fmtStrg": "w:B",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 43063,
			"fmtType": "TRICE0",
			"fmtStrg": "a:c",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 6176,
			"fmtType": "TRICE0",
			"fmtStrg": "wr:d",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 6665,
			"fmtType": "TRICE0",
			"fmtStrg": "rd:e",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 47959,
			"fmtType": "TRICE0",
			"fmtStrg": "diag:y",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 48390,
			"fmtType": "TRICE0",
			"fmtStrg": "d:G",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 64481,
			"fmtType": "TRICE0",
			"fmtStrg": "t:H",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 19868,
			"fmtType": "TRICE0",
			"fmtStrg": "time:i",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 9207,
			"fmtType": "TRICE0",
			"fmtStrg": "message:J",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 29302,
			"fmtType": "TRICE0",
			"fmtStrg": "inf:k\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 21100,
			"fmtType": "TRICE0",
			"fmtStrg": "att:Next few lines have no color and visible only when TRICE_RUNTIME_GENERATED_STRINGS_SUPPORT is defined\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 2502,
			"fmtType": "TRICE8_8",
			"fmtStrg": "msg:1:%03x %03x %03x %03x %03x %03x %03x %03x\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 20625,
			"fmtType": "TRICE16_1",
			"fmtStrg": "tim: pre encryption SysTick=%d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 64368,
			"fmtType": "TRICE16_1",
			"fmtStrg": "tim: post encryption SysTick=%d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 43335,
			"fmtType": "TRICE8_8",
			"fmtStrg": "att:1:%03x %03x %03x %03x %03x %03x %03x %03x\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 56697,
			"fmtType": "TRICE16_1",
			"fmtStrg": "tim: pre decryption SysTick=%d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 50168,
			"fmtType": "TRICE16_1",
			"fmtStrg": "tim: post decryption SysTick=%d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 990,
			"fmtType": "TRICE8_8",
			"fmtStrg": "msg:2:%03x %03x %03x %03x %03x %03x %03x %03x\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 7275,
			"fmtType": "TRICE0",
			"fmtStrg": "--------------------------------------------------\\r\\n\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 18577,
			"fmtType": "TRICE32_1",
			"fmtStrg": "ISR:alive time %d milliseconds\\n",
			"created": 1590414933,
			"removed": 0
		},		
		{
			"id": 44860,
			"fmtType": "TRICE0",
			"fmtStrg": "rd_:triceEscFifo.c",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 20913,
			"fmtType": "TRICE0",
			"fmtStrg": "s:                                        \\ns:   ARM-MDK_LL_UART_BARE_NUCLEO-F030RB   \\ns:                                        \\n\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 41488,
			"fmtType": "TRICE16_1",
			"fmtStrg": "MSG: triceBareFifoMaxDepth = %d\\n",
			"created": 1601240704,
			"removed": 0
		},
		{
			"id": 54442,
			"fmtType": "TRICE0",
			"fmtStrg": "s:                                        \\ns:   ARM-MDK_LL_UART_BARE_NUCLEO-F070RB   \\ns:                                        \\n\\n",
			"created": 1601240705,
			"removed": 0
		},
		{
			"id": 51601,
			"fmtType": "TRICE64_2",
			"fmtStrg": "MSG: triceBareFifoMaxDepth = %d (index %d)\\n",
			"created": 1601240705,
			"removed": 0
		},
		{
			"id": 37259,
			"fmtType": "TRICE0",
			"fmtStrg": "s:                                        \\ns:   ARM-MDK_LL_UART_BARE_TO_ESC_NUCLEO-F030R8   \\ns:                                        \\n\\n",
			"created": 1601240705,
			"removed": 0
		},
		{
			"id": 38143,
			"fmtType": "TRICE16_2",
			"fmtStrg": "MSG: triceFifoMaxDepth: Bare = %d, Esc = %d\\n",
			"created": 1601240705,
			"removed": 0
		},
		{
			"id": 33467,
			"fmtType": "TRICE0",
			"fmtStrg": "s:                                        \\ns:   ARM-MDK_LL_UART_BARE_TO_ESC_NUCLEO-F070RB   \\ns:                                        \\n\\n",
			"created": 1601240705,
			"removed": 0
		},
		{
			"id": 60242,
			"fmtType": "TRICE0",
			"fmtStrg": "\\ns:                                        \\ns:   ARM-MDK_LL_UART_ESC_NUCLEO-F030R8    \\ns:                                        \\n\\n",
			"created": 1601240705,
			"removed": 0
		},
		{
			"id": 45495,
			"fmtType": "TRICE64_2",
			"fmtStrg": "MSG: triceEscFifoMaxDepth = %d, index = %d\\n",
			"created": 1601240705,
			"removed": 0
		},
		{
			"id": 62429,
			"fmtType": "TRICE0",
			"fmtStrg": "s:                                        \\ns:   ARM-MDK_LL_UART_ESC_NUCLEO-F070RB    \\ns:                                        \\n\\n",
			"created": 1601240705,
			"removed": 0
		},
		{
			"id": 55132,
			"fmtType": "TRICE16_1",
			"fmtStrg": "MSG: triceEscFifoMaxDepth = %d\\n",
			"created": 1601240705,
			"removed": 0
		},
		{
			"id": 16627,
			"fmtType": "TRICE64_2",
			"fmtStrg": "tst:TRICE64_2 %x %16d\\n",
			"created": 1601240705,
			"removed": 0
		},
		{
			"id": 10566,
			"fmtType": "TRICE0",
			"fmtStrg": "s:                                        \\ns:   ARM-MDK_LL_BARE_STM32F0308-DISCO   \\ns:                                        \\n\\n",
			"created": 1601307082,
			"removed": 0
		},
		{
			"id": 9549,
			"fmtType": "TRICE0",
			"fmtStrg": "s:                                        \\ns:    ARM-MDK_LL_BARE_STM32F0308-DISCO    \\ns:                                        \\n\\n",
			"created": 1601308021,
			"removed": 0
		},
		{
			"id": 15835,
			"fmtType": "TRICE0",
			"fmtStrg": "s:                                              \\ns:    ARM-MDK_BARE_STM32F03051R8Tx-DISCOVERY    \\ns:                                              \\n\\n",
			"created": 1602255477,
			"removed": 0
		},
		{
			"id": 23971,
			"fmtType": "TRICE0",
			"fmtStrg": "s:                                              \\ns:    ARM-MDK_BARE_STM32F03051R8Tx-DISCOVERY    \\ns:                                              \\n\\n",
			"created": 1602255678,
			"removed": 0
		},
		{
			"id": 44249,
			"fmtType": "TRICE0",
			"fmtStrg": "Hallo Ralf!\\n",
			"created": 1604512008,
			"removed": 0
		},
		{
			"id": 10777,
			"fmtType": "TRICE0",
			"fmtStrg": "s:                                              \\ns:   MDK-ARM_LL_UART_RTT0_BARE_STM32F091_NUCLEO-64   \\ns:                                              \\n\\n",
			"created": 1607373562,
			"removed": 0
		}
	]`
)
