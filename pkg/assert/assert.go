// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

// Package assert provides hepler functions for testing.
package assert

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/udhos/equalfile"
)

// test helper ///////////////////////////////////////////////////////////////////////
//

// True fails the test if condition is false.
func True(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(msg, filepath.Base(file), line, " ")
		fmt.Println(v...)
		tb.FailNow()
	}
}

// False fails the test if condition is true.
func False(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Print(msg, filepath.Base(file), line, " ")
		fmt.Println(v...)
		tb.FailNow()
	}
}

// ErrorNil fails the test if an err is not nil.
func ErrorNil(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(err.Error(), filepath.Base(file), line)
		tb.FailNow()
	}
}

// Equal fails the test if exp is not equal to act.
func Equal(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		log.Println("expect:", exp)
		log.Println("actual:", act)
		es := fmt.Sprintf("%v", exp)
		as := fmt.Sprintf("%v", act)
		log.Println("expect:", es)
		log.Println("actual:", as)
		log.Println("expect:", []byte(es))
		log.Println("actual:", []byte(as))
		fmt.Println(filepath.Base(file), line)
		tb.FailNow()
	}
}

// EqualLines compares exp and act line by line ignoring lineendings and leading/trailing spaces.
func EqualLines(tb testing.TB, exp, act string) {

	// remove windows line endings
	exp0 := strings.ReplaceAll(exp, "\r\n", "\n")
	act0 := strings.ReplaceAll(act, "\r\n", "\n")

	exp1 := strings.Split(exp0, "\r")
	act1 := strings.Split(act0, "\r")

	if len(exp1) != len(act1) {
		fmt.Println(len(exp1), len(act1))
		tb.Fail()
	}
	for i, se := range exp1 {
		expLine := strings.TrimSpace(se)
		actLine := strings.TrimSpace(act1[i])
		Equal(tb, expLine, actLine)
	}
}

// EqualTextFiles fails test if lines in fn0 & fn1 NOT equal.
// Line endings are ignored
func EqualTextFiles(t *testing.T, fn0, fn1 string) {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	b0, e0 := ioutil.ReadFile(fn0)
	ErrorNil(t, e0)
	b1, e1 := ioutil.ReadFile(fn1)
	ErrorNil(t, e1)

	// Convert []byte to string and print to screen
	s0 := string(b0)
	s1 := string(b1)

	EqualLines(t, s0, s1)
}

// EqualFiles fails test if contence is NOT equal
func EqualFiles(t *testing.T, fn0, fn1 string) {
	ok := equalFileContent(fn0, fn1)
	if false == ok {
		t.FailNow()
	}
}

// equalFileContent returns true if contece is equal
func equalFileContent(fn0, fn1 string) bool {
	cmp := equalfile.New(nil, equalfile.Options{}) // compare using single mode
	ok, err := cmp.CompareFile(fn0, fn1)
	if nil != err {
		ok = false
	}
	return ok
}
