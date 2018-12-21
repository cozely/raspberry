package gpio

import (
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

//------------------------------------------------------------------------------

func init() {
	f, err := unix.Open("/dev/gpiomem", unix.O_RDWR|unix.O_SYNC, 0)
	if err != nil {
		panic(err)
	}
	defer unix.Close(f)
	m, err := unix.Mmap(
		f,
		0,
		int(unsafe.Sizeof(*Registers)),
		unix.PROT_READ|unix.PROT_WRITE,
		unix.MAP_SHARED,
	)
	Registers = (*[41]uint32)(unsafe.Pointer(&m[0]))
}

//------------------------------------------------------------------------------

// A Pin is specified using Broadcom numbering (aka BCM, GPIO)
type Pin uint8

// GPIO pins, using physical numbering
const (
	/* Physical1 is 3V3 Power */
	/* Physical2 is 5V Power */
	Physical3 = Pin(2)
	/* Physical4 is 5V Power */
	Physical5 = Pin(3)
	/* Physical6 is Ground */
	Physical7 = Pin(4)
	Physical8 = Pin(14)
	/* Physical9 is Ground */
	Physical10 = Pin(15)
	Physical11 = Pin(17)
	Physical12 = Pin(18)
	Physical13 = Pin(27)
	/* Physical14 is Ground */
	Physical15 = Pin(22)
	Physical16 = Pin(23)
)

// GPIO pins, using logical names of their alternate functions
const (
	// I2C
	I2C_SDA = Pin(2)
	I2C_SCL = Pin(3)
	// EEPROM
	ID_SD = Pin(0)
	ID_SC = Pin(1)
	//
	GPCLK0 = Pin(4)
	// Serial (UART)
	TXD = Pin(14)
	RXD = Pin(15)
	// SPI0
	SPI0_MOSI = Pin(10)
	SPI0_MISO = Pin(9)
	SPI0_SCLK = Pin(11)
	SPI0_CE0  = Pin(8)
	SPI0_CE1  = Pin(7)
	// SPI1
	SPI1_MOSI = Pin(20)
	SPI1_MISO = Pin(19)
	SPI1_SCLK = Pin(21)
	SPI1_CE0  = Pin(18)
	SPI1_CE1  = Pin(17)
	SPI1_CE2  = Pin(16)
	// PWM
	PWM0    = Pin(12)
	PWM1    = Pin(13)
	PWM0Alt = Pin(18)
	PWM1Alt = Pin(19)
)

//------------------------------------------------------------------------------

// Registers gives direct access to all the GPIO registers of the BCM2835 chip.
var Registers *[41]uint32

// GPIO registers used on the Pi, as indices for the Registers array.
const (
	GPSEL0    = 0x00 / 4
	GPSEL1    = 0x04 / 4
	GPSEL2    = 0x08 / 4
	GPSET0    = 0x1c / 4
	GPCLR0    = 0x28 / 4
	GPLEV0    = 0x34 / 4
	GPEDS0    = 0x40 / 4
	GPREN0    = 0x4C / 4
	GPFEN0    = 0x58 / 4
	GPHEN0    = 0x64 / 4
	GPLEN0    = 0x70 / 4
	GPAREN0   = 0x7C / 4
	GPAFEN0   = 0x88 / 4
	GPPUD     = 0x94 / 4
	GPPUDCLK0 = 0x98 / 4
)

//------------------------------------------------------------------------------

// Mode returns the current mode of a pin.
func Mode(p Pin) string {
	switch (Registers[p/10] >> ((p % 10) * 3)) & 0x7 {
	case 0:
		return "input"
	case 1:
		return "output"
	case 4:
		return "alt0"
	case 5:
		return "alt1"
	case 6:
		return "alt2"
	case 7:
		return "alt3"
	case 3:
		return "alt4"
	case 2:
		return "alt5"
	default:
		return "invalid"
	}
}

// Input configures a pin to input mode.
func Input(p Pin) {
	Registers[p/10] &= ^(0x7 << ((p % 10) * 3))
}

// Output configures a pin to output mode.
func Output(p Pin) {
	Registers[p/10] &= ^(0x7 << ((p % 10) * 3))
	Registers[p/10] |= 1 << ((p % 10) * 3)
}

// Function0 sets the mode of a pin.
func Function0(p Pin) {
	Registers[p/10] &= ^(0x7 << ((p % 10) * 3))
	Registers[p/10] |= 4 << ((p % 10) * 3)
}

// Function1 sets the mode of a pin.
func Function1(p Pin) {
	Registers[p/10] &= ^(0x7 << ((p % 10) * 3))
	Registers[p/10] |= 5 << ((p % 10) * 3)
}

// Function2 sets the mode of a pin.
func Function2(p Pin) {
	Registers[p/10] &= ^(0x7 << ((p % 10) * 3))
	Registers[p/10] |= 6 << ((p % 10) * 3)
}

// Function3 sets the mode of a pin.
func Function3(p Pin) {
	Registers[p/10] &= ^(0x7 << ((p % 10) * 3))
	Registers[p/10] |= 7 << ((p % 10) * 3)
}

// Function4 sets the mode of a pin.
func Function4(p Pin) {
	Registers[p/10] &= ^(0x7 << ((p % 10) * 3))
	Registers[p/10] |= 3 << ((p % 10) * 3)
}

// Function5 sets the mode of a pin.
func Function5(p Pin) {
	Registers[p/10] &= ^(0x7 << ((p % 10) * 3))
	Registers[p/10] |= 2 << ((p % 10) * 3)
}

//------------------------------------------------------------------------------

// PullOff disables any internal pull-up or pull-down for a pin.
func PullOff(p Pin) {
	Registers[GPPUD] = 0
	<-time.After(150 * time.Microsecond)
	Registers[GPPUDCLK0] = uint32(1 << p)
	<-time.After(150 * time.Microsecond)
	Registers[GPPUD] = 0
	Registers[GPPUDCLK0] = 0
}

// PullDown activates the internal pull-down for a ppin.
func PullDown(p Pin) {
	Registers[GPPUD] = 1
	<-time.After(150 * time.Microsecond)
	Registers[GPPUDCLK0] = uint32(1 << p)
	<-time.After(150 * time.Microsecond)
	Registers[GPPUD] = 0
	Registers[GPPUDCLK0] = 0
}

// PullUp activates the internal pull-up for a pin.
func PullUp(p Pin) {
	Registers[GPPUD] = 2
	<-time.After(150 * time.Microsecond)
	Registers[GPPUDCLK0] = uint32(1 << p)
	<-time.After(150 * time.Microsecond)
	Registers[GPPUD] = 0
	Registers[GPPUDCLK0] = 0
}

//------------------------------------------------------------------------------

// Set the pin, i.e. forces its value to 1.
func Set(p Pin) {
	Registers[GPSET0] = 1 << p
}

// Clear the pin, i.e. forces its value it to 0.
func Clear(p Pin) {
	Registers[GPCLR0] = 1 << p
}

// Write changes the pin to a specified value.
func Write(p Pin, b bool) {
	if b {
		Registers[GPSET0] = 1 << p
	} else {
		Registers[GPSET0] = 1 << p
	}
}

// Get returns the pin value.
func Get(p Pin) bool {
	return Registers[GPLEV0]&(1<<p) != 0
}

// Toggle inverts the pin value.
func Toggle(p Pin) {
	if Registers[GPLEV0]&(1<<p) != 0 {
		Registers[GPSET0] = 1 << p
	} else {
		Registers[GPCLR0] = 1 << p
	}
}

//------------------------------------------------------------------------------

func Event(p Pin) bool {
	result := Registers[GPEDS0]&(1<<p) != 0
	// Registers[GPEDS0] |= 1 << p
	return result
}

func PeekEvent(p Pin) bool {
	return Registers[GPEDS0]&(1<<p) != 0
}

func Detect(high, low, rising, falling, arising, afalling bool, p Pin) {
	if high {
		Registers[GPHEN0] |= 1 << p
	} else {
		Registers[GPHEN0] &= ^(1 << p)
	}
	if low {
		Registers[GPLEN0] |= 1 << p
	} else {
		Registers[GPLEN0] &= ^(1 << p)
	}
	if rising {
		Registers[GPREN0] |= 1 << p
	} else {
		Registers[GPREN0] &= ^(1 << p)
	}
	if falling {
		Registers[GPFEN0] |= 1 << p
	} else {
		Registers[GPFEN0] &= ^(1 << p)
	}
	if arising {
		Registers[GPAREN0] |= 1 << p
	} else {
		Registers[GPAREN0] &= ^(1 << p)
	}
	if afalling {
		Registers[GPAFEN0] |= 1 << p
	} else {
		Registers[GPAFEN0] &= ^(1 << p)
	}
}
