package kiss

import (
	"testing"
)

func TestDecoder(t *testing.T) {
	t.Run("valid 1B data frame", func(t *testing.T) {
		var d Decoder
		var f []byte
		d.OnDataFrame(func(b []byte) { f = b })
		d.FeedBytes([]byte{FEND, 0x00, 0x11, FEND})
		expecting := []byte{0x11}
		if !arraysEqual(expecting, f) {
			t.Errorf("got %v, want %v", f, expecting)
		}
	})
	t.Run("double FENDs", func(t *testing.T) {
		var d Decoder
		var f []byte
		d.OnDataFrame(func(b []byte) { f = b })
		d.FeedBytes([]byte{FEND, FEND, 0x00, 0x23, FEND, FEND})
		expecting := []byte{0x23}
		if !arraysEqual(expecting, f) {
			t.Errorf("got %v, want %v", f, expecting)
		}
	})
	t.Run("only FENDs", func(t *testing.T) {
		var d Decoder
		var f []byte
		d.OnDataFrame(func(b []byte) { f = b })
		d.FeedBytes([]byte{FEND, FEND, FEND, FEND})
		expecting := []byte{}
		if !arraysEqual(expecting, f) {
			t.Errorf("got %v, want %v", f, expecting)
		}
	})
	t.Run("deescape 2B > 1B", func(t *testing.T) {
		var d Decoder
		var f []byte
		d.OnDataFrame(func(b []byte) { f = b })
		d.FeedBytes([]byte{FEND, 0x00, FESC, TFEND, FEND})
		expecting := []byte{FEND}
		if !arraysEqual(expecting, f) {
			t.Errorf("got %v, want %v", f, expecting)
		}
	})
	t.Run("mismatched FESC", func(t *testing.T) {
		var d Decoder
		var f []byte
		d.OnDataFrame(func(b []byte) { f = b })
		d.FeedBytes([]byte{FEND, 0x00, FESC, FEND})
		expecting := []byte{}
		if !arraysEqual(expecting, f) {
			t.Errorf("got %v, want %v", f, expecting)
		}
	})
	t.Run("mismatched FESCs", func(t *testing.T) {
		var d Decoder
		var f []byte
		d.OnDataFrame(func(b []byte) { f = b })
		d.FeedBytes([]byte{FEND, 0x00, FESC, FESC, FESC, 0x00, 0x55, FEND})
		expecting := []byte{}
		if !arraysEqual(expecting, f) {
			t.Errorf("got %v, want %v", f, expecting)
		}
	})
}

func arraysEqual(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aElement := range a {
		if b[i] != aElement {
			return false
		}
	}
	return true
}
