// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os"
	"sync"
	"testing"

	"github.com/rokath/trice/internalStim/stimargs"
	"github.com/rokath/trice/pkg/tst"
)

var (
	m *sync.RWMutex
)

func init() {
	m = new(sync.RWMutex)
}

func TestWrongSwitch(t *testing.T) {
	os.Args = []string{"stim", "wrong"}
	expect := `unknown sub-command 'wrong'. try: 'stim help|h'
	`
	execHelper1(t, expect)
}

func TestVersion(t *testing.T) {
	version = "1.2.3"
	commit = "myCommit"
	date = "2006-01-02_1504-05"
	os.Args = []string{"stim", "version"}
	expect := `version=1.2.3, commit=myCommit, built at 2006-01-02_1504-05
	`
	execHelper1(t, expect)
	version = ""
	commit = ""
	date = ""
}

func execHelper1(t *testing.T, expect string) {
	m.Lock()
	defer m.Unlock()
	stimargs.FlagsInit() // maybe needed for clearance of previous tests (global vars)
	var out bytes.Buffer
	doit(&out)
	act := out.String()
	tst.EqualLines(t, expect, act)
}
