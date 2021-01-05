// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

// Package decoder provides several decoders for differently encoded trice streams.
package decoder

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/rokath/trice/internal/emitter"
	"github.com/rokath/trice/internal/id"
)

const (
	// receive and sync buffer size
	defaultSize  = 4096
	littleEndian = true
	bigEndian    = false
)

var (
	// Verbose gives mor information on output if set. The value is injected from main packages.
	Verbose bool

	// Encoding describes the way the byte stream is coded.
	Encoding string
)

// StringsReader as Decoder method uses an underlying (byte) Reader for reading and returns max len(p) strings inside p.
// n is the received amount of strings. err is the underlying Reader error or an internal error.
type StringsReader interface {
	StringsRead(p []string) (n int, err error)
}

// Decoder is the interface (method set) a decoder needs to provide
type Decoder interface {
	StringsReader
}

// decoding is the common data struct for all decoders
type decoding struct {
	in         io.Reader // inner reader
	syncBuffer []byte    // unprocessed bytes hold for next cycle
	lut        IDLookUp  // id look-up map for translation
	endian     bool      // littleEndian or bigEndian
	trice      idFmt     // received trice
	b          []byte    // read buffer
	bc         int       // trice specific bytes count
}

// idFmt contains the ID mapped information needed for decoding.
type idFmt struct {
	Type string
	Strg string
}

// IDLookUp is the ID to format info translation map
type IDLookUp map[int]idFmt

// UnmarshalTriceIDList extracts the trice ID list byte slice to an items slice.
// til is a result of a read til.json file or is a converted test string.
// til can change during runtime, when an "trice update" occurs.
// Just in case til is not consistent the err value is not nil.
func UnmarshalTriceIDList(til []byte) (list []id.Item, err error) {
	if 0 < len(til) {
		err = json.Unmarshal(til, &list)
	}
	return
}

// MakeLut returns a trice ID lookup map.
func MakeLut(list []id.Item) (lut IDLookUp) {
	// create look-up map
	lut = make(IDLookUp)
	// to do: add timestamp evaluation
	for _, item := range list {
		key := item.ID
		value := idFmt{Type: item.FmtType, Strg: item.FmtStrg}
		lut[key] = value
	}
	return
}

// newIDLut assumes til as JSON formatted input and returns a map for trice ID to fmt string translation.
func newIDLut(til []byte) (IDLookUp, error) {
	list, err := UnmarshalTriceIDList(til)
	if nil != err {
		return nil, err
	}
	lut := MakeLut(list) // create look-up map
	return lut, nil
}

// Translate performs the trice log task.
// Bytes are read with rc. Then according decoder.Encoding are translated into strings.
// Each read returns the amount of bytes for one trice.
// Translate returns true on io.EOF or false on hard read error or sigterm.
func Translate(sw *emitter.TriceLineComposer, list *id.List, rc io.ReadCloser) bool {

	var dec io.Reader
	switch Encoding {
	case "leg":
		dec := NewEscLegacyDecoder(list.ItemList, rc)
		for {
			err := run0(sw, dec)
			if io.EOF == err {
				continue // try again
			}
			if nil != err {
				return false // stop
			}
		}
	case "esc":
		dec = NewEscDecoder(list.ItemList, rc, bigEndian)
	case "pack":
		dec = NewPackDecoder(list.ItemList, rc, bigEndian)
	case "packl", "packL":
		dec = NewPackDecoder(list.ItemList, rc, littleEndian)
	case "bare":
		dec = NewBareDecoder(list.ItemList, rc, bigEndian)
	case "barel", "bareL":
		dec = NewBareDecoder(list.ItemList, rc, littleEndian)
	case "wrap":
		dec = NewBareDecoder(list.ItemList, NewBareReaderFromWrap(rc), bigEndian)
	case "wrapl", "wrapL":
		dec = NewBareDecoder(list.ItemList, NewBareReaderFromWrap(rc), littleEndian)
	//case "bareXTEACrypted", "wrapXTEACrypted":
	//	errorFatal(cipher.SetUp())
	//	fallthrough
	default:
		fmt.Println("unknown encoding ", Encoding)
		return false // stop
	}

	// prepare CTRL-C shutdown reaction
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// intermediate trice string buffer for a single trice
	b := make([]byte, defaultSize)
	for {
		select {
		case sig := <-sigs: // wait for a signal
			if Verbose {
				fmt.Println("####################################", sig, "####################################")
			}
			return false // end
		default:
			n, err := dec.Read(b)
			if io.EOF == err {
				if Verbose {
					fmt.Println(err)
				}
				continue // read again
			}
			if nil != err {
				if Verbose {
					fmt.Println(err)
				}
				return true // try again
			}
			sw.Write(b[:n])
			//time.Sleep(1 * time.Millisecond) // limit speed
		}
	}
}

func run0(sw *emitter.TriceLineComposer, sr StringsReader) error {
	var sssiz int // to do: 1 for pack, 100 for esc
	if Encoding == "pack" {
		sssiz = 1
	} else {
		sssiz = 100
	}
	ss := make([]string, sssiz)
	n, err := sr.StringsRead(ss)
	if nil != err && io.EOF != err {
		return err
	}
	for i := range ss[:n] {
		sw.WriteString(ss[i])
	}
	return nil
}

func shovel(sw io.Writer, b []byte, sr io.Reader) error {
	n, err := sr.Read(b)
	if nil != err && io.EOF != err {
		return err
	}
	sw.Write(b[:n])
	return nil
}

// errorFatal ends in osExit(1) if err not nil.
func errorFatal(err error) {
	if nil == err {
		return
	}
	if Verbose {
		_, file, line, _ := runtime.Caller(1)
		log.Fatal(err, " "+filepath.Base(file)+" ", line)
	}
	log.Fatal(err)
}

// readU16 returns the 2 b bytes as uint16 according the specified endianess
func (p *decoding) readU16(b []byte) uint16 {
	if littleEndian == p.endian {
		return binary.LittleEndian.Uint16(b)
	}
	return binary.BigEndian.Uint16(b)
}

// readU32 returns the 4 b bytes as uint32 according the specified endianess
func (p *decoding) readU32(b []byte) uint32 {
	if littleEndian == p.endian {
		return binary.LittleEndian.Uint32(b)
	}
	return binary.BigEndian.Uint32(b)
}

// readU64 returns the 8 b bytes as uint64 according the specified endianess
func (p *decoding) readU64(b []byte) uint64 {
	if littleEndian == p.endian {
		return binary.LittleEndian.Uint64(b)
	}
	return binary.BigEndian.Uint64(b)
}

// rub removes leading bytes from sync buffer
func (p *decoding) rub(n int) {
	p.syncBuffer = p.syncBuffer[n:]
}

func (p *decoding) outOfSync(msg string) (n int, e error) {
	n = copy(p.b, fmt.Sprintln("error:", msg, "ignoring first byte", p.syncBuffer[0:p.bc]))
	p.rub(1)
	return
}
