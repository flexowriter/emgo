// +build l476xx

// Peripheral: FLASH_Periph  FLASH Registers.
// Instances:
//  FLASH  mmap.FLASH_R_BASE
// Registers:
//  0x00 32  ACR       Access control register.
//  0x04 32  PDKEYR    Power down key register.
//  0x08 32  KEYR      Key register.
//  0x0C 32  OPTKEYR   Option key register.
//  0x10 32  SR        Status register.
//  0x14 32  CR        Control register.
//  0x18 32  ECCR      ECC register.
//  0x1C 32  RESERVED1 Reserved1.
//  0x20 32  OPTR      Option register.
//  0x24 32  PCROP1SR  Bank1 PCROP start address register.
//  0x28 32  PCROP1ER  Bank1 PCROP end address register.
//  0x2C 32  WRP1AR    Bank1 WRP area A address register.
//  0x30 32  WRP1BR    Bank1 WRP area B address register.
//  0x44 32  PCROP2SR  Bank2 PCROP start address register.
//  0x48 32  PCROP2ER  Bank2 PCROP end address register.
//  0x4C 32  WRP2AR    Bank2 WRP area A address register.
//  0x50 32  WRP2BR    Bank2 WRP area B address register.
// Import:
//  stm32/o/l476xx/mmap
package flash

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	LATENCY     ACR_Bits = 0x07 << 0 //+
	LATENCY_0WS ACR_Bits = 0x00 << 0
	LATENCY_1WS ACR_Bits = 0x01 << 0
	LATENCY_2WS ACR_Bits = 0x02 << 0
	LATENCY_3WS ACR_Bits = 0x03 << 0
	LATENCY_4WS ACR_Bits = 0x04 << 0
	PRFTEN      ACR_Bits = 0x01 << 8  //+
	ICEN        ACR_Bits = 0x01 << 9  //+
	DCEN        ACR_Bits = 0x01 << 10 //+
	ICRST       ACR_Bits = 0x01 << 11 //+
	DCRST       ACR_Bits = 0x01 << 12 //+
	RUN_PD      ACR_Bits = 0x01 << 13 //+ Flash power down mode during run.
	SLEEP_PD    ACR_Bits = 0x01 << 14 //+ Flash power down mode during sleep.
)

const (
	LATENCYn  = 0
	PRFTENn   = 8
	ICENn     = 9
	DCENn     = 10
	ICRSTn    = 11
	DCRSTn    = 12
	RUN_PDn   = 13
	SLEEP_PDn = 14
)

const (
	EOP     SR_Bits = 0x01 << 0  //+
	OPERR   SR_Bits = 0x01 << 1  //+
	PROGERR SR_Bits = 0x01 << 3  //+
	WRPERR  SR_Bits = 0x01 << 4  //+
	PGAERR  SR_Bits = 0x01 << 5  //+
	SIZERR  SR_Bits = 0x01 << 6  //+
	PGSERR  SR_Bits = 0x01 << 7  //+
	MISERR  SR_Bits = 0x01 << 8  //+
	FASTERR SR_Bits = 0x01 << 9  //+
	RDERR   SR_Bits = 0x01 << 14 //+
	OPTVERR SR_Bits = 0x01 << 15 //+
	BSY     SR_Bits = 0x01 << 16 //+
)

const (
	EOPn     = 0
	OPERRn   = 1
	PROGERRn = 3
	WRPERRn  = 4
	PGAERRn  = 5
	SIZERRn  = 6
	PGSERRn  = 7
	MISERRn  = 8
	FASTERRn = 9
	RDERRn   = 14
	OPTVERRn = 15
	BSYn     = 16
)

const (
	PG         CR_Bits = 0x01 << 0  //+
	PER        CR_Bits = 0x01 << 1  //+
	MER1       CR_Bits = 0x01 << 2  //+
	PNB        CR_Bits = 0xFF << 3  //+
	BKER       CR_Bits = 0x01 << 11 //+
	MER2       CR_Bits = 0x01 << 15 //+
	STRT       CR_Bits = 0x01 << 16 //+
	OPTSTRT    CR_Bits = 0x01 << 17 //+
	FSTPG      CR_Bits = 0x01 << 18 //+
	EOPIE      CR_Bits = 0x01 << 24 //+
	ERRIE      CR_Bits = 0x01 << 25 //+
	RDERRIE    CR_Bits = 0x01 << 26 //+
	OBL_LAUNCH CR_Bits = 0x01 << 27 //+
	OPTLOCK    CR_Bits = 0x01 << 30 //+
	LOCK       CR_Bits = 0x01 << 31 //+
)

const (
	PGn         = 0
	PERn        = 1
	MER1n       = 2
	PNBn        = 3
	BKERn       = 11
	MER2n       = 15
	STRTn       = 16
	OPTSTRTn    = 17
	FSTPGn      = 18
	EOPIEn      = 24
	ERRIEn      = 25
	RDERRIEn    = 26
	OBL_LAUNCHn = 27
	OPTLOCKn    = 30
	LOCKn       = 31
)

const (
	ADDR_ECC ECCR_Bits = 0x7FFFF << 0 //+
	BK_ECC   ECCR_Bits = 0x01 << 19   //+
	SYSF_ECC ECCR_Bits = 0x01 << 20   //+
	ECCIE    ECCR_Bits = 0x01 << 24   //+
	ECCC     ECCR_Bits = 0x01 << 30   //+
	ECCD     ECCR_Bits = 0x01 << 31   //+
)

const (
	ADDR_ECCn = 0
	BK_ECCn   = 19
	SYSF_ECCn = 20
	ECCIEn    = 24
	ECCCn     = 30
	ECCDn     = 31
)

const (
	RDP        OPTR_Bits = 0xFF << 0  //+
	BOR_LEV    OPTR_Bits = 0x07 << 8  //+
	nRST_STOP  OPTR_Bits = 0x01 << 12 //+
	nRST_STDBY OPTR_Bits = 0x01 << 13 //+
	nRST_SHDW  OPTR_Bits = 0x01 << 14 //+
	IWDG_SW    OPTR_Bits = 0x01 << 16 //+
	IWDG_STOP  OPTR_Bits = 0x01 << 17 //+
	IWDG_STDBY OPTR_Bits = 0x01 << 18 //+
	WWDG_SW    OPTR_Bits = 0x01 << 19 //+
	BFB2       OPTR_Bits = 0x01 << 20 //+
	DUALBANK   OPTR_Bits = 0x01 << 21 //+
	nBOOT1     OPTR_Bits = 0x01 << 23 //+
	SRAM2_PE   OPTR_Bits = 0x01 << 24 //+
	SRAM2_RST  OPTR_Bits = 0x01 << 25 //+
)

const (
	RDPn        = 0
	BOR_LEVn    = 8
	nRST_STOPn  = 12
	nRST_STDBYn = 13
	nRST_SHDWn  = 14
	IWDG_SWn    = 16
	IWDG_STOPn  = 17
	IWDG_STDBYn = 18
	WWDG_SWn    = 19
	BFB2n       = 20
	DUALBANKn   = 21
	nBOOT1n     = 23
	SRAM2_PEn   = 24
	SRAM2_RSTn  = 25
)

const (
	PCROP1_STRT PCROP1SR_Bits = 0xFFFF << 0 //+
)

const (
	PCROP1_STRTn = 0
)

const (
	PCROP1_END PCROP1ER_Bits = 0xFFFF << 0 //+
	PCROP_RDP  PCROP1ER_Bits = 0x01 << 31  //+
)

const (
	PCROP1_ENDn = 0
	PCROP_RDPn  = 31
)

const (
	WRP1A_STRT WRP1AR_Bits = 0xFF << 0  //+
	WRP1A_END  WRP1AR_Bits = 0xFF << 16 //+
)

const (
	WRP1A_STRTn = 0
	WRP1A_ENDn  = 16
)

const (
	WRP1B_STRT WRP1BR_Bits = 0xFF << 0  //+
	WRP1B_END  WRP1BR_Bits = 0xFF << 16 //+
)

const (
	WRP1B_STRTn = 0
	WRP1B_ENDn  = 16
)

const (
	PCROP2_STRT PCROP2SR_Bits = 0xFFFF << 0 //+
)

const (
	PCROP2_STRTn = 0
)

const (
	PCROP2_END PCROP2ER_Bits = 0xFFFF << 0 //+
)

const (
	PCROP2_ENDn = 0
)

const (
	WRP2A_STRT WRP2AR_Bits = 0xFF << 0  //+
	WRP2A_END  WRP2AR_Bits = 0xFF << 16 //+
)

const (
	WRP2A_STRTn = 0
	WRP2A_ENDn  = 16
)

const (
	WRP2B_STRT WRP2BR_Bits = 0xFF << 0  //+
	WRP2B_END  WRP2BR_Bits = 0xFF << 16 //+
)

const (
	WRP2B_STRTn = 0
	WRP2B_ENDn  = 16
)
