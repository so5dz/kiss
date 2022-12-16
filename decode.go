package kiss

type Decoder struct {
	buffer        []byte
	transposition bool
	onDataFrame   func([]byte)
}

func (d *Decoder) OnDataFrame(callback func([]byte)) {
	d.onDataFrame = callback
}

func (d *Decoder) FeedBytes(bytes []byte) {
	for _, b := range bytes {
		d.Feed(b)
	}
}

func (d *Decoder) Feed(b byte) {
	if b == FEND {
		if len(d.buffer) > 0 && (getCommand(d.buffer[0]) == Command_DataFrame) {
			d.onDataFrame(d.buffer[1:])
		}
		d.buffer = d.buffer[:0]
	} else if b == FESC {
		if d.transposition {
			d.buffer = d.buffer[:0]
			d.transposition = false
		} else {
			d.transposition = true
		}
	} else if d.transposition {
		if b == TFEND {
			d.buffer = append(d.buffer, FEND)
		} else if b == TFESC {
			d.buffer = append(d.buffer, FESC)
		} else {
			d.buffer = d.buffer[:0]
		}
		d.transposition = false
	} else {
		d.buffer = append(d.buffer, b)
	}
}
