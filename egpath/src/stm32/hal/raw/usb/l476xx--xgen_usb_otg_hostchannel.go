// +build l476xx

package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type USB_OTG_HostChannel_Periph struct {
	HCCHAR   HCCHAR
	HCSPLT   HCSPLT
	HCINT    HCINT
	HCINTMSK HCINTMSK
	HCTSIZ   HCTSIZ
	HCDMA    HCDMA
}

func (p *USB_OTG_HostChannel_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type HCCHAR_Bits uint32

func (b HCCHAR_Bits) Field(mask HCCHAR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCCHAR_Bits) J(v int) HCCHAR_Bits {
	return HCCHAR_Bits(bits.Make32(v, uint32(mask)))
}

type HCCHAR struct{ mmio.U32 }

func (r *HCCHAR) Bits(mask HCCHAR_Bits) HCCHAR_Bits { return HCCHAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *HCCHAR) StoreBits(mask, b HCCHAR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HCCHAR) SetBits(mask HCCHAR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *HCCHAR) ClearBits(mask HCCHAR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *HCCHAR) Load() HCCHAR_Bits                 { return HCCHAR_Bits(r.U32.Load()) }
func (r *HCCHAR) Store(b HCCHAR_Bits)               { r.U32.Store(uint32(b)) }

type HCCHAR_Mask struct{ mmio.UM32 }

func (rm HCCHAR_Mask) Load() HCCHAR_Bits   { return HCCHAR_Bits(rm.UM32.Load()) }
func (rm HCCHAR_Mask) Store(b HCCHAR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) MPSIZ() HCCHAR_Mask {
	return HCCHAR_Mask{mmio.UM32{&p.HCCHAR.U32, uint32(MPSIZ)}}
}

func (p *USB_OTG_HostChannel_Periph) EPNUM() HCCHAR_Mask {
	return HCCHAR_Mask{mmio.UM32{&p.HCCHAR.U32, uint32(EPNUM)}}
}

func (p *USB_OTG_HostChannel_Periph) EPDIR() HCCHAR_Mask {
	return HCCHAR_Mask{mmio.UM32{&p.HCCHAR.U32, uint32(EPDIR)}}
}

func (p *USB_OTG_HostChannel_Periph) LSDEV() HCCHAR_Mask {
	return HCCHAR_Mask{mmio.UM32{&p.HCCHAR.U32, uint32(LSDEV)}}
}

func (p *USB_OTG_HostChannel_Periph) EPTYP() HCCHAR_Mask {
	return HCCHAR_Mask{mmio.UM32{&p.HCCHAR.U32, uint32(EPTYP)}}
}

func (p *USB_OTG_HostChannel_Periph) MC() HCCHAR_Mask {
	return HCCHAR_Mask{mmio.UM32{&p.HCCHAR.U32, uint32(MC)}}
}

func (p *USB_OTG_HostChannel_Periph) DAD() HCCHAR_Mask {
	return HCCHAR_Mask{mmio.UM32{&p.HCCHAR.U32, uint32(DAD)}}
}

func (p *USB_OTG_HostChannel_Periph) ODDFRM() HCCHAR_Mask {
	return HCCHAR_Mask{mmio.UM32{&p.HCCHAR.U32, uint32(ODDFRM)}}
}

func (p *USB_OTG_HostChannel_Periph) CHDIS() HCCHAR_Mask {
	return HCCHAR_Mask{mmio.UM32{&p.HCCHAR.U32, uint32(CHDIS)}}
}

func (p *USB_OTG_HostChannel_Periph) CHENA() HCCHAR_Mask {
	return HCCHAR_Mask{mmio.UM32{&p.HCCHAR.U32, uint32(CHENA)}}
}

type HCSPLT_Bits uint32

func (b HCSPLT_Bits) Field(mask HCSPLT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCSPLT_Bits) J(v int) HCSPLT_Bits {
	return HCSPLT_Bits(bits.Make32(v, uint32(mask)))
}

type HCSPLT struct{ mmio.U32 }

func (r *HCSPLT) Bits(mask HCSPLT_Bits) HCSPLT_Bits { return HCSPLT_Bits(r.U32.Bits(uint32(mask))) }
func (r *HCSPLT) StoreBits(mask, b HCSPLT_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HCSPLT) SetBits(mask HCSPLT_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *HCSPLT) ClearBits(mask HCSPLT_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *HCSPLT) Load() HCSPLT_Bits                 { return HCSPLT_Bits(r.U32.Load()) }
func (r *HCSPLT) Store(b HCSPLT_Bits)               { r.U32.Store(uint32(b)) }

type HCSPLT_Mask struct{ mmio.UM32 }

func (rm HCSPLT_Mask) Load() HCSPLT_Bits   { return HCSPLT_Bits(rm.UM32.Load()) }
func (rm HCSPLT_Mask) Store(b HCSPLT_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) PRTADDR() HCSPLT_Mask {
	return HCSPLT_Mask{mmio.UM32{&p.HCSPLT.U32, uint32(PRTADDR)}}
}

func (p *USB_OTG_HostChannel_Periph) HUBADDR() HCSPLT_Mask {
	return HCSPLT_Mask{mmio.UM32{&p.HCSPLT.U32, uint32(HUBADDR)}}
}

func (p *USB_OTG_HostChannel_Periph) XACTPOS() HCSPLT_Mask {
	return HCSPLT_Mask{mmio.UM32{&p.HCSPLT.U32, uint32(XACTPOS)}}
}

func (p *USB_OTG_HostChannel_Periph) COMPLSPLT() HCSPLT_Mask {
	return HCSPLT_Mask{mmio.UM32{&p.HCSPLT.U32, uint32(COMPLSPLT)}}
}

func (p *USB_OTG_HostChannel_Periph) SPLITEN() HCSPLT_Mask {
	return HCSPLT_Mask{mmio.UM32{&p.HCSPLT.U32, uint32(SPLITEN)}}
}

type HCINT_Bits uint32

func (b HCINT_Bits) Field(mask HCINT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCINT_Bits) J(v int) HCINT_Bits {
	return HCINT_Bits(bits.Make32(v, uint32(mask)))
}

type HCINT struct{ mmio.U32 }

func (r *HCINT) Bits(mask HCINT_Bits) HCINT_Bits { return HCINT_Bits(r.U32.Bits(uint32(mask))) }
func (r *HCINT) StoreBits(mask, b HCINT_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HCINT) SetBits(mask HCINT_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *HCINT) ClearBits(mask HCINT_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *HCINT) Load() HCINT_Bits                { return HCINT_Bits(r.U32.Load()) }
func (r *HCINT) Store(b HCINT_Bits)              { r.U32.Store(uint32(b)) }

type HCINT_Mask struct{ mmio.UM32 }

func (rm HCINT_Mask) Load() HCINT_Bits   { return HCINT_Bits(rm.UM32.Load()) }
func (rm HCINT_Mask) Store(b HCINT_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) XFRC() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(XFRC)}}
}

func (p *USB_OTG_HostChannel_Periph) CHH() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(CHH)}}
}

func (p *USB_OTG_HostChannel_Periph) AHBERR() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(AHBERR)}}
}

func (p *USB_OTG_HostChannel_Periph) STALL() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(STALL)}}
}

func (p *USB_OTG_HostChannel_Periph) NAK() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(NAK)}}
}

func (p *USB_OTG_HostChannel_Periph) ACK() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(ACK)}}
}

func (p *USB_OTG_HostChannel_Periph) NYET() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(NYET)}}
}

func (p *USB_OTG_HostChannel_Periph) TXERR() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(TXERR)}}
}

func (p *USB_OTG_HostChannel_Periph) BBERR() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(BBERR)}}
}

func (p *USB_OTG_HostChannel_Periph) FRMOR() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(FRMOR)}}
}

func (p *USB_OTG_HostChannel_Periph) DTERR() HCINT_Mask {
	return HCINT_Mask{mmio.UM32{&p.HCINT.U32, uint32(DTERR)}}
}

type HCINTMSK_Bits uint32

func (b HCINTMSK_Bits) Field(mask HCINTMSK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCINTMSK_Bits) J(v int) HCINTMSK_Bits {
	return HCINTMSK_Bits(bits.Make32(v, uint32(mask)))
}

type HCINTMSK struct{ mmio.U32 }

func (r *HCINTMSK) Bits(mask HCINTMSK_Bits) HCINTMSK_Bits {
	return HCINTMSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *HCINTMSK) StoreBits(mask, b HCINTMSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HCINTMSK) SetBits(mask HCINTMSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *HCINTMSK) ClearBits(mask HCINTMSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *HCINTMSK) Load() HCINTMSK_Bits             { return HCINTMSK_Bits(r.U32.Load()) }
func (r *HCINTMSK) Store(b HCINTMSK_Bits)           { r.U32.Store(uint32(b)) }

type HCINTMSK_Mask struct{ mmio.UM32 }

func (rm HCINTMSK_Mask) Load() HCINTMSK_Bits   { return HCINTMSK_Bits(rm.UM32.Load()) }
func (rm HCINTMSK_Mask) Store(b HCINTMSK_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) XFRCM() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(XFRCM)}}
}

func (p *USB_OTG_HostChannel_Periph) CHHM() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(CHHM)}}
}

func (p *USB_OTG_HostChannel_Periph) AHBERR() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(AHBERR)}}
}

func (p *USB_OTG_HostChannel_Periph) STALLM() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(STALLM)}}
}

func (p *USB_OTG_HostChannel_Periph) NAKM() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(NAKM)}}
}

func (p *USB_OTG_HostChannel_Periph) ACKM() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(ACKM)}}
}

func (p *USB_OTG_HostChannel_Periph) NYET() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(NYET)}}
}

func (p *USB_OTG_HostChannel_Periph) TXERRM() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(TXERRM)}}
}

func (p *USB_OTG_HostChannel_Periph) BBERRM() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(BBERRM)}}
}

func (p *USB_OTG_HostChannel_Periph) FRMORM() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(FRMORM)}}
}

func (p *USB_OTG_HostChannel_Periph) DTERRM() HCINTMSK_Mask {
	return HCINTMSK_Mask{mmio.UM32{&p.HCINTMSK.U32, uint32(DTERRM)}}
}

type HCTSIZ_Bits uint32

func (b HCTSIZ_Bits) Field(mask HCTSIZ_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCTSIZ_Bits) J(v int) HCTSIZ_Bits {
	return HCTSIZ_Bits(bits.Make32(v, uint32(mask)))
}

type HCTSIZ struct{ mmio.U32 }

func (r *HCTSIZ) Bits(mask HCTSIZ_Bits) HCTSIZ_Bits { return HCTSIZ_Bits(r.U32.Bits(uint32(mask))) }
func (r *HCTSIZ) StoreBits(mask, b HCTSIZ_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HCTSIZ) SetBits(mask HCTSIZ_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *HCTSIZ) ClearBits(mask HCTSIZ_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *HCTSIZ) Load() HCTSIZ_Bits                 { return HCTSIZ_Bits(r.U32.Load()) }
func (r *HCTSIZ) Store(b HCTSIZ_Bits)               { r.U32.Store(uint32(b)) }

type HCTSIZ_Mask struct{ mmio.UM32 }

func (rm HCTSIZ_Mask) Load() HCTSIZ_Bits   { return HCTSIZ_Bits(rm.UM32.Load()) }
func (rm HCTSIZ_Mask) Store(b HCTSIZ_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) XFRSIZ() HCTSIZ_Mask {
	return HCTSIZ_Mask{mmio.UM32{&p.HCTSIZ.U32, uint32(XFRSIZ)}}
}

func (p *USB_OTG_HostChannel_Periph) PKTCNT() HCTSIZ_Mask {
	return HCTSIZ_Mask{mmio.UM32{&p.HCTSIZ.U32, uint32(PKTCNT)}}
}

func (p *USB_OTG_HostChannel_Periph) DOPING() HCTSIZ_Mask {
	return HCTSIZ_Mask{mmio.UM32{&p.HCTSIZ.U32, uint32(DOPING)}}
}

func (p *USB_OTG_HostChannel_Periph) DPID() HCTSIZ_Mask {
	return HCTSIZ_Mask{mmio.UM32{&p.HCTSIZ.U32, uint32(DPID)}}
}

type HCDMA_Bits uint32

func (b HCDMA_Bits) Field(mask HCDMA_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCDMA_Bits) J(v int) HCDMA_Bits {
	return HCDMA_Bits(bits.Make32(v, uint32(mask)))
}

type HCDMA struct{ mmio.U32 }

func (r *HCDMA) Bits(mask HCDMA_Bits) HCDMA_Bits { return HCDMA_Bits(r.U32.Bits(uint32(mask))) }
func (r *HCDMA) StoreBits(mask, b HCDMA_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HCDMA) SetBits(mask HCDMA_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *HCDMA) ClearBits(mask HCDMA_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *HCDMA) Load() HCDMA_Bits                { return HCDMA_Bits(r.U32.Load()) }
func (r *HCDMA) Store(b HCDMA_Bits)              { r.U32.Store(uint32(b)) }

type HCDMA_Mask struct{ mmio.UM32 }

func (rm HCDMA_Mask) Load() HCDMA_Bits   { return HCDMA_Bits(rm.UM32.Load()) }
func (rm HCDMA_Mask) Store(b HCDMA_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) DMAADDR() HCDMA_Mask {
	return HCDMA_Mask{mmio.UM32{&p.HCDMA.U32, uint32(DMAADDR)}}
}
