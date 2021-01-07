// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package decoder

/*
var (
	// byteStreamEsc are the raw read input bytes for esc format tests
	byteStreamEsc = string([]byte{
		236, 228, 113, 16, 0, 0, 0, 1, 127, 255, 255, 255, 128, 0, 0, 0, 255, 255, 255, 255, // TRICE32_4 %10d ->              1     2147483647     -2147483648            -1
		236, 227, 74, 105, 17, 34, 51, 68, 85, 102, 119, 136, // 64bit 0b1000100100010001100110100010001010101011001100111011110001000
		236, 228, 177, 183, 0, 0, 0, 0, 0, 0, 0, 129, 0, 0, 0, 0, 0, 0, 0, 3, // MSG: triceEscFifoMaxDepth = 129, index = 3
		236, 224, 161, 51, 255, // TRICE8_1 -1
		236, 228, 177, 183, 0, 0, 0, 0, 0, 0, 0, 98, 0, 0, 0, 0, 0, 0, 0, 2, // MSG: triceEscFifoMaxDepth = 98, index = 2
		236, 225, 173, 219, 0, 4, // MSG: triceFifoMaxDepth = 4\n
		236, 223, 95, 161,
		236, 223, 52, 116,
		236, 225, 70, 213, 48, 57,
		236, 223, 14, 207,
		236, 223, 66, 229,
		236, 223, 104, 173,
		236, 225, 173, 219, 0, 32,
		236, 225, 70, 193, 175, 231,
		236, 225, 110, 129, 167, 192,
		236, 225, 209, 56, 159, 147,
		236, 225, 65, 32, 151, 108,
		236, 225, 164, 222, 143, 63,
		236, 225, 93, 165, 135, 19,
		236, 225, 173, 219, 0, 42,
		236, 225, 53, 162, 175, 237,
		236, 225, 94, 238, 167, 192,
		236, 225, 55, 58, 159, 153,
		236, 225, 216, 149, 151, 108,
		236, 225, 162, 23, 143, 69,
		236, 225, 178, 129, 135, 24,
		236, 225, 96, 13, 126, 241,
		236, 225, 173, 219, 0, 48,
		236, 226, 9, 223, 1, 127, 128, 255,
		236, 226, 247, 191, 1, 127, 128, 255,
		236, 226, 156, 127, 1, 127, 128, 255,
		236, 227, 33, 226, 0, 1, 127, 255, 128, 0, 255, 255, // tst:TRICE16  %05x ->   00001   07fff   -8000   -0001\n
		236, 227, 233, 221, 0, 1, 127, 255, 128, 0, 255, 255, // tst:TRICE16   %6d ->       1   32767  -32768      -1\n
		236, 227, 96, 205, 0, 1, 127, 255, 128, 0, 255, 255, //  tst:TRICE16   %7o ->       1   77777 -100000      -1\n
		236, 225, 173, 219, 0, 66,
		236, 228, 65, 171, 0, 0, 0, 1, 127, 255, 255, 255, 128, 0, 0, 0, 255, 255, 255, 255, //  tst:TRICE32_4 %09x ->      000000001      07fffffff       -80000000     -00000001\n
		236, 228, 113, 16, 0, 0, 0, 1, 127, 255, 255, 255, 128, 0, 0, 0, 255, 255, 255, 255, //  tst:TRICE32_4 %10d ->              1     2147483647     -2147483648            -1\n
		236, 227, 74, 105, 17, 34, 51, 68, 85, 102, 119, 136, // att:64bit 0b1000100100010001100110100010001010101011001100111011110001000\n
		236, 225, 173, 219, 0, 66,
		236, 224, 161, 51, 255,
		236, 225, 112, 201, 255, 254,
		236, 226, 72, 69, 255, 254, 253, 0,
		236, 226, 136, 76, 255, 254, 253, 252,
		236, 227, 7, 251, 255, 254, 253, 252, 251, 0, 0, 0,
		236, 227, 141, 146, 255, 254, 253, 252, 251, 250, 0, 0,
		236, 227, 90, 147, 255, 254, 253, 252, 251, 250, 249, 0,
		236, 227, 20, 159, 255, 254, 253, 252, 251, 250, 249, 248,
		236, 225, 199, 33, 255, 255,
		236, 226, 254, 112, 255, 255, 255, 254, // tst:TRICE16_2 -1 -2\n
		236, 227, 230, 35, 255, 255, 255, 254, 255, 253, 0, 0, // tst:TRICE16_3 -1 -2 -3\n
		236, 227, 176, 61, 255, 255, 255, 254, 255, 253, 255, 252,
		236, 225, 173, 219, 0, 81,
		236, 226, 24, 240, 255, 255, 255, 255,
		236, 227, 44, 241, 255, 255, 255, 255, 255, 255, 255, 254,
		236, 228, 219, 125, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 253, 0, 0, 0, 0,
		236, 228, 142, 221, 255, 255, 255, 255, 255, 255, 255, 254, 255, 255, 255, 253, 255, 255, 255, 252,
		236, 225, 173, 219, 0, 81,
		236, 227, 168, 163, 255, 255, 255, 255, 255, 255, 255, 255,
		236, 228, 32, 210, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 254,
		236, 225, 173, 219, 0, 81,
		236, 223, 174, 146,
		236, 223, 5, 191,
		236, 223, 205, 83,
		236, 223, 240, 164,
		236, 223, 154, 163,
		236, 223, 208, 188,
		236, 223, 127, 166,
		236, 223, 53, 19,
		236, 223, 31, 123,
		236, 223, 102, 245,
		236, 223, 10, 252,
		236, 225, 173, 219, 0, 81,
		236, 223, 171, 6,
		236, 223, 124, 201,
		236, 223, 225, 33,
		236, 223, 125, 117,
		236, 223, 15, 156,
		236, 223, 25, 159,
		236, 223, 119, 15,
		236, 225, 173, 219, 0, 81,
		236, 225, 173, 219, 0, 81,
		236, 225, 173, 219, 0, 81,
		236, 224, 0, 18, 0,
		236, 225, 0, 18, 10, 0,
		236, 226, 0, 18, 49, 10, 0, 0,
		236, 226, 0, 18, 49, 50, 10, 0,
		236, 227, 0, 18, 49, 50, 51, 10, 0, 0, 0, 0,
		236, 227, 0, 18, 49, 50, 51, 52, 10, 0, 0, 0,
		236, 227, 0, 18, 49, 50, 51, 52, 53, 10, 0, 0,
		236, 227, 0, 18, 49, 50, 51, 52, 53, 54, 10, 0,
		236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 10, 0, 0, 0, 0, 0, 0, 0, 0,
		236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 10, 0, 0, 0, 0, 0, 0, 0,
		236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 10, 0, 0, 0, 0, 0, 0,
		236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 10, 0, 0, 0, 0, 0,
		236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 10, 0, 0, 0, 0,
		236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 10, 0, 0, 0,
		236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 10, 0, 0,
		236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 10, 0,
	})
)

func TestEsc(t *testing.T) {

	// rc is created ReadCloser
	rc, err := receiver.NewReader("BUFFER", byteStreamEsc)
	if err != nil {
		t.Fail()
	}

	list, err := UnmarshalTriceIDList([]byte(til))
	if err != nil {
		t.Fail()
	}

	p := NewEscDecoder(list, rc, bigEndian) // p is a new decoder instance

	ss := make([]string, 100)
	n, err := p.StringsRead(ss)
	if err != nil {
		t.Fail()
	}
	ss = ss[:n]
	act := fmt.Sprintln(ss)
	exp := "[tst:TRICE32_4 %10d ->              1     2147483647     -2147483648            -1\\n att:64bit 0b1000100100010001100110100010001010101011001100111011110001000\\n MSG: triceEscFifoMaxDepth = 129, index = 3\\n tst:TRICE8_1 -1\\n MSG: triceEscFifoMaxDepth = 98, index = 2\\n MSG: triceFifoMaxDepth = 4\\n --------------------------------------------------\\r\\n --------------------------------------------------\\r\\n dbg:12345 as 16bit is 0b0011000000111001\\n --------------------------------------------------\\r\\n sig:This ASSERT error is just a demo and no real error:\\n --------------------------------------------------\\r\\n MSG: triceFifoMaxDepth = 32\\n ERR:error       message, SysTick is -20505\\n WRN:warning     message, SysTick is -22592\\n ATT:attention   message, SysTick is -24685\\n DIA:diagnostics message, SysTick is -26772\\n TIM:timing      message, SysTick is -28865\\n DBG:debug       message, SysTick is -30957\\n MSG: triceFifoMaxDepth = 42\\n SIG:signal      message, SysTick is -20499\\n RD:read         message, SysTick is -22592\\n WR:write        message, SysTick is -24679\\n ISR:interrupt   message, SysTick is -26772\\n TST:test        message, SysTick is -28859\\n MSG:normal      message, SysTick is -30952\\n INFO:informal   message, SysTick is  32497\\n MSG: triceFifoMaxDepth = 48\\n tst:TRICE8  %03x ->  001  07f  -80  -01\\n tst:TRICE8   %4d ->    1  127 -128   -1\\n tst:TRICE8   %4o ->    1  177 -200   -1\\n tst:TRICE16  %05x ->   00001   07fff   -8000   -0001\\n tst:TRICE16   %6d ->       1   32767  -32768      -1\\n tst:TRICE16   %7o ->       1   77777 -100000      -1\\n MSG: triceFifoMaxDepth = 66\\n tst:TRICE32_4 %09x ->      000000001      07fffffff       -80000000     -00000001\\n tst:TRICE32_4 %10d ->              1     2147483647     -2147483648            -1\\n att:64bit 0b1000100100010001100110100010001010101011001100111011110001000\\n MSG: triceFifoMaxDepth = 66\\n tst:TRICE8_1 -1\\n tst:TRICE8_2 -1 -2\\n tst:TRICE8_3 -1 -2 -3\\n tst:TRICE8_4 -1 -2 -3 -4\\n tst:TRICE8_5 -1 -2 -3 -4 -5\\n tst:TRICE8_6 -1 -2 -3 -4 -5 -6\\n tst:TRICE8_7 -1 -2 -3 -4 -5 -6 -7\\n]\n"
	assert.Equal(t, exp, act)
}

func checkDynString(t *testing.T, list []id.Item, in, exp string) {
	// rc is created ReadCloser
	rc, err := receiver.NewReader("BUFFER", in)
	if err != nil {
		t.Fail()
	}

	p := NewEscDecoder(list, rc, bigEndian) // p is a new decoder instance

	ss := make([]string, 100)
	n, err := p.StringsRead(ss)
	if err != nil {
		t.Fail()
	}
	ss = ss[:n]
	act := fmt.Sprintln(ss)

	assert.Equal(t, exp, act)
}

func TestEscDynStrings(t *testing.T) {
	list, err := UnmarshalTriceIDList([]byte(til))
	if err != nil {
		t.Fail()
	}

	var in, exp string

	in = string([]byte{236, 223, 0, 18})
	exp = "[]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 224, 0, 18, 10})
	exp = "[\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 225, 0, 18, 49, 10})
	exp = "[1\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 223, 0, 18, 236, 224, 0, 18, 10, 236, 225, 0, 18, 49, 10})
	exp = "[\n1\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 226, 0, 18, 49, 50, 10, 0})
	exp = "[12\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 226, 0, 18, 49, 50, 51, 10})
	exp = "[123\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 227, 0, 18, 49, 50, 51, 52, 10, 0, 0, 0})
	exp = "[1234\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 227, 0, 18, 49, 50, 51, 52, 53, 10, 0, 0})
	exp = "[12345\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 227, 0, 18, 49, 50, 51, 52, 53, 54, 10, 0})
	exp = "[123456\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 227, 0, 18, 49, 50, 51, 52, 53, 54, 55, 10})
	exp = "[1234567\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 10, 0, 0, 0, 0, 0, 0, 0})
	exp = "[12345678\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 10, 0, 0, 0, 0, 0, 0})
	exp = "[123456789\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 10, 0, 0, 0, 0, 0})
	exp = "[123456789a\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 98, 10, 0, 0, 0, 0})
	exp = "[123456789ab\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 98, 99, 10, 0, 0, 0})
	exp = "[123456789abc\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 98, 99, 100, 10, 0, 0})
	exp = "[123456789abcd\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 98, 99, 100, 101, 10, 0})
	exp = "[123456789abcde\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 10, 0})
	exp = "[123456789ABCDE\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 228, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 10})
	exp = "[123456789ABCDEF\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 229, 0, 18, 49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 10, 0})
	exp = "[123456789ABCDEFGHIJKLMNOPQRSTU\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 229, 0, 18,
		49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 10})
	exp = "[123456789ABCDEFGHIJKLMNOPQRSTUV\n]\n"
	checkDynString(t, list, in, exp)

	in = string([]byte{236, 230, 0, 18,
		49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
		10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	exp = "[123456789ABCDEFGHIJKLMNOPQRSTUVW\n]\n"
	checkDynString(t, list, in, exp)

		in = string([]byte{236, 231, 0, 18,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		exp = "[123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW\n]\n"
		checkDynString(t, list, in, exp)

		in = string([]byte{236, 231, 0, 18,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 0, 0})
		exp = "[123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTU]\n"
		checkDynString(t, list, in, exp)

		in = string([]byte{236, 231, 0, 18,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 0})
		exp = "[123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUV]\n"
		checkDynString(t, list, in, exp)

		in = string([]byte{236, 232, 0, 18,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 0})
		exp = "[123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUV]\n"
		checkDynString(t, list, in, exp)

		in = string([]byte{236, 232, 0, 18,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87,
			49, 50, 51, 52, 53, 54, 55, 56, 57, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 0, 0})
		exp = "[123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTUVW123456789ABCDEFGHIJKLMNOPQRSTU]\n"
		checkDynString(t, list, in, exp)
}
*/
