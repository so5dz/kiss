package kiss

func encode(port byte, command byte, data []byte) []byte {
	newData := make([]byte, 0, 4+2*len(data))
	newData = append(newData, FEND)
	portCommand := ((port & 0xf) << 4) | getCommand(command)
	if portCommand == FEND {
		newData = append(newData, FESC)
		newData = append(newData, TFEND)
	} else if portCommand == FESC {
		newData = append(newData, FESC)
		newData = append(newData, TFESC)
	} else {
		newData = append(newData, portCommand)
	}
	for _, b := range data {
		if b == FEND {
			newData = append(newData, FESC)
			newData = append(newData, TFEND)
		} else if b == FESC {
			newData = append(newData, FESC)
			newData = append(newData, TFESC)
		} else {
			newData = append(newData, b)
		}
	}
	newData = append(newData, FEND)
	return newData
}

func DataFrame(port byte, data []byte) []byte {
	return encode(port, Command_DataFrame, data)
}

func SetTXDelay(port byte, units byte) []byte {
	return encode(port, Command_TXDelay, []byte{units})
}

func SetPersistence(port byte, units byte) []byte {
	return encode(port, Command_Persistence, []byte{units})
}

func SetSlotTime(port byte, units byte) []byte {
	return encode(port, Command_SlotTime, []byte{units})
}

func SetTXTail(port byte, units byte) []byte {
	return encode(port, Command_TXTail, []byte{units})
}

func SetFullDuplex(port byte, value bool) []byte {
	var tValue byte = 0
	if value {
		tValue = 1
	}
	return encode(port, Command_FullDuplex, []byte{tValue})
}

func SetHardware(port byte, data []byte) []byte {
	return encode(port, Command_SetHardware, data)
}

func Return() []byte {
	return encode(ReturnCommandPort, Command_Return, []byte{})
}
