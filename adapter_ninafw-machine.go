//go:build ninafw

package bluetooth

import (
	"machine"
)

func init() {
	NINA_SCK = machine.NINA_SCK
	NINA_SDO = machine.NINA_SDO
	NINA_SDI = machine.NINA_SDI
	NINA_CS = machine.NINA_CS
	NINA_ACK = machine.NINA_ACK
	NINA_GPIO0 = machine.NINA_GPIO0
	NINA_RESETN = machine.NINA_RESETN
	NINA_BAUDRATE = machine.NINA_BAUDRATE
	NINA_RESET_INVERTED = machine.NINA_RESET_INVERTED
	NINA_SOFT_FLOWCONTROL = machine.NINA_SOFT_FLOWCONTROL
}
