package cgot

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/rokath/trice/internal/args"
	"github.com/spf13/afero"
	"github.com/tj/assert"
)

func TestLogs(t *testing.T) {

	// triceLog is the log function for executing the trice logging on binary log data in buffer as space separated numbers.
	// It uses the inside fSys specified til.json and returns the log output.
	triceLog := func(t *testing.T, fSys *afero.Afero, buffer string) string {
		var o bytes.Buffer
		fmt.Println("-i", triceDir+"/test/testdata/til.json")
		assert.Nil(t, args.Handler(io.Writer(&o), fSys, []string{"trice", "log", "-i", triceDir + "/test/testdata/til.json", "-p", "BUFFER", "-args", buffer, "-ts", "off", "-prefix", "off", "-tsf0", "time:        ", "-tsf16", "time:%8x", "-tsf32", "time:%8x", "-li", "off", "-color", "off"}))
		return o.String()
	}

	triceLogTest(t, triceLog)
}