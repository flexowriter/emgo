// Peripheral: EXTI_Periph  External Interrupt/Event Controller.
// Instances:
//  EXTI  mmap.EXTI_BASE
// Registers:
//  0x00 32  IMR1   Interrupt mask register 1.
//  0x04 32  EMR1   Event mask register 1.
//  0x08 32  RTSR1  Rising trigger selection register 1.
//  0x0C 32  FTSR1  Falling trigger selection register 1.
//  0x10 32  SWIER1 Software interrupt event register 1.
//  0x14 32  PR1    Pending register 1.
//  0x20 32  IMR2   Interrupt mask register 2.
//  0x24 32  EMR2   Event mask register 2.
//  0x28 32  RTSR2  Rising trigger selection register 2.
//  0x2C 32  FTSR2  Falling trigger selection register 2.
//  0x30 32  SWIER2 Software interrupt event register 2.
//  0x34 32  PR2    Pending register 2.
// Import:
//  stm32/o/l476xx/mmap
package exti

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	IM0  IMR1_Bits = 0x01 << 0       //+ Interrupt Mask on line 0.
	IM1  IMR1_Bits = 0x01 << 1       //+ Interrupt Mask on line 1.
	IM2  IMR1_Bits = 0x01 << 2       //+ Interrupt Mask on line 2.
	IM3  IMR1_Bits = 0x01 << 3       //+ Interrupt Mask on line 3.
	IM4  IMR1_Bits = 0x01 << 4       //+ Interrupt Mask on line 4.
	IM5  IMR1_Bits = 0x01 << 5       //+ Interrupt Mask on line 5.
	IM6  IMR1_Bits = 0x01 << 6       //+ Interrupt Mask on line 6.
	IM7  IMR1_Bits = 0x01 << 7       //+ Interrupt Mask on line 7.
	IM8  IMR1_Bits = 0x01 << 8       //+ Interrupt Mask on line 8.
	IM9  IMR1_Bits = 0x01 << 9       //+ Interrupt Mask on line 9.
	IM10 IMR1_Bits = 0x01 << 10      //+ Interrupt Mask on line 10.
	IM11 IMR1_Bits = 0x01 << 11      //+ Interrupt Mask on line 11.
	IM12 IMR1_Bits = 0x01 << 12      //+ Interrupt Mask on line 12.
	IM13 IMR1_Bits = 0x01 << 13      //+ Interrupt Mask on line 13.
	IM14 IMR1_Bits = 0x01 << 14      //+ Interrupt Mask on line 14.
	IM15 IMR1_Bits = 0x01 << 15      //+ Interrupt Mask on line 15.
	IM16 IMR1_Bits = 0x01 << 16      //+ Interrupt Mask on line 16.
	IM17 IMR1_Bits = 0x01 << 17      //+ Interrupt Mask on line 17.
	IM18 IMR1_Bits = 0x01 << 18      //+ Interrupt Mask on line 18.
	IM19 IMR1_Bits = 0x01 << 19      //+ Interrupt Mask on line 19.
	IM20 IMR1_Bits = 0x01 << 20      //+ Interrupt Mask on line 20.
	IM21 IMR1_Bits = 0x01 << 21      //+ Interrupt Mask on line 21.
	IM22 IMR1_Bits = 0x01 << 22      //+ Interrupt Mask on line 22.
	IM23 IMR1_Bits = 0x01 << 23      //+ Interrupt Mask on line 23.
	IM24 IMR1_Bits = 0x01 << 24      //+ Interrupt Mask on line 24.
	IM25 IMR1_Bits = 0x01 << 25      //+ Interrupt Mask on line 25.
	IM26 IMR1_Bits = 0x01 << 26      //+ Interrupt Mask on line 26.
	IM27 IMR1_Bits = 0x01 << 27      //+ Interrupt Mask on line 27.
	IM28 IMR1_Bits = 0x01 << 28      //+ Interrupt Mask on line 28.
	IM29 IMR1_Bits = 0x01 << 29      //+ Interrupt Mask on line 29.
	IM30 IMR1_Bits = 0x01 << 30      //+ Interrupt Mask on line 30.
	IM31 IMR1_Bits = 0x01 << 31      //+ Interrupt Mask on line 31.
	IM   IMR1_Bits = 0xFFFFFFFF << 0 //  Interrupt Mask All.
)

const (
	IM0n  = 0
	IM1n  = 1
	IM2n  = 2
	IM3n  = 3
	IM4n  = 4
	IM5n  = 5
	IM6n  = 6
	IM7n  = 7
	IM8n  = 8
	IM9n  = 9
	IM10n = 10
	IM11n = 11
	IM12n = 12
	IM13n = 13
	IM14n = 14
	IM15n = 15
	IM16n = 16
	IM17n = 17
	IM18n = 18
	IM19n = 19
	IM20n = 20
	IM21n = 21
	IM22n = 22
	IM23n = 23
	IM24n = 24
	IM25n = 25
	IM26n = 26
	IM27n = 27
	IM28n = 28
	IM29n = 29
	IM30n = 30
	IM31n = 31
)

const (
	EM0  EMR1_Bits = 0x01 << 0  //+ Event Mask on line 0.
	EM1  EMR1_Bits = 0x01 << 1  //+ Event Mask on line 1.
	EM2  EMR1_Bits = 0x01 << 2  //+ Event Mask on line 2.
	EM3  EMR1_Bits = 0x01 << 3  //+ Event Mask on line 3.
	EM4  EMR1_Bits = 0x01 << 4  //+ Event Mask on line 4.
	EM5  EMR1_Bits = 0x01 << 5  //+ Event Mask on line 5.
	EM6  EMR1_Bits = 0x01 << 6  //+ Event Mask on line 6.
	EM7  EMR1_Bits = 0x01 << 7  //+ Event Mask on line 7.
	EM8  EMR1_Bits = 0x01 << 8  //+ Event Mask on line 8.
	EM9  EMR1_Bits = 0x01 << 9  //+ Event Mask on line 9.
	EM10 EMR1_Bits = 0x01 << 10 //+ Event Mask on line 10.
	EM11 EMR1_Bits = 0x01 << 11 //+ Event Mask on line 11.
	EM12 EMR1_Bits = 0x01 << 12 //+ Event Mask on line 12.
	EM13 EMR1_Bits = 0x01 << 13 //+ Event Mask on line 13.
	EM14 EMR1_Bits = 0x01 << 14 //+ Event Mask on line 14.
	EM15 EMR1_Bits = 0x01 << 15 //+ Event Mask on line 15.
	EM16 EMR1_Bits = 0x01 << 16 //+ Event Mask on line 16.
	EM17 EMR1_Bits = 0x01 << 17 //+ Event Mask on line 17.
	EM18 EMR1_Bits = 0x01 << 18 //+ Event Mask on line 18.
	EM19 EMR1_Bits = 0x01 << 19 //+ Event Mask on line 19.
	EM20 EMR1_Bits = 0x01 << 20 //+ Event Mask on line 20.
	EM21 EMR1_Bits = 0x01 << 21 //+ Event Mask on line 21.
	EM22 EMR1_Bits = 0x01 << 22 //+ Event Mask on line 22.
	EM23 EMR1_Bits = 0x01 << 23 //+ Event Mask on line 23.
	EM24 EMR1_Bits = 0x01 << 24 //+ Event Mask on line 24.
	EM25 EMR1_Bits = 0x01 << 25 //+ Event Mask on line 25.
	EM26 EMR1_Bits = 0x01 << 26 //+ Event Mask on line 26.
	EM27 EMR1_Bits = 0x01 << 27 //+ Event Mask on line 27.
	EM28 EMR1_Bits = 0x01 << 28 //+ Event Mask on line 28.
	EM29 EMR1_Bits = 0x01 << 29 //+ Event Mask on line 29.
	EM30 EMR1_Bits = 0x01 << 30 //+ Event Mask on line 30.
	EM31 EMR1_Bits = 0x01 << 31 //+ Event Mask on line 31.
)

const (
	EM0n  = 0
	EM1n  = 1
	EM2n  = 2
	EM3n  = 3
	EM4n  = 4
	EM5n  = 5
	EM6n  = 6
	EM7n  = 7
	EM8n  = 8
	EM9n  = 9
	EM10n = 10
	EM11n = 11
	EM12n = 12
	EM13n = 13
	EM14n = 14
	EM15n = 15
	EM16n = 16
	EM17n = 17
	EM18n = 18
	EM19n = 19
	EM20n = 20
	EM21n = 21
	EM22n = 22
	EM23n = 23
	EM24n = 24
	EM25n = 25
	EM26n = 26
	EM27n = 27
	EM28n = 28
	EM29n = 29
	EM30n = 30
	EM31n = 31
)

const (
	RT0  RTSR1_Bits = 0x01 << 0  //+ Rising trigger event configuration bit of line 0.
	RT1  RTSR1_Bits = 0x01 << 1  //+ Rising trigger event configuration bit of line 1.
	RT2  RTSR1_Bits = 0x01 << 2  //+ Rising trigger event configuration bit of line 2.
	RT3  RTSR1_Bits = 0x01 << 3  //+ Rising trigger event configuration bit of line 3.
	RT4  RTSR1_Bits = 0x01 << 4  //+ Rising trigger event configuration bit of line 4.
	RT5  RTSR1_Bits = 0x01 << 5  //+ Rising trigger event configuration bit of line 5.
	RT6  RTSR1_Bits = 0x01 << 6  //+ Rising trigger event configuration bit of line 6.
	RT7  RTSR1_Bits = 0x01 << 7  //+ Rising trigger event configuration bit of line 7.
	RT8  RTSR1_Bits = 0x01 << 8  //+ Rising trigger event configuration bit of line 8.
	RT9  RTSR1_Bits = 0x01 << 9  //+ Rising trigger event configuration bit of line 9.
	RT10 RTSR1_Bits = 0x01 << 10 //+ Rising trigger event configuration bit of line 10.
	RT11 RTSR1_Bits = 0x01 << 11 //+ Rising trigger event configuration bit of line 11.
	RT12 RTSR1_Bits = 0x01 << 12 //+ Rising trigger event configuration bit of line 12.
	RT13 RTSR1_Bits = 0x01 << 13 //+ Rising trigger event configuration bit of line 13.
	RT14 RTSR1_Bits = 0x01 << 14 //+ Rising trigger event configuration bit of line 14.
	RT15 RTSR1_Bits = 0x01 << 15 //+ Rising trigger event configuration bit of line 15.
	RT16 RTSR1_Bits = 0x01 << 16 //+ Rising trigger event configuration bit of line 16.
	RT18 RTSR1_Bits = 0x01 << 18 //+ Rising trigger event configuration bit of line 18.
	RT19 RTSR1_Bits = 0x01 << 19 //+ Rising trigger event configuration bit of line 19.
	RT20 RTSR1_Bits = 0x01 << 20 //+ Rising trigger event configuration bit of line 20.
	RT21 RTSR1_Bits = 0x01 << 21 //+ Rising trigger event configuration bit of line 21.
	RT22 RTSR1_Bits = 0x01 << 22 //+ Rising trigger event configuration bit of line 22.
)

const (
	RT0n  = 0
	RT1n  = 1
	RT2n  = 2
	RT3n  = 3
	RT4n  = 4
	RT5n  = 5
	RT6n  = 6
	RT7n  = 7
	RT8n  = 8
	RT9n  = 9
	RT10n = 10
	RT11n = 11
	RT12n = 12
	RT13n = 13
	RT14n = 14
	RT15n = 15
	RT16n = 16
	RT18n = 18
	RT19n = 19
	RT20n = 20
	RT21n = 21
	RT22n = 22
)

const (
	FT0  FTSR1_Bits = 0x01 << 0  //+ Falling trigger event configuration bit of line 0.
	FT1  FTSR1_Bits = 0x01 << 1  //+ Falling trigger event configuration bit of line 1.
	FT2  FTSR1_Bits = 0x01 << 2  //+ Falling trigger event configuration bit of line 2.
	FT3  FTSR1_Bits = 0x01 << 3  //+ Falling trigger event configuration bit of line 3.
	FT4  FTSR1_Bits = 0x01 << 4  //+ Falling trigger event configuration bit of line 4.
	FT5  FTSR1_Bits = 0x01 << 5  //+ Falling trigger event configuration bit of line 5.
	FT6  FTSR1_Bits = 0x01 << 6  //+ Falling trigger event configuration bit of line 6.
	FT7  FTSR1_Bits = 0x01 << 7  //+ Falling trigger event configuration bit of line 7.
	FT8  FTSR1_Bits = 0x01 << 8  //+ Falling trigger event configuration bit of line 8.
	FT9  FTSR1_Bits = 0x01 << 9  //+ Falling trigger event configuration bit of line 9.
	FT10 FTSR1_Bits = 0x01 << 10 //+ Falling trigger event configuration bit of line 10.
	FT11 FTSR1_Bits = 0x01 << 11 //+ Falling trigger event configuration bit of line 11.
	FT12 FTSR1_Bits = 0x01 << 12 //+ Falling trigger event configuration bit of line 12.
	FT13 FTSR1_Bits = 0x01 << 13 //+ Falling trigger event configuration bit of line 13.
	FT14 FTSR1_Bits = 0x01 << 14 //+ Falling trigger event configuration bit of line 14.
	FT15 FTSR1_Bits = 0x01 << 15 //+ Falling trigger event configuration bit of line 15.
	FT16 FTSR1_Bits = 0x01 << 16 //+ Falling trigger event configuration bit of line 16.
	FT18 FTSR1_Bits = 0x01 << 18 //+ Falling trigger event configuration bit of line 18.
	FT19 FTSR1_Bits = 0x01 << 19 //+ Falling trigger event configuration bit of line 19.
	FT20 FTSR1_Bits = 0x01 << 20 //+ Falling trigger event configuration bit of line 20.
	FT21 FTSR1_Bits = 0x01 << 21 //+ Falling trigger event configuration bit of line 21.
	FT22 FTSR1_Bits = 0x01 << 22 //+ Falling trigger event configuration bit of line 22.
)

const (
	FT0n  = 0
	FT1n  = 1
	FT2n  = 2
	FT3n  = 3
	FT4n  = 4
	FT5n  = 5
	FT6n  = 6
	FT7n  = 7
	FT8n  = 8
	FT9n  = 9
	FT10n = 10
	FT11n = 11
	FT12n = 12
	FT13n = 13
	FT14n = 14
	FT15n = 15
	FT16n = 16
	FT18n = 18
	FT19n = 19
	FT20n = 20
	FT21n = 21
	FT22n = 22
)

const (
	SWI0  SWIER1_Bits = 0x01 << 0  //+ Software Interrupt on line 0.
	SWI1  SWIER1_Bits = 0x01 << 1  //+ Software Interrupt on line 1.
	SWI2  SWIER1_Bits = 0x01 << 2  //+ Software Interrupt on line 2.
	SWI3  SWIER1_Bits = 0x01 << 3  //+ Software Interrupt on line 3.
	SWI4  SWIER1_Bits = 0x01 << 4  //+ Software Interrupt on line 4.
	SWI5  SWIER1_Bits = 0x01 << 5  //+ Software Interrupt on line 5.
	SWI6  SWIER1_Bits = 0x01 << 6  //+ Software Interrupt on line 6.
	SWI7  SWIER1_Bits = 0x01 << 7  //+ Software Interrupt on line 7.
	SWI8  SWIER1_Bits = 0x01 << 8  //+ Software Interrupt on line 8.
	SWI9  SWIER1_Bits = 0x01 << 9  //+ Software Interrupt on line 9.
	SWI10 SWIER1_Bits = 0x01 << 10 //+ Software Interrupt on line 10.
	SWI11 SWIER1_Bits = 0x01 << 11 //+ Software Interrupt on line 11.
	SWI12 SWIER1_Bits = 0x01 << 12 //+ Software Interrupt on line 12.
	SWI13 SWIER1_Bits = 0x01 << 13 //+ Software Interrupt on line 13.
	SWI14 SWIER1_Bits = 0x01 << 14 //+ Software Interrupt on line 14.
	SWI15 SWIER1_Bits = 0x01 << 15 //+ Software Interrupt on line 15.
	SWI16 SWIER1_Bits = 0x01 << 16 //+ Software Interrupt on line 16.
	SWI18 SWIER1_Bits = 0x01 << 18 //+ Software Interrupt on line 18.
	SWI19 SWIER1_Bits = 0x01 << 19 //+ Software Interrupt on line 19.
	SWI20 SWIER1_Bits = 0x01 << 20 //+ Software Interrupt on line 20.
	SWI21 SWIER1_Bits = 0x01 << 21 //+ Software Interrupt on line 21.
	SWI22 SWIER1_Bits = 0x01 << 22 //+ Software Interrupt on line 22.
)

const (
	SWI0n  = 0
	SWI1n  = 1
	SWI2n  = 2
	SWI3n  = 3
	SWI4n  = 4
	SWI5n  = 5
	SWI6n  = 6
	SWI7n  = 7
	SWI8n  = 8
	SWI9n  = 9
	SWI10n = 10
	SWI11n = 11
	SWI12n = 12
	SWI13n = 13
	SWI14n = 14
	SWI15n = 15
	SWI16n = 16
	SWI18n = 18
	SWI19n = 19
	SWI20n = 20
	SWI21n = 21
	SWI22n = 22
)

const (
	PIF0  PR1_Bits = 0x01 << 0  //+ Pending bit for line 0.
	PIF1  PR1_Bits = 0x01 << 1  //+ Pending bit for line 1.
	PIF2  PR1_Bits = 0x01 << 2  //+ Pending bit for line 2.
	PIF3  PR1_Bits = 0x01 << 3  //+ Pending bit for line 3.
	PIF4  PR1_Bits = 0x01 << 4  //+ Pending bit for line 4.
	PIF5  PR1_Bits = 0x01 << 5  //+ Pending bit for line 5.
	PIF6  PR1_Bits = 0x01 << 6  //+ Pending bit for line 6.
	PIF7  PR1_Bits = 0x01 << 7  //+ Pending bit for line 7.
	PIF8  PR1_Bits = 0x01 << 8  //+ Pending bit for line 8.
	PIF9  PR1_Bits = 0x01 << 9  //+ Pending bit for line 9.
	PIF10 PR1_Bits = 0x01 << 10 //+ Pending bit for line 10.
	PIF11 PR1_Bits = 0x01 << 11 //+ Pending bit for line 11.
	PIF12 PR1_Bits = 0x01 << 12 //+ Pending bit for line 12.
	PIF13 PR1_Bits = 0x01 << 13 //+ Pending bit for line 13.
	PIF14 PR1_Bits = 0x01 << 14 //+ Pending bit for line 14.
	PIF15 PR1_Bits = 0x01 << 15 //+ Pending bit for line 15.
	PIF16 PR1_Bits = 0x01 << 16 //+ Pending bit for line 16.
	PIF18 PR1_Bits = 0x01 << 18 //+ Pending bit for line 18.
	PIF19 PR1_Bits = 0x01 << 19 //+ Pending bit for line 19.
	PIF20 PR1_Bits = 0x01 << 20 //+ Pending bit for line 20.
	PIF21 PR1_Bits = 0x01 << 21 //+ Pending bit for line 21.
	PIF22 PR1_Bits = 0x01 << 22 //+ Pending bit for line 22.
)

const (
	PIF0n  = 0
	PIF1n  = 1
	PIF2n  = 2
	PIF3n  = 3
	PIF4n  = 4
	PIF5n  = 5
	PIF6n  = 6
	PIF7n  = 7
	PIF8n  = 8
	PIF9n  = 9
	PIF10n = 10
	PIF11n = 11
	PIF12n = 12
	PIF13n = 13
	PIF14n = 14
	PIF15n = 15
	PIF16n = 16
	PIF18n = 18
	PIF19n = 19
	PIF20n = 20
	PIF21n = 21
	PIF22n = 22
)

const (
	IM32 IMR2_Bits = 0x01 << 0 //+ Interrupt Mask on line 32.
	IM33 IMR2_Bits = 0x01 << 1 //+ Interrupt Mask on line 33.
	IM34 IMR2_Bits = 0x01 << 2 //+ Interrupt Mask on line 34.
	IM35 IMR2_Bits = 0x01 << 3 //+ Interrupt Mask on line 35.
	IM36 IMR2_Bits = 0x01 << 4 //+ Interrupt Mask on line 36.
	IM37 IMR2_Bits = 0x01 << 5 //+ Interrupt Mask on line 37.
	IM38 IMR2_Bits = 0x01 << 6 //+ Interrupt Mask on line 38.
	IM39 IMR2_Bits = 0x01 << 7 //+ Interrupt Mask on line 39.
	IM   IMR2_Bits = 0xFF << 0 //  Interrupt Mask all.
)

const (
	IM32n = 0
	IM33n = 1
	IM34n = 2
	IM35n = 3
	IM36n = 4
	IM37n = 5
	IM38n = 6
	IM39n = 7
)

const (
	EM32 EMR2_Bits = 0x01 << 0 //+ Event Mask on line 32.
	EM33 EMR2_Bits = 0x01 << 1 //+ Event Mask on line 33.
	EM34 EMR2_Bits = 0x01 << 2 //+ Event Mask on line 34.
	EM35 EMR2_Bits = 0x01 << 3 //+ Event Mask on line 35.
	EM36 EMR2_Bits = 0x01 << 4 //+ Event Mask on line 36.
	EM37 EMR2_Bits = 0x01 << 5 //+ Event Mask on line 37.
	EM38 EMR2_Bits = 0x01 << 6 //+ Event Mask on line 38.
	EM39 EMR2_Bits = 0x01 << 7 //+ Event Mask on line 39.
)

const (
	EM32n = 0
	EM33n = 1
	EM34n = 2
	EM35n = 3
	EM36n = 4
	EM37n = 5
	EM38n = 6
	EM39n = 7
)

const (
	RT35 RTSR2_Bits = 0x01 << 3 //+ Rising trigger event configuration bit of line 35.
	RT36 RTSR2_Bits = 0x01 << 4 //+ Rising trigger event configuration bit of line 36.
	RT37 RTSR2_Bits = 0x01 << 5 //+ Rising trigger event configuration bit of line 37.
	RT38 RTSR2_Bits = 0x01 << 6 //+ Rising trigger event configuration bit of line 38.
)

const (
	RT35n = 3
	RT36n = 4
	RT37n = 5
	RT38n = 6
)

const (
	FT35 FTSR2_Bits = 0x01 << 3 //+ Falling trigger event configuration bit of line 35.
	FT36 FTSR2_Bits = 0x01 << 4 //+ Falling trigger event configuration bit of line 36.
	FT37 FTSR2_Bits = 0x01 << 5 //+ Falling trigger event configuration bit of line 37.
	FT38 FTSR2_Bits = 0x01 << 6 //+ Falling trigger event configuration bit of line 38.
)

const (
	FT35n = 3
	FT36n = 4
	FT37n = 5
	FT38n = 6
)

const (
	SWI35 SWIER2_Bits = 0x01 << 3 //+ Software Interrupt on line 35.
	SWI36 SWIER2_Bits = 0x01 << 4 //+ Software Interrupt on line 36.
	SWI37 SWIER2_Bits = 0x01 << 5 //+ Software Interrupt on line 37.
	SWI38 SWIER2_Bits = 0x01 << 6 //+ Software Interrupt on line 38.
)

const (
	SWI35n = 3
	SWI36n = 4
	SWI37n = 5
	SWI38n = 6
)

const (
	PIF35 PR2_Bits = 0x01 << 3 //+ Pending bit for line 35.
	PIF36 PR2_Bits = 0x01 << 4 //+ Pending bit for line 36.
	PIF37 PR2_Bits = 0x01 << 5 //+ Pending bit for line 37.
	PIF38 PR2_Bits = 0x01 << 6 //+ Pending bit for line 38.
)

const (
	PIF35n = 3
	PIF36n = 4
	PIF37n = 5
	PIF38n = 6
)
