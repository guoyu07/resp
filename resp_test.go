package resp

import (
	"bytes"
	"testing"
)

type testResult struct {
	T byte
	Value interface{}
}

var validBody map[string]testResult
var validCommand map[string]string

func TestValidData(t *testing.T) {
	for body, result := range validBody {
		buf := bytes.NewReader([]byte(body))
		ret, err := ReadData(buf)
		if nil != err {
			t.Error(err)
		}

		switch ret.T {
			case T_SimpleString, T_BulkString:
				if 0 != bytes.Compare(result.Value.([]byte), ret.String) {
					t.Error("not eq")
				}
			case T_Integer:
				if result.Value.(int64) != ret.Integer {
					t.Error("not eq")
				}
		}
	}
}

func TestValidCommand(t *testing.T) {
	for input, cmd := range validCommand {
		reader := bytes.NewReader([]byte(input))
		c, err := ReadCommand(reader)
		if nil != err {
			t.Error("read command error", err)
		} else if c.Name() != cmd {
			t.Error("read command error", c.Name(), cmd)
		}
	}
}

func _validCommand(tb testing.TB) {
	for input, cmd := range validCommand {
		reader := bytes.NewReader([]byte(input))
		c, err := ReadCommand(reader)
		if nil != err {
			tb.Error("read command error", err)
		} else if c.Name() != cmd {
			tb.Error("read command error", c.Name(), cmd)
		}
	}

}

func BenchmarkValidCommand(b *testing.B) {
	for i:=0; i<b.N; i++ {
		_validCommand(b)
	}
}

func init() {
	validBody = map[string]testResult {
		"+OK\r\n" : {T_SimpleString, []byte("OK")},
		"-Errors\r\n" : {T_Error, []byte("Errors")},
		":100\r\n" : {T_Integer, int64(100)},
		"$7\r\nwalu.cc\r\n" : {T_BulkString, []byte("walu.cc")},
	}

	validCommand = map[string]string{
		"PING" : "PING",
		"PING\n" : "PING",
		"PING\r" : "PING",
		"  PING ": "PING",
		"*2\r\n$4\r\nLLEN\r\n$6\r\nmysist\r\n" : "LLEN",
	}
}
