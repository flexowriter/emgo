// +build l476xx

// Peripheral: I2C_Periph  Inter-integrated Circuit Interface.
// Instances:
//  I2C1  mmap.I2C1_BASE
//  I2C2  mmap.I2C2_BASE
//  I2C3  mmap.I2C3_BASE
// Registers:
//  0x00 32  CR1      Control register 1.
//  0x04 32  CR2      Control register 2.
//  0x08 32  OAR1     Own address 1 register.
//  0x0C 32  OAR2     Own address 2 register.
//  0x10 32  TIMINGR  Timing register.
//  0x14 32  TIMEOUTR Timeout register.
//  0x18 32  ISR      Interrupt and status register.
//  0x1C 32  ICR      Interrupt clear register.
//  0x20 32  PECR     PEC register.
//  0x24 32  RXDR     Receive data register.
//  0x28 32  TXDR     Transmit data register.
// Import:
//  stm32/o/l476xx/mmap
package i2c

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PE        CR1_Bits = 0x01 << 0  //+ Peripheral enable.
	TXIE      CR1_Bits = 0x01 << 1  //+ TX interrupt enable.
	RXIE      CR1_Bits = 0x01 << 2  //+ RX interrupt enable.
	ADDRIE    CR1_Bits = 0x01 << 3  //+ Address match interrupt enable.
	NACKIE    CR1_Bits = 0x01 << 4  //+ NACK received interrupt enable.
	STOPIE    CR1_Bits = 0x01 << 5  //+ STOP detection interrupt enable.
	TCIE      CR1_Bits = 0x01 << 6  //+ Transfer complete interrupt enable.
	ERRIE     CR1_Bits = 0x01 << 7  //+ Errors interrupt enable.
	DNF       CR1_Bits = 0x0F << 8  //+ Digital noise filter.
	ANFOFF    CR1_Bits = 0x01 << 12 //+ Analog noise filter OFF.
	SWRST     CR1_Bits = 0x01 << 13 //+ Software reset.
	TXDMAEN   CR1_Bits = 0x01 << 14 //+ DMA transmission requests enable.
	RXDMAEN   CR1_Bits = 0x01 << 15 //+ DMA reception requests enable.
	SBC       CR1_Bits = 0x01 << 16 //+ Slave byte control.
	NOSTRETCH CR1_Bits = 0x01 << 17 //+ Clock stretching disable.
	WUPEN     CR1_Bits = 0x01 << 18 //+ Wakeup from STOP enable.
	GCEN      CR1_Bits = 0x01 << 19 //+ General call enable.
	SMBHEN    CR1_Bits = 0x01 << 20 //+ SMBus host address enable.
	SMBDEN    CR1_Bits = 0x01 << 21 //+ SMBus device default address enable.
	ALERTEN   CR1_Bits = 0x01 << 22 //+ SMBus alert enable.
	PECEN     CR1_Bits = 0x01 << 23 //+ PEC enable.
)

const (
	PEn        = 0
	TXIEn      = 1
	RXIEn      = 2
	ADDRIEn    = 3
	NACKIEn    = 4
	STOPIEn    = 5
	TCIEn      = 6
	ERRIEn     = 7
	DNFn       = 8
	ANFOFFn    = 12
	SWRSTn     = 13
	TXDMAENn   = 14
	RXDMAENn   = 15
	SBCn       = 16
	NOSTRETCHn = 17
	WUPENn     = 18
	GCENn      = 19
	SMBHENn    = 20
	SMBDENn    = 21
	ALERTENn   = 22
	PECENn     = 23
)

const (
	SADD    CR2_Bits = 0x3FF << 0 //+ Slave address (master mode).
	RD_WRN  CR2_Bits = 0x01 << 10 //+ Transfer direction (master mode).
	ADD10   CR2_Bits = 0x01 << 11 //+ 10-bit addressing mode (master mode).
	HEAD10R CR2_Bits = 0x01 << 12 //+ 10-bit address header only read direction (master mode).
	START   CR2_Bits = 0x01 << 13 //+ START generation.
	STOP    CR2_Bits = 0x01 << 14 //+ STOP generation (master mode).
	NACK    CR2_Bits = 0x01 << 15 //+ NACK generation (slave mode).
	NBYTES  CR2_Bits = 0xFF << 16 //+ Number of bytes.
	RELOAD  CR2_Bits = 0x01 << 24 //+ NBYTES reload mode.
	AUTOEND CR2_Bits = 0x01 << 25 //+ Automatic end mode (master mode).
	PECBYTE CR2_Bits = 0x01 << 26 //+ Packet error checking byte.
)

const (
	SADDn    = 0
	RD_WRNn  = 10
	ADD10n   = 11
	HEAD10Rn = 12
	STARTn   = 13
	STOPn    = 14
	NACKn    = 15
	NBYTESn  = 16
	RELOADn  = 24
	AUTOENDn = 25
	PECBYTEn = 26
)

const (
	OA1     OAR1_Bits = 0x3FF << 0 //+ Interface own address 1.
	OA1MODE OAR1_Bits = 0x01 << 10 //+ Own address 1 10-bit mode.
	OA1EN   OAR1_Bits = 0x01 << 15 //+ Own address 1 enable.
)

const (
	OA1n     = 0
	OA1MODEn = 10
	OA1ENn   = 15
)

const (
	OA2       OAR2_Bits = 0x7F << 1  //+ Interface own address 2.
	OA2MSK    OAR2_Bits = 0x07 << 8  //+ Own address 2 masks.
	OA2NOMASK OAR2_Bits = 0x00 << 8  //  No mask.
	OA2MASK01 OAR2_Bits = 0x01 << 8  //  OA2[1] is masked, Only OA2[7:2] are compared.
	OA2MASK02 OAR2_Bits = 0x02 << 8  //  OA2[2:1] is masked, Only OA2[7:3] are compared.
	OA2MASK03 OAR2_Bits = 0x03 << 8  //  OA2[3:1] is masked, Only OA2[7:4] are compared.
	OA2MASK04 OAR2_Bits = 0x04 << 8  //  OA2[4:1] is masked, Only OA2[7:5] are compared.
	OA2MASK05 OAR2_Bits = 0x05 << 8  //  OA2[5:1] is masked, Only OA2[7:6] are compared.
	OA2MASK06 OAR2_Bits = 0x06 << 8  //  OA2[6:1] is masked, Only OA2[7] are compared.
	OA2MASK07 OAR2_Bits = 0x07 << 8  //  OA2[7:1] is masked, No comparison is done.
	OA2EN     OAR2_Bits = 0x01 << 15 //+ Own address 2 enable.
)

const (
	OA2n    = 1
	OA2MSKn = 8
	OA2ENn  = 15
)

const (
	SCLL   TIMINGR_Bits = 0xFF << 0  //+ SCL low period (master mode).
	SCLH   TIMINGR_Bits = 0xFF << 8  //+ SCL high period (master mode).
	SDADEL TIMINGR_Bits = 0x0F << 16 //+ Data hold time.
	SCLDEL TIMINGR_Bits = 0x0F << 20 //+ Data setup time.
	PRESC  TIMINGR_Bits = 0x0F << 28 //+ Timings prescaler.
)

const (
	SCLLn   = 0
	SCLHn   = 8
	SDADELn = 16
	SCLDELn = 20
	PRESCn  = 28
)

const (
	TIMEOUTA TIMEOUTR_Bits = 0xFFF << 0  //+ Bus timeout A.
	TIDLE    TIMEOUTR_Bits = 0x01 << 12  //+ Idle clock timeout detection.
	TIMOUTEN TIMEOUTR_Bits = 0x01 << 15  //+ Clock timeout enable.
	TIMEOUTB TIMEOUTR_Bits = 0xFFF << 16 //+ Bus timeout B.
	TEXTEN   TIMEOUTR_Bits = 0x01 << 31  //+ Extended clock timeout enable.
)

const (
	TIMEOUTAn = 0
	TIDLEn    = 12
	TIMOUTENn = 15
	TIMEOUTBn = 16
	TEXTENn   = 31
)

const (
	TXE     ISR_Bits = 0x01 << 0  //+ Transmit data register empty.
	TXIS    ISR_Bits = 0x01 << 1  //+ Transmit interrupt status.
	RXNE    ISR_Bits = 0x01 << 2  //+ Receive data register not empty.
	ADDR    ISR_Bits = 0x01 << 3  //+ Address matched (slave mode).
	NACKF   ISR_Bits = 0x01 << 4  //+ NACK received flag.
	STOPF   ISR_Bits = 0x01 << 5  //+ STOP detection flag.
	TC      ISR_Bits = 0x01 << 6  //+ Transfer complete (master mode).
	TCR     ISR_Bits = 0x01 << 7  //+ Transfer complete reload.
	BERR    ISR_Bits = 0x01 << 8  //+ Bus error.
	ARLO    ISR_Bits = 0x01 << 9  //+ Arbitration lost.
	OVR     ISR_Bits = 0x01 << 10 //+ Overrun/Underrun.
	PECERR  ISR_Bits = 0x01 << 11 //+ PEC error in reception.
	TIMEOUT ISR_Bits = 0x01 << 12 //+ Timeout or Tlow detection flag.
	ALERT   ISR_Bits = 0x01 << 13 //+ SMBus alert.
	BUSY    ISR_Bits = 0x01 << 15 //+ Bus busy.
	DIR     ISR_Bits = 0x01 << 16 //+ Transfer direction (slave mode).
	ADDCODE ISR_Bits = 0x7F << 17 //+ Address match code (slave mode).
)

const (
	TXEn     = 0
	TXISn    = 1
	RXNEn    = 2
	ADDRn    = 3
	NACKFn   = 4
	STOPFn   = 5
	TCn      = 6
	TCRn     = 7
	BERRn    = 8
	ARLOn    = 9
	OVRn     = 10
	PECERRn  = 11
	TIMEOUTn = 12
	ALERTn   = 13
	BUSYn    = 15
	DIRn     = 16
	ADDCODEn = 17
)

const (
	ADDRCF   ICR_Bits = 0x01 << 3  //+ Address matched clear flag.
	NACKCF   ICR_Bits = 0x01 << 4  //+ NACK clear flag.
	STOPCF   ICR_Bits = 0x01 << 5  //+ STOP detection clear flag.
	BERRCF   ICR_Bits = 0x01 << 8  //+ Bus error clear flag.
	ARLOCF   ICR_Bits = 0x01 << 9  //+ Arbitration lost clear flag.
	OVRCF    ICR_Bits = 0x01 << 10 //+ Overrun/Underrun clear flag.
	PECCF    ICR_Bits = 0x01 << 11 //+ PAC error clear flag.
	TIMOUTCF ICR_Bits = 0x01 << 12 //+ Timeout clear flag.
	ALERTCF  ICR_Bits = 0x01 << 13 //+ Alert clear flag.
)

const (
	ADDRCFn   = 3
	NACKCFn   = 4
	STOPCFn   = 5
	BERRCFn   = 8
	ARLOCFn   = 9
	OVRCFn    = 10
	PECCFn    = 11
	TIMOUTCFn = 12
	ALERTCFn  = 13
)

const (
	PEC PECR_Bits = 0xFF << 0 //+ PEC register.
)

const (
	PECn = 0
)

const (
	RXDATA RXDR_Bits = 0xFF << 0 //+ 8-bit receive data.
)

const (
	RXDATAn = 0
)

const (
	TXDATA TXDR_Bits = 0xFF << 0 //+ 8-bit transmit data.
)

const (
	TXDATAn = 0
)
