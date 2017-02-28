// Peripheral: RTC_Periph  Real-Time Clock.
// Instances:
//  RTC  mmap.RTC_BASE
// Registers:
//  0x00 32  TR         Time register.
//  0x04 32  DR         Date register.
//  0x08 32  CR         Control register.
//  0x0C 32  ISR        Initialization and status register.
//  0x10 32  PRER       Prescaler register.
//  0x14 32  WUTR       Wakeup timer register.
//  0x1C 32  ALRMR[2]   Alarm A, B registers.
//  0x24 32  WPR        Write protection register.
//  0x28 32  SSR        Sub second register.
//  0x2C 32  SHIFTR     Shift control register.
//  0x30 32  TSTR       Time stamp time register.
//  0x34 32  TSDR       Time stamp date register.
//  0x38 32  TSSSR      Time-stamp sub second register.
//  0x3C 32  CALR       Calibration register.
//  0x40 32  TAMPCR     Tamper configuration register.
//  0x44 32  ALRMSSR[2] Alarm A, B subsecond registers.
//  0x4C 32  OR         Option register.
//  0x50 32  BKPR[32]   Backup registers.
// Import:
//  stm32/o/l476xx/mmap
package rtc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	PM  TR_Bits = 0x01 << 22 //+
	HT  TR_Bits = 0x03 << 20 //+
	HU  TR_Bits = 0x0F << 16 //+
	MNT TR_Bits = 0x07 << 12 //+
	MNU TR_Bits = 0x0F << 8  //+
	ST  TR_Bits = 0x07 << 4  //+
	SU  TR_Bits = 0x0F << 0  //+
)

const (
	PMn  = 22
	HTn  = 20
	HUn  = 16
	MNTn = 12
	MNUn = 8
	STn  = 4
	SUn  = 0
)

const (
	YT  DR_Bits = 0x0F << 20 //+
	YU  DR_Bits = 0x0F << 16 //+
	WDU DR_Bits = 0x07 << 13 //+
	MT  DR_Bits = 0x01 << 12 //+
	MU  DR_Bits = 0x0F << 8  //+
	DT  DR_Bits = 0x03 << 4  //+
	DU  DR_Bits = 0x0F << 0  //+
)

const (
	YTn  = 20
	YUn  = 16
	WDUn = 13
	MTn  = 12
	MUn  = 8
	DTn  = 4
	DUn  = 0
)

const (
	ITSE    CR_Bits = 0x01 << 24 //+
	COE     CR_Bits = 0x01 << 23 //+
	OSEL    CR_Bits = 0x03 << 21 //+
	POL     CR_Bits = 0x01 << 20 //+
	COSEL   CR_Bits = 0x01 << 19 //+
	BCK     CR_Bits = 0x01 << 18 //+
	SUB1H   CR_Bits = 0x01 << 17 //+
	ADD1H   CR_Bits = 0x01 << 16 //+
	TSIE    CR_Bits = 0x01 << 15 //+
	WUTIE   CR_Bits = 0x01 << 14 //+
	ALRBIE  CR_Bits = 0x01 << 13 //+
	ALRAIE  CR_Bits = 0x01 << 12 //+
	TSE     CR_Bits = 0x01 << 11 //+
	WUTE    CR_Bits = 0x01 << 10 //+
	ALRBE   CR_Bits = 0x01 << 9  //+
	ALRAE   CR_Bits = 0x01 << 8  //+
	FMT     CR_Bits = 0x01 << 6  //+
	BYPSHAD CR_Bits = 0x01 << 5  //+
	REFCKON CR_Bits = 0x01 << 4  //+
	TSEDGE  CR_Bits = 0x01 << 3  //+
	WUCKSEL CR_Bits = 0x07 << 0  //+
)

const (
	ITSEn    = 24
	COEn     = 23
	OSELn    = 21
	POLn     = 20
	COSELn   = 19
	BCKn     = 18
	SUB1Hn   = 17
	ADD1Hn   = 16
	TSIEn    = 15
	WUTIEn   = 14
	ALRBIEn  = 13
	ALRAIEn  = 12
	TSEn     = 11
	WUTEn    = 10
	ALRBEn   = 9
	ALRAEn   = 8
	FMTn     = 6
	BYPSHADn = 5
	REFCKONn = 4
	TSEDGEn  = 3
	WUCKSELn = 0
)

const (
	ITSF    ISR_Bits = 0x01 << 17 //+
	RECALPF ISR_Bits = 0x01 << 16 //+
	TAMP3F  ISR_Bits = 0x01 << 15 //+
	TAMP2F  ISR_Bits = 0x01 << 14 //+
	TAMP1F  ISR_Bits = 0x01 << 13 //+
	TSOVF   ISR_Bits = 0x01 << 12 //+
	TSF     ISR_Bits = 0x01 << 11 //+
	WUTF    ISR_Bits = 0x01 << 10 //+
	ALRBF   ISR_Bits = 0x01 << 9  //+
	ALRAF   ISR_Bits = 0x01 << 8  //+
	INIT    ISR_Bits = 0x01 << 7  //+
	INITF   ISR_Bits = 0x01 << 6  //+
	RSF     ISR_Bits = 0x01 << 5  //+
	INITS   ISR_Bits = 0x01 << 4  //+
	SHPF    ISR_Bits = 0x01 << 3  //+
	WUTWF   ISR_Bits = 0x01 << 2  //+
	ALRBWF  ISR_Bits = 0x01 << 1  //+
	ALRAWF  ISR_Bits = 0x01 << 0  //+
)

const (
	ITSFn    = 17
	RECALPFn = 16
	TAMP3Fn  = 15
	TAMP2Fn  = 14
	TAMP1Fn  = 13
	TSOVFn   = 12
	TSFn     = 11
	WUTFn    = 10
	ALRBFn   = 9
	ALRAFn   = 8
	INITn    = 7
	INITFn   = 6
	RSFn     = 5
	INITSn   = 4
	SHPFn    = 3
	WUTWFn   = 2
	ALRBWFn  = 1
	ALRAWFn  = 0
)

const (
	PREDIV_A PRER_Bits = 0x7F << 16  //+
	PREDIV_S PRER_Bits = 0x7FFF << 0 //+
)

const (
	PREDIV_An = 16
	PREDIV_Sn = 0
)

const (
	WUT WUTR_Bits = 0xFFFF << 0 //+
)

const (
	WUTn = 0
)

const (
	AMSK4  ALRMR_Bits = 0x01 << 31 //+
	AWDSEL ALRMR_Bits = 0x01 << 30 //+
	ADT    ALRMR_Bits = 0x03 << 28 //+
	ADU    ALRMR_Bits = 0x0F << 24 //+
	AMSK3  ALRMR_Bits = 0x01 << 23 //+
	APM    ALRMR_Bits = 0x01 << 22 //+
	AHT    ALRMR_Bits = 0x03 << 20 //+
	AHU    ALRMR_Bits = 0x0F << 16 //+
	AMSK2  ALRMR_Bits = 0x01 << 15 //+
	AMNT   ALRMR_Bits = 0x07 << 12 //+
	AMNU   ALRMR_Bits = 0x0F << 8  //+
	AMSK1  ALRMR_Bits = 0x01 << 7  //+
	AST    ALRMR_Bits = 0x07 << 4  //+
	ASU    ALRMR_Bits = 0x0F << 0  //+
)

const (
	AMSK4n  = 31
	AWDSELn = 30
	ADTn    = 28
	ADUn    = 24
	AMSK3n  = 23
	APMn    = 22
	AHTn    = 20
	AHUn    = 16
	AMSK2n  = 15
	AMNTn   = 12
	AMNUn   = 8
	AMSK1n  = 7
	ASTn    = 4
	ASUn    = 0
)

const (
	KEY WPR_Bits = 0xFF << 0 //+
)

const (
	KEYn = 0
)

const (
	SS SSR_Bits = 0xFFFF << 0 //+
)

const (
	SSn = 0
)

const (
	SUBFS SHIFTR_Bits = 0x7FFF << 0 //+
	ADD1S SHIFTR_Bits = 0x01 << 31  //+
)

const (
	SUBFSn = 0
	ADD1Sn = 31
)

const (
	TPM  TSTR_Bits = 0x01 << 22 //+
	THT  TSTR_Bits = 0x03 << 20 //+
	THU  TSTR_Bits = 0x0F << 16 //+
	TMNT TSTR_Bits = 0x07 << 12 //+
	TMNU TSTR_Bits = 0x0F << 8  //+
	TST  TSTR_Bits = 0x07 << 4  //+
	TSU  TSTR_Bits = 0x0F << 0  //+
)

const (
	TPMn  = 22
	THTn  = 20
	THUn  = 16
	TMNTn = 12
	TMNUn = 8
	TSTn  = 4
	TSUn  = 0
)

const (
	TWDU TSDR_Bits = 0x07 << 13 //+
	TMT  TSDR_Bits = 0x01 << 12 //+
	TMU  TSDR_Bits = 0x0F << 8  //+
	TDT  TSDR_Bits = 0x03 << 4  //+
	TDU  TSDR_Bits = 0x0F << 0  //+
)

const (
	TWDUn = 13
	TMTn  = 12
	TMUn  = 8
	TDTn  = 4
	TDUn  = 0
)

const (
	TSS TSSSR_Bits = 0xFFFF << 0 //+
)

const (
	TSSn = 0
)

const (
	CALP   CALR_Bits = 0x01 << 15 //+
	CALW8  CALR_Bits = 0x01 << 14 //+
	CALW16 CALR_Bits = 0x01 << 13 //+
	CALM   CALR_Bits = 0x1FF << 0 //+
)

const (
	CALPn   = 15
	CALW8n  = 14
	CALW16n = 13
	CALMn   = 0
)

const (
	TAMP3MF      TAMPCR_Bits = 0x01 << 24 //+
	TAMP3NOERASE TAMPCR_Bits = 0x01 << 23 //+
	TAMP3IE      TAMPCR_Bits = 0x01 << 22 //+
	TAMP2MF      TAMPCR_Bits = 0x01 << 21 //+
	TAMP2NOERASE TAMPCR_Bits = 0x01 << 20 //+
	TAMP2IE      TAMPCR_Bits = 0x01 << 19 //+
	TAMP1MF      TAMPCR_Bits = 0x01 << 18 //+
	TAMP1NOERASE TAMPCR_Bits = 0x01 << 17 //+
	TAMP1IE      TAMPCR_Bits = 0x01 << 16 //+
	TAMPPUDIS    TAMPCR_Bits = 0x01 << 15 //+
	TAMPPRCH     TAMPCR_Bits = 0x03 << 13 //+
	TAMPFLT      TAMPCR_Bits = 0x03 << 11 //+
	TAMPFREQ     TAMPCR_Bits = 0x07 << 8  //+
	TAMPTS       TAMPCR_Bits = 0x01 << 7  //+
	TAMP3TRG     TAMPCR_Bits = 0x01 << 6  //+
	TAMP3E       TAMPCR_Bits = 0x01 << 5  //+
	TAMP2TRG     TAMPCR_Bits = 0x01 << 4  //+
	TAMP2E       TAMPCR_Bits = 0x01 << 3  //+
	TAMPIE       TAMPCR_Bits = 0x01 << 2  //+
	TAMP1TRG     TAMPCR_Bits = 0x01 << 1  //+
	TAMP1E       TAMPCR_Bits = 0x01 << 0  //+
)

const (
	TAMP3MFn      = 24
	TAMP3NOERASEn = 23
	TAMP3IEn      = 22
	TAMP2MFn      = 21
	TAMP2NOERASEn = 20
	TAMP2IEn      = 19
	TAMP1MFn      = 18
	TAMP1NOERASEn = 17
	TAMP1IEn      = 16
	TAMPPUDISn    = 15
	TAMPPRCHn     = 13
	TAMPFLTn      = 11
	TAMPFREQn     = 8
	TAMPTSn       = 7
	TAMP3TRGn     = 6
	TAMP3En       = 5
	TAMP2TRGn     = 4
	TAMP2En       = 3
	TAMPIEn       = 2
	TAMP1TRGn     = 1
	TAMP1En       = 0
)

const (
	AMASKSS ALRMSSR_Bits = 0x0F << 24  //+
	ASS     ALRMSSR_Bits = 0x7FFF << 0 //+
)

const (
	AMASKSSn = 24
	ASSn     = 0
)

const (
	OUT_RMP      OR_Bits = 0x01 << 1 //+
	ALARMOUTTYPE OR_Bits = 0x01 << 0 //+
)

const (
	OUT_RMPn      = 1
	ALARMOUTTYPEn = 0
)
