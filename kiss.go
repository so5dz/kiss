package kiss

// Special bytes
const (
	FEND  byte = 0xc0
	FESC  byte = 0xdb
	TFEND byte = 0xdc
	TFESC byte = 0xdd
)

// Commands
const (
	Command_DataFrame   = 0x0
	Command_TXDelay     = 0x1
	Command_Persistence = 0x2
	Command_SlotTime    = 0x3
	Command_TXTail      = 0x4
	Command_FullDuplex  = 0x5
	Command_SetHardware = 0x6
	Command_Return      = 0xf
)

const ReturnCommandPort = 0xf

func getCommand(commandByte byte) byte {
	return commandByte & 0xf
}
