package kiss

import (
	"reflect"
	"testing"
)

func Test_encode(t *testing.T) {
	type args struct {
		port    byte
		command byte
		data    []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "setting port and command",
			args: args{
				port:    0x1,
				command: 0x2,
				data:    []byte{0x11},
			},
			want: []byte{0xc0, 0x12, 0x11, 0xc0},
		},
		{
			name: "escaping FEND",
			args: args{
				port:    0x3,
				command: 0x4,
				data:    []byte{FEND},
			},
			want: []byte{FEND, 0x34, FESC, TFEND, FEND},
		}, {
			name: "escaping FESC",
			args: args{
				port:    0x5,
				command: 0x6,
				data:    []byte{FESC},
			},
			want: []byte{FEND, 0x56, FESC, TFESC, FEND},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encode(tt.args.port, tt.args.command, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataFrame(t *testing.T) {
	type args struct {
		port byte
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "send TEST out of port 0",
			args: args{
				port: 0x0,
				data: []byte("TEST"),
			},
			want: []byte{0xc0, 0x00, 0x54, 0x45, 0x53, 0x54, 0xc0},
		},
		{
			name: "send Hello out of port 5",
			args: args{
				port: 0x5,
				data: []byte("Hello"),
			},
			want: []byte{0xc0, 0x50, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0xc0},
		},
		{
			name: "send 0xc0 0xdb out of port 15",
			args: args{
				port: 0xf,
				data: []byte{0xc0, 0xdb},
			},
			want: []byte{0xc0, 0xf0, 0xdb, 0xdc, 0xdb, 0xdd, 0xc0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DataFrame(tt.args.port, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturn(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{
			name: "only case",
			want: []byte{0xc0, 0xff, 0xc0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Return(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Return() = %v, want %v", got, tt.want)
			}
		})
	}
}
