package ni488

/*
#cgo linux LDFLAGS: -lgpibapi
#cgo darwin CFLAGS: -I.
#cgo darwin LDFLAGS: -framework NI488
#cgo windows CFLAGS: -I.
#cgo windows LDFLAGS: -lgpib-32 -LC:/WINDOWS/system32
#include <stdlib.h>
#include <ni488.h>
*/
import "C"
import "unsafe"

const (
	UNL = C.UNL // GPIB unlisten command
	UNT = C.UNT // GPIB untalk command
	GTL = C.GTL // GPIB go to local
	SDC = C.SDC // GPIB selected device clear
	PPC = C.PPC // GPIB parallel poll configure
	GET = C.GET // GPIB group execute trigger
	TCT = C.TCT // GPIB take control
	LLO = C.LLO // GPIB local lock out
	DCL = C.DCL // GPIB device clear
	PPU = C.PPU // GPIB parallel poll unconfigure
	SPE = C.SPE // GPIB serial poll enable
	SPD = C.SPD // GPIB serial poll disable
	PPE = C.PPE // GPIB parallel poll enable
	PPD = C.PPD // GPIB parallel poll disable

	// GPIB status bit vector: global variable ibsta and wait mask
	ERR  = C.ERR  // Error detected
	TIMO = C.TIMO // Timeout
	END  = C.END  // EOI or EOS detected
	SRQI = C.SRQI // SRQ detected by CIC
	RQS  = C.RQS  // Device needs service
	CMPL = C.CMPL // I/O completed
	LOK  = C.LOK  // Local lockout state
	REM  = C.REM  // Remote state
	CIC  = C.CIC  // Controller-in-Charge
	ATN  = C.ATN  // Attention asserted
	TACS = C.TACS // Talker active
	LACS = C.LACS // Listener active
	DTAS = C.DTAS // Device trigger state
	DCAS = C.DCAS // Device clear state

	// Error messages returned in global variable iberr
	EDVR = C.EDVR // System error
	ECIC = C.ECIC // Function requires GPIB board to be CIC
	ENOL = C.ENOL // Write function detected no Listeners
	EADR = C.EADR // Interface board not addressed correctly
	EARG = C.EARG // Invalid argument to function call
	ESAC = C.ESAC // Function requires GPIB board to be SAC
	EABO = C.EABO // I/O operation aborted
	ENEB = C.ENEB // Non-existent interface board
	EDMA = C.EDMA // Error performing DMA
	EOIP = C.EOIP // I/O operation started before previous

	// Operation completed
	ECAP = C.ECAP // No capability for intended operation
	EFSO = C.EFSO // File system operation error
	EBUS = C.EBUS // Command error during device call
	ESTB = C.ESTB // Serial poll status byte lost
	ESRQ = C.ESRQ // SRQ remains asserted
	ETAB = C.ETAB // The return buffer is full.
	ELCK = C.ELCK // Address or board is locked.
	EARM = C.EARM // The ibnotify Callback failed to rearm
	EHDL = C.EHDL // The input handle is invalid
	EWIP = C.EWIP // Wait already in progress on input ud
	ERST = C.ERST // The event notification was cancelled due to a reset of the interface
	EPWR = C.EPWR // The system or board has lost power or gone to standby

	// Warning messages returned in global variable iberr
	WCFG = C.WCFG // Configuration warning
	ECFG = C.ECFG

	// EOS mode bits
	BIN  = C.BIN  // Eight bit compare
	XEOS = C.XEOS // Send END with EOS byte
	REOS = C.REOS // Terminate read on EOS

	// Timeout values and meanings
	TNONE  = C.TNONE  // Infinite timeout (disabled)
	T10us  = C.T10us  // Timeout of 10 us (ideal)
	T30us  = C.T30us  // Timeout of 30 us (ideal)
	T100us = C.T100us // Timeout of 100 us (ideal)
	T300us = C.T300us // Timeout of 300 us (ideal)
	T1ms   = C.T1ms   // Timeout of 1 ms (ideal)
	T3ms   = C.T3ms   // Timeout of 3 ms (ideal)
	T10ms  = C.T10ms  // Timeout of 10 ms (ideal)
	T30ms  = C.T30ms  // Timeout of 30 ms (ideal)
	T100ms = C.T100ms // Timeout of 100 ms (ideal)
	T300ms = C.T300ms // Timeout of 300 ms (ideal)
	T1s    = C.T1s    // Timeout of 1 s (ideal)
	T3s    = C.T3s    // Timeout of 3 s (ideal)
	T10s   = C.T10s   // Timeout of 10 s (ideal)
	T30s   = C.T30s   //  Timeout of 30 s (ideal)
	T100s  = C.T100s  // Timeout of 100 s (ideal)
	T300s  = C.T300s  // Timeout of 300 s (ideal)
	T1000s = C.T1000s // Timeout of 1000 s (ideal)

	// IBLN Constants
	NO_SAD  = C.NO_SAD
	ALL_SAD = C.ALL_SAD

	// Constants used for the second parameter of the ibconfig function.
	// They are the "option" selection codes.
	IbcPAD            = C.IbcPAD            // Primary Address
	IbcSAD            = C.IbcSAD            // Secondary Address
	IbcTMO            = C.IbcTMO            // Timeout Value
	IbcEOT            = C.IbcEOT            // Send EOI with last data byte?
	IbcPPC            = C.IbcPPC            // Parallel Poll Configure
	IbcREADDR         = C.IbcREADDR         // Repeat Addressing
	IbcAUTOPOLL       = C.IbcAUTOPOLL       // Disable Auto Serial Polling
	IbcCICPROT        = C.IbcCICPROT        // Use the CIC Protocol?
	IbcIRQ            = C.IbcIRQ            // Use PIO for I/O
	IbcSC             = C.IbcSC             // Board is System Controller?
	IbcSRE            = C.IbcSRE            // Assert SRE on device calls?
	IbcEOSrd          = C.IbcEOSrd          // Terminate reads on EOS
	IbcEOSwrt         = C.IbcEOSwrt         // Send EOI with EOS character
	IbcEOScmp         = C.IbcEOScmp         // Use 7 or 8-bit EOS compare
	IbcEOSchar        = C.IbcEOSchar        // The EOS character.
	IbcPP2            = C.IbcPP2            // Use Parallel Poll Mode 2.
	IbcTIMING         = C.IbcTIMING         // NORMAL, HIGH, or VERY_HIGH timing.
	IbcDMA            = C.IbcDMA            // Use DMA for I/O
	IbcReadAdjust     = C.IbcReadAdjust     // Swap bytes during an ibrd.
	IbcWriteAdjust    = C.IbcWriteAdjust    // Swap bytes during an ibwrt.
	IbcSendLLO        = C.IbcSendLLO        // Enable/disable the sending of LLO.
	IbcSPollTime      = C.IbcSPollTime      // Set the timeout value for serial polls.
	IbcPPollTime      = C.IbcPPollTime      // Set the parallel poll length period.
	IbcEndBitIsNormal = C.IbcEndBitIsNormal // Remove EOS from END bit of IBSTA.
	IbcUnAddr         = C.IbcUnAddr         // Enable/disable device unaddressing.
	IbcSignalNumber   = C.IbcSignalNumber   // Set UNIX signal number - unsupported
	IbcBlockIfLocked  = C.IbcBlockIfLocked  // Enable/disable blocking for locked boards/
	IbcHSCableLength  = C.IbcHSCableLength  // Length of cable specified for high speed t
	IbcIst            = C.IbcIst            // Set the IST bit.
	IbcRsv            = C.IbcRsv            // Set the RSV byte.
	IbcLON            = C.IbcLON            // Enter listen only mode

	// Constants that can be used (in addition to the ibconfig constants)
	// when calling the ibask() function.
	IbaPAD            = C.IbaPAD
	IbaSAD            = C.IbaSAD
	IbaTMO            = C.IbaTMO
	IbaEOT            = C.IbaEOT
	IbaPPC            = C.IbaPPC
	IbaREADDR         = C.IbaREADDR
	IbaAUTOPOLL       = C.IbaAUTOPOLL
	IbaCICPROT        = C.IbaCICPROT
	IbaIRQ            = C.IbaIRQ
	IbaSC             = C.IbaSC
	IbaSRE            = C.IbaSRE
	IbaEOSrd          = C.IbaEOSrd
	IbaEOSwrt         = C.IbaEOSwrt
	IbaEOScmp         = C.IbaEOScmp
	IbaEOSchar        = C.IbaEOSchar
	IbaPP2            = C.IbaPP2
	IbaTIMING         = C.IbaTIMING
	IbaDMA            = C.IbaDMA
	IbaReadAdjust     = C.IbaReadAdjust
	IbaWriteAdjust    = C.IbaWriteAdjust
	IbaSendLLO        = C.IbaSendLLO
	IbaSPollTime      = C.IbaSPollTime
	IbaPPollTime      = C.IbaPPollTime
	IbaEndBitIsNormal = C.IbaEndBitIsNormal
	IbaUnAddr         = C.IbaUnAddr
	IbaSignalNumber   = C.IbaSignalNumber
	IbaBlockIfLocked  = C.IbaBlockIfLocked
	IbaHSCableLength  = C.IbaHSCableLength
	IbaIst            = C.IbaIst
	IbaRsv            = C.IbaRsv
	IbaLON            = C.IbaLON
	IbaSerialNumber   = C.IbaSerialNumber
	IbaBNA            = C.IbaBNA // A device's access board.

	// Values used by the Send 488.2 command.
	NULLend = C.NULLend // Do nothing at the end of a transfer.
	NLend   = C.NLend   // Send NL with EOI after a transfer.
	DABend  = C.DABend  // Send EOI with the last DAB.

	// Value used by the 488.2 Receive command.
	STOPend = C.STOPend

	// Terminates an address list
	NOADDR = C.NOADDR

	ValidEOI                 = C.ValidEOI
	ValidATNV                = C.ValidATN
	ValidSRQ                 = C.ValidSRQ
	ValidREN                 = C.ValidREN
	ValidIFC                 = C.ValidIFC
	ValidNRFD                = C.ValidNRFD
	ValidNDAC                = C.ValidNDAC
	ValidDAV                 = C.ValidDAV
	BusEOI                   = C.BusEOI
	BusSRQ                   = C.BusSRQ
	BusREN                   = C.BusREN
	BusIFC                   = C.BusIFC
	BusNRFD                  = C.BusNRFD
	BusNDAC                  = C.BusNDAC
	BusDAV                   = C.BusDAV
	TIMMEDIATE               = C.TIMMEDIATE
	TINFINITE                = C.TINFINITE
	MAX_LOCKSHARENAME_LENGTH = C.MAX_LOCKSHARENAME_LENGTH
)

// Ibbn assigns the device described by ud to the access board described by
// bname.
//
// All subsequent bus activity with device ud occurs through the access board
// bname. If the call succeeds iberr contains the previous access board index.
func Ibbn(ud int, udname string) (ibsta int) {
	n := C.CString(udname)
	defer C.free(unsafe.Pointer(n))
	ibsta = int(C.ibbnaA(C.int(ud), n))
	return
}

// Ibeos configures the EOS termination mode or EOS character for the board
// or device.
//
// The parameter v describes the new end-of-string (EOS)
// configuration to use. If v is zero, then the EOS configuration is
// disabled. Otherwise, the low byte is the EOS character and the upper
// byte contains flags which define the EOS mode.
func Ibeos(ud, v int) (ibsta int) {
	ibsta = int(C.ibeos(C.int(ud), C.int(v)))
	return
}
