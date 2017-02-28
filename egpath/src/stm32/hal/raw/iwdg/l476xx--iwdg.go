// +build l476xx

// Peripheral: IWDG_Periph  Independent WATCHDOG.
// Instances:
//  IWDG  mmap.IWDG_BASE
// Registers:
//  0x00 32  KR   Key register.
//  0x04 32  PR   Prescaler register.
//  0x08 32  RLR  Reload register.
//  0x0C 32  SR   Status register.
//  0x10 32  WINR Window register.
// Import:
//  stm32/o/l476xx/mmap
package iwdg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	KEY KR_Bits = 0xFFFF << 0 //+ Key value (write only, read 0000h).
)

const (
	KEYn = 0
)

const (
	RL RLR_Bits = 0xFFF << 0 //+ Watchdog counter reload value.
)

const (
	RLn = 0
)

const (
	PVU SR_Bits = 0x01 << 0 //+ Watchdog prescaler value update.
	RVU SR_Bits = 0x01 << 1 //+ Watchdog counter reload value update.
	WVU SR_Bits = 0x01 << 2 //+ Watchdog counter window value update.
)

const (
	PVUn = 0
	RVUn = 1
	WVUn = 2
)

const (
	WIN WINR_Bits = 0xFFF << 0 //+ Watchdog counter window value.
)

const (
	WINn = 0
)
