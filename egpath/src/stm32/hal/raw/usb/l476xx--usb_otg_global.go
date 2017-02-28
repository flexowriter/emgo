// +build l476xx

// Peripheral: USB_OTG_Global_Periph  USB_OTG_Core_register.
// Instances:
// Registers:
//  0x00  32  GOTGCTL            USB_OTG Control and Status Register          000h.
//  0x04  32  GOTGINT            USB_OTG Interrupt Register                   004h.
//  0x08  32  GAHBCFG            Core AHB Configuration Register              008h.
//  0x0C  32  GUSBCFG            Core USB Configuration Register              00Ch.
//  0x10  32  GRSTCTL            Core Reset Register                          010h.
//  0x14  32  GINTSTS            Core Interrupt Register                      014h.
//  0x18  32  GINTMSK            Core Interrupt Mask Register                 018h.
//  0x1C  32  GRXSTSR            Receive Sts Q Read Register                  01Ch.
//  0x20  32  GRXSTSP            Receive Sts Q Read & POP Register            020h.
//  0x24  32  GRXFSIZ            Receive FIFO Size Register                      024h.
//  0x28  32  DIEPTXF0_HNPTXFSIZ EP0 / Non Periodic Tx FIFO Size Register     028h.
//  0x2C  32  HNPTXSTS           Non Periodic Tx FIFO/Queue Sts reg           02Ch.
//  0x38  32  GCCFG              General Purpose IO Register                     038h.
//  0x3C  32  CID                User ID Register                                03Ch.
//  0x40  32  GSNPSID            USB_OTG core ID                                 040h.
//  0x44  32  GHWCFG1            User HW config1                                 044h.
//  0x48  32  GHWCFG2            User HW config2                                 048h.
//  0x4C  32  GHWCFG3            User HW config3                                 04Ch.
//  0x54  32  GLPMCFG            LPM Register                                    054h.
//  0x58  32  GPWRDN             Power Down Register                             058h.
//  0x5C  32  GDFIFOCFG          DFIFO Software Config Register                  05Ch.
//  0x60  32  GADPCTL            ADP Timer, Control and Status Register          60Ch.
//  0x100 32  HPTXFSIZ           Host Periodic Tx FIFO Size Reg                  100h.
//  0x104 32  DIEPTXF[15]        dev Periodic Transmit FIFO.
// Import:
//  stm32/o/l476xx/mmap
package usb

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	SRQSCS    GOTGCTL_Bits = 0x01 << 0  //+ Session request success.
	SRQ       GOTGCTL_Bits = 0x01 << 1  //+ Session request.
	VBVALOEN  GOTGCTL_Bits = 0x01 << 2  //+ VBUS valid override enable.
	VBVALOVAL GOTGCTL_Bits = 0x01 << 3  //+ VBUS valid override value.
	AVALOEN   GOTGCTL_Bits = 0x01 << 4  //+ A-peripheral session valid override enable.
	AVALOVAL  GOTGCTL_Bits = 0x01 << 5  //+ A-peripheral session valid override value.
	BVALOEN   GOTGCTL_Bits = 0x01 << 6  //+ B-peripheral session valid override enable.
	BVALOVAL  GOTGCTL_Bits = 0x01 << 7  //+ B-peripheral session valid override value.
	BSESVLD   GOTGCTL_Bits = 0x01 << 19 //+ B-session valid.
)

const (
	SRQSCSn    = 0
	SRQn       = 1
	VBVALOENn  = 2
	VBVALOVALn = 3
	AVALOENn   = 4
	AVALOVALn  = 5
	BVALOENn   = 6
	BVALOVALn  = 7
	BSESVLDn   = 19
)

const (
	SEDET   GOTGINT_Bits = 0x01 << 2  //+ Session end detected.
	SRSSCHG GOTGINT_Bits = 0x01 << 8  //+ Session request success status change.
	HNSSCHG GOTGINT_Bits = 0x01 << 9  //+ Host negotiation success status change.
	HNGDET  GOTGINT_Bits = 0x01 << 17 //+ Host negotiation detected.
	ADTOCHG GOTGINT_Bits = 0x01 << 18 //+ A-device timeout change.
	DBCDNE  GOTGINT_Bits = 0x01 << 19 //+ Debounce done.
)

const (
	SEDETn   = 2
	SRSSCHGn = 8
	HNSSCHGn = 9
	HNGDETn  = 17
	ADTOCHGn = 18
	DBCDNEn  = 19
)

const (
	GINT     GAHBCFG_Bits = 0x01 << 0 //+ Global interrupt mask.
	HBSTLEN  GAHBCFG_Bits = 0x0F << 1 //+ Burst length/type.
	DMAEN    GAHBCFG_Bits = 0x01 << 5 //+ DMA enable.
	TXFELVL  GAHBCFG_Bits = 0x01 << 7 //+ TxFIFO empty level.
	PTXFELVL GAHBCFG_Bits = 0x01 << 8 //+ Periodic TxFIFO empty level.
)

const (
	GINTn     = 0
	HBSTLENn  = 1
	DMAENn    = 5
	TXFELVLn  = 7
	PTXFELVLn = 8
)

const (
	TOCAL      GUSBCFG_Bits = 0x07 << 0  //+ FS timeout calibration.
	PHYSEL     GUSBCFG_Bits = 0x01 << 6  //+ USB 2.0 high-speed ULPI PHY or USB 1.1 full-speed serial transceiver select.
	SRPCAP     GUSBCFG_Bits = 0x01 << 8  //+ SRP-capable.
	HNPCAP     GUSBCFG_Bits = 0x01 << 9  //+ HNP-capable.
	TRDT       GUSBCFG_Bits = 0x0F << 10 //+ USB turnaround time.
	PHYLPCS    GUSBCFG_Bits = 0x01 << 15 //+ PHY Low-power clock select.
	ULPIFSLS   GUSBCFG_Bits = 0x01 << 17 //+ ULPI FS/LS select.
	ULPIAR     GUSBCFG_Bits = 0x01 << 18 //+ ULPI Auto-resume.
	ULPICSM    GUSBCFG_Bits = 0x01 << 19 //+ ULPI Clock SuspendM.
	ULPIEVBUSD GUSBCFG_Bits = 0x01 << 20 //+ ULPI External VBUS Drive.
	ULPIEVBUSI GUSBCFG_Bits = 0x01 << 21 //+ ULPI external VBUS indicator.
	TSDPS      GUSBCFG_Bits = 0x01 << 22 //+ TermSel DLine pulsing selection.
	PCCI       GUSBCFG_Bits = 0x01 << 23 //+ Indicator complement.
	PTCI       GUSBCFG_Bits = 0x01 << 24 //+ Indicator pass through.
	ULPIIPD    GUSBCFG_Bits = 0x01 << 25 //+ ULPI interface protect disable.
	FHMOD      GUSBCFG_Bits = 0x01 << 29 //+ Forced host mode.
	FDMOD      GUSBCFG_Bits = 0x01 << 30 //+ Forced peripheral mode.
	CTXPKT     GUSBCFG_Bits = 0x01 << 31 //+ Corrupt Tx packet.
)

const (
	TOCALn      = 0
	PHYSELn     = 6
	SRPCAPn     = 8
	HNPCAPn     = 9
	TRDTn       = 10
	PHYLPCSn    = 15
	ULPIFSLSn   = 17
	ULPIARn     = 18
	ULPICSMn    = 19
	ULPIEVBUSDn = 20
	ULPIEVBUSIn = 21
	TSDPSn      = 22
	PCCIn       = 23
	PTCIn       = 24
	ULPIIPDn    = 25
	FHMODn      = 29
	FDMODn      = 30
	CTXPKTn     = 31
)

const (
	CSRST   GRSTCTL_Bits = 0x01 << 0  //+ Core soft reset.
	HSRST   GRSTCTL_Bits = 0x01 << 1  //+ HCLK soft reset.
	FCRST   GRSTCTL_Bits = 0x01 << 2  //+ Host frame counter reset.
	RXFFLSH GRSTCTL_Bits = 0x01 << 4  //+ RxFIFO flush.
	TXFFLSH GRSTCTL_Bits = 0x01 << 5  //+ TxFIFO flush.
	TXFNUM  GRSTCTL_Bits = 0x1F << 6  //+ TxFIFO number.
	DMAREQ  GRSTCTL_Bits = 0x01 << 30 //+ DMA request signal.
	AHBIDL  GRSTCTL_Bits = 0x01 << 31 //+ AHB master idle.
)

const (
	CSRSTn   = 0
	HSRSTn   = 1
	FCRSTn   = 2
	RXFFLSHn = 4
	TXFFLSHn = 5
	TXFNUMn  = 6
	DMAREQn  = 30
	AHBIDLn  = 31
)

const (
	CMOD              GINTSTS_Bits = 0x01 << 0  //+ Current mode of operation.
	MMIS              GINTSTS_Bits = 0x01 << 1  //+ Mode mismatch interrupt.
	OTGINT            GINTSTS_Bits = 0x01 << 2  //+ OTG interrupt.
	SOF               GINTSTS_Bits = 0x01 << 3  //+ Start of frame.
	RXFLVL            GINTSTS_Bits = 0x01 << 4  //+ RxFIFO nonempty.
	NPTXFE            GINTSTS_Bits = 0x01 << 5  //+ Nonperiodic TxFIFO empty.
	GINAKEFF          GINTSTS_Bits = 0x01 << 6  //+ Global IN nonperiodic NAK effective.
	BOUTNAKEFF        GINTSTS_Bits = 0x01 << 7  //+ Global OUT NAK effective.
	ESUSP             GINTSTS_Bits = 0x01 << 10 //+ Early suspend.
	USBSUSP           GINTSTS_Bits = 0x01 << 11 //+ USB suspend.
	USBRST            GINTSTS_Bits = 0x01 << 12 //+ USB reset.
	ENUMDNE           GINTSTS_Bits = 0x01 << 13 //+ Enumeration done.
	ISOODRP           GINTSTS_Bits = 0x01 << 14 //+ Isochronous OUT packet dropped interrupt.
	EOPF              GINTSTS_Bits = 0x01 << 15 //+ End of periodic frame interrupt.
	IEPINT            GINTSTS_Bits = 0x01 << 18 //+ IN endpoint interrupt.
	OEPINT            GINTSTS_Bits = 0x01 << 19 //+ OUT endpoint interrupt.
	IISOIXFR          GINTSTS_Bits = 0x01 << 20 //+ Incomplete isochronous IN transfer.
	PXFR_INCOMPISOOUT GINTSTS_Bits = 0x01 << 21 //+ Incomplete periodic transfer.
	DATAFSUSP         GINTSTS_Bits = 0x01 << 22 //+ Data fetch suspended.
	HPRTINT           GINTSTS_Bits = 0x01 << 24 //+ Host port interrupt.
	HCINT             GINTSTS_Bits = 0x01 << 25 //+ Host channels interrupt.
	PTXFE             GINTSTS_Bits = 0x01 << 26 //+ Periodic TxFIFO empty.
	LPMINT            GINTSTS_Bits = 0x01 << 27 //+ LPM interrupt.
	CIDSCHG           GINTSTS_Bits = 0x01 << 28 //+ Connector ID status change.
	DISCINT           GINTSTS_Bits = 0x01 << 29 //+ Disconnect detected interrupt.
	SRQINT            GINTSTS_Bits = 0x01 << 30 //+ Session request/new session detected interrupt.
	WKUINT            GINTSTS_Bits = 0x01 << 31 //+ Resume/remote wakeup detected interrupt.
)

const (
	CMODn              = 0
	MMISn              = 1
	OTGINTn            = 2
	SOFn               = 3
	RXFLVLn            = 4
	NPTXFEn            = 5
	GINAKEFFn          = 6
	BOUTNAKEFFn        = 7
	ESUSPn             = 10
	USBSUSPn           = 11
	USBRSTn            = 12
	ENUMDNEn           = 13
	ISOODRPn           = 14
	EOPFn              = 15
	IEPINTn            = 18
	OEPINTn            = 19
	IISOIXFRn          = 20
	PXFR_INCOMPISOOUTn = 21
	DATAFSUSPn         = 22
	HPRTINTn           = 24
	HCINTn             = 25
	PTXFEn             = 26
	LPMINTn            = 27
	CIDSCHGn           = 28
	DISCINTn           = 29
	SRQINTn            = 30
	WKUINTn            = 31
)

const (
	MMISM           GINTMSK_Bits = 0x01 << 1  //+ Mode mismatch interrupt mask.
	OTGINT          GINTMSK_Bits = 0x01 << 2  //+ OTG interrupt mask.
	SOFM            GINTMSK_Bits = 0x01 << 3  //+ Start of frame mask.
	RXFLVLM         GINTMSK_Bits = 0x01 << 4  //+ Receive FIFO nonempty mask.
	NPTXFEM         GINTMSK_Bits = 0x01 << 5  //+ Nonperiodic TxFIFO empty mask.
	GINAKEFFM       GINTMSK_Bits = 0x01 << 6  //+ Global nonperiodic IN NAK effective mask.
	GONAKEFFM       GINTMSK_Bits = 0x01 << 7  //+ Global OUT NAK effective mask.
	ESUSPM          GINTMSK_Bits = 0x01 << 10 //+ Early suspend mask.
	USBSUSPM        GINTMSK_Bits = 0x01 << 11 //+ USB suspend mask.
	USBRST          GINTMSK_Bits = 0x01 << 12 //+ USB reset mask.
	ENUMDNEM        GINTMSK_Bits = 0x01 << 13 //+ Enumeration done mask.
	ISOODRPM        GINTMSK_Bits = 0x01 << 14 //+ Isochronous OUT packet dropped interrupt mask.
	EOPFM           GINTMSK_Bits = 0x01 << 15 //+ End of periodic frame interrupt mask.
	EPMISM          GINTMSK_Bits = 0x01 << 17 //+ Endpoint mismatch interrupt mask.
	IEPINT          GINTMSK_Bits = 0x01 << 18 //+ IN endpoints interrupt mask.
	OEPINT          GINTMSK_Bits = 0x01 << 19 //+ OUT endpoints interrupt mask.
	IISOIXFRM       GINTMSK_Bits = 0x01 << 20 //+ Incomplete isochronous IN transfer mask.
	PXFRM_IISOOXFRM GINTMSK_Bits = 0x01 << 21 //+ Incomplete periodic transfer mask.
	FSUSPM          GINTMSK_Bits = 0x01 << 22 //+ Data fetch suspended mask.
	PRTIM           GINTMSK_Bits = 0x01 << 24 //+ Host port interrupt mask.
	HCIM            GINTMSK_Bits = 0x01 << 25 //+ Host channels interrupt mask.
	PTXFEM          GINTMSK_Bits = 0x01 << 26 //+ Periodic TxFIFO empty mask.
	LPMINTM         GINTMSK_Bits = 0x01 << 27 //+ LPM interrupt Mask.
	CIDSCHGM        GINTMSK_Bits = 0x01 << 28 //+ Connector ID status change mask.
	DISCINT         GINTMSK_Bits = 0x01 << 29 //+ Disconnect detected interrupt mask.
	SRQIM           GINTMSK_Bits = 0x01 << 30 //+ Session request/new session detected interrupt mask.
	WUIM            GINTMSK_Bits = 0x01 << 31 //+ Resume/remote wakeup detected interrupt mask.
)

const (
	MMISMn           = 1
	OTGINTn          = 2
	SOFMn            = 3
	RXFLVLMn         = 4
	NPTXFEMn         = 5
	GINAKEFFMn       = 6
	GONAKEFFMn       = 7
	ESUSPMn          = 10
	USBSUSPMn        = 11
	USBRSTn          = 12
	ENUMDNEMn        = 13
	ISOODRPMn        = 14
	EOPFMn           = 15
	EPMISMn          = 17
	IEPINTn          = 18
	OEPINTn          = 19
	IISOIXFRMn       = 20
	PXFRM_IISOOXFRMn = 21
	FSUSPMn          = 22
	PRTIMn           = 24
	HCIMn            = 25
	PTXFEMn          = 26
	LPMINTMn         = 27
	CIDSCHGMn        = 28
	DISCINTn         = 29
	SRQIMn           = 30
	WUIMn            = 31
)

const (
	EPNUM  GRXSTSP_Bits = 0x0F << 0  //+ IN EP interrupt mask bits.
	BCNT   GRXSTSP_Bits = 0x7FF << 4 //+ OUT EP interrupt mask bits.
	DPID   GRXSTSP_Bits = 0x03 << 15 //+ OUT EP interrupt mask bits.
	PKTSTS GRXSTSP_Bits = 0x0F << 17 //+ OUT EP interrupt mask bits.
)

const (
	EPNUMn  = 0
	BCNTn   = 4
	DPIDn   = 15
	PKTSTSn = 17
)

const (
	RXFD GRXFSIZ_Bits = 0xFFFF << 0 //+ RxFIFO depth.
)

const (
	RXFDn = 0
)

const (
	DCDET  GCCFG_Bits = 0x01 << 0  //+ Data contact detection (DCD) status.
	PDET   GCCFG_Bits = 0x01 << 1  //+ Primary detection (PD) status.
	SDET   GCCFG_Bits = 0x01 << 2  //+ Secondary detection (SD) status.
	PS2DET GCCFG_Bits = 0x01 << 3  //+ DM pull-up detection status.
	PWRDWN GCCFG_Bits = 0x01 << 16 //+ Power down.
	BCDEN  GCCFG_Bits = 0x01 << 17 //+ Battery charging detector (BCD) enable.
	DCDEN  GCCFG_Bits = 0x01 << 18 //+ Data contact detection (DCD) mode enable.
	PDEN   GCCFG_Bits = 0x01 << 19 //+ Primary detection (PD) mode enable.
	SDEN   GCCFG_Bits = 0x01 << 20 //+ Secondary detection (SD) mode enable.
	VBDEN  GCCFG_Bits = 0x01 << 21 //+ Secondary detection (SD) mode enable.
)

const (
	DCDETn  = 0
	PDETn   = 1
	SDETn   = 2
	PS2DETn = 3
	PWRDWNn = 16
	BCDENn  = 17
	DCDENn  = 18
	PDENn   = 19
	SDENn   = 20
	VBDENn  = 21
)

const (
	PRODUCT_ID CID_Bits = 0xFFFFFFFF << 0 //+ Product ID field.
)

const (
	PRODUCT_IDn = 0
)

const (
	LPMMode GHWCFG3_Bits = 0x01 << 14 //+ LPM mode specified for Mode of Operation.
)

const (
	LPMModen = 14
)

const (
	ENBESL     GLPMCFG_Bits = 0x01 << 28 //+ Enable best effort service latency.
	LPMRCNTSTS GLPMCFG_Bits = 0x07 << 25 //+ LPM retry count status.
	SNDLPM     GLPMCFG_Bits = 0x01 << 24 //+ Send LPM transaction.
	LPMRCNT    GLPMCFG_Bits = 0x07 << 21 //+ LPM retry count.
	LPMCHIDX   GLPMCFG_Bits = 0x0F << 17 //+ LPMCHIDX:.
	L1ResumeOK GLPMCFG_Bits = 0x01 << 16 //+ Sleep State Resume OK.
	SLPSTS     GLPMCFG_Bits = 0x01 << 15 //+ Port sleep status.
	LPMRSP     GLPMCFG_Bits = 0x03 << 13 //+ LPM response.
	L1DSEN     GLPMCFG_Bits = 0x01 << 12 //+ L1 deep sleep enable.
	BESLTHRS   GLPMCFG_Bits = 0x0F << 8  //+ BESL threshold.
	L1SSEN     GLPMCFG_Bits = 0x01 << 7  //+ L1 shallow sleep enable.
	REMWAKE    GLPMCFG_Bits = 0x01 << 6  //+ bRemoteWake value received with last ACKed LPM Token.
	BESL       GLPMCFG_Bits = 0x0F << 2  //+ BESL value received with last ACKed LPM Token.
	LPMACK     GLPMCFG_Bits = 0x01 << 1  //+ LPM Token acknowledge enable.
	LPMEN      GLPMCFG_Bits = 0x01 << 0  //+ LPM support enable.
)

const (
	ENBESLn     = 28
	LPMRCNTSTSn = 25
	SNDLPMn     = 24
	LPMRCNTn    = 21
	LPMCHIDXn   = 17
	L1ResumeOKn = 16
	SLPSTSn     = 15
	LPMRSPn     = 13
	L1DSENn     = 12
	BESLTHRSn   = 8
	L1SSENn     = 7
	REMWAKEn    = 6
	BESLn       = 2
	LPMACKn     = 1
	LPMENn      = 0
)

const (
	DISABLEVBUS GPWRDN_Bits = 0x01 << 6 //+ Power down.
)

const (
	DISABLEVBUSn = 6
)

const (
	PTXSA HPTXFSIZ_Bits = 0xFFFF << 0  //+ Host periodic TxFIFO start address.
	PTXFD HPTXFSIZ_Bits = 0xFFFF << 16 //+ Host periodic TxFIFO depth.
)

const (
	PTXSAn = 0
	PTXFDn = 16
)

const (
	INEPTXSA DIEPTXF_Bits = 0xFFFF << 0  //+ IN endpoint FIFOx transmit RAM start address.
	INEPTXFD DIEPTXF_Bits = 0xFFFF << 16 //+ IN endpoint TxFIFO depth.
)

const (
	INEPTXSAn = 0
	INEPTXFDn = 16
)
