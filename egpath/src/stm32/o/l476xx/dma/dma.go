// Peripheral: DMA_Periph  DMA Controller.
// Instances:
//  DMA1  mmap.DMA1_BASE
//  DMA2  mmap.DMA2_BASE
// Registers:
//  0x00 32  ISR  Interrupt status register.
//  0x04 32  IFCR Interrupt flag clear register.
// Import:
//  stm32/o/l476xx/mmap
package dma

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	GIF1  ISR_Bits = 0x01 << 0  //+ Channel 1 Global interrupt flag.
	TCIF1 ISR_Bits = 0x01 << 1  //+ Channel 1 Transfer Complete flag.
	HTIF1 ISR_Bits = 0x01 << 2  //+ Channel 1 Half Transfer flag.
	TEIF1 ISR_Bits = 0x01 << 3  //+ Channel 1 Transfer Error flag.
	GIF2  ISR_Bits = 0x01 << 4  //+ Channel 2 Global interrupt flag.
	TCIF2 ISR_Bits = 0x01 << 5  //+ Channel 2 Transfer Complete flag.
	HTIF2 ISR_Bits = 0x01 << 6  //+ Channel 2 Half Transfer flag.
	TEIF2 ISR_Bits = 0x01 << 7  //+ Channel 2 Transfer Error flag.
	GIF3  ISR_Bits = 0x01 << 8  //+ Channel 3 Global interrupt flag.
	TCIF3 ISR_Bits = 0x01 << 9  //+ Channel 3 Transfer Complete flag.
	HTIF3 ISR_Bits = 0x01 << 10 //+ Channel 3 Half Transfer flag.
	TEIF3 ISR_Bits = 0x01 << 11 //+ Channel 3 Transfer Error flag.
	GIF4  ISR_Bits = 0x01 << 12 //+ Channel 4 Global interrupt flag.
	TCIF4 ISR_Bits = 0x01 << 13 //+ Channel 4 Transfer Complete flag.
	HTIF4 ISR_Bits = 0x01 << 14 //+ Channel 4 Half Transfer flag.
	TEIF4 ISR_Bits = 0x01 << 15 //+ Channel 4 Transfer Error flag.
	GIF5  ISR_Bits = 0x01 << 16 //+ Channel 5 Global interrupt flag.
	TCIF5 ISR_Bits = 0x01 << 17 //+ Channel 5 Transfer Complete flag.
	HTIF5 ISR_Bits = 0x01 << 18 //+ Channel 5 Half Transfer flag.
	TEIF5 ISR_Bits = 0x01 << 19 //+ Channel 5 Transfer Error flag.
	GIF6  ISR_Bits = 0x01 << 20 //+ Channel 6 Global interrupt flag.
	TCIF6 ISR_Bits = 0x01 << 21 //+ Channel 6 Transfer Complete flag.
	HTIF6 ISR_Bits = 0x01 << 22 //+ Channel 6 Half Transfer flag.
	TEIF6 ISR_Bits = 0x01 << 23 //+ Channel 6 Transfer Error flag.
	GIF7  ISR_Bits = 0x01 << 24 //+ Channel 7 Global interrupt flag.
	TCIF7 ISR_Bits = 0x01 << 25 //+ Channel 7 Transfer Complete flag.
	HTIF7 ISR_Bits = 0x01 << 26 //+ Channel 7 Half Transfer flag.
	TEIF7 ISR_Bits = 0x01 << 27 //+ Channel 7 Transfer Error flag.
)

const (
	GIF1n  = 0
	TCIF1n = 1
	HTIF1n = 2
	TEIF1n = 3
	GIF2n  = 4
	TCIF2n = 5
	HTIF2n = 6
	TEIF2n = 7
	GIF3n  = 8
	TCIF3n = 9
	HTIF3n = 10
	TEIF3n = 11
	GIF4n  = 12
	TCIF4n = 13
	HTIF4n = 14
	TEIF4n = 15
	GIF5n  = 16
	TCIF5n = 17
	HTIF5n = 18
	TEIF5n = 19
	GIF6n  = 20
	TCIF6n = 21
	HTIF6n = 22
	TEIF6n = 23
	GIF7n  = 24
	TCIF7n = 25
	HTIF7n = 26
	TEIF7n = 27
)

const (
	CGIF1  IFCR_Bits = 0x01 << 0  //+ Channel 1 Global interrupt clearr.
	CTCIF1 IFCR_Bits = 0x01 << 1  //+ Channel 1 Transfer Complete clear.
	CHTIF1 IFCR_Bits = 0x01 << 2  //+ Channel 1 Half Transfer clear.
	CTEIF1 IFCR_Bits = 0x01 << 3  //+ Channel 1 Transfer Error clear.
	CGIF2  IFCR_Bits = 0x01 << 4  //+ Channel 2 Global interrupt clear.
	CTCIF2 IFCR_Bits = 0x01 << 5  //+ Channel 2 Transfer Complete clear.
	CHTIF2 IFCR_Bits = 0x01 << 6  //+ Channel 2 Half Transfer clear.
	CTEIF2 IFCR_Bits = 0x01 << 7  //+ Channel 2 Transfer Error clear.
	CGIF3  IFCR_Bits = 0x01 << 8  //+ Channel 3 Global interrupt clear.
	CTCIF3 IFCR_Bits = 0x01 << 9  //+ Channel 3 Transfer Complete clear.
	CHTIF3 IFCR_Bits = 0x01 << 10 //+ Channel 3 Half Transfer clear.
	CTEIF3 IFCR_Bits = 0x01 << 11 //+ Channel 3 Transfer Error clear.
	CGIF4  IFCR_Bits = 0x01 << 12 //+ Channel 4 Global interrupt clear.
	CTCIF4 IFCR_Bits = 0x01 << 13 //+ Channel 4 Transfer Complete clear.
	CHTIF4 IFCR_Bits = 0x01 << 14 //+ Channel 4 Half Transfer clear.
	CTEIF4 IFCR_Bits = 0x01 << 15 //+ Channel 4 Transfer Error clear.
	CGIF5  IFCR_Bits = 0x01 << 16 //+ Channel 5 Global interrupt clear.
	CTCIF5 IFCR_Bits = 0x01 << 17 //+ Channel 5 Transfer Complete clear.
	CHTIF5 IFCR_Bits = 0x01 << 18 //+ Channel 5 Half Transfer clear.
	CTEIF5 IFCR_Bits = 0x01 << 19 //+ Channel 5 Transfer Error clear.
	CGIF6  IFCR_Bits = 0x01 << 20 //+ Channel 6 Global interrupt clear.
	CTCIF6 IFCR_Bits = 0x01 << 21 //+ Channel 6 Transfer Complete clear.
	CHTIF6 IFCR_Bits = 0x01 << 22 //+ Channel 6 Half Transfer clear.
	CTEIF6 IFCR_Bits = 0x01 << 23 //+ Channel 6 Transfer Error clear.
	CGIF7  IFCR_Bits = 0x01 << 24 //+ Channel 7 Global interrupt clear.
	CTCIF7 IFCR_Bits = 0x01 << 25 //+ Channel 7 Transfer Complete clear.
	CHTIF7 IFCR_Bits = 0x01 << 26 //+ Channel 7 Half Transfer clear.
	CTEIF7 IFCR_Bits = 0x01 << 27 //+ Channel 7 Transfer Error clear.
)

const (
	CGIF1n  = 0
	CTCIF1n = 1
	CHTIF1n = 2
	CTEIF1n = 3
	CGIF2n  = 4
	CTCIF2n = 5
	CHTIF2n = 6
	CTEIF2n = 7
	CGIF3n  = 8
	CTCIF3n = 9
	CHTIF3n = 10
	CTEIF3n = 11
	CGIF4n  = 12
	CTCIF4n = 13
	CHTIF4n = 14
	CTEIF4n = 15
	CGIF5n  = 16
	CTCIF5n = 17
	CHTIF5n = 18
	CTEIF5n = 19
	CGIF6n  = 20
	CTCIF6n = 21
	CHTIF6n = 22
	CTEIF6n = 23
	CGIF7n  = 24
	CTCIF7n = 25
	CHTIF7n = 26
	CTEIF7n = 27
)
