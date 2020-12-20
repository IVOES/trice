# Trice encodings

There are different possibilities for encoding trices. The trice encoding inside the triceFifo can differ from the trice transmit format.

## bareL internal storage format

- This is the fastest storage option on little endian mashines.

```b
      0     1     2     3 | macro
--------------------------|--------
     IL    IH     0     0 | TRICE0( Id(I), "..." );
     IL    IH    b0     0 | TRICE8_1( Id(I), "...", b0 );
     IL    IH    b0    b1 | TRICE8_2( Id(I), "...", b1 );
     IL    IH   w0L   w0H | TRICE16_1( Id(I), "...", w0 );
```

```b
     0      1     2     3     4     5     6     7 | macro
--------------------------------------------------|-----------------------------------------------
     0      0    b0    b1    IL    IH    b2     0 | TRICE8_3( Id(I), "...", b0, b1, b2 );
     0      0    b0    b1    IL    IH    b2    b3 | TRICE8_4( Id(I), "...", b0, b1, b2, b3 );
     0      0   w0L   w0H    IL    IH   w1L   w1H | TRICE16_2( Id(I), "...", w0, w1 );
     0      0  d0LL  d0LH    IL    IH  d0HL  d0HH | TRICE32_1( Id(I), "...", d0 );
```

```b
     0      1     2     3     4     5     6     7     8     9     10     11 | macro
----------------------------------------------------------------------------|------------
     0      0    b0    b1     0     0    b2    b3    IL    IH     b4      0 | TRICE8_5( Id(I), "...", b0, b1, b2, b3, b4 );
     0      0    b0    b1     0     0    b2    b3    IL    IH     b4     b5 | TRICE8_6( Id(I), "...", b0, b1, b2, b3, b4, b5 );
     0      0   w0L   w0H     0     0   w1L   w1H    IL    IH    w2L    w2H | TRICE16_3( Id(I), "...", w0, w1, w2);
```

```b
     0      1     2     3     4     5     6     7     8     9     10     11     12     13     14     15 | macro
--------------------------------------------------------------------------------------------------------|------------
     0      0    b0    b1     0     0    b2    b3     0     0     b4     b5     IL     IH     b6      0 | TRICE8_7( Id(I), "...", b0, b1, b2, b3, b4, b5, b6 );
     0      0    b0    b1     0     0    b2    b3     0     0     b4     b5     IL     IH     b6     b7 | TRICE8_8( Id(I), "...", b0, b1, b2, b3, b4, b5, b6, b7 );
     0      0   w0L   w0H     0     0   w1L   w1H     0     0    w2L    w2H     IL     IH    w3L    w3H | TRICE16_3( Id(I), "...", w0, w1, w2, w3);
     0      0  d0LL  d0LH     0     0  d0HL  d0HH     0     0   d1LL   d1LH     IL     IH   d1HL   d1HH | TRICE32_2( Id(I), "...", d0, d1 );
     0      0 l0LLL l0LLH     0     0 l0LHL l0LHH     0     0  l0HLL  l0HLH     IL     IH  l0HHL  l0HHH | TRICE64_1( Id(I), "...", l0 );
```

and so on...

## bareL transmit format

- This is exactly the same as the bareL internal storage format with one and only one difference: There are sometimes 4 byte [sync packages](#sync-packages) mixed in on 4 byte offsets.
- Transmitting `bareL` encoded trice messages reduces the code and increases the speed on little-endian mashines.


## bare internal storage format

- This is the fastest storage option on big endian mashines.

```b
      0     1     2     3 | macro
--------------------------|--------
     IH    IL     0     0 | TRICE0( Id(I), "..." );
     IH    IL    b0     0 | TRICE8_1( Id(I), "...", b0 );
     IH    IL    b0    b1 | TRICE8_2( Id(I), "...", b1 );
     IH    IL   w0H   w0L | TRICE16_1( Id(I), "...", w0 );
```

```b
     0      1     2     3     4     5     6     7 | macro
--------------------------------------------------|-----------------------------------------------
     0      0    b0    b1    IH    IL    b2     0 | TRICE8_3( Id(I), "...", b0, b1, b2 );
     0      0    b0    b1    IH    IL    b2    b3 | TRICE8_4( Id(I), "...", b0, b1, b2, b3 );
     0      0   w0H   w0L    IH    IL   w1H   w1L | TRICE16_2( Id(I), "...", w0, w1 );
     0      0  d0HH  d0HL    IH    IL  d0LH  d0LL | TRICE32_1( Id(I), "...", d0 );
```

```b
     0      1     2     3     4     5     6     7     8     9     10     11 | macro
----------------------------------------------------------------------------|------------
     0      0    b0    b1     0     0    b2    b3    IH    IL     b4      0 | TRICE8_5( Id(I), "...", b0, b1, b2, b3, b4 );
     0      0    b0    b1     0     0    b2    b3    IH    IL     b4     b5 | TRICE8_6( Id(I), "...", b0, b1, b2, b3, b4, b5 );
     0      0   w0H   w0H     0     0   w1H   w1L    IH    IL    w2H    w2L | TRICE16_3( Id(I), "...", w0, w1, w2);
```

```b
     0      1     2     3     4     5     6     7     8     9     10     11     12     13     14     15 | macro
--------------------------------------------------------------------------------------------------------|------------
     0      0    b0    b1     0     0    b2    b3     0     0     b4     b5     IH     IL     b6      0 | TRICE8_7( Id(I), "...", b0, b1, b2, b3, b4, b5, b6 );
     0      0    b0    b1     0     0    b2    b3     0     0     b4     b5     IH     IL     b6     b7 | TRICE8_8( Id(I), "...", b0, b1, b2, b3, b4, b5, b6, b7 );
     0      0   w0H   w0L     0     0   w1H   w1L     0     0    w2H    w2L     IH     IL    w3H    w3L | TRICE16_3( Id(I), "...", w0, w1, w2, w3);
     0      0  d0HH  d0HL     0     0  d0LH  d0LL     0     0   d1HH   d1HL     IH     IL   d1LH   d1LL | TRICE32_2( Id(I), "...", d0, d1 );
     0      0 l0HHH l0HHL     0     0 l0HLH l0HLL     0     0  l0LHH  l0LHL     IH     IL  l0LLH  l0LLL | TRICE64_1( Id(I), "...", l0 );
```

and so on...

## bare transmit format

- This is exactly the same as the bare internal storage format with one and only one difference: There are sometimes 4 byte [sync packages](#sync-packages) mixed in on 4 byte offsets.
- Transmitting `bare` encoded trice messages reduces the code and increases the speed on little-endian mashines.







internal fifo buffer storage format:

- bare code format
  
  During runtime normally only the 16-bit ID 12345 (together with the parameters like hour, min, sec) is copied to a buffer. Execution time for a TRICE16_1 (as example) on a 48 MHz ARM can be about 16 systicks resulting in 250 nanoseconds duration, so you can use `trice` also inside interrupts or the RTOS scheduler to analyze task timings. The needed buffer space is one 32 bit word per normal trice (for up to 2 data bytes). A direct out transfer is possible but not recommended for serial output because of possible issues to re-sync in case of data loss. Just in case the internal bare fifo overflows, the data are still in sync.

  If the wrap format is desired as output the buffered 4 byte trice is transmitted as an 8 byte packet allowing start byte, sender and receiver addresses and CRC8 check to be used later in parallel with different software protocols.

  The bare output format contains exactly the bare bytes but is enriched with sync packages to achieve syncing. The sync package interval is adjustable.

- (direct) esc(ape) code format

  During untime the esc code format is generated immediately during the TRICE macro execution. This results in a slightly longer TRICE macro execution but allows the direct background transfer to the output device (UART or RTT memory) because re-sync is easy. One advantage of the esc format compared to bare is the more efficient coding of dynamic strings if you have lots of them.


Targest can send trices in different encodings

## Encoding `esc`

The `esc` encoding uses an escape character for syncing after some data loss. It is extendable and recommended if plenty of runtime strings need to be transmitted because these are more compact in `esc` encoding compared to the bare format.

- All numbers are transmitted in network order (big endian).
- All values are in left-right order - first value comes first.

An `esc` trice transfer packet consists of an 4-byte header followed by an optional payload.

|     Start Byte     |    Second Byte    |  Third Byte  | Fourth Byte
|--------------------|-------------------|--------------|--------------
|  Escape char `EC`  | Length Code `LC`  | triceID `IH` | triceID `IL`

### Start byte `EC`

```c
#define TRICE_ESC  0xEC //!< Escape char is control char to start a package.
```

The ESCaped encoding allowes a syncing to the next trice message on the behalf of the escape character `0xec`. It is always followed by a length code in the range `0xdf ... 0xe8`. All other bytes are reserved for future usage despite the `0xde` byte. 

```c
#define TRICE_DEL  0xDE //!< Delete char, if follower of TRICE_ESC, deletes the meaning os TRICE_ESC making it an ordinary TRICE_ESC char.
```

This is inserted as not counted value into the bytestream after an `0xec` to signal that this is an ordinary `0xec` byte inside the data stream. As byte `0xec` is not used so often is is defined as ESC character:

### Length Code `LC`

The LC is a 1-byte logarithmic length code. This is a copy from [trice.h lines 44-58](https://github.com/rokath/trice/blob/master/srcTrice.C/trice.h) and shows the length code meaning:

```c
#define TRICE_P0   0xDF //!< No param char = If follower of TRICE_ESC only a 16 bit ID is inside the payload.
#define TRICE_P1   0xE0 //!< 1 byte param char = If follower of TRICE_ESC a 16 bit ID and 1 byte are inside the payload.
#define TRICE_P2   0xE1 //!< 2 byte param char = If follower of TRICE_ESC a 16 bit ID and 2 byte are inside the payload.
#define TRICE_P4   0xE2 //!< 4 byte param char = If follower of TRICE_ESC a 16 bit ID and 4 byte are inside the payload.
#define TRICE_P8   0xE3 //!< 8 byte param char = If follower of TRICE_ESC a 16 bit ID and 8 byte are inside the payload.
#define TRICE_P16  0xE4 //!< 16 byte param char = If follower of TRICE_ESC a 16 bit ID and 8 byte are inside the payload.
//                 0xE5 // dynamically used for runtime strings with size 17-32
//                 0xE6 // dynamically used for runtime strings with size 33-64
//                 0xE7 // dynamically used for runtime strings with size 63-128
//                 0xE8 // dynamically used for runtime strings with size 127-256
//                 0xE9 // reserved
//                 0xEA // reserved
//                 0xEB // reserved

```

### TriceID `IH` and `IL`

- The third and fourth byte are the 16 bit trice ID: IH & IL.
- The trice ID encodes one of the allowed trice macros and a format string.
- The format string has some format specifiers accordingly to the trice macro.
- In the case of `TRICE_S` the format string contains one and only one `%s`.

### Payload

A number of bytes according LC is optionally following the header. If within the data to be transmitted an 0xEC occurs it stays on its place and is followed by a not counted 0xDE byte to signal that this is no start byte.

- Generic description

| Code                   | Meaning                    | max padding | Remark
|------------------------|----------------------------|-------------|---------------------------------------------------------------------------------------------------------------
| EC LC IH IL ...        | payload = 2^(LC-E0) bytes  |             | LC is a length code
| EC 00 ...              | reserved                   |             | All packages starting with EC 00 until starting with EC DD are reserved.
| EC .. ...              | reserved                   |             | All packages starting with EC 00 until starting with EC DD are reserved.
| EC DD ...              | reserved                   |             | All packages starting with EC 00 until starting with EC DD are reserved.
| EC DE = EC             | real EC character          |             | If inside the payload occures EC an uncounted DE is injected afterwards.
| EC DF IH IL            | 16 bit ID no payload       |        0    | TRICE0, special case: 2^-1 = 0 byte payload
| EC E0 IH IL B0         | 16 bit ID   1 byte payload |        0    | TRICE8_1, TRICE_S(""): 2^0 = 1 byte payload
| EC E1 IH IL B0 B1      | 16 bit ID   2 byte payload |        0    | TRICE8_2, TRICE16_1, TRICE_S("0")
| EC E2 IH IL B0 .. B3   | 16 bit ID   4 byte payload |        1    | TRICE8_3, TRICE8_4, TRICE16_2, TRICE32_1, TRICE_S("01"), TRICE_S("012")
| EC E3 IH IL B0 .. B7   | 16 bit ID   8 byte payload |        3    | TRICE8_5,... TRICE8_8, TRICE16_3, TRICE16_4, TRICE32_2, TRICE64_1, TRICE_S("0123"), ..., TRICE_S("01234567")
| EC E4 IH IL B0 .. B15  | 16 bit ID  16 byte payload |        7    | TRICE32_3, TRICE32_4, TRICE64_2, TRICE_S("012345678"), ..., TRICE_S("0123456789abcde")
| EC E5 IH IL B0 .. B31  | 16 bit ID  32 byte payload |       15    | TRICE_S("0123456789abcdef"), ... TRICE_S(strlen(31))
| EC E6 IH IL B0 .. B63  | 16 bit ID  64 byte payload |       31    | TRICE_S(strlen(32)), ...,TRICE_S(strlen(63))
| EC E7 IH IL B0 .. B127 | 16 bit ID 128 byte payload |       63    | TRICE_S(strlen(64)), ...,TRICE_S(strlen(127))
| EC E8 IH IL B0 .. B255 | 16 bit ID 256 byte payload |      127    | TRICE_S(strlen(128)), ...,TRICE_S(strlen(255))
| EC E9 ...              | reserved                   |             | All packages starting with EC E9 until starting with EC FF are reserved.
| EC .. ...              | reserved                   |             | All packages starting with EC E9 until starting with EC FF are reserved.
| EC FF ...              | reserved                   |             | All packages starting with EC E9 until starting with EC FF are reserved.

- Examples
See function `TestEsc` and `TestEscDynStrings` in file [decoder_test.go](https://github.com/rokath/trice/blob/master/internal/decoder/decoder_test.go).

## Encoding `bare`


## Sync packages
- The frequency is adjustable and could be every 100ms or 40 bytes.
- The PC `trice` tool removes them silently.
- A sync package is `0x89 0xab 0xcd 0ef`. 
- The PC `trice` tool does not use 


```c
//! TRICE_SYNC is an optional trice sync message for syncing, when bare transmission is used.
//! The value 35243 (0x89ab) is a reserved pattern used as ID with value DA 0xcdef.
//! The hex notation protects against unwanted automatic changes.
//! The byte sequence of the sync message is 0x89 0xab 0xcd 0xef.
//! It cannot occure in the trice stream in another way due to ID generaion policy.
//! Sync package is IDDA=89abcdef
//!
//! To avoid wrong syncing these ID's are excluded: xx89, abcd, cdef, efxx (514 pieces)
//!
//! Possible:    IH IL DH DL IH IL DH DL IH IL DH DL (1 right)
//!              xx xx xx xx xx 89 ab cd ef xx xx xx -> avoid IL=89, IH=ef
//!
//! Possible:    IH IL DH DL IH IL DH DL IH IL DH DL (2 right)
//!              xx xx xx xx xx xx 89 ab cd ef xx xx -> avoid ID=cdef
//!
//! Possible:    IH IL DH DL IH IL DH DL IH IL DH DL (3 right)
//!              xx xx xx xx xx xx xx 89 ab cd ef xx -> avoid ID=abcd
//!
//! Sync packet: IH IL DH DL IH IL DH DL IH IL DH DL
//!              xx xx xx xx 89 ab cd ef xx xx xx xx -> use ID=89ab with DA=cdef as sync packet
//!
//!  Possible:   IH IL DH DL IH IL DH DL IH IL DH DL (1 left)
//!              xx xx xx 89 ab cd ef xx xx xx xx xx -> avoid ID=abcd
//!
//!  Possible:   IH IL DH DL IH IL DH DL IH IL DH DL (2 left)
//!              xx xx 89 ab cd ef xx xx xx xx xx xx -> avoid ID=cdef
//!
//!  Possible:   IH IL DH DL IH IL DH DL IH IL DH DL (3 left)
//!              xx 89 ab cd ef xx xx xx xx xx xx xx ->  avoid IL=nn89, IH=ef
//!
//! If an ID=89ab with DA!=cdef is detected -> out of sync!
//! If an IH=ef is detected -> out of sync, all 256 IDs starting with 0xef are excluded
//! If an IL=89 is detected -> out of sync, all 256 IDs ending with 0x89 are excluded
//! If an ID=abcd is detected -> out of sync, ID 0xabcd is excluded
//! If an ID=cdef is detected -> out of sync, ID 0xcdef is excluded
//! ID 0x89ab is reserved for this trice sync package.
//! The trice sync message payload must be 0xcdef.
//! You must not change any of the above demands. Otherwise the syncing will not work.
//! The Id(0x89ab) is here as hex value, so it is ignored by ID management.
//! The trice sync string makes the trice sync info invisible just in case,
//! but the trice tool will filter them out anyway. The trice tool automatic id generation
//! follows these rules.
//#define TRICE_SYNC do{ TRICE16_1( Id(0x89ab), "%x\b\b\b\b", 0xcdef ); }while(0)
```

- Byte order is network order (bigendian)
- Each trice is coded in one to eight 4-byte trice atoms.
- A trice atom consists of a 2 byte id and 2 bytes data.
- When a trice consists of several trice atoms only the last one carries the trice id. The others have a trice id 0.

## Encoding `wrap`

- This is the same as bare, but each trice atom is prefixed with a 4 byte wrap information:
  - 0xEB = start byte
  - 0x60 = source address
  - 0x60 = destination address
  - crc8 = 8 bit checksum over start byte, source and destination address and the 4 bare bytes.


## `esc` encoding (to do: improve this)