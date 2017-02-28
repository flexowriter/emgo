// Peripheral: VREFBUF_Periph  VREFBUF.
// Instances:
//  VREFBUF  mmap.VREFBUF_BASE
// Registers:
//  0x00 32  CSR Control and status register.
//  0x04 32  CCR Calibration and control register.
// Import:
//  stm32/o/l476xx/mmap
package vrefbuf

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	ENVR CSR_Bits = 0x01 << 0 //+ Voltage reference buffer enable.
	HIZ  CSR_Bits = 0x01 << 1 //+ High impedance mode.
	VRS  CSR_Bits = 0x01 << 2 //+ Voltage reference scale.
	VRR  CSR_Bits = 0x01 << 3 //+ Voltage reference buffer ready.
)

const (
	ENVRn = 0
	HIZn  = 1
	VRSn  = 2
	VRRn  = 3
)

const (
	TRIM CCR_Bits = 0x3F << 0 //+ TRIM[5:0] bits (Trimming code).
)

const (
	TRIMn = 0
)
