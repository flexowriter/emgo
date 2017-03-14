// Peripheral: ADC_Common_Periph  Analog to Digital Converter.
// Instances:
//  ADC123_COMMON  mmap.ADC123_COMMON_BASE
// Registers:
//  0x00 32  CSR ADC common status register.
//  0x08 32  CCR ADC common configuration register.
//  0x0C 32  CDR ADC common group regular data register.
// Import:
//  stm32/o/l476xx/mmap
package adc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	ADRDY_MST CSR_Bits = 0x01 << 0  //+ ADC multimode master ready flag.
	EOSMP_MST CSR_Bits = 0x01 << 1  //+ ADC multimode master group regular end of sampling flag.
	EOC_MST   CSR_Bits = 0x01 << 2  //+ ADC multimode master group regular end of unitary conversion flag.
	EOS_MST   CSR_Bits = 0x01 << 3  //+ ADC multimode master group regular end of sequence conversions flag.
	OVR_MST   CSR_Bits = 0x01 << 4  //+ ADC multimode master group regular overrun flag.
	JEOC_MST  CSR_Bits = 0x01 << 5  //+ ADC multimode master group injected end of unitary conversion flag.
	JEOS_MST  CSR_Bits = 0x01 << 6  //+ ADC multimode master group injected end of sequence conversions flag.
	AWD1_MST  CSR_Bits = 0x01 << 7  //+ ADC multimode master analog watchdog 1 flag.
	AWD2_MST  CSR_Bits = 0x01 << 8  //+ ADC multimode master analog watchdog 2 flag.
	AWD3_MST  CSR_Bits = 0x01 << 9  //+ ADC multimode master analog watchdog 3 flag.
	JQOVF_MST CSR_Bits = 0x01 << 10 //+ ADC multimode master group injected contexts queue overflow flag.
	ADRDY_SLV CSR_Bits = 0x01 << 16 //+ ADC multimode slave ready flag.
	EOSMP_SLV CSR_Bits = 0x01 << 17 //+ ADC multimode slave group regular end of sampling flag.
	EOC_SLV   CSR_Bits = 0x01 << 18 //+ ADC multimode slave group regular end of unitary conversion flag.
	EOS_SLV   CSR_Bits = 0x01 << 19 //+ ADC multimode slave group regular end of sequence conversions flag.
	OVR_SLV   CSR_Bits = 0x01 << 20 //+ ADC multimode slave group regular overrun flag.
	JEOC_SLV  CSR_Bits = 0x01 << 21 //+ ADC multimode slave group injected end of unitary conversion flag.
	JEOS_SLV  CSR_Bits = 0x01 << 22 //+ ADC multimode slave group injected end of sequence conversions flag.
	AWD1_SLV  CSR_Bits = 0x01 << 23 //+ ADC multimode slave analog watchdog 1 flag.
	AWD2_SLV  CSR_Bits = 0x01 << 24 //+ ADC multimode slave analog watchdog 2 flag.
	AWD3_SLV  CSR_Bits = 0x01 << 25 //+ ADC multimode slave analog watchdog 3 flag.
	JQOVF_SLV CSR_Bits = 0x01 << 26 //+ ADC multimode slave group injected contexts queue overflow flag.
)

const (
	ADRDY_MSTn = 0
	EOSMP_MSTn = 1
	EOC_MSTn   = 2
	EOS_MSTn   = 3
	OVR_MSTn   = 4
	JEOC_MSTn  = 5
	JEOS_MSTn  = 6
	AWD1_MSTn  = 7
	AWD2_MSTn  = 8
	AWD3_MSTn  = 9
	JQOVF_MSTn = 10
	ADRDY_SLVn = 16
	EOSMP_SLVn = 17
	EOC_SLVn   = 18
	EOS_SLVn   = 19
	OVR_SLVn   = 20
	JEOC_SLVn  = 21
	JEOS_SLVn  = 22
	AWD1_SLVn  = 23
	AWD2_SLVn  = 24
	AWD3_SLVn  = 25
	JQOVF_SLVn = 26
)

const (
	DUAL    CCR_Bits = 0x1F << 0  //+ ADC multimode mode selection.
	DELAY   CCR_Bits = 0x0F << 8  //+ ADC multimode delay between 2 sampling phases.
	MDMACFG CCR_Bits = 0x01 << 13 //+ ADC multimode DMA transfer configuration.
	MDMA    CCR_Bits = 0x03 << 14 //+ ADC multimode DMA transfer enable.
	CKMODE  CCR_Bits = 0x03 << 16 //+ ADC common clock source and prescaler (prescaler only for clock source synchronous).
	PRESC   CCR_Bits = 0x0F << 18 //+ ADC common clock prescaler, only for clock source asynchronous.
	VREFEN  CCR_Bits = 0x01 << 22 //+ ADC internal path to VrefInt enable.
	TSEN    CCR_Bits = 0x01 << 23 //+ ADC internal path to temperature sensor enable.
	VBATEN  CCR_Bits = 0x01 << 24 //+ ADC internal path to battery voltage enable.
)

const (
	DUALn    = 0
	DELAYn   = 8
	MDMACFGn = 13
	MDMAn    = 14
	CKMODEn  = 16
	PRESCn   = 18
	VREFENn  = 22
	TSENn    = 23
	VBATENn  = 24
)

const (
	RDATA_MST CDR_Bits = 0xFFFF << 0  //+ ADC multimode master group regular conversion data.
	RDATA_SLV CDR_Bits = 0xFFFF << 16 //+ ADC multimode slave group regular conversion data.
)

const (
	RDATA_MSTn = 0
	RDATA_SLVn = 16
)