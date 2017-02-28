// +build l476xx

package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type USB_OTG_Device_Periph struct {
	DCFG       DCFG
	DCTL       DCTL
	DSTS       DSTS
	_          uint32
	DIEPMSK    DIEPMSK
	DOEPMSK    DOEPMSK
	DAINT      DAINT
	DAINTMSK   DAINTMSK
	_          [2]uint32
	DVBUSDIS   DVBUSDIS
	DVBUSPULSE DVBUSPULSE
	DTHRCTL    DTHRCTL
	DIEPEMPMSK DIEPEMPMSK
	DEACHINT   DEACHINT
	DEACHMSK   DEACHMSK
	_          uint32
	DINEP1MSK  DINEP1MSK
	_          [15]uint32
	DOUTEP1MSK DOUTEP1MSK
}

func (p *USB_OTG_Device_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type DCFG_Bits uint32

func (b DCFG_Bits) Field(mask DCFG_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCFG_Bits) J(v int) DCFG_Bits {
	return DCFG_Bits(bits.Make32(v, uint32(mask)))
}

type DCFG struct{ mmio.U32 }

func (r *DCFG) Bits(mask DCFG_Bits) DCFG_Bits { return DCFG_Bits(r.U32.Bits(uint32(mask))) }
func (r *DCFG) StoreBits(mask, b DCFG_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DCFG) SetBits(mask DCFG_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DCFG) ClearBits(mask DCFG_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DCFG) Load() DCFG_Bits               { return DCFG_Bits(r.U32.Load()) }
func (r *DCFG) Store(b DCFG_Bits)             { r.U32.Store(uint32(b)) }

type DCFG_Mask struct{ mmio.UM32 }

func (rm DCFG_Mask) Load() DCFG_Bits   { return DCFG_Bits(rm.UM32.Load()) }
func (rm DCFG_Mask) Store(b DCFG_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) DSPD() DCFG_Mask {
	return DCFG_Mask{mmio.UM32{&p.DCFG.U32, uint32(DSPD)}}
}

func (p *USB_OTG_Device_Periph) NZLSOHSK() DCFG_Mask {
	return DCFG_Mask{mmio.UM32{&p.DCFG.U32, uint32(NZLSOHSK)}}
}

func (p *USB_OTG_Device_Periph) DAD() DCFG_Mask {
	return DCFG_Mask{mmio.UM32{&p.DCFG.U32, uint32(DAD)}}
}

func (p *USB_OTG_Device_Periph) PFIVL() DCFG_Mask {
	return DCFG_Mask{mmio.UM32{&p.DCFG.U32, uint32(PFIVL)}}
}

func (p *USB_OTG_Device_Periph) PERSCHIVL() DCFG_Mask {
	return DCFG_Mask{mmio.UM32{&p.DCFG.U32, uint32(PERSCHIVL)}}
}

type DCTL_Bits uint32

func (b DCTL_Bits) Field(mask DCTL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCTL_Bits) J(v int) DCTL_Bits {
	return DCTL_Bits(bits.Make32(v, uint32(mask)))
}

type DCTL struct{ mmio.U32 }

func (r *DCTL) Bits(mask DCTL_Bits) DCTL_Bits { return DCTL_Bits(r.U32.Bits(uint32(mask))) }
func (r *DCTL) StoreBits(mask, b DCTL_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DCTL) SetBits(mask DCTL_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DCTL) ClearBits(mask DCTL_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DCTL) Load() DCTL_Bits               { return DCTL_Bits(r.U32.Load()) }
func (r *DCTL) Store(b DCTL_Bits)             { r.U32.Store(uint32(b)) }

type DCTL_Mask struct{ mmio.UM32 }

func (rm DCTL_Mask) Load() DCTL_Bits   { return DCTL_Bits(rm.UM32.Load()) }
func (rm DCTL_Mask) Store(b DCTL_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) RWUSIG() DCTL_Mask {
	return DCTL_Mask{mmio.UM32{&p.DCTL.U32, uint32(RWUSIG)}}
}

func (p *USB_OTG_Device_Periph) SDIS() DCTL_Mask {
	return DCTL_Mask{mmio.UM32{&p.DCTL.U32, uint32(SDIS)}}
}

func (p *USB_OTG_Device_Periph) GINSTS() DCTL_Mask {
	return DCTL_Mask{mmio.UM32{&p.DCTL.U32, uint32(GINSTS)}}
}

func (p *USB_OTG_Device_Periph) GONSTS() DCTL_Mask {
	return DCTL_Mask{mmio.UM32{&p.DCTL.U32, uint32(GONSTS)}}
}

func (p *USB_OTG_Device_Periph) TCTL() DCTL_Mask {
	return DCTL_Mask{mmio.UM32{&p.DCTL.U32, uint32(TCTL)}}
}

func (p *USB_OTG_Device_Periph) SGINAK() DCTL_Mask {
	return DCTL_Mask{mmio.UM32{&p.DCTL.U32, uint32(SGINAK)}}
}

func (p *USB_OTG_Device_Periph) CGINAK() DCTL_Mask {
	return DCTL_Mask{mmio.UM32{&p.DCTL.U32, uint32(CGINAK)}}
}

func (p *USB_OTG_Device_Periph) SGONAK() DCTL_Mask {
	return DCTL_Mask{mmio.UM32{&p.DCTL.U32, uint32(SGONAK)}}
}

func (p *USB_OTG_Device_Periph) CGONAK() DCTL_Mask {
	return DCTL_Mask{mmio.UM32{&p.DCTL.U32, uint32(CGONAK)}}
}

func (p *USB_OTG_Device_Periph) POPRGDNE() DCTL_Mask {
	return DCTL_Mask{mmio.UM32{&p.DCTL.U32, uint32(POPRGDNE)}}
}

type DSTS_Bits uint32

func (b DSTS_Bits) Field(mask DSTS_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DSTS_Bits) J(v int) DSTS_Bits {
	return DSTS_Bits(bits.Make32(v, uint32(mask)))
}

type DSTS struct{ mmio.U32 }

func (r *DSTS) Bits(mask DSTS_Bits) DSTS_Bits { return DSTS_Bits(r.U32.Bits(uint32(mask))) }
func (r *DSTS) StoreBits(mask, b DSTS_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DSTS) SetBits(mask DSTS_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DSTS) ClearBits(mask DSTS_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DSTS) Load() DSTS_Bits               { return DSTS_Bits(r.U32.Load()) }
func (r *DSTS) Store(b DSTS_Bits)             { r.U32.Store(uint32(b)) }

type DSTS_Mask struct{ mmio.UM32 }

func (rm DSTS_Mask) Load() DSTS_Bits   { return DSTS_Bits(rm.UM32.Load()) }
func (rm DSTS_Mask) Store(b DSTS_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) SUSPSTS() DSTS_Mask {
	return DSTS_Mask{mmio.UM32{&p.DSTS.U32, uint32(SUSPSTS)}}
}

func (p *USB_OTG_Device_Periph) ENUMSPD() DSTS_Mask {
	return DSTS_Mask{mmio.UM32{&p.DSTS.U32, uint32(ENUMSPD)}}
}

func (p *USB_OTG_Device_Periph) EERR() DSTS_Mask {
	return DSTS_Mask{mmio.UM32{&p.DSTS.U32, uint32(EERR)}}
}

func (p *USB_OTG_Device_Periph) FNSOF() DSTS_Mask {
	return DSTS_Mask{mmio.UM32{&p.DSTS.U32, uint32(FNSOF)}}
}

type DIEPMSK_Bits uint32

func (b DIEPMSK_Bits) Field(mask DIEPMSK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPMSK_Bits) J(v int) DIEPMSK_Bits {
	return DIEPMSK_Bits(bits.Make32(v, uint32(mask)))
}

type DIEPMSK struct{ mmio.U32 }

func (r *DIEPMSK) Bits(mask DIEPMSK_Bits) DIEPMSK_Bits { return DIEPMSK_Bits(r.U32.Bits(uint32(mask))) }
func (r *DIEPMSK) StoreBits(mask, b DIEPMSK_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIEPMSK) SetBits(mask DIEPMSK_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DIEPMSK) ClearBits(mask DIEPMSK_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DIEPMSK) Load() DIEPMSK_Bits                  { return DIEPMSK_Bits(r.U32.Load()) }
func (r *DIEPMSK) Store(b DIEPMSK_Bits)                { r.U32.Store(uint32(b)) }

type DIEPMSK_Mask struct{ mmio.UM32 }

func (rm DIEPMSK_Mask) Load() DIEPMSK_Bits   { return DIEPMSK_Bits(rm.UM32.Load()) }
func (rm DIEPMSK_Mask) Store(b DIEPMSK_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) XFRCM() DIEPMSK_Mask {
	return DIEPMSK_Mask{mmio.UM32{&p.DIEPMSK.U32, uint32(XFRCM)}}
}

func (p *USB_OTG_Device_Periph) EPDM() DIEPMSK_Mask {
	return DIEPMSK_Mask{mmio.UM32{&p.DIEPMSK.U32, uint32(EPDM)}}
}

func (p *USB_OTG_Device_Periph) TOM() DIEPMSK_Mask {
	return DIEPMSK_Mask{mmio.UM32{&p.DIEPMSK.U32, uint32(TOM)}}
}

func (p *USB_OTG_Device_Periph) ITTXFEMSK() DIEPMSK_Mask {
	return DIEPMSK_Mask{mmio.UM32{&p.DIEPMSK.U32, uint32(ITTXFEMSK)}}
}

func (p *USB_OTG_Device_Periph) INEPNMM() DIEPMSK_Mask {
	return DIEPMSK_Mask{mmio.UM32{&p.DIEPMSK.U32, uint32(INEPNMM)}}
}

func (p *USB_OTG_Device_Periph) INEPNEM() DIEPMSK_Mask {
	return DIEPMSK_Mask{mmio.UM32{&p.DIEPMSK.U32, uint32(INEPNEM)}}
}

func (p *USB_OTG_Device_Periph) TXFURM() DIEPMSK_Mask {
	return DIEPMSK_Mask{mmio.UM32{&p.DIEPMSK.U32, uint32(TXFURM)}}
}

func (p *USB_OTG_Device_Periph) BIM() DIEPMSK_Mask {
	return DIEPMSK_Mask{mmio.UM32{&p.DIEPMSK.U32, uint32(BIM)}}
}

type DOEPMSK_Bits uint32

func (b DOEPMSK_Bits) Field(mask DOEPMSK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOEPMSK_Bits) J(v int) DOEPMSK_Bits {
	return DOEPMSK_Bits(bits.Make32(v, uint32(mask)))
}

type DOEPMSK struct{ mmio.U32 }

func (r *DOEPMSK) Bits(mask DOEPMSK_Bits) DOEPMSK_Bits { return DOEPMSK_Bits(r.U32.Bits(uint32(mask))) }
func (r *DOEPMSK) StoreBits(mask, b DOEPMSK_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOEPMSK) SetBits(mask DOEPMSK_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DOEPMSK) ClearBits(mask DOEPMSK_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DOEPMSK) Load() DOEPMSK_Bits                  { return DOEPMSK_Bits(r.U32.Load()) }
func (r *DOEPMSK) Store(b DOEPMSK_Bits)                { r.U32.Store(uint32(b)) }

type DOEPMSK_Mask struct{ mmio.UM32 }

func (rm DOEPMSK_Mask) Load() DOEPMSK_Bits   { return DOEPMSK_Bits(rm.UM32.Load()) }
func (rm DOEPMSK_Mask) Store(b DOEPMSK_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) XFRCM() DOEPMSK_Mask {
	return DOEPMSK_Mask{mmio.UM32{&p.DOEPMSK.U32, uint32(XFRCM)}}
}

func (p *USB_OTG_Device_Periph) EPDM() DOEPMSK_Mask {
	return DOEPMSK_Mask{mmio.UM32{&p.DOEPMSK.U32, uint32(EPDM)}}
}

func (p *USB_OTG_Device_Periph) STUPM() DOEPMSK_Mask {
	return DOEPMSK_Mask{mmio.UM32{&p.DOEPMSK.U32, uint32(STUPM)}}
}

func (p *USB_OTG_Device_Periph) OTEPDM() DOEPMSK_Mask {
	return DOEPMSK_Mask{mmio.UM32{&p.DOEPMSK.U32, uint32(OTEPDM)}}
}

func (p *USB_OTG_Device_Periph) B2BSTUP() DOEPMSK_Mask {
	return DOEPMSK_Mask{mmio.UM32{&p.DOEPMSK.U32, uint32(B2BSTUP)}}
}

func (p *USB_OTG_Device_Periph) OPEM() DOEPMSK_Mask {
	return DOEPMSK_Mask{mmio.UM32{&p.DOEPMSK.U32, uint32(OPEM)}}
}

func (p *USB_OTG_Device_Periph) BOIM() DOEPMSK_Mask {
	return DOEPMSK_Mask{mmio.UM32{&p.DOEPMSK.U32, uint32(BOIM)}}
}

type DAINT_Bits uint32

func (b DAINT_Bits) Field(mask DAINT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DAINT_Bits) J(v int) DAINT_Bits {
	return DAINT_Bits(bits.Make32(v, uint32(mask)))
}

type DAINT struct{ mmio.U32 }

func (r *DAINT) Bits(mask DAINT_Bits) DAINT_Bits { return DAINT_Bits(r.U32.Bits(uint32(mask))) }
func (r *DAINT) StoreBits(mask, b DAINT_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DAINT) SetBits(mask DAINT_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *DAINT) ClearBits(mask DAINT_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *DAINT) Load() DAINT_Bits                { return DAINT_Bits(r.U32.Load()) }
func (r *DAINT) Store(b DAINT_Bits)              { r.U32.Store(uint32(b)) }

type DAINT_Mask struct{ mmio.UM32 }

func (rm DAINT_Mask) Load() DAINT_Bits   { return DAINT_Bits(rm.UM32.Load()) }
func (rm DAINT_Mask) Store(b DAINT_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) IEPINT() DAINT_Mask {
	return DAINT_Mask{mmio.UM32{&p.DAINT.U32, uint32(IEPINT)}}
}

func (p *USB_OTG_Device_Periph) OEPINT() DAINT_Mask {
	return DAINT_Mask{mmio.UM32{&p.DAINT.U32, uint32(OEPINT)}}
}

type DAINTMSK_Bits uint32

func (b DAINTMSK_Bits) Field(mask DAINTMSK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DAINTMSK_Bits) J(v int) DAINTMSK_Bits {
	return DAINTMSK_Bits(bits.Make32(v, uint32(mask)))
}

type DAINTMSK struct{ mmio.U32 }

func (r *DAINTMSK) Bits(mask DAINTMSK_Bits) DAINTMSK_Bits {
	return DAINTMSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DAINTMSK) StoreBits(mask, b DAINTMSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DAINTMSK) SetBits(mask DAINTMSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DAINTMSK) ClearBits(mask DAINTMSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DAINTMSK) Load() DAINTMSK_Bits             { return DAINTMSK_Bits(r.U32.Load()) }
func (r *DAINTMSK) Store(b DAINTMSK_Bits)           { r.U32.Store(uint32(b)) }

type DAINTMSK_Mask struct{ mmio.UM32 }

func (rm DAINTMSK_Mask) Load() DAINTMSK_Bits   { return DAINTMSK_Bits(rm.UM32.Load()) }
func (rm DAINTMSK_Mask) Store(b DAINTMSK_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) IEPM() DAINTMSK_Mask {
	return DAINTMSK_Mask{mmio.UM32{&p.DAINTMSK.U32, uint32(IEPM)}}
}

func (p *USB_OTG_Device_Periph) OEPM() DAINTMSK_Mask {
	return DAINTMSK_Mask{mmio.UM32{&p.DAINTMSK.U32, uint32(OEPM)}}
}

type DVBUSDIS_Bits uint32

func (b DVBUSDIS_Bits) Field(mask DVBUSDIS_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DVBUSDIS_Bits) J(v int) DVBUSDIS_Bits {
	return DVBUSDIS_Bits(bits.Make32(v, uint32(mask)))
}

type DVBUSDIS struct{ mmio.U32 }

func (r *DVBUSDIS) Bits(mask DVBUSDIS_Bits) DVBUSDIS_Bits {
	return DVBUSDIS_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DVBUSDIS) StoreBits(mask, b DVBUSDIS_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DVBUSDIS) SetBits(mask DVBUSDIS_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DVBUSDIS) ClearBits(mask DVBUSDIS_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DVBUSDIS) Load() DVBUSDIS_Bits             { return DVBUSDIS_Bits(r.U32.Load()) }
func (r *DVBUSDIS) Store(b DVBUSDIS_Bits)           { r.U32.Store(uint32(b)) }

type DVBUSDIS_Mask struct{ mmio.UM32 }

func (rm DVBUSDIS_Mask) Load() DVBUSDIS_Bits   { return DVBUSDIS_Bits(rm.UM32.Load()) }
func (rm DVBUSDIS_Mask) Store(b DVBUSDIS_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) VBUSDT() DVBUSDIS_Mask {
	return DVBUSDIS_Mask{mmio.UM32{&p.DVBUSDIS.U32, uint32(VBUSDT)}}
}

type DVBUSPULSE_Bits uint32

func (b DVBUSPULSE_Bits) Field(mask DVBUSPULSE_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DVBUSPULSE_Bits) J(v int) DVBUSPULSE_Bits {
	return DVBUSPULSE_Bits(bits.Make32(v, uint32(mask)))
}

type DVBUSPULSE struct{ mmio.U32 }

func (r *DVBUSPULSE) Bits(mask DVBUSPULSE_Bits) DVBUSPULSE_Bits {
	return DVBUSPULSE_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DVBUSPULSE) StoreBits(mask, b DVBUSPULSE_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DVBUSPULSE) SetBits(mask DVBUSPULSE_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DVBUSPULSE) ClearBits(mask DVBUSPULSE_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DVBUSPULSE) Load() DVBUSPULSE_Bits             { return DVBUSPULSE_Bits(r.U32.Load()) }
func (r *DVBUSPULSE) Store(b DVBUSPULSE_Bits)           { r.U32.Store(uint32(b)) }

type DVBUSPULSE_Mask struct{ mmio.UM32 }

func (rm DVBUSPULSE_Mask) Load() DVBUSPULSE_Bits   { return DVBUSPULSE_Bits(rm.UM32.Load()) }
func (rm DVBUSPULSE_Mask) Store(b DVBUSPULSE_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) DVBUSP() DVBUSPULSE_Mask {
	return DVBUSPULSE_Mask{mmio.UM32{&p.DVBUSPULSE.U32, uint32(DVBUSP)}}
}

type DTHRCTL_Bits uint32

func (b DTHRCTL_Bits) Field(mask DTHRCTL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DTHRCTL_Bits) J(v int) DTHRCTL_Bits {
	return DTHRCTL_Bits(bits.Make32(v, uint32(mask)))
}

type DTHRCTL struct{ mmio.U32 }

func (r *DTHRCTL) Bits(mask DTHRCTL_Bits) DTHRCTL_Bits { return DTHRCTL_Bits(r.U32.Bits(uint32(mask))) }
func (r *DTHRCTL) StoreBits(mask, b DTHRCTL_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DTHRCTL) SetBits(mask DTHRCTL_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DTHRCTL) ClearBits(mask DTHRCTL_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DTHRCTL) Load() DTHRCTL_Bits                  { return DTHRCTL_Bits(r.U32.Load()) }
func (r *DTHRCTL) Store(b DTHRCTL_Bits)                { r.U32.Store(uint32(b)) }

type DTHRCTL_Mask struct{ mmio.UM32 }

func (rm DTHRCTL_Mask) Load() DTHRCTL_Bits   { return DTHRCTL_Bits(rm.UM32.Load()) }
func (rm DTHRCTL_Mask) Store(b DTHRCTL_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) NONISOTHREN() DTHRCTL_Mask {
	return DTHRCTL_Mask{mmio.UM32{&p.DTHRCTL.U32, uint32(NONISOTHREN)}}
}

func (p *USB_OTG_Device_Periph) ISOTHREN() DTHRCTL_Mask {
	return DTHRCTL_Mask{mmio.UM32{&p.DTHRCTL.U32, uint32(ISOTHREN)}}
}

func (p *USB_OTG_Device_Periph) TXTHRLEN() DTHRCTL_Mask {
	return DTHRCTL_Mask{mmio.UM32{&p.DTHRCTL.U32, uint32(TXTHRLEN)}}
}

func (p *USB_OTG_Device_Periph) RXTHREN() DTHRCTL_Mask {
	return DTHRCTL_Mask{mmio.UM32{&p.DTHRCTL.U32, uint32(RXTHREN)}}
}

func (p *USB_OTG_Device_Periph) RXTHRLEN() DTHRCTL_Mask {
	return DTHRCTL_Mask{mmio.UM32{&p.DTHRCTL.U32, uint32(RXTHRLEN)}}
}

func (p *USB_OTG_Device_Periph) ARPEN() DTHRCTL_Mask {
	return DTHRCTL_Mask{mmio.UM32{&p.DTHRCTL.U32, uint32(ARPEN)}}
}

type DIEPEMPMSK_Bits uint32

func (b DIEPEMPMSK_Bits) Field(mask DIEPEMPMSK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DIEPEMPMSK_Bits) J(v int) DIEPEMPMSK_Bits {
	return DIEPEMPMSK_Bits(bits.Make32(v, uint32(mask)))
}

type DIEPEMPMSK struct{ mmio.U32 }

func (r *DIEPEMPMSK) Bits(mask DIEPEMPMSK_Bits) DIEPEMPMSK_Bits {
	return DIEPEMPMSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DIEPEMPMSK) StoreBits(mask, b DIEPEMPMSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIEPEMPMSK) SetBits(mask DIEPEMPMSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DIEPEMPMSK) ClearBits(mask DIEPEMPMSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DIEPEMPMSK) Load() DIEPEMPMSK_Bits             { return DIEPEMPMSK_Bits(r.U32.Load()) }
func (r *DIEPEMPMSK) Store(b DIEPEMPMSK_Bits)           { r.U32.Store(uint32(b)) }

type DIEPEMPMSK_Mask struct{ mmio.UM32 }

func (rm DIEPEMPMSK_Mask) Load() DIEPEMPMSK_Bits   { return DIEPEMPMSK_Bits(rm.UM32.Load()) }
func (rm DIEPEMPMSK_Mask) Store(b DIEPEMPMSK_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) INEPTXFEM() DIEPEMPMSK_Mask {
	return DIEPEMPMSK_Mask{mmio.UM32{&p.DIEPEMPMSK.U32, uint32(INEPTXFEM)}}
}

type DEACHINT_Bits uint32

func (b DEACHINT_Bits) Field(mask DEACHINT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DEACHINT_Bits) J(v int) DEACHINT_Bits {
	return DEACHINT_Bits(bits.Make32(v, uint32(mask)))
}

type DEACHINT struct{ mmio.U32 }

func (r *DEACHINT) Bits(mask DEACHINT_Bits) DEACHINT_Bits {
	return DEACHINT_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DEACHINT) StoreBits(mask, b DEACHINT_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DEACHINT) SetBits(mask DEACHINT_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DEACHINT) ClearBits(mask DEACHINT_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DEACHINT) Load() DEACHINT_Bits             { return DEACHINT_Bits(r.U32.Load()) }
func (r *DEACHINT) Store(b DEACHINT_Bits)           { r.U32.Store(uint32(b)) }

type DEACHINT_Mask struct{ mmio.UM32 }

func (rm DEACHINT_Mask) Load() DEACHINT_Bits   { return DEACHINT_Bits(rm.UM32.Load()) }
func (rm DEACHINT_Mask) Store(b DEACHINT_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Device_Periph) IEP1INT() DEACHINT_Mask {
	return DEACHINT_Mask{mmio.UM32{&p.DEACHINT.U32, uint32(IEP1INT)}}
}

func (p *USB_OTG_Device_Periph) OEP1INT() DEACHINT_Mask {
	return DEACHINT_Mask{mmio.UM32{&p.DEACHINT.U32, uint32(OEP1INT)}}
}

type DEACHMSK_Bits uint32

func (b DEACHMSK_Bits) Field(mask DEACHMSK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DEACHMSK_Bits) J(v int) DEACHMSK_Bits {
	return DEACHMSK_Bits(bits.Make32(v, uint32(mask)))
}

type DEACHMSK struct{ mmio.U32 }

func (r *DEACHMSK) Bits(mask DEACHMSK_Bits) DEACHMSK_Bits {
	return DEACHMSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DEACHMSK) StoreBits(mask, b DEACHMSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DEACHMSK) SetBits(mask DEACHMSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DEACHMSK) ClearBits(mask DEACHMSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DEACHMSK) Load() DEACHMSK_Bits             { return DEACHMSK_Bits(r.U32.Load()) }
func (r *DEACHMSK) Store(b DEACHMSK_Bits)           { r.U32.Store(uint32(b)) }

type DEACHMSK_Mask struct{ mmio.UM32 }

func (rm DEACHMSK_Mask) Load() DEACHMSK_Bits   { return DEACHMSK_Bits(rm.UM32.Load()) }
func (rm DEACHMSK_Mask) Store(b DEACHMSK_Bits) { rm.UM32.Store(uint32(b)) }

type DINEP1MSK_Bits uint32

func (b DINEP1MSK_Bits) Field(mask DINEP1MSK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DINEP1MSK_Bits) J(v int) DINEP1MSK_Bits {
	return DINEP1MSK_Bits(bits.Make32(v, uint32(mask)))
}

type DINEP1MSK struct{ mmio.U32 }

func (r *DINEP1MSK) Bits(mask DINEP1MSK_Bits) DINEP1MSK_Bits {
	return DINEP1MSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DINEP1MSK) StoreBits(mask, b DINEP1MSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DINEP1MSK) SetBits(mask DINEP1MSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DINEP1MSK) ClearBits(mask DINEP1MSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DINEP1MSK) Load() DINEP1MSK_Bits             { return DINEP1MSK_Bits(r.U32.Load()) }
func (r *DINEP1MSK) Store(b DINEP1MSK_Bits)           { r.U32.Store(uint32(b)) }

type DINEP1MSK_Mask struct{ mmio.UM32 }

func (rm DINEP1MSK_Mask) Load() DINEP1MSK_Bits   { return DINEP1MSK_Bits(rm.UM32.Load()) }
func (rm DINEP1MSK_Mask) Store(b DINEP1MSK_Bits) { rm.UM32.Store(uint32(b)) }

type DOUTEP1MSK_Bits uint32

func (b DOUTEP1MSK_Bits) Field(mask DOUTEP1MSK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOUTEP1MSK_Bits) J(v int) DOUTEP1MSK_Bits {
	return DOUTEP1MSK_Bits(bits.Make32(v, uint32(mask)))
}

type DOUTEP1MSK struct{ mmio.U32 }

func (r *DOUTEP1MSK) Bits(mask DOUTEP1MSK_Bits) DOUTEP1MSK_Bits {
	return DOUTEP1MSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DOUTEP1MSK) StoreBits(mask, b DOUTEP1MSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOUTEP1MSK) SetBits(mask DOUTEP1MSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DOUTEP1MSK) ClearBits(mask DOUTEP1MSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DOUTEP1MSK) Load() DOUTEP1MSK_Bits             { return DOUTEP1MSK_Bits(r.U32.Load()) }
func (r *DOUTEP1MSK) Store(b DOUTEP1MSK_Bits)           { r.U32.Store(uint32(b)) }

type DOUTEP1MSK_Mask struct{ mmio.UM32 }

func (rm DOUTEP1MSK_Mask) Load() DOUTEP1MSK_Bits   { return DOUTEP1MSK_Bits(rm.UM32.Load()) }
func (rm DOUTEP1MSK_Mask) Store(b DOUTEP1MSK_Bits) { rm.UM32.Store(uint32(b)) }
