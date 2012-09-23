// Copyright (c) 2011 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

// Package ni488 wraps National Instruments 488.2 General Purpose Interface
// Bus (GPIB) driver. The driver allows a client application to communicate
// with a GPIB enabled piece of test equipment remotely and/or programmatically.
// NI-488.2 is an industry standard for GPIB communications.
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
// Direct download: http://download.ni.com/support/softlib/gpib/
//
package ni488

// TODO:
// -

/*
#cgo linux LDFLAGS: -lgpibapi
#cgo darwin CFLAGS: -I.
#cgo darwin LDFLAGS: -framework NI488
#cgo windows CFLAGS: -I.
#cgo windows LDFLAGS: -lgpib-32 -LC:/WINDOWS/system32
#include <stdlib.h>
#if defined(__amd64) || defined(__amd64__) || defined(__x86_64) || defined(__x86_64__) && !defined(__APPLE__)
#define size_g size_t
#include <ni4882.h>
#else
#define size_g long
#include <ni488.h>
#endif
*/
import "C"
import "unsafe"

var PackageVersion string = "v0.4"

// GetPad extracts and returns the primary instrument
// address from a base instrument address.
func GetPad(addr uint16) int {
	return int(addr & 0xFF)
}

// GetSad extracts and returns the secondary instrument
// address from a base instrument address.
func GetSad(addr uint16) int {
	return int(((addr) >> 8) & 0xFF)
}

//  Functions to access Thread-Specific copies of the GPIB global vars

// The global variables ibsta, iberr, ibcnt, and ibcntl are maintained on a
// process-specific rather than a thread-specific basis. If you call GPIB
// functions in more than one thread, the values in these global variables
// are not always reliable.

// Status variables analogous to ibsta, iberr, ibcnt, and ibcntl are maintained
// for each thread. ThreadIbcntl returns the value of the thread-specific ibcntl
// variable.

// ThreadIbsta returns the thread-specific ibsta value for the current thread.
//
// The return value is the value for the current thread of execution. The
// value describes the state of the GPIB and the result of the most recent
// GPIB function call in the thread.  Call ThreadIberr for a specific error
// code.
func ThreadIbsta() uint32 {
	return uint32(C.ThreadIbsta())
}

// ThreadIberr returns the thread-specific iberr value for the current thread.
//
// The return value is the most recent GPIB error code for the current
// thread of execution. The value is meaningful only when ThreadIbsta returns
// a value with the ERR bit set.
func ThreadIberr() uint32 {
	return uint32(C.ThreadIberr())
}

// ThreadIbcnt returns the thread-specific ibcnt value for the current thread.
//
// The return value is either the number of bytes actually transferred by
// the most recent GPIB read, write, or command operation for the current
// thread of execution or an error code if an error occured.
func ThreadIbcnt() uint32 {
	return uint32(C.ThreadIbcnt())
}

// ThreadIbcntl returns the thread-specific ibcntl value for the current thread.
//
// The return value is either the number of bytes actually transferred by
// the most recent GPIB read, write, or command operation for the current
// thread of execution or an error code if an error occured.
func ThreadIbcntl() uint32 {
	return uint32(C.ThreadIbcnt())
}

//  NI-488 Functions

// Ibrdf reads data asynchronously from a device into a user buffer.
//
// If ud is a device descriptor, ibrdf addresses the GPIB, reads data from
// a GPIB device, and places the data into the file specified by filename.
// If ud is a board descriptor, ibrdf reads data from a GPIB device and
// places the data into the file specified by filename
func Ibrdf(ud int, filename string) (ibsta uint32) {
	n := C.CString(filename)
	defer C.free(unsafe.Pointer(n))
	return uint32(C.ibrdfA(C.int(ud), n))

}

// Ibask returns the current value of various configuration parameters for the
// specified board or device.
//
// The current value of the selected configuration item is returned in v.
func Ibask(ud, option int) (v, ibsta uint32) {
	return uint32(C.ibask(C.int(ud), C.int(option),
		(*C.int)(unsafe.Pointer(&v))))
}

// Ibcac uses the designated GPIB board to attempt to become the Active
// Controller by asserting ATN.
//
// If v is zero, the GPIB board takes control asynchronously and if v
// is non-zero the GPIB board takes control synchronously. Before calling
// ibcac, the GPIB board must already be CIC. To make the board CIC, use
// the ibsic function.
func Ibcac(ud, v int) (ibsta uint32) {
	ibsta = uint32(C.ibcac(C.int(ud), C.int(v)))
	return
}

// Ibclr sends the GPIB Selected Device Clear (SDC) message to the device
// described by ud.
func Ibclr(ud int) (ibsta uint32) {
	return uint32(C.ibclr(C.int(ud)))
}

// Ibcmd sends GPIB commands.
//
// Sends cmds over the GPIB as command bytes (interface messages). The actual
// transferred byte count is returned in the global variable ibcntl.
func Ibcmd(ud int, cmds string) (ibsta uint32) {
	n := C.CString(cmds)
	defer C.free(unsafe.Pointer(n))
	return uint32(C.ibcmd(C.int(ud), unsafe.Pointer(n), C.size_g(len(cmds))))
}

// Ibcmda sends GPIB commands asynchronously.
//
// Sends cmds asynchronously over the GPIB as command bytes (interface
// messages). The actual transferred byte count is returned in the global
// variable ibcntl.
func Ibcmda(ud int, cmds string) (ibsta uint32) {
	n := C.CString(cmds)
	defer C.free(unsafe.Pointer(n))
	ibsta = uint32(C.ibcmda(C.int(ud), unsafe.Pointer(n), C.size_g(len(cmds))))
	return
}

// Ibconfig changes software configuration parameters.
//
// Changes a configuration item in option to the specified value in
// v for the selected board or device.
func Ibconfig(ud, option, v int) (ibsta uint32) {
	return uint32(C.ibconfig(C.int(ud), C.int(option), C.int(v)))
}

// Ibdev opens and initialize a device.
//
// Acquires a device descriptor to use in subsequent device-level NI-488
// functions. It opens and initializes a device descriptor, and configures
// it according to the input parameters. Returns the device descriptor or 1.
func Ibdev(boardID, pad, sad, tmo, eot, eos int) (dev int) {
	return int(C.ibdev(C.int(boardID), C.int(pad), C.int(sad),
		C.int(tmo), C.int(eot), C.int(eos)))
}

// TODO
//extern int  ibexpert (int ud, int option, void * input, void * Output);
//func Ibexpert(ud, option int, input, output string) uint32 {
//	in := C.CString(input)
//	defer C.free(unsafe.Pointer(in))
//	return uint32(C.ibexpert(C.int(ud), C.int(option),
//			unsafe.Pointer(input), unsafe.Pointer(output)))
//}

// Ibfind opens and initialize a board or a user-configured device descriptor.
//
// Performs the equivalent of an ibonl 1 to initialize the board or
// device descriptor. The unit descriptor returned by ibfind remains valid
// until the board or device is put offline using ibonl 0.
//
// If ibfind is unable to get a valid descriptor, a -1 is returned; the ERR
// bit is set in ibsta and iberr contains EDVR.
func Ibfind(udname string) (ud int) {
	n := C.CString(udname)
	defer C.free(unsafe.Pointer(n))
	return int(C.ibfindA(n))
}

// Ibgts causes the GPIB board at ud to go to Standby Controller and
// the GPIB ATN line to be unasserted.
//
// v determines whether to perform acceptor handshaking
func Ibgts(ud, v int) (ibsta uint32) {
	return uint32(C.ibgts(C.int(ud), C.int(v)))
}

// TODO
//extern int  iblck    (int ud, int v, unsigned int LockWaitTime, void * Reserved);
//// Acquire or release an exclusive interface lock
//func Iblck(ud, v int, LockWaitTime uint, Reserved) int {
//	return int(C.iblck(C.int(ud), C.int(v), C.uint(LockWaitTime), ))
//}

// Iblines returns the status of the eight GPIB control lines.
func Iblines(ud int) (ibsta, result uint32) {
	return uint32(C.iblines(C.int(ud), (*C.short)(unsafe.Pointer(&result))))
}

// Ibln checks for the presence of a device on the bus.
//
// Determines whether there is a listening device at the GPIB address
// designated by the pad and sad parameters. If ud is a board descriptor,
// then the bus associated with that board is tested for Listeners.
// If ud is a device descriptor, then ibln uses the access board
// associated with that device to test for Listeners. If a Listener is
// detected, a non-zero value is returned in listen. If no Listener is
// found, zero is returned.
func Ibln(ud, pad, sad int) (ibsta, listen uint32) {
	return uint32(C.ibln(C.int(ud), C.int(pad), C.int(sad),
		(*C.short)(unsafe.Pointer(&listen))))
}

// Ibloc places the board in local mode if it is not in a lockout state.
func Ibloc(ud int) (ibsta uint32) {
	return uint32(C.ibloc(C.int(ud)))
}

// Ibnotify notifies user of one or more GPIB events by invoking the user
// callback.
//
// Installs an asynchronous callback function for a specified
// board or device. If mask is non-zero, ibnotify monitors the events
// specified by mask, and when one or more of the events is true, the
// Callback is invoked. refData User-defined reference data for the callback.
// int mycallback(int ud, int ibsta, int iberr, long ibcntl, void *RefData)
func Ibnotify(ud, mask int, f func(), redData []uint32) (ibsta uint32) {
	return uint32(C.ibnotify(C.int(ud), C.int(mask),
		(C.GpibNotifyCallback_t)(unsafe.Pointer(&f)),
		unsafe.Pointer(&redData[0])))
}

// Ibonl places the device online or offline.
//
// Resets the board or device and places all its software configuration
// parameters in their pre-configured state. In addition, if v is zero,
// the device or interface board is taken offline. If v is non-zero, the
// device or interface board is left operational, or online.
// ud Board or device descriptor
func Ibonl(ud, v int) (ibsta int) {
	return int(C.ibonl(C.int(ud), C.int(v)))
}

// Ibpct passes control to another GPIB device with Controller capability.
//
// Passes Controller-in-Charge status to the device indicated by ud.
func Ibpct(ud int) (ibsta uint32) {
	return uint32(C.ibpct(C.int(ud)))
}

// Ibppc configures parallel polling.
//
// If ud is a device descriptor, ibppc enables or disables the device
// from responding to parallel polls.
func Ibppc(ud, v int) (ibsta uint32) {
	return uint32(C.ibppc(C.int(ud), C.int(v)))
}

// Ibrd reads data asynchronously from a device into a user buffer.
//
// If ud is a device descriptor, ibrd addresses the GPIB, reads up to
// len(buf) bytes of data, and places the data into the buffer specified
// buf.
func Ibrd(ud int, buf []byte) (ibsta uint32) {
	return uint32(C.ibrd(C.int(ud), unsafe.Pointer(&buf[0]),
		C.size_g(len(buf))))
}

// Ibrpp conducts a parallel poll.
//
// If this routine is called specifying a GPIB Interface Board, the board
// parallel polls all previously configured devices. If the routine is called
// specifying a device, the GPIB Interface board associated with the device
// conducts the parallel poll. Note that if the GPIB Interface Board to conduct
// the parallel poll is not the Controller- In-Charge, an ECIC error is generated.
func Ibrpp(ud int) (ibsta, resp uint32) {
	return uint32(C.ibrpp(C.int(ud), (*C.char)(unsafe.Pointer(&resp))))
}

// Ibrsp conducts a serial poll on the device ud.
func Ibrsp(ud int) (ibsta, resp uint32) {
	return uint32(C.ibrsp(C.int(ud), (*C.char)(unsafe.Pointer(&resp))))
}

// Ibsic asserts an interface clear.
//
// Asserts the GPIB interfaces clear (IFC) line for at least 100s
// if the GPIB board is System Controller.
func Ibsic(ud int) (ibsta uint32) {
	return uint32(C.ibsic(C.int(ud)))
}

// Ibstop aborts an asynchronous I/O operation.
//
// Aborts any asynchronous read, write, or command operation that is in
// progress and resynchronizes the application with the driver.
func Ibstop(ud int) (ibsta uint32) {
	return uint32(C.ibstop(C.int(ud)))
}

// Ibtrg triggers the selected device.
//
// Sends the Group Execute Trigger (GET) message to the device
// described by ud.
func Ibtrg(ud int) (ibsta uint32) {
	return uint32(C.ibtrg(C.int(ud)))
}

// Ibwait waits for GPIB events.
//
// Monitors the events specified by mask and delays processing until
// one or more of the events occurs.
func Ibwait(ud, mask int) (ibsta uint32) {
	return uint32(C.ibwait(C.int(ud), C.int(mask)))
}

// Ibwrt writes data to a device from a user buffer.
//
// If ud is a device descriptor, ibwrt addresses the GPIB and writes
// len(buf) bytes from the memory location specified by buf to a GPIB device.
// If ud is a board descriptor, ibwrt writes len(buf) bytes of data from the
// buffer specified by buf to a GPIB device; a board-level ibwrt assumes that
// the GPIB is already properly addressed.
func Ibwrt(ud int, buf string) (ibsta uint32) {
	n := C.CString(buf)
	defer C.free(unsafe.Pointer(n))
	return uint32(C.ibwrt(C.int(ud), unsafe.Pointer(n), C.size_g(len(buf))))
}

// Ibwrta writes data asynchronously to a device from a user buffer.
//
// If ud is a device descriptor, ibwrta addresses the GPIB and writes
// len(buf) bytes from the memory location specified by buf to a GPIB device.
// If ud is a board descriptor, ibwrt writes len(buf) bytes of data from the
// buffer specified by buf to a GPIB device; a board-level ibwrt assumes that
// the GPIB is already properly addressed.
func Ibwrta(ud int, buf string) (ibsta uint32) {
	n := C.CString(buf)
	defer C.free(unsafe.Pointer(n))
	return uint32(C.ibwrta(C.int(ud), unsafe.Pointer(n), C.size_g(len(buf))))
}

// Ibwrtf writes data to a device from a file.
//
// If ud is a device descriptor, ibwrtf addresses the GPIB and writes all
// of the bytes from the file filename to a GPIB device. If ud is a board
// descriptor, ibwrtf writes all of the bytes of data from the file filename
// to a GPIB device.
func Ibwrtf(ud int, filename string) (ibsta uint32) {
	n := C.CString(filename)
	defer C.free(unsafe.Pointer(n))
	return uint32(C.ibwrtfA(C.int(ud), n))
}

// Ibdma enables or disables DMA.
//
// Enables or disables DMA transfers for the board, according to v.
// If v is zero, then DMA is not used for GPIB I/O transfers, and if v
// is non-zero, then DMA is used for GPIB I/O transfers.
func Ibdma(ud, v int) (ibsta int) {
	ibsta = int(C.ibconfig(C.int(ud), C.int(C.IbcDMA), C.int(v)))
	return
}

// Ibeot enables or disables the assertion of the EOI line at the end of
// write I/O operations for the board or device described by ud.
//
// If v is non-zero, then EOI is asserted when the last byte of a GPIB
// write is sent.
func Ibeot(ud, v int) (ibsta int) {
	return int(C.ibconfig(C.int(ud), C.int(C.IbcEOT), C.int(v)))
}

// Ibist sets or clears the board individual status bit for parallel polls.
func Ibist(ud, v int) (ibsta int) {
	return int(C.ibconfig(C.int(ud), C.int(C.IbcIst), C.int(v)))
}

// Ibpad changes the primary address.
//
// Sets the primary GPIB address of the board or device to v, an
// integer ranging from 0 to 30.
func Ibpad(ud, v int) (ibsta int) {
	return int(C.ibconfig(C.int(ud), C.int(C.IbcPAD), C.int(v)))
}

// Ibrsc requests or releases system control.
//
// Requests or releases the capability to send Interface Clear (IFC)
// and Remote Enable (REN) messages to devices.
func Ibrsc(ud, v int) (ibsta int) {
	return int(C.ibconfig(C.int(ud), C.int(C.IbcSC), C.int(v)))
}

// Ibrsv requests service and change the serial poll status byte.
//
// Is used to request service from the Controller and to provide the
// Controller with an application-dependent status byte when the
// Controller serial polls the GPIB board.
func Ibrsv(ud, status int) (ibsta int) {
	return int(C.ibconfig(C.int(ud), C.int(C.IbcRsv), C.int(status)))
}

// Ibsad changes or disables the secondary address.
//
// Changes the secondary GPIB address of the given board or device
// to v, an integer in the range 96 to 126 (hex 60 to hex 7E) or zero.
func Ibsad(ud, v int) (ibsta int) {
	return int(C.ibconfig(C.int(ud), C.int(C.IbcSAD), C.int(v)))
}

// Ibsre sets or clears the Remote Enable (REN) line.
//
// If v is non-zero, the GPIB Remote Enable (REN) line is asserted.
// If v is zero, REN is unasserted.
func Ibsre(ud, v int) (ibsta int) {
	return int(C.ibconfig(C.int(ud), C.int(C.IbcSRE), C.int(v)))
}

// Ibtmo changes or disables the timeout period.
//
// Sets the timeout period of the board or device to v.
func Ibtmo(ud, v int) (ibsta int) {
	return int(C.ibconfig(C.int(ud), C.int(C.IbcTMO), C.int(v)))
}

//  NI-488.2 Functions

// XXX: AllSpoll is undefined when linking against ni4882.dll
//      and gpib-32.dll (32-bit compilation via 8g&8l)

// AllSpoll performs a serial poll on all devices.
//
// Serial polls all of the devices described by addrlist. It
// stores the poll responses in resultlist and the number of responses
// in ibcntl.
//func AllSpoll(boardID int, addrlist []int16) (results []int16) {
//	n := append(addrlist, NOADDR)
//	results = make([]int16, len(addrlist))
//	C.AllSpoll(C.int(boardID), (*C.short)(&n[0]),
//		(*C.short)(&results[0]))
//	return
//}

// DevClear clears a single device.
//
// Sends the Selected Device Clear (SDC) GPIB message to the device
// described by address. If address is the constant NOADDR, then the Universal
// Device Clear (DCL) message is sent to all devices.
func DevClear(boardID, address int) {
	C.DevClear(C.int(boardID), C.short(address))
}

// DevClearList clears multiple devices.
//
// Sends the Selected Device Clear (SDC) GPIB message to all
// the device addresses described by addrlist. If addrlist contains only the
// constant NOADDR, then the Universal Device Clear (DCL) message is sent to
// all the devices on the bus.
func DevClearList(boardID int, addrlist []int16) {
	n := append(addrlist, NOADDR)
	C.DevClearList(C.int(boardID), (*C.short)(&n[0]))
}

// EnableLocal enables operations from the front panel of devices (leave
// remote programming mode).
//
// Sends the Go To Local (GTL) GPIB message to all the devices
// described by addrlist. This places the devices into local mode. If addrlist
// contains only the constant NOADDR, then the Remote Enable (REN) GPIB line
// is unasserted.
func EnableLocal(boardID int, addrlist []int16) {
	n := append(addrlist, NOADDR)
	C.EnableLocal(C.int(boardID), (*C.short)(&n[0]))
}

// EnableRemote enables a remote GPIB programming for devices.
//
// Asserts the Remote Enable (REN) GPIB line. All devices
// described by addrlist are put into a listen-active state.
func EnableRemote(boardID int, addrlist []int16) {
	n := append(addrlist, NOADDR)
	C.EnableLocal(C.int(boardID), (*C.short)(&n[0]))
}

// FindLstn finds listening devices on GPIB.
//
// Tests all of the primary addresses in addrlist as follows: If a
// device is present at a primary address given in addrlist, then the primary
// address is stored in results, otherwise, all secondary addresses of the
// primary address are tested, and the addresses of any devices found are
// stored in results. No more than limit addresses are stored in results.
// ibcntl contains the actual number of addresses stored in results.
func FindLstn(boardID int, addrlist []int16, limit int) (results []int16) {
	n := append(addrlist, NOADDR)
	results = make([]int16, limit)
	C.FindLstn(C.int(boardID), (*C.short)(&n[0]),
		(*C.short)(&results[0]), C.int(limit))
	return
}

// FindRQS determines which device is requesting service.
//
// Serial polls the devices described by addrlist, in order, until it
// finds a device which is requesting service. The serial poll response byte
// is then placed in status. ibcntl contains the index of the device requesting
// service in addrlist. If none of the devices are requesting service, then
// the index corresponding to NOADDR in addrlist is returned in ibcntl and
// ETAB is returned in iberr.
func FindRQS(boardID int, padList []int16) (status int16) {
	n := append(padList, NOADDR)
	C.FindRQS(C.int(boardID), (*C.short)(&n[0]),
		(*C.short)(&status))
	return
}

// PPoll perform a parallel poll on the GPIB.
//
// Conducts a parallel poll and the result is placed in status. Each of
// the eight bits of result represents the status information for each device
// configured for a parallel poll.
func PPoll(boardID int) (status int16) {
	C.PPoll(C.int(boardID), (*C.short)(&status))
	return
}

// PPollConfig configures a device for parallel polls.
//
// Configures the device described by addr to respond to parallel
// polls by asserting or not asserting the GPIB data line, dataline. If
// lineSense equals the individual status (ist) bit of the device, then the
// assigned GPIB data line is asserted during a parallel poll, otherwise, the
// data line is not asserted during a parallel poll.
func PPollConfig(boardID, dataLine, lineSense, addr int16) {
	C.PPollConfig(C.int(boardID), C.short(addr),
		C.int(dataLine), C.int(lineSense))
}

// PPollUnconfig unconfigures devices for parallel polls.
//
// Unconfigures all the devices described by addrlist for parallel
// polls. If addrlist contains only the constant NOADDR, then the Parallel Poll
// Unconfigure (PPU) GPIB message is sent to all GPIB devices. The devices
// unconfigured by this function do not participate in subsequent parallel polls.
// boardID The interface board number.
func PPollUnconfig(boardID int, addrlist []int16) {
	n := append(addrlist, NOADDR)
	C.PPollUnconfig(C.int(boardID), (*C.short)(&n[0]))
}

// PassControl passes control to another device with Controller capability.
//
// Sends the Take Control (TCT) GPIB message to the device
// described by addr. The device becomes Controller-In-Charge and the
// interface board is no longer CIC.
func PassControl(boardID, addr int16) {
	C.PassControl(C.int(boardID), C.short(addr))
}

// RcvRespMsg reads data bytes from a device that is already addressed to talk.
//
// Reads up to count bytes from the GPIB and places these bytes
// into data. Data bytes are read until either count data bytes have been
// read or the termination condition is detected. If the termination condition
// is STOPend, the read is stopped when a byte is received with the EOI line
// asserted. Otherwise, the read is stopped when the 8-bit EOS character is
// detected. The actual number of bytes transferred is returned in the global
// variable, ibcntl.
//
// Assumes that the interface board is already in its listen-active
// state and a device is already addressed to be a Talker (see ReceiveSetup
// or Receive).
func RcvRespMsg(boardID, count, Termination int) (data []byte) {
	data = make([]byte, count)
	C.RcvRespMsg(C.int(boardID), unsafe.Pointer(&data[0]),
		C.size_g(count), C.int(Termination))
	return
}

// ReadStatusByte serial poll a single device.
//
// Serial polls the device described by addr. The response
// byte is stored in result.
func ReadStatusByte(boardID, addr int16) (result int16) {
	C.ReadStatusByte(C.int(boardID), C.short(addr),
		(*C.short)(&result))
	return
}

// Receive reads data bytes from a device.
//
// Addresses the device described by addr to talk and the interface
// board to listen. Then up to count bytes are read and placed into the data.
// Data bytes are read until either count bytes have been read or the
// termination condition is detected. If the termination condition is STOPend,
// the read is stopped when a byte is received with the EOI line asserted.
// Otherwise, the read is stopped when an 8-bit EOS character is detected.
// The actual number of bytes transferred is returned in the global variable,
// ibcntl.
func Receive(boardID, count, Termination, addr int16) (data []byte) {
	data = make([]byte, count)
	C.Receive(C.int(boardID), C.short(addr), unsafe.Pointer(&data[0]),
		C.size_g(count), C.int(Termination))
	return
}

// ReceiveSetup addresses a device to be a Talker and the interface board
// to be a Listener in preparation for RcvRespMsg.
//
// Makes the device described by addr talk-active, and makes
// the interface board listen-active. This call is usually followed by a call
// to RcvRespMsg to transfer data from the device to the interface board.
func ReceiveSetup(boardID, addr int16) {
	C.ReceiveSetup(C.int(boardID), C.short(addr))
}

// ResetSys resets and initializes IEEE 488.2-compliant devices.
//
// Reset and initialization takes place in three steps. The first step
// resets the GPIB by asserting the Remote Enable (REN) line and then the
// Interface Clear (IFC) line. The second step clears all of the devices by
// sending the Universal Device Clear (DCL) GPIB message. The final step
// causes IEEE 488.2-compliant devices to perform device-specific reset and
// initialization. This step is accomplished by sending the message "*RST\n"
// to the devices described by addrlist.
func ResetSys(boardID int, addrlist []int16) {
	n := append(addrlist, NOADDR)
	C.ResetSys(C.int(boardID), (*C.short)(&n[0]))
}

// Send sends data bytes to a device.
//
// Sends count len(cmds) bytes from cmds to the device.
// The last byte is sent with the EOI line asserted if eotmode is
// DABend. The last byte is sent without the EOI line asserted if eotmode is
// NULLend. If eotmode is NLend then a new line character ('\n') is sent with
// the EOI line asserted after the last byte of buffer. The actual number of
// bytes transferred is returned in the global variable, ibcntl.
func Send(boardID, eotMode int, addr int16, cmds string) {
	n := C.CString(cmds)
	defer C.free(unsafe.Pointer(n))
	C.Send(C.int(boardID), C.short(addr), unsafe.Pointer(n),
		C.size_g(len(cmds)), C.int(eotMode))
}

// SendCmds sends GPIB command bytes.
//
// Sends len(cmds) command bytes from cmds over the GPIB as command
// bytes (interface messages). The number of command bytes transferred is
// returned in the global variable ibcntl.
//
// Use command bytes to configure the state of the GPIB, not to send
// instructions to GPIB devices. Use Send or SendList to send device-specific
// instructions.
func SendCmds(boardID int, cmds string) {
	n := C.CString(cmds)
	defer C.free(unsafe.Pointer(n))
	C.SendCmds(C.int(boardID), unsafe.Pointer(n), C.size_g(len(cmds)))
}

// SendDataBytes sends cmd bytes to devices that are already addressed to listen.
//
// Sends count number of bytes from cmds to devices which
// are already addressed to listen. The last byte is sent with the EOI line
// asserted if eotMode is DABend; the last byte is sent without the EOI line
// asserted if eotMode is NULLend. If eotMode is NLend then a new line character
// ('\n') is sent with the EOI line asserted after the last byte. The actual
// number of bytes transferred is returned in the global variable, ibcntl.
//
// Assumes that the interface board is in talk-active state and
// that devices are already addressed as Listeners on the GPIB (see SendSetup,
// Send, or SendList).
func SendDataBytes(boardID, eotMode int, cmds string) {
	n := C.CString(cmds)
	defer C.free(unsafe.Pointer(n))
	C.SendDataBytes(C.int(boardID), unsafe.Pointer(n), C.size_g(len(cmds)),
		C.int(eotMode))
	return
}

// SendIFC resets the GPIB by sending interface clear.
//
// Is used as part of GPIB initialization. It forces the interface
// board to be Controller-In-Charge of the GPIB. It also ensures that the
// connected devices are all unaddressed and that the interface functions of
// the devices are in their idle states.
func SendIFC(boardID int) {
	C.SendIFC(C.int(boardID))
}

// SendLLO sends the Local Lockout (LLO) message to all devices.
//
// Sends the GPIB Local Lockout (LLO) message to all devices. While
// Local Lockout is in effect, only the Controller-In-Charge can alter the
// state of the devices by sending appropriate GPIB messages.
func SendLLO(boardID int) {
	C.SendLLO(C.int(boardID))
}

// SendList sends data bytes to multiple GPIB devices.
//
// Addresses the devices described by addrlist to listen and the
// interface board to talk. Then, count bytes from data are sent to the
// devices. The last byte is sent with the EOI line asserted if eotMode is
// DABend. The last byte is sent without the EOI line asserted if eotMode is
// NULLend. If eotMode is NLend, then a new line character ('\n') is sent with
// the EOI line asserted after the last byte. The actual number of bytes
// transferred is returned in the global variable, ibcntl.
func SendList(boardID, count, eotMode int, addrlist []int16, data []byte) {
	n := append(addrlist, NOADDR)
	C.SendList(C.int(boardID), (*C.short)(&n[0]),
		unsafe.Pointer(&data[0]), C.size_g(count), C.int(eotMode))
}

// SendSetup sets up devices to receive data in preparation for SendDataBytes.
//
// Makes the devices described by addrlist listen-active and makes
// the interface board talk-active. This call is usually followed by
// SendDataBytes to actually transfer data from the interface board to the
// devices.
func SendSetup(boardID int, addrlist []int16) {
	n := append(addrlist, NOADDR)
	C.SendSetup(C.int(boardID), (*C.short)(&n[0]))
}

// SetRWLS places devices in remote with lockout state.
//
// Places the devices described by addrlist in remote mode by
// asserting the Remote Enable (REN) GPIB line. Then those devices are placed
// in lockout state by the Local Lockout (LLO) GPIB message. You cannot program
// those devices locally until the Controller-In-Charge releases the Local
// Lockout by way of the EnableLocal NI-488.2 routine.
func SetRWLS(boardID int, addrlist []int16) {
	n := append(addrlist, NOADDR)
	C.SetRWLS(C.int(boardID), (*C.short)(&n[0]))
}

// TestSRQ determines the current state of the GPIB Service Request (SRQ) line.
//
// Returns the current state of the GPIB SRQ line in result. If SRQ is
// asserted, then result contains a non-zero value, otherwise, result is
// zero.
func TestSRQ(boardID int) (result int16) {
	C.TestSRQ(C.int(boardID), (*C.short)(&result))
	return
}

// TestSys causes the IEEE 488.2-compliant devices to conduct self tests.
//
// Sends the "*TST?" message to the IEEE 488.2-compliant devices
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
func TestSys(boardID int, addrlist []int16) (results []int16) {
	n := append(addrlist, NOADDR)
	results = make([]int16, len(addrlist))
	C.TestSys(C.int(boardID), (*C.short)(&n[0]),
		(*C.short)(&results[0]))
	return
}

// Trigger triggers a device.
//
// Sends the Group Execute Trigger (GET) GPIB message to the device
// described by addr. If address is the constant NOADDR, then the GET message
// is sent to all devices that are currently listen-active on the GPIB.
func Trigger(boardID int, addr int16) {
	C.Trigger(C.int(boardID), C.short(addr))
}

// TriggerList triggers multiple devices.
//
// Sends the Group Execute Trigger (GET) GPIB message to the devices
// described by addrlist. If the only address in addrlist is the constant NOADDR,
// then no addressing is performed and the GET message is sent to all devices
// that are currently listen-active on the GPIB.
func TriggerList(boardID int, addrlist []int16) {
	n := append(addrlist, NOADDR)
	C.TriggerList(C.int(boardID), (*C.short)(&n[0]))
}

// WaitSRQ waita until a device asserts the GPIB Service Request (SRQ) line.
//
// Waits until either the GPIB SRQ line is asserted or the timeout
// period has expired (see ibtmo). When WaitSRQ returns, result is non-zero
// if SRQ is asserted, otherwise, result is zero.
func WaitSRQ(boardID int) (result int16) {
	C.WaitSRQ(C.int(boardID), (*C.short)(&result))
	return
}
