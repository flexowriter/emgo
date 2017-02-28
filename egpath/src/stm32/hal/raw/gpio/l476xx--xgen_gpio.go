// +build l476xx

package gpio

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type GPIO_Periph struct {
	MODER   MODER
	OTYPER  OTYPER
	OSPEEDR OSPEEDR
	PUPDR   PUPDR
	IDR     IDR
	ODR     ODR
	BSRR    BSRR
	LCKR    LCKR
	AFR     [2]AFR
	BRR     BRR
	ASCR    ASCR
}

func (p *GPIO_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var GPIOA = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOA_BASE)))

var GPIOB = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOB_BASE)))

var GPIOC = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOC_BASE)))

var GPIOD = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOD_BASE)))

var GPIOE = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOE_BASE)))

var GPIOF = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOF_BASE)))

var GPIOG = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOG_BASE)))

var GPIOH = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOH_BASE)))

type MODER_Bits uint32

func (b MODER_Bits) Field(mask MODER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MODER_Bits) J(v int) MODER_Bits {
	return MODER_Bits(bits.Make32(v, uint32(mask)))
}

type MODER struct{ mmio.U32 }

func (r *MODER) Bits(mask MODER_Bits) MODER_Bits { return MODER_Bits(r.U32.Bits(uint32(mask))) }
func (r *MODER) StoreBits(mask, b MODER_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MODER) SetBits(mask MODER_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *MODER) ClearBits(mask MODER_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *MODER) Load() MODER_Bits                { return MODER_Bits(r.U32.Load()) }
func (r *MODER) Store(b MODER_Bits)              { r.U32.Store(uint32(b)) }

type MODER_Mask struct{ mmio.UM32 }

func (rm MODER_Mask) Load() MODER_Bits   { return MODER_Bits(rm.UM32.Load()) }
func (rm MODER_Mask) Store(b MODER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) MODE0() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE0)}}
}

func (p *GPIO_Periph) MODE1() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE1)}}
}

func (p *GPIO_Periph) MODE2() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE2)}}
}

func (p *GPIO_Periph) MODE3() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE3)}}
}

func (p *GPIO_Periph) MODE4() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE4)}}
}

func (p *GPIO_Periph) MODE5() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE5)}}
}

func (p *GPIO_Periph) MODE6() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE6)}}
}

func (p *GPIO_Periph) MODE7() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE7)}}
}

func (p *GPIO_Periph) MODE8() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE8)}}
}

func (p *GPIO_Periph) MODE9() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE9)}}
}

func (p *GPIO_Periph) MODE10() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE10)}}
}

func (p *GPIO_Periph) MODE11() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE11)}}
}

func (p *GPIO_Periph) MODE12() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE12)}}
}

func (p *GPIO_Periph) MODE13() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE13)}}
}

func (p *GPIO_Periph) MODE14() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE14)}}
}

func (p *GPIO_Periph) MODE15() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODE15)}}
}

type OTYPER_Bits uint32

func (b OTYPER_Bits) Field(mask OTYPER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OTYPER_Bits) J(v int) OTYPER_Bits {
	return OTYPER_Bits(bits.Make32(v, uint32(mask)))
}

type OTYPER struct{ mmio.U32 }

func (r *OTYPER) Bits(mask OTYPER_Bits) OTYPER_Bits { return OTYPER_Bits(r.U32.Bits(uint32(mask))) }
func (r *OTYPER) StoreBits(mask, b OTYPER_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OTYPER) SetBits(mask OTYPER_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *OTYPER) ClearBits(mask OTYPER_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *OTYPER) Load() OTYPER_Bits                 { return OTYPER_Bits(r.U32.Load()) }
func (r *OTYPER) Store(b OTYPER_Bits)               { r.U32.Store(uint32(b)) }

type OTYPER_Mask struct{ mmio.UM32 }

func (rm OTYPER_Mask) Load() OTYPER_Bits   { return OTYPER_Bits(rm.UM32.Load()) }
func (rm OTYPER_Mask) Store(b OTYPER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OT0() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT0)}}
}

func (p *GPIO_Periph) OT1() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT1)}}
}

func (p *GPIO_Periph) OT2() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT2)}}
}

func (p *GPIO_Periph) OT3() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT3)}}
}

func (p *GPIO_Periph) OT4() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT4)}}
}

func (p *GPIO_Periph) OT5() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT5)}}
}

func (p *GPIO_Periph) OT6() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT6)}}
}

func (p *GPIO_Periph) OT7() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT7)}}
}

func (p *GPIO_Periph) OT8() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT8)}}
}

func (p *GPIO_Periph) OT9() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT9)}}
}

func (p *GPIO_Periph) OT10() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT10)}}
}

func (p *GPIO_Periph) OT11() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT11)}}
}

func (p *GPIO_Periph) OT12() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT12)}}
}

func (p *GPIO_Periph) OT13() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT13)}}
}

func (p *GPIO_Periph) OT14() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT14)}}
}

func (p *GPIO_Periph) OT15() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT15)}}
}

type OSPEEDR_Bits uint32

func (b OSPEEDR_Bits) Field(mask OSPEEDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OSPEEDR_Bits) J(v int) OSPEEDR_Bits {
	return OSPEEDR_Bits(bits.Make32(v, uint32(mask)))
}

type OSPEEDR struct{ mmio.U32 }

func (r *OSPEEDR) Bits(mask OSPEEDR_Bits) OSPEEDR_Bits { return OSPEEDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OSPEEDR) StoreBits(mask, b OSPEEDR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OSPEEDR) SetBits(mask OSPEEDR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *OSPEEDR) ClearBits(mask OSPEEDR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *OSPEEDR) Load() OSPEEDR_Bits                  { return OSPEEDR_Bits(r.U32.Load()) }
func (r *OSPEEDR) Store(b OSPEEDR_Bits)                { r.U32.Store(uint32(b)) }

type OSPEEDR_Mask struct{ mmio.UM32 }

func (rm OSPEEDR_Mask) Load() OSPEEDR_Bits   { return OSPEEDR_Bits(rm.UM32.Load()) }
func (rm OSPEEDR_Mask) Store(b OSPEEDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OSPEED0() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED0)}}
}

func (p *GPIO_Periph) OSPEED1() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED1)}}
}

func (p *GPIO_Periph) OSPEED2() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED2)}}
}

func (p *GPIO_Periph) OSPEED3() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED3)}}
}

func (p *GPIO_Periph) OSPEED4() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED4)}}
}

func (p *GPIO_Periph) OSPEED5() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED5)}}
}

func (p *GPIO_Periph) OSPEED6() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED6)}}
}

func (p *GPIO_Periph) OSPEED7() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED7)}}
}

func (p *GPIO_Periph) OSPEED8() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED8)}}
}

func (p *GPIO_Periph) OSPEED9() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED9)}}
}

func (p *GPIO_Periph) OSPEED10() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED10)}}
}

func (p *GPIO_Periph) OSPEED11() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED11)}}
}

func (p *GPIO_Periph) OSPEED12() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED12)}}
}

func (p *GPIO_Periph) OSPEED13() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED13)}}
}

func (p *GPIO_Periph) OSPEED14() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED14)}}
}

func (p *GPIO_Periph) OSPEED15() OSPEEDR_Mask {
	return OSPEEDR_Mask{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEED15)}}
}

type PUPDR_Bits uint32

func (b PUPDR_Bits) Field(mask PUPDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PUPDR_Bits) J(v int) PUPDR_Bits {
	return PUPDR_Bits(bits.Make32(v, uint32(mask)))
}

type PUPDR struct{ mmio.U32 }

func (r *PUPDR) Bits(mask PUPDR_Bits) PUPDR_Bits { return PUPDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PUPDR) StoreBits(mask, b PUPDR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PUPDR) SetBits(mask PUPDR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PUPDR) ClearBits(mask PUPDR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PUPDR) Load() PUPDR_Bits                { return PUPDR_Bits(r.U32.Load()) }
func (r *PUPDR) Store(b PUPDR_Bits)              { r.U32.Store(uint32(b)) }

type PUPDR_Mask struct{ mmio.UM32 }

func (rm PUPDR_Mask) Load() PUPDR_Bits   { return PUPDR_Bits(rm.UM32.Load()) }
func (rm PUPDR_Mask) Store(b PUPDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) PUPD0() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD0)}}
}

func (p *GPIO_Periph) PUPD1() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD1)}}
}

func (p *GPIO_Periph) PUPD2() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD2)}}
}

func (p *GPIO_Periph) PUPD3() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD3)}}
}

func (p *GPIO_Periph) PUPD4() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD4)}}
}

func (p *GPIO_Periph) PUPD5() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD5)}}
}

func (p *GPIO_Periph) PUPD6() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD6)}}
}

func (p *GPIO_Periph) PUPD7() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD7)}}
}

func (p *GPIO_Periph) PUPD8() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD8)}}
}

func (p *GPIO_Periph) PUPD9() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD9)}}
}

func (p *GPIO_Periph) PUPD10() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD10)}}
}

func (p *GPIO_Periph) PUPD11() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD11)}}
}

func (p *GPIO_Periph) PUPD12() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD12)}}
}

func (p *GPIO_Periph) PUPD13() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD13)}}
}

func (p *GPIO_Periph) PUPD14() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD14)}}
}

func (p *GPIO_Periph) PUPD15() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPD15)}}
}

type IDR_Bits uint32

func (b IDR_Bits) Field(mask IDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDR_Bits) J(v int) IDR_Bits {
	return IDR_Bits(bits.Make32(v, uint32(mask)))
}

type IDR struct{ mmio.U32 }

func (r *IDR) Bits(mask IDR_Bits) IDR_Bits { return IDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IDR) StoreBits(mask, b IDR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IDR) SetBits(mask IDR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IDR) ClearBits(mask IDR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IDR) Load() IDR_Bits              { return IDR_Bits(r.U32.Load()) }
func (r *IDR) Store(b IDR_Bits)            { r.U32.Store(uint32(b)) }

type IDR_Mask struct{ mmio.UM32 }

func (rm IDR_Mask) Load() IDR_Bits   { return IDR_Bits(rm.UM32.Load()) }
func (rm IDR_Mask) Store(b IDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) ID0() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID0)}}
}

func (p *GPIO_Periph) ID1() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID1)}}
}

func (p *GPIO_Periph) ID2() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID2)}}
}

func (p *GPIO_Periph) ID3() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID3)}}
}

func (p *GPIO_Periph) ID4() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID4)}}
}

func (p *GPIO_Periph) ID5() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID5)}}
}

func (p *GPIO_Periph) ID6() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID6)}}
}

func (p *GPIO_Periph) ID7() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID7)}}
}

func (p *GPIO_Periph) ID8() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID8)}}
}

func (p *GPIO_Periph) ID9() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID9)}}
}

func (p *GPIO_Periph) ID10() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID10)}}
}

func (p *GPIO_Periph) ID11() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID11)}}
}

func (p *GPIO_Periph) ID12() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID12)}}
}

func (p *GPIO_Periph) ID13() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID13)}}
}

func (p *GPIO_Periph) ID14() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID14)}}
}

func (p *GPIO_Periph) ID15() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(ID15)}}
}

type ODR_Bits uint32

func (b ODR_Bits) Field(mask ODR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ODR_Bits) J(v int) ODR_Bits {
	return ODR_Bits(bits.Make32(v, uint32(mask)))
}

type ODR struct{ mmio.U32 }

func (r *ODR) Bits(mask ODR_Bits) ODR_Bits { return ODR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ODR) StoreBits(mask, b ODR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ODR) SetBits(mask ODR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ODR) ClearBits(mask ODR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ODR) Load() ODR_Bits              { return ODR_Bits(r.U32.Load()) }
func (r *ODR) Store(b ODR_Bits)            { r.U32.Store(uint32(b)) }

type ODR_Mask struct{ mmio.UM32 }

func (rm ODR_Mask) Load() ODR_Bits   { return ODR_Bits(rm.UM32.Load()) }
func (rm ODR_Mask) Store(b ODR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OD0() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD0)}}
}

func (p *GPIO_Periph) OD1() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD1)}}
}

func (p *GPIO_Periph) OD2() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD2)}}
}

func (p *GPIO_Periph) OD3() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD3)}}
}

func (p *GPIO_Periph) OD4() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD4)}}
}

func (p *GPIO_Periph) OD5() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD5)}}
}

func (p *GPIO_Periph) OD6() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD6)}}
}

func (p *GPIO_Periph) OD7() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD7)}}
}

func (p *GPIO_Periph) OD8() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD8)}}
}

func (p *GPIO_Periph) OD9() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD9)}}
}

func (p *GPIO_Periph) OD10() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD10)}}
}

func (p *GPIO_Periph) OD11() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD11)}}
}

func (p *GPIO_Periph) OD12() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD12)}}
}

func (p *GPIO_Periph) OD13() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD13)}}
}

func (p *GPIO_Periph) OD14() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD14)}}
}

func (p *GPIO_Periph) OD15() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(OD15)}}
}

type BSRR_Bits uint32

func (b BSRR_Bits) Field(mask BSRR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BSRR_Bits) J(v int) BSRR_Bits {
	return BSRR_Bits(bits.Make32(v, uint32(mask)))
}

type BSRR struct{ mmio.U32 }

func (r *BSRR) Bits(mask BSRR_Bits) BSRR_Bits { return BSRR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BSRR) StoreBits(mask, b BSRR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BSRR) SetBits(mask BSRR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BSRR) ClearBits(mask BSRR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BSRR) Load() BSRR_Bits               { return BSRR_Bits(r.U32.Load()) }
func (r *BSRR) Store(b BSRR_Bits)             { r.U32.Store(uint32(b)) }

type BSRR_Mask struct{ mmio.UM32 }

func (rm BSRR_Mask) Load() BSRR_Bits   { return BSRR_Bits(rm.UM32.Load()) }
func (rm BSRR_Mask) Store(b BSRR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) BS0() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS0)}}
}

func (p *GPIO_Periph) BS1() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS1)}}
}

func (p *GPIO_Periph) BS2() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS2)}}
}

func (p *GPIO_Periph) BS3() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS3)}}
}

func (p *GPIO_Periph) BS4() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS4)}}
}

func (p *GPIO_Periph) BS5() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS5)}}
}

func (p *GPIO_Periph) BS6() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS6)}}
}

func (p *GPIO_Periph) BS7() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS7)}}
}

func (p *GPIO_Periph) BS8() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS8)}}
}

func (p *GPIO_Periph) BS9() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS9)}}
}

func (p *GPIO_Periph) BS10() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS10)}}
}

func (p *GPIO_Periph) BS11() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS11)}}
}

func (p *GPIO_Periph) BS12() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS12)}}
}

func (p *GPIO_Periph) BS13() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS13)}}
}

func (p *GPIO_Periph) BS14() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS14)}}
}

func (p *GPIO_Periph) BS15() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS15)}}
}

func (p *GPIO_Periph) BR0() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR0)}}
}

func (p *GPIO_Periph) BR1() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR1)}}
}

func (p *GPIO_Periph) BR2() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR2)}}
}

func (p *GPIO_Periph) BR3() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR3)}}
}

func (p *GPIO_Periph) BR4() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR4)}}
}

func (p *GPIO_Periph) BR5() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR5)}}
}

func (p *GPIO_Periph) BR6() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR6)}}
}

func (p *GPIO_Periph) BR7() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR7)}}
}

func (p *GPIO_Periph) BR8() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR8)}}
}

func (p *GPIO_Periph) BR9() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR9)}}
}

func (p *GPIO_Periph) BR10() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR10)}}
}

func (p *GPIO_Periph) BR11() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR11)}}
}

func (p *GPIO_Periph) BR12() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR12)}}
}

func (p *GPIO_Periph) BR13() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR13)}}
}

func (p *GPIO_Periph) BR14() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR14)}}
}

func (p *GPIO_Periph) BR15() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR15)}}
}

type LCKR_Bits uint32

func (b LCKR_Bits) Field(mask LCKR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LCKR_Bits) J(v int) LCKR_Bits {
	return LCKR_Bits(bits.Make32(v, uint32(mask)))
}

type LCKR struct{ mmio.U32 }

func (r *LCKR) Bits(mask LCKR_Bits) LCKR_Bits { return LCKR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LCKR) StoreBits(mask, b LCKR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LCKR) SetBits(mask LCKR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *LCKR) ClearBits(mask LCKR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *LCKR) Load() LCKR_Bits               { return LCKR_Bits(r.U32.Load()) }
func (r *LCKR) Store(b LCKR_Bits)             { r.U32.Store(uint32(b)) }

type LCKR_Mask struct{ mmio.UM32 }

func (rm LCKR_Mask) Load() LCKR_Bits   { return LCKR_Bits(rm.UM32.Load()) }
func (rm LCKR_Mask) Store(b LCKR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) LCK0() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK0)}}
}

func (p *GPIO_Periph) LCK1() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK1)}}
}

func (p *GPIO_Periph) LCK2() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK2)}}
}

func (p *GPIO_Periph) LCK3() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK3)}}
}

func (p *GPIO_Periph) LCK4() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK4)}}
}

func (p *GPIO_Periph) LCK5() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK5)}}
}

func (p *GPIO_Periph) LCK6() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK6)}}
}

func (p *GPIO_Periph) LCK7() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK7)}}
}

func (p *GPIO_Periph) LCK8() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK8)}}
}

func (p *GPIO_Periph) LCK9() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK9)}}
}

func (p *GPIO_Periph) LCK10() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK10)}}
}

func (p *GPIO_Periph) LCK11() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK11)}}
}

func (p *GPIO_Periph) LCK12() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK12)}}
}

func (p *GPIO_Periph) LCK13() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK13)}}
}

func (p *GPIO_Periph) LCK14() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK14)}}
}

func (p *GPIO_Periph) LCK15() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK15)}}
}

func (p *GPIO_Periph) LCKK() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCKK)}}
}

type AFR_Bits uint32

func (b AFR_Bits) Field(mask AFR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AFR_Bits) J(v int) AFR_Bits {
	return AFR_Bits(bits.Make32(v, uint32(mask)))
}

type AFR struct{ mmio.U32 }

func (r *AFR) Bits(mask AFR_Bits) AFR_Bits { return AFR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AFR) StoreBits(mask, b AFR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AFR) SetBits(mask AFR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *AFR) ClearBits(mask AFR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *AFR) Load() AFR_Bits              { return AFR_Bits(r.U32.Load()) }
func (r *AFR) Store(b AFR_Bits)            { r.U32.Store(uint32(b)) }

type AFR_Mask struct{ mmio.UM32 }

func (rm AFR_Mask) Load() AFR_Bits   { return AFR_Bits(rm.UM32.Load()) }
func (rm AFR_Mask) Store(b AFR_Bits) { rm.UM32.Store(uint32(b)) }

type BRR_Bits uint32

func (b BRR_Bits) Field(mask BRR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BRR_Bits) J(v int) BRR_Bits {
	return BRR_Bits(bits.Make32(v, uint32(mask)))
}

type BRR struct{ mmio.U32 }

func (r *BRR) Bits(mask BRR_Bits) BRR_Bits { return BRR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BRR) StoreBits(mask, b BRR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BRR) SetBits(mask BRR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *BRR) ClearBits(mask BRR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *BRR) Load() BRR_Bits              { return BRR_Bits(r.U32.Load()) }
func (r *BRR) Store(b BRR_Bits)            { r.U32.Store(uint32(b)) }

type BRR_Mask struct{ mmio.UM32 }

func (rm BRR_Mask) Load() BRR_Bits   { return BRR_Bits(rm.UM32.Load()) }
func (rm BRR_Mask) Store(b BRR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) BR0() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR0)}}
}

func (p *GPIO_Periph) BR1() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR1)}}
}

func (p *GPIO_Periph) BR2() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR2)}}
}

func (p *GPIO_Periph) BR3() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR3)}}
}

func (p *GPIO_Periph) BR4() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR4)}}
}

func (p *GPIO_Periph) BR5() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR5)}}
}

func (p *GPIO_Periph) BR6() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR6)}}
}

func (p *GPIO_Periph) BR7() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR7)}}
}

func (p *GPIO_Periph) BR8() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR8)}}
}

func (p *GPIO_Periph) BR9() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR9)}}
}

func (p *GPIO_Periph) BR10() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR10)}}
}

func (p *GPIO_Periph) BR11() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR11)}}
}

func (p *GPIO_Periph) BR12() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR12)}}
}

func (p *GPIO_Periph) BR13() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR13)}}
}

func (p *GPIO_Periph) BR14() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR14)}}
}

func (p *GPIO_Periph) BR15() BRR_Mask {
	return BRR_Mask{mmio.UM32{&p.BRR.U32, uint32(BR15)}}
}

type ASCR_Bits uint32

func (b ASCR_Bits) Field(mask ASCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ASCR_Bits) J(v int) ASCR_Bits {
	return ASCR_Bits(bits.Make32(v, uint32(mask)))
}

type ASCR struct{ mmio.U32 }

func (r *ASCR) Bits(mask ASCR_Bits) ASCR_Bits { return ASCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ASCR) StoreBits(mask, b ASCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ASCR) SetBits(mask ASCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *ASCR) ClearBits(mask ASCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *ASCR) Load() ASCR_Bits               { return ASCR_Bits(r.U32.Load()) }
func (r *ASCR) Store(b ASCR_Bits)             { r.U32.Store(uint32(b)) }

type ASCR_Mask struct{ mmio.UM32 }

func (rm ASCR_Mask) Load() ASCR_Bits   { return ASCR_Bits(rm.UM32.Load()) }
func (rm ASCR_Mask) Store(b ASCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) ASC0() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC0)}}
}

func (p *GPIO_Periph) ASC1() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC1)}}
}

func (p *GPIO_Periph) ASC2() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC2)}}
}

func (p *GPIO_Periph) ASC3() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC3)}}
}

func (p *GPIO_Periph) ASC4() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC4)}}
}

func (p *GPIO_Periph) ASC5() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC5)}}
}

func (p *GPIO_Periph) ASC6() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC6)}}
}

func (p *GPIO_Periph) ASC7() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC7)}}
}

func (p *GPIO_Periph) ASC8() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC8)}}
}

func (p *GPIO_Periph) ASC9() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC9)}}
}

func (p *GPIO_Periph) ASC10() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC10)}}
}

func (p *GPIO_Periph) ASC11() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC11)}}
}

func (p *GPIO_Periph) ASC12() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC12)}}
}

func (p *GPIO_Periph) ASC13() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC13)}}
}

func (p *GPIO_Periph) ASC14() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC14)}}
}

func (p *GPIO_Periph) ASC15() ASCR_Mask {
	return ASCR_Mask{mmio.UM32{&p.ASCR.U32, uint32(ASC15)}}
}
