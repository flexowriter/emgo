// +build l476xx

package quadspi

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type QUADSPI_Periph struct {
	CR    CR
	DCR   DCR
	SR    SR
	FCR   FCR
	DLR   DLR
	CCR   CCR
	AR    AR
	ABR   ABR
	DR    DR
	PSMKR PSMKR
	PSMAR PSMAR
	PIR   PIR
	LPTR  LPTR
}

func (p *QUADSPI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var QUADSPI = (*QUADSPI_Periph)(unsafe.Pointer(uintptr(mmap.QSPI_R_BASE)))

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) EN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(EN)}}
}

func (p *QUADSPI_Periph) ABORT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ABORT)}}
}

func (p *QUADSPI_Periph) DMAEN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(DMAEN)}}
}

func (p *QUADSPI_Periph) TCEN() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TCEN)}}
}

func (p *QUADSPI_Periph) SSHIFT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(SSHIFT)}}
}

func (p *QUADSPI_Periph) FTHRES() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(FTHRES)}}
}

func (p *QUADSPI_Periph) TEIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TEIE)}}
}

func (p *QUADSPI_Periph) TCIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TCIE)}}
}

func (p *QUADSPI_Periph) FTIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(FTIE)}}
}

func (p *QUADSPI_Periph) SMIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(SMIE)}}
}

func (p *QUADSPI_Periph) TOIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TOIE)}}
}

func (p *QUADSPI_Periph) APMS() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(APMS)}}
}

func (p *QUADSPI_Periph) PMM() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PMM)}}
}

func (p *QUADSPI_Periph) PRESCALER() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(PRESCALER)}}
}

type DCR_Bits uint32

func (b DCR_Bits) Field(mask DCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCR_Bits) J(v int) DCR_Bits {
	return DCR_Bits(bits.Make32(v, uint32(mask)))
}

type DCR struct{ mmio.U32 }

func (r *DCR) Bits(mask DCR_Bits) DCR_Bits { return DCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DCR) StoreBits(mask, b DCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DCR) SetBits(mask DCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *DCR) ClearBits(mask DCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *DCR) Load() DCR_Bits              { return DCR_Bits(r.U32.Load()) }
func (r *DCR) Store(b DCR_Bits)            { r.U32.Store(uint32(b)) }

type DCR_Mask struct{ mmio.UM32 }

func (rm DCR_Mask) Load() DCR_Bits   { return DCR_Bits(rm.UM32.Load()) }
func (rm DCR_Mask) Store(b DCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) CKMODE() DCR_Mask {
	return DCR_Mask{mmio.UM32{&p.DCR.U32, uint32(CKMODE)}}
}

func (p *QUADSPI_Periph) CSHT() DCR_Mask {
	return DCR_Mask{mmio.UM32{&p.DCR.U32, uint32(CSHT)}}
}

func (p *QUADSPI_Periph) FSIZE() DCR_Mask {
	return DCR_Mask{mmio.UM32{&p.DCR.U32, uint32(FSIZE)}}
}

type SR_Bits uint32

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) TEF() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(TEF)}}
}

func (p *QUADSPI_Periph) TCF() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(TCF)}}
}

func (p *QUADSPI_Periph) FTF() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(FTF)}}
}

func (p *QUADSPI_Periph) SMF() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(SMF)}}
}

func (p *QUADSPI_Periph) TOF() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(TOF)}}
}

func (p *QUADSPI_Periph) BUSY() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(BUSY)}}
}

func (p *QUADSPI_Periph) FLEVEL() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(FLEVEL)}}
}

type FCR_Bits uint32

func (b FCR_Bits) Field(mask FCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FCR_Bits) J(v int) FCR_Bits {
	return FCR_Bits(bits.Make32(v, uint32(mask)))
}

type FCR struct{ mmio.U32 }

func (r *FCR) Bits(mask FCR_Bits) FCR_Bits { return FCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FCR) StoreBits(mask, b FCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FCR) SetBits(mask FCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *FCR) ClearBits(mask FCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *FCR) Load() FCR_Bits              { return FCR_Bits(r.U32.Load()) }
func (r *FCR) Store(b FCR_Bits)            { r.U32.Store(uint32(b)) }

type FCR_Mask struct{ mmio.UM32 }

func (rm FCR_Mask) Load() FCR_Bits   { return FCR_Bits(rm.UM32.Load()) }
func (rm FCR_Mask) Store(b FCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) CTEF() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(CTEF)}}
}

func (p *QUADSPI_Periph) CTCF() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(CTCF)}}
}

func (p *QUADSPI_Periph) CSMF() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(CSMF)}}
}

func (p *QUADSPI_Periph) CTOF() FCR_Mask {
	return FCR_Mask{mmio.UM32{&p.FCR.U32, uint32(CTOF)}}
}

type DLR_Bits uint32

func (b DLR_Bits) Field(mask DLR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DLR_Bits) J(v int) DLR_Bits {
	return DLR_Bits(bits.Make32(v, uint32(mask)))
}

type DLR struct{ mmio.U32 }

func (r *DLR) Bits(mask DLR_Bits) DLR_Bits { return DLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DLR) StoreBits(mask, b DLR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DLR) SetBits(mask DLR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *DLR) ClearBits(mask DLR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *DLR) Load() DLR_Bits              { return DLR_Bits(r.U32.Load()) }
func (r *DLR) Store(b DLR_Bits)            { r.U32.Store(uint32(b)) }

type DLR_Mask struct{ mmio.UM32 }

func (rm DLR_Mask) Load() DLR_Bits   { return DLR_Bits(rm.UM32.Load()) }
func (rm DLR_Mask) Store(b DLR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) DL() DLR_Mask {
	return DLR_Mask{mmio.UM32{&p.DLR.U32, uint32(DL)}}
}

type CCR_Bits uint32

func (b CCR_Bits) Field(mask CCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR_Bits) J(v int) CCR_Bits {
	return CCR_Bits(bits.Make32(v, uint32(mask)))
}

type CCR struct{ mmio.U32 }

func (r *CCR) Bits(mask CCR_Bits) CCR_Bits { return CCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CCR) StoreBits(mask, b CCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CCR) SetBits(mask CCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CCR) ClearBits(mask CCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CCR) Load() CCR_Bits              { return CCR_Bits(r.U32.Load()) }
func (r *CCR) Store(b CCR_Bits)            { r.U32.Store(uint32(b)) }

type CCR_Mask struct{ mmio.UM32 }

func (rm CCR_Mask) Load() CCR_Bits   { return CCR_Bits(rm.UM32.Load()) }
func (rm CCR_Mask) Store(b CCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) INSTRUCTION() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(INSTRUCTION)}}
}

func (p *QUADSPI_Periph) IMODE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(IMODE)}}
}

func (p *QUADSPI_Periph) ADMODE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(ADMODE)}}
}

func (p *QUADSPI_Periph) ADSIZE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(ADSIZE)}}
}

func (p *QUADSPI_Periph) ABMODE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(ABMODE)}}
}

func (p *QUADSPI_Periph) ABSIZE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(ABSIZE)}}
}

func (p *QUADSPI_Periph) DCYC() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(DCYC)}}
}

func (p *QUADSPI_Periph) DMODE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(DMODE)}}
}

func (p *QUADSPI_Periph) FMODE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(FMODE)}}
}

func (p *QUADSPI_Periph) SIOO() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(SIOO)}}
}

func (p *QUADSPI_Periph) DDRM() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(DDRM)}}
}

type AR_Bits uint32

func (b AR_Bits) Field(mask AR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AR_Bits) J(v int) AR_Bits {
	return AR_Bits(bits.Make32(v, uint32(mask)))
}

type AR struct{ mmio.U32 }

func (r *AR) Bits(mask AR_Bits) AR_Bits { return AR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AR) StoreBits(mask, b AR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AR) SetBits(mask AR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *AR) ClearBits(mask AR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *AR) Load() AR_Bits             { return AR_Bits(r.U32.Load()) }
func (r *AR) Store(b AR_Bits)           { r.U32.Store(uint32(b)) }

type AR_Mask struct{ mmio.UM32 }

func (rm AR_Mask) Load() AR_Bits   { return AR_Bits(rm.UM32.Load()) }
func (rm AR_Mask) Store(b AR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) ADDRESS() AR_Mask {
	return AR_Mask{mmio.UM32{&p.AR.U32, uint32(ADDRESS)}}
}

type ABR_Bits uint32

func (b ABR_Bits) Field(mask ABR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ABR_Bits) J(v int) ABR_Bits {
	return ABR_Bits(bits.Make32(v, uint32(mask)))
}

type ABR struct{ mmio.U32 }

func (r *ABR) Bits(mask ABR_Bits) ABR_Bits { return ABR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ABR) StoreBits(mask, b ABR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ABR) SetBits(mask ABR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ABR) ClearBits(mask ABR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ABR) Load() ABR_Bits              { return ABR_Bits(r.U32.Load()) }
func (r *ABR) Store(b ABR_Bits)            { r.U32.Store(uint32(b)) }

type ABR_Mask struct{ mmio.UM32 }

func (rm ABR_Mask) Load() ABR_Bits   { return ABR_Bits(rm.UM32.Load()) }
func (rm ABR_Mask) Store(b ABR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) ALTERNATE() ABR_Mask {
	return ABR_Mask{mmio.UM32{&p.ABR.U32, uint32(ALTERNATE)}}
}

type DR_Bits uint32

func (b DR_Bits) Field(mask DR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR_Bits) J(v int) DR_Bits {
	return DR_Bits(bits.Make32(v, uint32(mask)))
}

type DR struct{ mmio.U32 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U32.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U32.Store(uint32(b)) }

type DR_Mask struct{ mmio.UM32 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM32.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) DATA() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(DATA)}}
}

type PSMKR_Bits uint32

func (b PSMKR_Bits) Field(mask PSMKR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PSMKR_Bits) J(v int) PSMKR_Bits {
	return PSMKR_Bits(bits.Make32(v, uint32(mask)))
}

type PSMKR struct{ mmio.U32 }

func (r *PSMKR) Bits(mask PSMKR_Bits) PSMKR_Bits { return PSMKR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PSMKR) StoreBits(mask, b PSMKR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PSMKR) SetBits(mask PSMKR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PSMKR) ClearBits(mask PSMKR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PSMKR) Load() PSMKR_Bits                { return PSMKR_Bits(r.U32.Load()) }
func (r *PSMKR) Store(b PSMKR_Bits)              { r.U32.Store(uint32(b)) }

type PSMKR_Mask struct{ mmio.UM32 }

func (rm PSMKR_Mask) Load() PSMKR_Bits   { return PSMKR_Bits(rm.UM32.Load()) }
func (rm PSMKR_Mask) Store(b PSMKR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) MASK() PSMKR_Mask {
	return PSMKR_Mask{mmio.UM32{&p.PSMKR.U32, uint32(MASK)}}
}

type PSMAR_Bits uint32

func (b PSMAR_Bits) Field(mask PSMAR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PSMAR_Bits) J(v int) PSMAR_Bits {
	return PSMAR_Bits(bits.Make32(v, uint32(mask)))
}

type PSMAR struct{ mmio.U32 }

func (r *PSMAR) Bits(mask PSMAR_Bits) PSMAR_Bits { return PSMAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PSMAR) StoreBits(mask, b PSMAR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PSMAR) SetBits(mask PSMAR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PSMAR) ClearBits(mask PSMAR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PSMAR) Load() PSMAR_Bits                { return PSMAR_Bits(r.U32.Load()) }
func (r *PSMAR) Store(b PSMAR_Bits)              { r.U32.Store(uint32(b)) }

type PSMAR_Mask struct{ mmio.UM32 }

func (rm PSMAR_Mask) Load() PSMAR_Bits   { return PSMAR_Bits(rm.UM32.Load()) }
func (rm PSMAR_Mask) Store(b PSMAR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) MATCH() PSMAR_Mask {
	return PSMAR_Mask{mmio.UM32{&p.PSMAR.U32, uint32(MATCH)}}
}

type PIR_Bits uint32

func (b PIR_Bits) Field(mask PIR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PIR_Bits) J(v int) PIR_Bits {
	return PIR_Bits(bits.Make32(v, uint32(mask)))
}

type PIR struct{ mmio.U32 }

func (r *PIR) Bits(mask PIR_Bits) PIR_Bits { return PIR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PIR) StoreBits(mask, b PIR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PIR) SetBits(mask PIR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *PIR) ClearBits(mask PIR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *PIR) Load() PIR_Bits              { return PIR_Bits(r.U32.Load()) }
func (r *PIR) Store(b PIR_Bits)            { r.U32.Store(uint32(b)) }

type PIR_Mask struct{ mmio.UM32 }

func (rm PIR_Mask) Load() PIR_Bits   { return PIR_Bits(rm.UM32.Load()) }
func (rm PIR_Mask) Store(b PIR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) INTERVAL() PIR_Mask {
	return PIR_Mask{mmio.UM32{&p.PIR.U32, uint32(INTERVAL)}}
}

type LPTR_Bits uint32

func (b LPTR_Bits) Field(mask LPTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LPTR_Bits) J(v int) LPTR_Bits {
	return LPTR_Bits(bits.Make32(v, uint32(mask)))
}

type LPTR struct{ mmio.U32 }

func (r *LPTR) Bits(mask LPTR_Bits) LPTR_Bits { return LPTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LPTR) StoreBits(mask, b LPTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LPTR) SetBits(mask LPTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *LPTR) ClearBits(mask LPTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *LPTR) Load() LPTR_Bits               { return LPTR_Bits(r.U32.Load()) }
func (r *LPTR) Store(b LPTR_Bits)             { r.U32.Store(uint32(b)) }

type LPTR_Mask struct{ mmio.UM32 }

func (rm LPTR_Mask) Load() LPTR_Bits   { return LPTR_Bits(rm.UM32.Load()) }
func (rm LPTR_Mask) Store(b LPTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *QUADSPI_Periph) TIMEOUT() LPTR_Mask {
	return LPTR_Mask{mmio.UM32{&p.LPTR.U32, uint32(TIMEOUT)}}
}
