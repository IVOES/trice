// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
//                basti@blackoutcloud.de
// Use of this source code is governed by a license that can be found in the LICENSE file.

// Package com_test is a blackbox test.
package com_test

import (
	"testing"

	"github.com/rokath/trice/internal/com"
	"github.com/rokath/trice/pkg/tst"
)

func Test1(t *testing.T) {
	ss, err := com.GetSerialPorts()
	tst.AssertNoErr(t, err)
	// To do: handle special cases:
	//
	//	PS C:\repos\trice> trice s
	//	Could not enumerate serial ports
	//	PS C:\repos\trice> trice s
	//	No serial ports found!
	//
	// A normal case:
	//	PS C:\repos\trice> trice s
	//	Found port:  COM4
	for i := range ss {
		port := ss[i]
		p := com.NewCOMPortGoBugSt(port)
		if p.Open() {
			tst.AssertNoErr(t, p.Close())
		}
	}
}
