package sdmmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type SDMMC_Periph struct {
	POWER  POWER
	CLKCR  CLKCR
	ARG    ARG
	CMD    CMD
	DTIMER DTIMER
	DLEN   DLEN
	DCTRL  DCTRL
	ICR    ICR
	MASK   MASK
	_      [15]uint32
	FIFO   FIFO
}

func (p *SDMMC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var SDMMC1 = (*SDMMC_Periph)(unsafe.Pointer(uintptr(mmap.SDMMC1_BASE)))

type POWER_Bits uint32

func (b POWER_Bits) Field(mask POWER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask POWER_Bits) J(v int) POWER_Bits {
	return POWER_Bits(bits.Make32(v, uint32(mask)))
}

type POWER struct{ mmio.U32 }

func (r *POWER) Bits(mask POWER_Bits) POWER_Bits { return POWER_Bits(r.U32.Bits(uint32(mask))) }
func (r *POWER) StoreBits(mask, b POWER_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *POWER) SetBits(mask POWER_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *POWER) ClearBits(mask POWER_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *POWER) Load() POWER_Bits                { return POWER_Bits(r.U32.Load()) }
func (r *POWER) Store(b POWER_Bits)              { r.U32.Store(uint32(b)) }

type POWER_Mask struct{ mmio.UM32 }

func (rm POWER_Mask) Load() POWER_Bits   { return POWER_Bits(rm.UM32.Load()) }
func (rm POWER_Mask) Store(b POWER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SDMMC_Periph) PWRCTRL() POWER_Mask {
	return POWER_Mask{mmio.UM32{&p.POWER.U32, uint32(PWRCTRL)}}
}

type CLKCR_Bits uint32

func (b CLKCR_Bits) Field(mask CLKCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CLKCR_Bits) J(v int) CLKCR_Bits {
	return CLKCR_Bits(bits.Make32(v, uint32(mask)))
}

type CLKCR struct{ mmio.U32 }

func (r *CLKCR) Bits(mask CLKCR_Bits) CLKCR_Bits { return CLKCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CLKCR) StoreBits(mask, b CLKCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CLKCR) SetBits(mask CLKCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CLKCR) ClearBits(mask CLKCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CLKCR) Load() CLKCR_Bits                { return CLKCR_Bits(r.U32.Load()) }
func (r *CLKCR) Store(b CLKCR_Bits)              { r.U32.Store(uint32(b)) }

type CLKCR_Mask struct{ mmio.UM32 }

func (rm CLKCR_Mask) Load() CLKCR_Bits   { return CLKCR_Bits(rm.UM32.Load()) }
func (rm CLKCR_Mask) Store(b CLKCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SDMMC_Periph) CLKDIV() CLKCR_Mask {
	return CLKCR_Mask{mmio.UM32{&p.CLKCR.U32, uint32(CLKDIV)}}
}

func (p *SDMMC_Periph) CLKEN() CLKCR_Mask {
	return CLKCR_Mask{mmio.UM32{&p.CLKCR.U32, uint32(CLKEN)}}
}

func (p *SDMMC_Periph) PWRSAV() CLKCR_Mask {
	return CLKCR_Mask{mmio.UM32{&p.CLKCR.U32, uint32(PWRSAV)}}
}

func (p *SDMMC_Periph) BYPASS() CLKCR_Mask {
	return CLKCR_Mask{mmio.UM32{&p.CLKCR.U32, uint32(BYPASS)}}
}

func (p *SDMMC_Periph) WIDBUS() CLKCR_Mask {
	return CLKCR_Mask{mmio.UM32{&p.CLKCR.U32, uint32(WIDBUS)}}
}

func (p *SDMMC_Periph) NEGEDGE() CLKCR_Mask {
	return CLKCR_Mask{mmio.UM32{&p.CLKCR.U32, uint32(NEGEDGE)}}
}

func (p *SDMMC_Periph) HWFC_EN() CLKCR_Mask {
	return CLKCR_Mask{mmio.UM32{&p.CLKCR.U32, uint32(HWFC_EN)}}
}

type ARG_Bits uint32

func (b ARG_Bits) Field(mask ARG_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ARG_Bits) J(v int) ARG_Bits {
	return ARG_Bits(bits.Make32(v, uint32(mask)))
}

type ARG struct{ mmio.U32 }

func (r *ARG) Bits(mask ARG_Bits) ARG_Bits { return ARG_Bits(r.U32.Bits(uint32(mask))) }
func (r *ARG) StoreBits(mask, b ARG_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ARG) SetBits(mask ARG_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ARG) ClearBits(mask ARG_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ARG) Load() ARG_Bits              { return ARG_Bits(r.U32.Load()) }
func (r *ARG) Store(b ARG_Bits)            { r.U32.Store(uint32(b)) }

type ARG_Mask struct{ mmio.UM32 }

func (rm ARG_Mask) Load() ARG_Bits   { return ARG_Bits(rm.UM32.Load()) }
func (rm ARG_Mask) Store(b ARG_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SDMMC_Periph) CMDARG() ARG_Mask {
	return ARG_Mask{mmio.UM32{&p.ARG.U32, uint32(CMDARG)}}
}

type CMD_Bits uint32

func (b CMD_Bits) Field(mask CMD_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMD_Bits) J(v int) CMD_Bits {
	return CMD_Bits(bits.Make32(v, uint32(mask)))
}

type CMD struct{ mmio.U32 }

func (r *CMD) Bits(mask CMD_Bits) CMD_Bits { return CMD_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMD) StoreBits(mask, b CMD_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMD) SetBits(mask CMD_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CMD) ClearBits(mask CMD_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CMD) Load() CMD_Bits              { return CMD_Bits(r.U32.Load()) }
func (r *CMD) Store(b CMD_Bits)            { r.U32.Store(uint32(b)) }

type CMD_Mask struct{ mmio.UM32 }

func (rm CMD_Mask) Load() CMD_Bits   { return CMD_Bits(rm.UM32.Load()) }
func (rm CMD_Mask) Store(b CMD_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SDMMC_Periph) CMDINDEX() CMD_Mask {
	return CMD_Mask{mmio.UM32{&p.CMD.U32, uint32(CMDINDEX)}}
}

func (p *SDMMC_Periph) WAITRESP() CMD_Mask {
	return CMD_Mask{mmio.UM32{&p.CMD.U32, uint32(WAITRESP)}}
}

func (p *SDMMC_Periph) WAITINT() CMD_Mask {
	return CMD_Mask{mmio.UM32{&p.CMD.U32, uint32(WAITINT)}}
}

func (p *SDMMC_Periph) WAITPEND() CMD_Mask {
	return CMD_Mask{mmio.UM32{&p.CMD.U32, uint32(WAITPEND)}}
}

func (p *SDMMC_Periph) CPSMEN() CMD_Mask {
	return CMD_Mask{mmio.UM32{&p.CMD.U32, uint32(CPSMEN)}}
}

func (p *SDMMC_Periph) SDIOSUSPEND() CMD_Mask {
	return CMD_Mask{mmio.UM32{&p.CMD.U32, uint32(SDIOSUSPEND)}}
}

type DTIMER_Bits uint32

func (b DTIMER_Bits) Field(mask DTIMER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DTIMER_Bits) J(v int) DTIMER_Bits {
	return DTIMER_Bits(bits.Make32(v, uint32(mask)))
}

type DTIMER struct{ mmio.U32 }

func (r *DTIMER) Bits(mask DTIMER_Bits) DTIMER_Bits { return DTIMER_Bits(r.U32.Bits(uint32(mask))) }
func (r *DTIMER) StoreBits(mask, b DTIMER_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DTIMER) SetBits(mask DTIMER_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *DTIMER) ClearBits(mask DTIMER_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *DTIMER) Load() DTIMER_Bits                 { return DTIMER_Bits(r.U32.Load()) }
func (r *DTIMER) Store(b DTIMER_Bits)               { r.U32.Store(uint32(b)) }

type DTIMER_Mask struct{ mmio.UM32 }

func (rm DTIMER_Mask) Load() DTIMER_Bits   { return DTIMER_Bits(rm.UM32.Load()) }
func (rm DTIMER_Mask) Store(b DTIMER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SDMMC_Periph) DATATIME() DTIMER_Mask {
	return DTIMER_Mask{mmio.UM32{&p.DTIMER.U32, uint32(DATATIME)}}
}

type DLEN_Bits uint32

func (b DLEN_Bits) Field(mask DLEN_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DLEN_Bits) J(v int) DLEN_Bits {
	return DLEN_Bits(bits.Make32(v, uint32(mask)))
}

type DLEN struct{ mmio.U32 }

func (r *DLEN) Bits(mask DLEN_Bits) DLEN_Bits { return DLEN_Bits(r.U32.Bits(uint32(mask))) }
func (r *DLEN) StoreBits(mask, b DLEN_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DLEN) SetBits(mask DLEN_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DLEN) ClearBits(mask DLEN_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DLEN) Load() DLEN_Bits               { return DLEN_Bits(r.U32.Load()) }
func (r *DLEN) Store(b DLEN_Bits)             { r.U32.Store(uint32(b)) }

type DLEN_Mask struct{ mmio.UM32 }

func (rm DLEN_Mask) Load() DLEN_Bits   { return DLEN_Bits(rm.UM32.Load()) }
func (rm DLEN_Mask) Store(b DLEN_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SDMMC_Periph) DATALENGTH() DLEN_Mask {
	return DLEN_Mask{mmio.UM32{&p.DLEN.U32, uint32(DATALENGTH)}}
}

type DCTRL_Bits uint32

func (b DCTRL_Bits) Field(mask DCTRL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCTRL_Bits) J(v int) DCTRL_Bits {
	return DCTRL_Bits(bits.Make32(v, uint32(mask)))
}

type DCTRL struct{ mmio.U32 }

func (r *DCTRL) Bits(mask DCTRL_Bits) DCTRL_Bits { return DCTRL_Bits(r.U32.Bits(uint32(mask))) }
func (r *DCTRL) StoreBits(mask, b DCTRL_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DCTRL) SetBits(mask DCTRL_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *DCTRL) ClearBits(mask DCTRL_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *DCTRL) Load() DCTRL_Bits                { return DCTRL_Bits(r.U32.Load()) }
func (r *DCTRL) Store(b DCTRL_Bits)              { r.U32.Store(uint32(b)) }

type DCTRL_Mask struct{ mmio.UM32 }

func (rm DCTRL_Mask) Load() DCTRL_Bits   { return DCTRL_Bits(rm.UM32.Load()) }
func (rm DCTRL_Mask) Store(b DCTRL_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SDMMC_Periph) DTEN() DCTRL_Mask {
	return DCTRL_Mask{mmio.UM32{&p.DCTRL.U32, uint32(DTEN)}}
}

func (p *SDMMC_Periph) DTDIR() DCTRL_Mask {
	return DCTRL_Mask{mmio.UM32{&p.DCTRL.U32, uint32(DTDIR)}}
}

func (p *SDMMC_Periph) DTMODE() DCTRL_Mask {
	return DCTRL_Mask{mmio.UM32{&p.DCTRL.U32, uint32(DTMODE)}}
}

func (p *SDMMC_Periph) DMAEN() DCTRL_Mask {
	return DCTRL_Mask{mmio.UM32{&p.DCTRL.U32, uint32(DMAEN)}}
}

func (p *SDMMC_Periph) DBLOCKSIZE() DCTRL_Mask {
	return DCTRL_Mask{mmio.UM32{&p.DCTRL.U32, uint32(DBLOCKSIZE)}}
}

func (p *SDMMC_Periph) RWSTART() DCTRL_Mask {
	return DCTRL_Mask{mmio.UM32{&p.DCTRL.U32, uint32(RWSTART)}}
}

func (p *SDMMC_Periph) RWSTOP() DCTRL_Mask {
	return DCTRL_Mask{mmio.UM32{&p.DCTRL.U32, uint32(RWSTOP)}}
}

func (p *SDMMC_Periph) RWMOD() DCTRL_Mask {
	return DCTRL_Mask{mmio.UM32{&p.DCTRL.U32, uint32(RWMOD)}}
}

func (p *SDMMC_Periph) SDIOEN() DCTRL_Mask {
	return DCTRL_Mask{mmio.UM32{&p.DCTRL.U32, uint32(SDIOEN)}}
}

type ICR_Bits uint32

func (b ICR_Bits) Field(mask ICR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR_Bits) J(v int) ICR_Bits {
	return ICR_Bits(bits.Make32(v, uint32(mask)))
}

type ICR struct{ mmio.U32 }

func (r *ICR) Bits(mask ICR_Bits) ICR_Bits { return ICR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ICR) StoreBits(mask, b ICR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ICR) SetBits(mask ICR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ICR) ClearBits(mask ICR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ICR) Load() ICR_Bits              { return ICR_Bits(r.U32.Load()) }
func (r *ICR) Store(b ICR_Bits)            { r.U32.Store(uint32(b)) }

type ICR_Mask struct{ mmio.UM32 }

func (rm ICR_Mask) Load() ICR_Bits   { return ICR_Bits(rm.UM32.Load()) }
func (rm ICR_Mask) Store(b ICR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SDMMC_Periph) CCRCFAILC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(CCRCFAILC)}}
}

func (p *SDMMC_Periph) DCRCFAILC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(DCRCFAILC)}}
}

func (p *SDMMC_Periph) CTIMEOUTC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(CTIMEOUTC)}}
}

func (p *SDMMC_Periph) DTIMEOUTC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(DTIMEOUTC)}}
}

func (p *SDMMC_Periph) TXUNDERRC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(TXUNDERRC)}}
}

func (p *SDMMC_Periph) RXOVERRC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(RXOVERRC)}}
}

func (p *SDMMC_Periph) CMDRENDC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(CMDRENDC)}}
}

func (p *SDMMC_Periph) CMDSENTC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(CMDSENTC)}}
}

func (p *SDMMC_Periph) DATAENDC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(DATAENDC)}}
}

func (p *SDMMC_Periph) STBITERRC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(STBITERRC)}}
}

func (p *SDMMC_Periph) DBCKENDC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(DBCKENDC)}}
}

func (p *SDMMC_Periph) SDIOITC() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(SDIOITC)}}
}

type MASK_Bits uint32

func (b MASK_Bits) Field(mask MASK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MASK_Bits) J(v int) MASK_Bits {
	return MASK_Bits(bits.Make32(v, uint32(mask)))
}

type MASK struct{ mmio.U32 }

func (r *MASK) Bits(mask MASK_Bits) MASK_Bits { return MASK_Bits(r.U32.Bits(uint32(mask))) }
func (r *MASK) StoreBits(mask, b MASK_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MASK) SetBits(mask MASK_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *MASK) ClearBits(mask MASK_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *MASK) Load() MASK_Bits               { return MASK_Bits(r.U32.Load()) }
func (r *MASK) Store(b MASK_Bits)             { r.U32.Store(uint32(b)) }

type MASK_Mask struct{ mmio.UM32 }

func (rm MASK_Mask) Load() MASK_Bits   { return MASK_Bits(rm.UM32.Load()) }
func (rm MASK_Mask) Store(b MASK_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SDMMC_Periph) CCRCFAILIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(CCRCFAILIE)}}
}

func (p *SDMMC_Periph) DCRCFAILIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(DCRCFAILIE)}}
}

func (p *SDMMC_Periph) CTIMEOUTIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(CTIMEOUTIE)}}
}

func (p *SDMMC_Periph) DTIMEOUTIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(DTIMEOUTIE)}}
}

func (p *SDMMC_Periph) TXUNDERRIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(TXUNDERRIE)}}
}

func (p *SDMMC_Periph) RXOVERRIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(RXOVERRIE)}}
}

func (p *SDMMC_Periph) CMDRENDIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(CMDRENDIE)}}
}

func (p *SDMMC_Periph) CMDSENTIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(CMDSENTIE)}}
}

func (p *SDMMC_Periph) DATAENDIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(DATAENDIE)}}
}

func (p *SDMMC_Periph) DBCKENDIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(DBCKENDIE)}}
}

func (p *SDMMC_Periph) CMDACTIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(CMDACTIE)}}
}

func (p *SDMMC_Periph) TXACTIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(TXACTIE)}}
}

func (p *SDMMC_Periph) RXACTIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(RXACTIE)}}
}

func (p *SDMMC_Periph) TXFIFOHEIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(TXFIFOHEIE)}}
}

func (p *SDMMC_Periph) RXFIFOHFIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(RXFIFOHFIE)}}
}

func (p *SDMMC_Periph) TXFIFOFIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(TXFIFOFIE)}}
}

func (p *SDMMC_Periph) RXFIFOFIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(RXFIFOFIE)}}
}

func (p *SDMMC_Periph) TXFIFOEIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(TXFIFOEIE)}}
}

func (p *SDMMC_Periph) RXFIFOEIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(RXFIFOEIE)}}
}

func (p *SDMMC_Periph) TXDAVLIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(TXDAVLIE)}}
}

func (p *SDMMC_Periph) RXDAVLIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(RXDAVLIE)}}
}

func (p *SDMMC_Periph) SDIOITIE() MASK_Mask {
	return MASK_Mask{mmio.UM32{&p.MASK.U32, uint32(SDIOITIE)}}
}

type FIFO_Bits uint32

func (b FIFO_Bits) Field(mask FIFO_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FIFO_Bits) J(v int) FIFO_Bits {
	return FIFO_Bits(bits.Make32(v, uint32(mask)))
}

type FIFO struct{ mmio.U32 }

func (r *FIFO) Bits(mask FIFO_Bits) FIFO_Bits { return FIFO_Bits(r.U32.Bits(uint32(mask))) }
func (r *FIFO) StoreBits(mask, b FIFO_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FIFO) SetBits(mask FIFO_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *FIFO) ClearBits(mask FIFO_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *FIFO) Load() FIFO_Bits               { return FIFO_Bits(r.U32.Load()) }
func (r *FIFO) Store(b FIFO_Bits)             { r.U32.Store(uint32(b)) }

type FIFO_Mask struct{ mmio.UM32 }

func (rm FIFO_Mask) Load() FIFO_Bits   { return FIFO_Bits(rm.UM32.Load()) }
func (rm FIFO_Mask) Store(b FIFO_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SDMMC_Periph) FIFODATA() FIFO_Mask {
	return FIFO_Mask{mmio.UM32{&p.FIFO.U32, uint32(FIFODATA)}}
}
