// +build l476xx

// Peripheral: SYSCFG_Periph  System configuration controller.
// Instances:
//  SYSCFG  mmap.SYSCFG_BASE
// Registers:
//  0x00 32  MEMRMP    Memory remap register.
//  0x04 32  CFGR1     Configuration register 1.
//  0x08 32  EXTICR[4] External interrupt configuration registers.
//  0x18 32  SCSR      SRAM2 control and status register.
//  0x1C 32  CFGR2     Configuration register 2.
//  0x20 32  SWPR      SRAM2 write protection register.
//  0x24 32  SKR       SRAM2 key register.
// Import:
//  stm32/o/l476xx/mmap
package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	MEM_MODE MEMRMP_Bits = 0x07 << 0 //+ SYSCFG_Memory Remap Config.
	FB_MODE  MEMRMP_Bits = 0x01 << 8 //+ Flash Bank mode selection.
)

const (
	MEM_MODEn = 0
	FB_MODEn  = 8
)

const (
	FWDIS       CFGR1_Bits = 0x01 << 0  //+ FIREWALL access enable.
	BOOSTEN     CFGR1_Bits = 0x01 << 8  //+ I/O analog switch voltage booster enable.
	I2C_PB6_FMP CFGR1_Bits = 0x01 << 16 //+ I2C PB6 Fast mode plus.
	I2C_PB7_FMP CFGR1_Bits = 0x01 << 17 //+ I2C PB7 Fast mode plus.
	I2C_PB8_FMP CFGR1_Bits = 0x01 << 18 //+ I2C PB8 Fast mode plus.
	I2C_PB9_FMP CFGR1_Bits = 0x01 << 19 //+ I2C PB9 Fast mode plus.
	I2C1_FMP    CFGR1_Bits = 0x01 << 20 //+ I2C1 Fast mode plus.
	I2C2_FMP    CFGR1_Bits = 0x01 << 21 //+ I2C2 Fast mode plus.
	I2C3_FMP    CFGR1_Bits = 0x01 << 22 //+ I2C3 Fast mode plus.
	FPU_IE_0    CFGR1_Bits = 0x01 << 26 //+ Invalid operation Interrupt enable.
	FPU_IE_1    CFGR1_Bits = 0x01 << 27 //+ Divide-by-zero Interrupt enable.
	FPU_IE_2    CFGR1_Bits = 0x01 << 28 //+ Underflow Interrupt enable.
	FPU_IE_3    CFGR1_Bits = 0x01 << 29 //+ Overflow Interrupt enable.
	FPU_IE_4    CFGR1_Bits = 0x01 << 30 //+ Input denormal Interrupt enable.
	FPU_IE_5    CFGR1_Bits = 0x01 << 31 //+ Inexact Interrupt enable (interrupt disabled at reset).
)

const (
	FWDISn       = 0
	BOOSTENn     = 8
	I2C_PB6_FMPn = 16
	I2C_PB7_FMPn = 17
	I2C_PB8_FMPn = 18
	I2C_PB9_FMPn = 19
	I2C1_FMPn    = 20
	I2C2_FMPn    = 21
	I2C3_FMPn    = 22
	FPU_IE_0n    = 26
	FPU_IE_1n    = 27
	FPU_IE_2n    = 28
	FPU_IE_3n    = 29
	FPU_IE_4n    = 30
	FPU_IE_5n    = 31
)

const (
	EXTI0    EXTICR_Bits = 0x07 << 0  //+ EXTI 0 configuration.
	EXTI1    EXTICR_Bits = 0x07 << 4  //+ EXTI 1 configuration.
	EXTI2    EXTICR_Bits = 0x07 << 8  //+ EXTI 2 configuration.
	EXTI3    EXTICR_Bits = 0x07 << 12 //+ EXTI 3 configuration.
	EXTI0_PA EXTICR_Bits = 0x00 << 12 //  PA[0] pin.
	EXTI0_PB EXTICR_Bits = 0x01 << 0  //  PB[0] pin.
	EXTI0_PC EXTICR_Bits = 0x02 << 0  //  PC[0] pin.
	EXTI0_PD EXTICR_Bits = 0x03 << 0  //  PD[0] pin.
	EXTI0_PE EXTICR_Bits = 0x04 << 0  //  PE[0] pin.
	EXTI0_PF EXTICR_Bits = 0x05 << 0  //  PF[0] pin.
	EXTI0_PG EXTICR_Bits = 0x06 << 0  //  PG[0] pin.
	EXTI0_PH EXTICR_Bits = 0x07 << 0  //  PH[0] pin.
	EXTI1_PA EXTICR_Bits = 0x00 << 12 //  PA[1] pin.
	EXTI1_PB EXTICR_Bits = 0x01 << 4  //  PB[1] pin.
	EXTI1_PC EXTICR_Bits = 0x02 << 4  //  PC[1] pin.
	EXTI1_PD EXTICR_Bits = 0x03 << 4  //  PD[1] pin.
	EXTI1_PE EXTICR_Bits = 0x04 << 4  //  PE[1] pin.
	EXTI1_PF EXTICR_Bits = 0x05 << 4  //  PF[1] pin.
	EXTI1_PG EXTICR_Bits = 0x06 << 4  //  PG[1] pin.
	EXTI1_PH EXTICR_Bits = 0x07 << 4  //  PH[1] pin.
	EXTI2_PA EXTICR_Bits = 0x00 << 12 //  PA[2] pin.
	EXTI2_PB EXTICR_Bits = 0x01 << 8  //  PB[2] pin.
	EXTI2_PC EXTICR_Bits = 0x02 << 8  //  PC[2] pin.
	EXTI2_PD EXTICR_Bits = 0x03 << 8  //  PD[2] pin.
	EXTI2_PE EXTICR_Bits = 0x04 << 8  //  PE[2] pin.
	EXTI2_PF EXTICR_Bits = 0x05 << 8  //  PF[2] pin.
	EXTI2_PG EXTICR_Bits = 0x06 << 8  //  PG[2] pin.
	EXTI3_PA EXTICR_Bits = 0x00 << 12 //  PA[3] pin.
	EXTI3_PB EXTICR_Bits = 0x01 << 12 //  PB[3] pin.
	EXTI3_PC EXTICR_Bits = 0x02 << 12 //  PC[3] pin.
	EXTI3_PD EXTICR_Bits = 0x03 << 12 //  PD[3] pin.
	EXTI3_PE EXTICR_Bits = 0x04 << 12 //  PE[3] pin.
	EXTI3_PF EXTICR_Bits = 0x05 << 12 //  PF[3] pin.
	EXTI3_PG EXTICR_Bits = 0x06 << 12 //  PG[3] pin.
)

const (
	EXTI0n = 0
	EXTI1n = 4
	EXTI2n = 8
	EXTI3n = 12
)

const (
	SRAM2ER  SCSR_Bits = 0x01 << 0 //+ SRAM2 Erase Request.
	SRAM2BSY SCSR_Bits = 0x01 << 1 //+ SRAM2 Erase Ongoing.
)

const (
	SRAM2ERn  = 0
	SRAM2BSYn = 1
)

const (
	CLL  CFGR2_Bits = 0x01 << 0 //+ Core Lockup Lock.
	SPL  CFGR2_Bits = 0x01 << 1 //+ SRAM Parity Lock.
	PVDL CFGR2_Bits = 0x01 << 2 //+ PVD Lock.
	ECCL CFGR2_Bits = 0x01 << 3 //+ ECC Lock.
	SPF  CFGR2_Bits = 0x01 << 8 //+ SRAM Parity Flag.
)

const (
	CLLn  = 0
	SPLn  = 1
	PVDLn = 2
	ECCLn = 3
	SPFn  = 8
)

const (
	PAGE0  SWPR_Bits = 0x01 << 0  //+ SRAM2 Write protection page 0.
	PAGE1  SWPR_Bits = 0x01 << 1  //+ SRAM2 Write protection page 1.
	PAGE2  SWPR_Bits = 0x01 << 2  //+ SRAM2 Write protection page 2.
	PAGE3  SWPR_Bits = 0x01 << 3  //+ SRAM2 Write protection page 3.
	PAGE4  SWPR_Bits = 0x01 << 4  //+ SRAM2 Write protection page 4.
	PAGE5  SWPR_Bits = 0x01 << 5  //+ SRAM2 Write protection page 5.
	PAGE6  SWPR_Bits = 0x01 << 6  //+ SRAM2 Write protection page 6.
	PAGE7  SWPR_Bits = 0x01 << 7  //+ SRAM2 Write protection page 7.
	PAGE8  SWPR_Bits = 0x01 << 8  //+ SRAM2 Write protection page 8.
	PAGE9  SWPR_Bits = 0x01 << 9  //+ SRAM2 Write protection page 9.
	PAGE10 SWPR_Bits = 0x01 << 10 //+ SRAM2 Write protection page 10.
	PAGE11 SWPR_Bits = 0x01 << 11 //+ SRAM2 Write protection page 11.
	PAGE12 SWPR_Bits = 0x01 << 12 //+ SRAM2 Write protection page 12.
	PAGE13 SWPR_Bits = 0x01 << 13 //+ SRAM2 Write protection page 13.
	PAGE14 SWPR_Bits = 0x01 << 14 //+ SRAM2 Write protection page 14.
	PAGE15 SWPR_Bits = 0x01 << 15 //+ SRAM2 Write protection page 15.
	PAGE16 SWPR_Bits = 0x01 << 16 //+ SRAM2 Write protection page 16.
	PAGE17 SWPR_Bits = 0x01 << 17 //+ SRAM2 Write protection page 17.
	PAGE18 SWPR_Bits = 0x01 << 18 //+ SRAM2 Write protection page 18.
	PAGE19 SWPR_Bits = 0x01 << 19 //+ SRAM2 Write protection page 19.
	PAGE20 SWPR_Bits = 0x01 << 20 //+ SRAM2 Write protection page 20.
	PAGE21 SWPR_Bits = 0x01 << 21 //+ SRAM2 Write protection page 21.
	PAGE22 SWPR_Bits = 0x01 << 22 //+ SRAM2 Write protection page 22.
	PAGE23 SWPR_Bits = 0x01 << 23 //+ SRAM2 Write protection page 23.
	PAGE24 SWPR_Bits = 0x01 << 24 //+ SRAM2 Write protection page 24.
	PAGE25 SWPR_Bits = 0x01 << 25 //+ SRAM2 Write protection page 25.
	PAGE26 SWPR_Bits = 0x01 << 26 //+ SRAM2 Write protection page 26.
	PAGE27 SWPR_Bits = 0x01 << 27 //+ SRAM2 Write protection page 27.
	PAGE28 SWPR_Bits = 0x01 << 28 //+ SRAM2 Write protection page 28.
	PAGE29 SWPR_Bits = 0x01 << 29 //+ SRAM2 Write protection page 29.
	PAGE30 SWPR_Bits = 0x01 << 30 //+ SRAM2 Write protection page 30.
	PAGE31 SWPR_Bits = 0x01 << 31 //+ SRAM2 Write protection page 31.
)

const (
	PAGE0n  = 0
	PAGE1n  = 1
	PAGE2n  = 2
	PAGE3n  = 3
	PAGE4n  = 4
	PAGE5n  = 5
	PAGE6n  = 6
	PAGE7n  = 7
	PAGE8n  = 8
	PAGE9n  = 9
	PAGE10n = 10
	PAGE11n = 11
	PAGE12n = 12
	PAGE13n = 13
	PAGE14n = 14
	PAGE15n = 15
	PAGE16n = 16
	PAGE17n = 17
	PAGE18n = 18
	PAGE19n = 19
	PAGE20n = 20
	PAGE21n = 21
	PAGE22n = 22
	PAGE23n = 23
	PAGE24n = 24
	PAGE25n = 25
	PAGE26n = 26
	PAGE27n = 27
	PAGE28n = 28
	PAGE29n = 29
	PAGE30n = 30
	PAGE31n = 31
)

const (
	KEY SKR_Bits = 0xFF << 0 //+ SRAM2 write protection key for software erase.
)

const (
	KEYn = 0
)
