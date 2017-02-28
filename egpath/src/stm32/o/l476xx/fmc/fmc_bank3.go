// Peripheral: FMC_Bank3_Periph  Flexible Memory Controller Bank3.
// Instances:
//  FMC_Bank3_R  mmap.FMC_Bank3_R_BASE
// Registers:
//  0x00 32  PCR  NAND Flash control register.
//  0x04 32  SR   NAND Flash FIFO status and interrupt register.
//  0x08 32  PMEM NAND Flash Common memory space timing register.
//  0x0C 32  PATT NAND Flash Attribute memory space timing register.
//  0x14 32  ECCR NAND Flash ECC result registers.
// Import:
//  stm32/o/l476xx/mmap
package fmc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PWAITEN PCR_Bits = 0x01 << 1  //+ Wait feature enable bit.
	PBKEN   PCR_Bits = 0x01 << 2  //+ NAND Flash memory bank enable bit.
	PTYP    PCR_Bits = 0x01 << 3  //+ Memory type.
	PWID    PCR_Bits = 0x03 << 4  //+ PWID[1:0] bits (NAND Flash databus width).
	ECCEN   PCR_Bits = 0x01 << 6  //+ ECC computation logic enable bit.
	TCLR    PCR_Bits = 0x0F << 9  //+ TCLR[3:0] bits (CLE to RE delay).
	TAR     PCR_Bits = 0x0F << 13 //+ TAR[3:0] bits (ALE to RE delay).
	ECCPS   PCR_Bits = 0x07 << 17 //+ ECCPS[1:0] bits (ECC page size).
)

const (
	PWAITENn = 1
	PBKENn   = 2
	PTYPn    = 3
	PWIDn    = 4
	ECCENn   = 6
	TCLRn    = 9
	TARn     = 13
	ECCPSn   = 17
)

const (
	IRS   SR_Bits = 0x01 << 0 //+ Interrupt Rising Edge status.
	ILS   SR_Bits = 0x01 << 1 //+ Interrupt Level status.
	IFS   SR_Bits = 0x01 << 2 //+ Interrupt Falling Edge status.
	IREN  SR_Bits = 0x01 << 3 //+ Interrupt Rising Edge detection Enable bit.
	ILEN  SR_Bits = 0x01 << 4 //+ Interrupt Level detection Enable bit.
	IFEN  SR_Bits = 0x01 << 5 //+ Interrupt Falling Edge detection Enable bit.
	FEMPT SR_Bits = 0x01 << 6 //+ FIFO empty.
)

const (
	IRSn   = 0
	ILSn   = 1
	IFSn   = 2
	IRENn  = 3
	ILENn  = 4
	IFENn  = 5
	FEMPTn = 6
)

const (
	MEMSET  PMEM_Bits = 0xFF << 0  //+ MEMSET[7:0] bits (Common memory setup time).
	MEMWAIT PMEM_Bits = 0xFF << 8  //+ MEMWAIT[7:0] bits (Common memory wait time).
	MEMHOLD PMEM_Bits = 0xFF << 16 //+ MEMHOLD[7:0] bits (Common memory hold time).
	MEMHIZ  PMEM_Bits = 0xFF << 24 //+ MEMHIZ[7:0] bits (Common memory databus HiZ time).
)

const (
	MEMSETn  = 0
	MEMWAITn = 8
	MEMHOLDn = 16
	MEMHIZn  = 24
)

const (
	ATTSET  PATT_Bits = 0xFF << 0  //+ ATTSET[7:0] bits (Attribute memory setup time).
	ATTWAIT PATT_Bits = 0xFF << 8  //+ ATTWAIT[7:0] bits (Attribute memory wait time).
	ATTHOLD PATT_Bits = 0xFF << 16 //+ ATTHOLD[7:0] bits (Attribute memory hold time).
	ATTHIZ  PATT_Bits = 0xFF << 24 //+ ATTHIZ[7:0] bits (Attribute memory databus HiZ time).
)

const (
	ATTSETn  = 0
	ATTWAITn = 8
	ATTHOLDn = 16
	ATTHIZn  = 24
)

const (
	ECC ECCR_Bits = 0xFFFFFFFF << 0 //+ ECC result.
)

const (
	ECCn = 0
)
