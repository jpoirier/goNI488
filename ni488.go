// Copyright (c) 2011 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

// Package ni488 is a wrapper around the NI-488.2 interface which allows
// communications with GPIB capable teting equipment. NI-488.2 is an industry
// standard for GPIB communications.
//
// The package is low level and, for the most part, is one-to-one with the
// exported C functions it wraps. Clients would typically build instrument
// drivers around the package but it can also be used directly.
//
// Lots of miscellaneous NI-488.2 information:
//     http://sine.ni.com/psp/app/doc/p/id/psp-356
//
// GPIB Driver Versions for Microsoft Windows and DOS:
//     http://zone.ni.com/devzone/cda/tut/p/id/5326#toc0
//
// GPIB Driver Versions for non-Microsoft Operating Systems:
//     http://zone.ni.com/devzone/cda/tut/p/id/5458
//
package ni488

// TODO:
// -

// #cgo linux CFLAGS: -arch i386
// #cgo linux LDFLAGS: -lgpibapi
// #cgo darwin CFLAGS: -arch i386 -I/Library/Frameworks/NI488.framework/Headers
// #cgo darwin LDFLAGS: -framework NI488
// #cgo windows CFLAGS: -m32
// #cgo windows LDFLAGS: -lgpib-32.dll
// #include <ni488.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

var PackageVersion string = "v0.2"

type Addr4882_t C.Addr4882_t

//    HANDY CONSTANTS FOR USE BY APPLICATION PROGRAMS ...

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

	// GPIB status bit vector :
	//       global variable ibsta and wait mask
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
	// operation completed
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
	ERST = C.ERST // The event notification was cancelled
	// due to a reset of the interface
	EPWR = C.EPWR // The system or board has lost power or
	// gone to standby

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

	// The following constants are used for the second parameter of the
	// ibconfig function.  They are the "option" selection codes.
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
	IbcBlockIfLocked  = C.IbcBlockIfLocked  // Enable/disable blocking for locked boards/devices
	IbcHSCableLength  = C.IbcHSCableLength  // Length of cable specified for high speed timing.
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

	NOADDR = C.NOADDR // Terminates an address list
)

//  Functions to access Thread-Specific copies of the GPIB global vars

// The global variables ibsta, iberr, ibcnt, and ibcntl are maintained on a
// process-specific rather than a thread-specific basis. If you call GPIB
// functions in more than one thread, the values in these global variables
// are not always reliable.

// Status variables analogous to ibsta, iberr, ibcnt, and ibcntl are maintained
// for each thread. ThreadIbcntl returns the value of the thread-specific ibcntl
// variable.

// Returns the thread-specific ibsta value for the current thread.
//
// The return value is the value for the current thread of execution. The
// value describes the state of the GPIB and the result of the most recent
// GPIB function call in the thread.  Call ThreadIberr for a specific error
// code.
func ThreadIbsta() int {
	return int(C.ThreadIbsta())
}

// Returns the thread-specific iberr value for the current thread.
//
// The return value is the most recent GPIB error code for the current
// thread of execution. The value is meaningful only when ThreadIbsta returns
// a value with the ERR bit set.
func ThreadIberr() int {
	return int(C.ThreadIberr())
}

// Returns the thread-specific ibcnt value for the current thread.
//
// The return value is either the number of bytes actually transferred by
// the most recent GPIB read, write, or command operation for the current
// thread of execution or an error code if an error occured.
func ThreadIbcnt() int {
	return int(C.ThreadIbcnt())
}

// Returns the thread-specific ibcntl value for the current thread.
//
// The return value is either the number of bytes actually transferred by
// the most recent GPIB read, write, or command operation for the current
// thread of execution or an error code if an error occured.
func ThreadIbcntl() int {
	return int(C.ThreadIbcntl())
}

//  NI-488 Functions

// Open and initialize a board or a user-configured device descriptor.
//
// Ibfind performs the equivalent of an ibonl 1 to initialize the board or
// device descriptor. The unit descriptor returned by ibfind remains valid
// until the board or device is put offline using ibonl 0.
//
// If ibfind is unable to get a valid descriptor, a â1 is returned; the ERR
// bit is set in ibsta and iberr contains EDVR.
func Ibfind(udname string) (ud int) {
	n := C.CString(udname)
	defer C.free(unsafe.Pointer(n))
	ud = int(C.ibfindA(n))
	return
}

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

// Ibrdf reads data asynchronously from a device into a user buffer.
//
// If ud is a device descriptor, ibrdf addresses the GPIB, reads data from
// a GPIB device, and places the data into the file specified by filename.
// If ud is a board descriptor, ibrdf reads data from a GPIB device and
// places the data into the file specified by filename
func Ibrdf(ud int, filename string) (ibsta int) {
	n := C.CString(filename)
	defer C.free(unsafe.Pointer(n))
	ibsta = int(C.ibrdfA(C.int(ud), n))
	return
}

// Ibwrtf writes data to a device from a file.
//
// If ud is a device descriptor, ibwrtf addresses the GPIB and writes all
// of the bytes from the file filename to a GPIB device. If ud is a board
// descriptor, ibwrtf writes all of the bytes of data from the file filename
// to a GPIB device.
func Ibwrtf(ud int, filename string) (ibsta int) {
	n := C.CString(filename)
	defer C.free(unsafe.Pointer(n))
	ibsta = int(C.ibwrtfA(C.int(ud), n))
	return
}

// Ibask returns the current value of various configuration parameters for the
// specified board or device.
//
// The current value of the selected configuration item is returned in v.
func Ibask(ud, option int) (ibsta, v int) {
	ibsta = int(C.ibask(C.int(ud), C.int(option),
		(*C.int)(unsafe.Pointer(&v))))
	return
}

// Ibcac uses the designated GPIB board to attempt to become the Active
// Controller by asserting ATN.
//
// If v is zero, the GPIB board takes control asynchronously and if v
// is non-zero the GPIB board takes control synchronously. Before calling
// ibcac, the GPIB board must already be CIC. To make the board CIC, use
// the ibsic function.
func Ibcac(ud, v int) (ibsta int) {
	ibsta = int(C.ibcac(C.int(ud), C.int(v)))
	return
}

// Ibclr sends the GPIB Selected Device Clear (SDC) message to the device
// described by ud.
func Ibclr(ud int) (ibsta int) {
	ibsta = int(C.ibclr(C.int(ud)))
	return
}

// Send GPIB commands.
//
// Sends cmds over the GPIB as command bytes (interface messages). The actual
// transferred byte count is returned in the global variable ibcntl.
func Ibcmd(ud int, cmds string) (ibsta int) {
	n := C.CString(cmds)
	defer C.free(unsafe.Pointer(n))
	ibsta = int(C.ibcmd(C.int(ud), unsafe.Pointer(n), C.long(len(cmds))))
	return
}

// Send GPIB commands asynchronously.
//
// Sends cmds asynchronously over the GPIB as command bytes (interface
// messages). The actual transferred byte count is returned in the global
// variable ibcntl.
func Ibcmda(ud int, cmds string) (ibsta int) {
	n := C.CString(cmds)
	defer C.free(unsafe.Pointer(n))
	ibsta = int(C.ibcmda(C.int(ud), unsafe.Pointer(n), C.long(len(cmds))))
	return
}

// Change software configuration parameters.
//
// Ibconfig changes a configuration item in option to the specified value in
// v for the selected board or device.
func Ibconfig(ud, option, v int) (ibsta int) {
	ibsta = int(C.ibconfig(C.int(ud), C.int(option), C.int(v)))
	return
}

// Open and initialize a device.
//
// Ibdev acquires a device descriptor to use in subsequent device-level NI-488
// functions. It opens and initializes a device descriptor, and configures
// it according to the input parameters. Returns the device descriptor or â1.
func Ibdev(boardID, pad, sad, tmo, eot, eos int) (dev int) {
	dev = int(C.ibdev(C.int(boardID), C.int(pad), C.int(sad),
		C.int(tmo), C.int(eot), C.int(eos)))
	return
}

// Enable or disable DMA.
//
// Ibdma enables or disables DMA transfers for the board, according to v.
// If v is zero, then DMA is not used for GPIB I/O transfers, and if v
// is non-zero, then DMA is used for GPIB I/O transfers.
func Ibdma(ud, v int) (ibsta int) {
	ibsta = int(C.ibdma(C.int(ud), C.int(v)))
	return
}

//extern int  ibexpert (int ud, int option, void * Input, void * Output);
//func Ibexpert() int {
//	return int(C.ibexpert())
//}

// Configure the end-of-string (EOS) termination mode or character.
//
// Ibeos configures the EOS termination mode or EOS character for the board
// or device. The parameter v describes the new end-of-string (EOS)
// configuration to use. If v is zero, then the EOS configuration is
// disabled. Otherwise, the low byte is the EOS character and the upper
// byte contains flags which define the EOS mode.
func Ibeos(ud, v int) (ibsta int) {
	ibsta = int(C.ibeos(C.int(ud), C.int(v)))
	return
}

// Enable or disable the automatic assertion of the GPIB EOI line at the end
// of write I/O operations.
//
// Ibeot enables or disables the assertion of the EOI line at the end of
// write I/O operations for the board or device described by ud. If v
// is non-zero, then EOI is asserted when the last byte of a GPIB
// write is sent.
func Ibeot(ud, v int) (ibsta int) {
	ibsta = int(C.ibeot(C.int(ud), C.int(v)))
	return
}

// Go from Active Controller to Standby.
//
// Ibgts causes the GPIB board at ud to go to Standby Controller and
// the GPIB ATN line to be unasserted. v Determines whether to perform
// acceptor handshaking
func Ibgts(ud, v int) (ibsta int) {
	ibsta = int(C.ibgts(C.int(ud), C.int(v)))
	return
}

// Set or clear the board individual status bit for parallel polls.
//
// Ibist sets the interface board ist (individual status) bit according to v.
func Ibist(ud, v int) (ibsta int) {
	ibsta = int(C.ibist(C.int(ud), C.int(v)))
	return
}

//extern int  iblck    (int ud, int v, unsigned int LockWaitTime, void * Reserved);
//// Acquire or release an exclusive interface lock
//func Iblck(ud, v int, LockWaitTime uint, Reserved) int {
//	return int(C.iblck(C.int(ud), C.int(v), C.uint(LockWaitTime), ))
//}

// Return the status of the eight GPIB control lines.
//
// Iblines returns the state of the GPIB control lines in clines.
// ud Board descriptor
// result Returns GPIB control line state information
// return The value of ibsta
func Iblines(ud int) (ibsta, result int) {
	ibsta = int(C.iblines(C.int(ud), (*C.short)(unsafe.Pointer(&result))))
	return
}

// Check for the presence of a device on the bus.
//
// Ibln determines whether there is a listening device at the GPIB address
// designated by the pad and sad parameters. If ud is a board descriptor,
// then the bus associated with that board is tested for Listeners.
// If ud is a device descriptor, then ibln uses the access board
// associated with that device to test for Listeners. If a Listener is
// detected, a non-zero value is returned in listen. If no Listener is
// found, zero is returned.
func Ibln(ud, pad, sad int) (ibsta, listen int) {
	ibsta = int(C.ibln(C.int(ud), C.int(pad), C.int(sad),
		(*C.short)(unsafe.Pointer(&listen))))
	return
}

// Go to local.
//
// Ibloc places the board in local mode if it is not in a lockout state.
func Ibloc(ud int) (ibsta int) {
	ibsta = int(C.ibloc(C.int(ud)))
	return
}

// Notify user of one or more GPIB events by invoking the user callback.
//
// Ibnotify installs an asynchronous callback function for a specified
// board or device. If mask is non-zero, ibnotify monitors the events
// specified by mask, and when one or more of the events is true, the
// Callback is invoked. refData User-defined reference data for the callback.
// int mycallback(int ud, int ibsta, int iberr, long ibcntl, void *RefData)
func Ibnotify(ud, mask int, f func(), redData []uint32) (ibsta int) {
	ibsta = int(C.ibnotify(C.int(ud), C.int(mask),
		(C.GpibNotifyCallback_t)(unsafe.Pointer(&f)),
		unsafe.Pointer(&redData[0])))
	return
}

// Place the device online or offline.
//
// Ibonl resets the board or device and places all its software configuration
// parameters in their pre-configured state. In addition, if v is zero,
// the device or interface board is taken offline. If v is non-zero, the
// device or interface board is left operational, or online.
// ud Board or device descriptor
func Ibonl(ud, v int) (ibsta int) {
	ibsta = int(C.ibonl(C.int(ud), C.int(v)))
	return
}

// Change the primary address.
//
// Ibpad sets the primary GPIB address of the board or device to v, an
// integer ranging from 0 to 30.
func Ibpad(ud, v int) (ibsta int) {
	ibsta = int(C.ibpad(C.int(ud), C.int(v)))
	return
}

// Pass control to another GPIB device with Controller capability.
//
// Ibpct passes Controller-in-Charge status to the device indicated by ud.
func Ibpct(ud int) (ibsta int) {
	ibsta = int(C.ibpct(C.int(ud)))
	return
}

// Parallel poll configure.
//
// If ud is a device descriptor, ibppc enables or disables the device
// from responding to parallel polls.
func Ibppc(ud, v int) (ibsta int) {
	ibsta = int(C.ibppc(C.int(ud), C.int(v)))
	return
}

// Read data asynchronously from a device into a user buffer.
//
// If ud is a device descriptor, ibrd addresses the GPIB, reads up to
// len(buf) bytes of data, and places the data into the buffer specified
// buf.
func Ibrd(ud int, buf []byte) (ibsta int) {
	ibsta = int(C.ibrd(C.int(ud), unsafe.Pointer(&buf[0]),
		C.long(len(buf))))
	return
}

// Conduct a parallel poll.
//
// If this routine is called specifying a GPIB Interface Board, the board
// parallel polls all previously configured devices. If the routine is called
// specifying a device, the GPIB Interface board associated with the device
// conducts the parallel poll. Note that if the GPIB Interface Board to conduct
// the parallel poll is not the Controller- In-Charge, an ECIC error is generated.
func Ibrpp(ud int) (ibsta, resp int) {
	ibsta = int(C.ibrpp(C.int(ud), (*C.char)(unsafe.Pointer(&resp))))
	return
}

// Request or release system control.
//
// Ibrsc requests or releases the capability to send Interface Clear (IFC)
// and Remote Enable (REN) messages to devices.
func Ibrsc(ud, v int) (ibsta int) {
	ibsta = int(C.ibrsc(C.int(ud), C.int(v)))
	return
}

// Conduct a serial poll.
//
// The ibrsp function is used to serial poll the device ud.
func Ibrsp(ud int) (ibsta, resp int) {
	ibsta = int(C.ibrsp(C.int(ud), (*C.char)(unsafe.Pointer(&resp))))
	return
}

// Request service and change the serial poll status byte.
//
// Ibrsv is used to request service from the Controller and to provide the
// Controller with an application-dependent status byte when the
// Controller serial polls the GPIB board.
func Ibrsv(ud, status int) (ibsta int) {
	ibsta = int(C.ibrsv(C.int(ud), C.int(status)))
	return
}

// Change or disable the secondary address.

// Ibsad changes the secondary GPIB address of the given board or device
// to v, an integer in the range 96 to 126 (hex 60 to hex 7E) or zero.
func Ibsad(ud, v int) (ibsta int) {
	ibsta = int(C.ibsad(C.int(ud), C.int(v)))
	return
}

// Assert interface clear.
//
// Ibsic asserts the GPIB interfaces clear (IFC) line for at least 100 ?s
// if the GPIB board is System Controller.
func Ibsic(ud int) (ibsta int) {
	ibsta = int(C.ibsic(C.int(ud)))
	return
}

// Set or clear the Remote Enable (REN) line.

// If v is non-zero, the GPIB Remote Enable (REN) line is asserted.
// If v is zero, REN is unasserted.
func Ibsre(ud, v int) (ibsta int) {
	ibsta = int(C.ibsre(C.int(ud), C.int(v)))
	return
}

// Abort asynchronous I/O operation.
//
// Ibstop aborts any asynchronous read, write, or command operation that is in
// progress and resynchronizes the application with the driver.
func Ibstop(ud int) (ibsta int) {
	ibsta = int(C.ibstop(C.int(ud)))
	return
}

// Change or disable the timeout period.

// Ibtmo sets the timeout period of the board or device to v.
func Ibtmo(ud, v int) (ibsta int) {
	ibsta = int(C.ibtmo(C.int(ud), C.int(v)))
	return
}

// Trigger selected device.
//
// Ibtrg sends the Group Execute Trigger (GET) message to the device
// described by ud.
func Ibtrg(ud int) (ibsta int) {
	ibsta = int(C.ibtrg(C.int(ud)))
	return
}

// Wait for GPIB events.
//
// Ibwait monitors the events specified by mask and delays processing until
// one or more of the events occurs.
func Ibwait(ud, mask int) (ibsta int) {
	ibsta = int(C.ibwait(C.int(ud), C.int(mask)))
	return
}

// Write data to a device from a user buffer.
//
// If ud is a device descriptor, ibwrt addresses the GPIB and writes
// len(buf) bytes from the memory location specified by buf to a GPIB device.
// If ud is a board descriptor, ibwrt writes len(buf) bytes of data from the
// buffer specified by buf to a GPIB device; a board-level ibwrt assumes that
// the GPIB is already properly addressed.
func Ibwrt(ud int, buf string) (ibsta int) {
	n := C.CString(buf)
	defer C.free(unsafe.Pointer(n))
	ibsta = int(C.ibwrt(C.int(ud), unsafe.Pointer(n), C.long(len(buf))))
	return
}

// Write data asynchronously to a device from a user buffer.
//
// If ud is a device descriptor, ibwrta addresses the GPIB and writes
// len(buf) bytes from the memory location specified by buf to a GPIB device.
// If ud is a board descriptor, ibwrt writes len(buf) bytes of data from the
// buffer specified by buf to a GPIB device; a board-level ibwrt assumes that
// the GPIB is already properly addressed.
func Ibwrta(ud int, buf string) (ibsta int) {
	n := C.CString(buf)
	defer C.free(unsafe.Pointer(n))
	ibsta = int(C.ibwrta(C.int(ud), unsafe.Pointer(n), C.long(len(buf))))
	return
}

//  NI-488.2 Functions

// Serial poll all devices.
//
// AllSpoll serial polls all of the devices described by addrlist. It
// stores the poll responses in resultlist and the number of responses
// in ibcntl.
func AllSpoll(boardID int, addrlist []Addr4882_t) (results []int16) {
        n := append(addrlist, NOADDR)
	results = make([]int16, len(addrlist))
	C.AllSpoll(C.int(boardID), (*C.Addr4882_t)(&n[0]),
		(*C.short)(&results[0]))
	return
}

// Clear a single device.
//
// DevClear sends the Selected Device Clear (SDC) GPIB message to the device
// described by address. If address is the constant NOADDR, then the Universal
// Device Clear (DCL) message is sent to all devices.
func DevClear(boardID, address int) {
	C.DevClear(C.int(boardID), C.Addr4882_t(address))
}

// Clear multiple devices.
//
// DevClearList sends the Selected Device Clear (SDC) GPIB message to all
// the device addresses described by addrlist. If addrlist contains only the
// constant NOADDR, then the Universal Device Clear (DCL) message is sent to
// all the devices on the bus.
func DevClearList(boardID int, addrlist []Addr4882_t) {
        n := append(addrlist, NOADDR)
	C.DevClearList(C.int(boardID), (*C.Addr4882_t)(&n[0]))
}

// Enable operations from the front panel of devices (leave remote programming mode).
//
// EnableLocal sends the Go To Local (GTL) GPIB message to all the devices
// described by addrlist. This places the devices into local mode. If addrlist
// contains only the constant NOADDR, then the Remote Enable (REN) GPIB line
// is unasserted.
func EnableLocal(boardID int, addrlist []Addr4882_t) {
        n := append(addrlist, NOADDR)
	C.EnableLocal(C.int(boardID), (*C.Addr4882_t)(&n[0]))
}

// Enable remote GPIB programming for devices.
//
// EnableRemote asserts the Remote Enable (REN) GPIB line. All devices
// described by addrlist are put into a listen-active state.
func EnableRemote(boardID int, addrlist []Addr4882_t) {
        n := append(addrlist, NOADDR)
	C.EnableLocal(C.int(boardID), (*C.Addr4882_t)(&n[0]))
}

// Find listening devices on GPIB.
//
// FindLstn tests all of the primary addresses in addrlist as follows: If a
// device is present at a primary address given in addrlist, then the primary
// address is stored in results Otherwise, all secondary addresses of the
// primary address are tested, and the addresses of any devices found are
// stored in results. No more than limit addresses are stored in results.
// ibcntl contains the actual number of addresses stored in results.
func FindLstn(boardID, limit int, addrlist []Addr4882_t) (results []Addr4882_t) {
        n := append(addrlist, NOADDR)
	results = make([]Addr4882_t, limit)
	C.FindLstn(C.int(boardID), (*C.Addr4882_t)(&n[0]),
		(*C.Addr4882_t)(&results[0]), C.int(limit))
	return
}

// Determines which device is requesting service.
//
// FindRQS serial polls the devices described by addrlist, in order, until it
// finds a device which is requesting service. The serial poll response byte
// is then placed in status. ibcntl contains the index of the device requesting
// service in addrlist. If none of the devices are requesting service, then
// the index corresponding to NOADDR in addrlist is returned in ibcntl and
// ETAB is returned in iberr.
func FindRQS(boardID int, addrlist []Addr4882_t) (status int16) {
        n := append(addrlist, NOADDR)
	C.FindRQS(C.int(boardID), (*C.Addr4882_t)(&n[0]),
		(*C.short)(&status))
	return
}

// Perform a parallel poll on the GPIB.
//
// PPoll conducts a parallel poll and the result is placed in status. Each of
// the eight bits of result represents the status information for each device
// configured for a parallel poll.
func PPoll(boardID int) (status int16) {
	C.PPoll(C.int(boardID), (*C.short)(&status))
	return
}

// Configure a device for parallel polls.
//
// PPollConfig configures the device described by addr to respond to parallel
// polls by asserting or not asserting the GPIB data line, dataline. If
// lineSense equals the individual status (ist) bit of the device, then the
// assigned GPIB data line is asserted during a parallel poll, otherwise, the
// data line is not asserted during a parallel poll.
func PPollConfig(boardID, dataLine, lineSense, addr Addr4882_t) {
	C.PPollConfig(C.int(boardID), C.Addr4882_t(addr),
		C.int(dataLine), C.int(lineSense))
}

// Unconfigure devices for parallel polls.
//
// PPollUnconfig unconfigures all the devices described by addrlist for parallel
// polls. If addrlist contains only the constant NOADDR, then the Parallel Poll
// Unconfigure (PPU) GPIB message is sent to all GPIB devices. The devices
// unconfigured by this function do not participate in subsequent parallel polls.
// boardID The interface board number.
func PPollUnconfig(boardID int, addrlist []Addr4882_t) {
        n := append(addrlist, NOADDR)
	C.PPollUnconfig(C.int(boardID), (*C.Addr4882_t)(&n[0]))
}

// Pass control to another device with Controller capability.
//
// PassControl sends the Take Control (TCT) GPIB message to the device
// described by addr. The device becomes Controller-In-Charge and the
// interface board is no longer CIC.
func PassControl(boardID, addr Addr4882_t) {
	C.PassControl(C.int(boardID), C.Addr4882_t(addr))
}

// Read data bytes from a device that is already addressed to talk.
//
// RcvRespMsg reads up to count bytes from the GPIB and places these bytes
// into data. Data bytes are read until either count data bytes have been
// read or the termination condition is detected. If the termination condition
// is STOPend, the read is stopped when a byte is received with the EOI line
// asserted. Otherwise, the read is stopped when the 8-bit EOS character is
// detected. The actual number of bytes transferred is returned in the global
// variable, ibcntl.
//
// RcvRespMsg assumes that the interface board is already in its listen-active
// state and a device is already addressed to be a Talker (see ReceiveSetup
// or Receive).
func RcvRespMsg(boardID, count, Termination int) (data []byte) {
	data = make([]byte, count)
	C.RcvRespMsg(C.int(boardID), unsafe.Pointer(&data[0]),
		C.long(count), C.int(Termination))
	return
}

// Serial poll a single device.
//
// ReadStatusByte serial polls the device described by addr. The response
// byte is stored in result.
func ReadStatusByte(boardID, addr Addr4882_t) (result int16) {
	C.ReadStatusByte(C.int(boardID), C.Addr4882_t(addr),
		(*C.short)(&result))
	return
}

// Read data bytes from a device.
//
// Receive addresses the device described by addr to talk and the interface
// board to listen. Then up to count bytes are read and placed into the data.
// Data bytes are read until either count bytes have been read or the
// termination condition is detected. If the termination condition is STOPend,
// the read is stopped when a byte is received with the EOI line asserted.
// Otherwise, the read is stopped when an 8-bit EOS character is detected.
// The actual number of bytes transferred is returned in the global variable,
// ibcntl.
func Receive(boardID, count, Termination, addr Addr4882_t) (data []byte) {
	data = make([]byte, count)
	C.Receive(C.int(boardID), C.Addr4882_t(addr), unsafe.Pointer(&data[0]),
		C.long(count), C.int(Termination))
	return
}

// Address a device to be a Talker and the interface board
// to be a Listener in preparation for RcvRespMsg.
//
// ReceiveSetup makes the device described by addr talk-active, and makes
// the interface board listen-active. This call is usually followed by a call
// to RcvRespMsg to transfer data from the device to the interface board.
func ReceiveSetup(boardID, addr Addr4882_t) {
	C.ReceiveSetup(C.int(boardID), C.Addr4882_t(addr))
}

// Reset and initialize IEEE 488.2-compliant devices.
//
// The reset and initialization takes place in three steps. The first step
// resets the GPIB by asserting the Remote Enable (REN) line and then the
// Interface Clear (IFC) line. The second step clears all of the devices by
// sending the Universal Device Clear (DCL) GPIB message. The final step
// causes IEEE 488.2-compliant devices to perform device-specific reset and
// initialization. This step is accomplished by sending the message "*RST\n"
// to the devices described by addrlist.
func ResetSys(boardID int, addrlist []Addr4882_t) {
        n := append(addrlist, NOADDR)
	C.ResetSys(C.int(boardID), (*C.Addr4882_t)(&n[0]))
}

// Send data bytes to a device.
//
// Send sends count len(cmds) bytes from cmds to the device.
// The last byte is sent with the EOI line asserted if eotmode is
// DABend. The last byte is sent without the EOI line asserted if eotmode is
// NULLend. If eotmode is NLend then a new line character ('\n') is sent with
// the EOI line asserted after the last byte of buffer. The actual number of
// bytes transferred is returned in the global variable, ibcntl.
func Send(boardID, eotMode int, addr Addr4882_t, cmds string) {
	n := C.CString(cmds)
	defer C.free(unsafe.Pointer(n))
	C.Send(C.int(boardID), C.Addr4882_t(addr), unsafe.Pointer(n),
		C.long(len(cmds)), C.int(eotMode))
}

// Send GPIB command bytes.
//
// SendCmds sends len(cmds) command bytes from cmds over the GPIB as command
// bytes (interface messages). The number of command bytes transferred is
// returned in the global variable ibcntl.
//
// Use command bytes to configure the state of the GPIB, not to send
// instructions to GPIB devices. Use Send or SendList to send device-specific
// instructions.
func SendCmds(boardID int, cmds string) {
	n := C.CString(cmds)
	defer C.free(unsafe.Pointer(n))
	C.SendCmds(C.int(boardID), unsafe.Pointer(n), C.long(len(cmds)))
}

// Send cmd bytes to devices that are already addressed to listen.
//
// SendDataBytes sends count number of bytes from cmds to devices which
// are already addressed to listen. The last byte is sent with the EOI line
// asserted if eotMode is DABend; the last byte is sent without the EOI line
// asserted if eotMode is NULLend. If eotMode is NLend then a new line character
// ('\n') is sent with the EOI line asserted after the last byte. The actual
// number of bytes transferred is returned in the global variable, ibcntl.
//
// SendDataBytes assumes that the interface board is in talk-active state and
// that devices are already addressed as Listeners on the GPIB (see SendSetup,
// Send, or SendList).
func SendDataBytes(boardID, eotMode int, cmds string) {
	n := C.CString(cmds)
	defer C.free(unsafe.Pointer(n))
	C.SendDataBytes(C.int(boardID), unsafe.Pointer(n), C.long(len(cmds)),
        	C.int(eotMode))
	return
}

// Reset the GPIB by sending interface clear.
//
// SendIFC is used as part of GPIB initialization. It forces the interface
// board to be Controller-In-Charge of the GPIB. It also ensures that the
// connected devices are all unaddressed and that the interface functions of
// the devices are in their idle states.
func SendIFC(boardID int) {
	C.SendIFC(C.int(boardID))
}

// Send the Local Lockout (LLO) message to all devices.
//
// SendLLO sends the GPIB Local Lockout (LLO) message to all devices. While
// Local Lockout is in effect, only the Controller-In-Charge can alter the
// state of the devices by sending appropriate GPIB messages.
func SendLLO(boardID int) {
	C.SendLLO(C.int(boardID))
}

// Send data bytes to multiple GPIB devices.
//
// SendList addresses the devices described by addrlist to listen and the
// interface board to talk. Then, count bytes from data are sent to the
// devices. The last byte is sent with the EOI line asserted if eotMode is
// DABend. The last byte is sent without the EOI line asserted if eotMode is
// NULLend. If eotMode is NLend, then a new line character ('\n') is sent with
// the EOI line asserted after the last byte. The actual number of bytes
// transferred is returned in the global variable, ibcntl.
func SendList(boardID, count, eotMode int, addrlist []Addr4882_t, data []byte) {
        n := append(addrlist, NOADDR)
	C.SendList(C.int(boardID), (*C.Addr4882_t)(&n[0]),
		unsafe.Pointer(&data[0]), C.long(count), C.int(eotMode))
}

// Set up devices to receive data in preparation for SendDataBytes.
//
// SendSetup makes the devices described by addrlist listen-active and makes
// the interface board talk-active. This call is usually followed by
// SendDataBytes to actually transfer data from the interface board to the
// devices.
func SendSetup(boardID int, addrlist []Addr4882_t) {
        n := append(addrlist, NOADDR)
	C.SendSetup(C.int(boardID), (*C.Addr4882_t)(&n[0]))
}

// Place devices in remote with lockout state.
//
// SetRWLS places the devices described by addrlist in remote mode by
// asserting the Remote Enable (REN) GPIB line. Then those devices are placed
// in lockout state by the Local Lockout (LLO) GPIB message. You cannot program
// those devices locally until the Controller-In-Charge releases the Local
// Lockout by way of the EnableLocal NI-488.2 routine.
func SetRWLS(boardID int, addrlist []Addr4882_t) {
	n := append(addrlist, NOADDR)
	C.SetRWLS(C.int(boardID), (*C.Addr4882_t)(&n[0]))
}

// Determine the current state of the GPIB Service Request (SRQ) line.
//
// TestSRQ returns the current state of the GPIB SRQ line in result. If SRQ is
// asserted, then result contains a non-zero value, otherwise, result is
// zero.
func TestSRQ(boardID int) (result int16) {
	C.TestSRQ(C.int(boardID), (*C.short)(&result))
	return
}

// Cause the IEEE 488.2-compliant devices to conduct self tests.
//
// TestSys sends the "*TST?" message to the IEEE 488.2-compliant devices
// described by addrlist. The "*TST?" message instructs them to conduct their
// self-test procedures. A 16-bit test result code is read from each device
// and stored in results. A test result of 0\n indicates that the device
// passed its self test. Refer to the manual that came with the device to
// determine the meaning of a failure code. Any other value indicates that
// the device failed its self test. If the function returns without an error
// (that is, the ERR bit is not set in ibsta), ibcntl contains the number of
// devices that failed. Otherwise, the meaning of ibcntl depends on the error
// returned. If a device fails to send a response before the timeout period
// expires, a test result of 1 is reported for it, and the error EABO is returned.
func TestSys(boardID int, addrlist []Addr4882_t) (results []int16) {
        n := append(addrlist, NOADDR)
	results = make([]int16, len(addrlist))
	C.TestSys(C.int(boardID), (*C.Addr4882_t)(&n[0]),
		(*C.short)(&results[0]))
	return
}

// Trigger a device.
//
// Trigger sends the Group Execute Trigger (GET) GPIB message to the device
// described by addr. If address is the constant NOADDR, then the GET message
// is sent to all devices that are currently listen-active on the GPIB.
func Trigger(boardID int, addr Addr4882_t) {
	C.Trigger(C.int(boardID), C.Addr4882_t(addr))
}

// Trigger multiple devices.
//
// TriggerList sends the Group Execute Trigger (GET) GPIB message to the devices
// described by addrlist. If the only address in addrlist is the constant NOADDR,
// then no addressing is performed and the GET message is sent to all devices
// that are currently listen-active on the GPIB.
func TriggerList(boardID int, addrlist []Addr4882_t) {
        n := append(addrlist, NOADDR)
	C.TriggerList(C.int(boardID), (*C.Addr4882_t)(&n[0]))
}

// Wait until a device asserts the GPIB Service Request (SRQ) line.
//
// WaitSRQ waits until either the GPIB SRQ line is asserted or the timeout
// period has expired (see ibtmo). When WaitSRQ returns, result is non-zero
// if SRQ is asserted, otherwise, result is zero.
func WaitSRQ(boardID int) (result int16) {
	C.WaitSRQ(C.int(boardID), (*C.short)(&result))
	return
}
