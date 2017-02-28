// Peripheral: DFSDM_Channel_Periph  DFSDM channel configuration registers.
// Instances:
//  DFSDM1_Channel0  mmap.DFSDM1_Channel0_BASE
//  DFSDM1_Channel1  mmap.DFSDM1_Channel1_BASE
//  DFSDM1_Channel2  mmap.DFSDM1_Channel2_BASE
//  DFSDM1_Channel3  mmap.DFSDM1_Channel3_BASE
//  DFSDM1_Channel4  mmap.DFSDM1_Channel4_BASE
//  DFSDM1_Channel5  mmap.DFSDM1_Channel5_BASE
//  DFSDM1_Channel6  mmap.DFSDM1_Channel6_BASE
//  DFSDM1_Channel7  mmap.DFSDM1_Channel7_BASE
// Registers:
//  0x00 32  CHCFGR1  DFSDM channel configuration register1.
//  0x04 32  CHCFGR2  DFSDM channel configuration register2.
//  0x08 32  CHAWSCDR DFSDM channel analog watchdog and.
//  0x0C 32  CHWDATAR DFSDM channel watchdog filter data register.
//  0x10 32  CHDATINR DFSDM channel data input register.
// Import:
//  stm32/o/l476xx/mmap
package dfsdm

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	DFSDMEN  CHCFGR1_Bits = 0x01 << 31 //+ Global enable for DFSDM interface.
	CKOUTSRC CHCFGR1_Bits = 0x01 << 30 //+ Output serial clock source selection.
	CKOUTDIV CHCFGR1_Bits = 0xFF << 16 //+ CKOUTDIV[7:0] output serial clock divider.
	DATPACK  CHCFGR1_Bits = 0x03 << 14 //+ DATPACK[1:0] Data packing mode.
	DATMPX   CHCFGR1_Bits = 0x03 << 12 //+ DATMPX[1:0] Input data multiplexer for channel y.
	CHINSEL  CHCFGR1_Bits = 0x01 << 8  //+ Serial inputs selection for channel y.
	CHEN     CHCFGR1_Bits = 0x01 << 7  //+ Channel y enable.
	CKABEN   CHCFGR1_Bits = 0x01 << 6  //+ Clock absence detector enable on channel y.
	SCDEN    CHCFGR1_Bits = 0x01 << 5  //+ Short circuit detector enable on channel y.
	SPICKSEL CHCFGR1_Bits = 0x03 << 2  //+ SPICKSEL[1:0] SPI clock select for channel y.
	SITP     CHCFGR1_Bits = 0x03 << 0  //+ SITP[1:0] Serial interface type for channel y.
)

const (
	DFSDMENn  = 31
	CKOUTSRCn = 30
	CKOUTDIVn = 16
	DATPACKn  = 14
	DATMPXn   = 12
	CHINSELn  = 8
	CHENn     = 7
	CKABENn   = 6
	SCDENn    = 5
	SPICKSELn = 2
	SITPn     = 0
)

const (
	OFFSET CHCFGR2_Bits = 0xFFFFFF << 8 //+ OFFSET[23:0] 24-bit calibration offset for channel y.
	DTRBS  CHCFGR2_Bits = 0x1F << 3     //+ DTRBS[4:0] Data right bit-shift for channel y.
)

const (
	OFFSETn = 8
	DTRBSn  = 3
)

const (
	AWFORD CHAWSCDR_Bits = 0x03 << 22 //+ AWFORD[1:0] Analog watchdog Sinc filter order on channel y.
	AWFOSR CHAWSCDR_Bits = 0x1F << 16 //+ AWFOSR[4:0] Analog watchdog filter oversampling ratio on channel y.
	BKSCD  CHAWSCDR_Bits = 0x0F << 12 //+ BKSCD[3:0] Break signal assignment for short circuit detector on channel y.
	SCDT   CHAWSCDR_Bits = 0xFF << 0  //+ SCDT[7:0] Short circuit detector threshold for channel y.
)

const (
	AWFORDn = 22
	AWFOSRn = 16
	BKSCDn  = 12
	SCDTn   = 0
)

const (
	INDAT0 CHDATINR_Bits = 0xFFFF << 0  //+ INDAT0[31:16] Input data for channel y or channel (y+1).
	INDAT1 CHDATINR_Bits = 0xFFFF << 16 //+ INDAT0[15:0] Input data for channel y.
)

const (
	INDAT0n = 0
	INDAT1n = 16
)
