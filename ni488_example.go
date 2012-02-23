// +build ignore

package main

import (
	"os"
	"fmt"
	. "github.com/jpoirier/ni488"
)

func GpibError(dev, ud int, msg string, exit bool) {
	fmt.Println(msg)
	ibsta := ThreadIbsta()
	fmt.Printf("ibsta: 0x%X  <", ibsta)

	if (ibsta & ERR) != 0 {
		fmt.Print(" ERR")
	}
	if (ibsta & TIMO) != 0 {
		fmt.Print(" TIMO")
	}
	if (ibsta & END) != 0 {
		fmt.Print(" END")
	}
	if (ibsta & SRQI) != 0 {
		fmt.Print(" SRQI")
	}
	if (ibsta & RQS) != 0 {
		fmt.Print(" RQS")
	}
	if (ibsta & CMPL) != 0 {
		fmt.Print(" CMPL")
	}
	if (ibsta & LOK) != 0 {
		fmt.Print(" LOK")
	}
	if (ibsta & REM) != 0 {
		fmt.Print(" REM")
	}
	if (ibsta & CIC) != 0 {
		fmt.Print(" CIC")
	}
	if (ibsta & ATN) != 0 {
		fmt.Print(" ATN")
	}
	if (ibsta & TACS) != 0 {
		fmt.Print(" TACS")
	}
	if (ibsta & LACS) != 0 {
		fmt.Print(" LACS")
	}
	if (ibsta & DTAS) != 0 {
		fmt.Print(" DTAS")
	}
	if (ibsta & DCAS) != 0 {
		fmt.Print(" DCAS")
	}
	fmt.Println(" >")

	iberr := ThreadIberr()
	switch iberr {
	default:
		fmt.Println("iberr ???? <Unknown Error>")
	case EDVR:
		fmt.Println("iberr EDVR <System Error>")
	case ECIC:
		fmt.Println("iberr ECIC <Not Controller-In-Charge>")
	case ENOL:
		fmt.Println("iberr ENOL <No Listener>")
	case EADR:
		fmt.Println("iberr EADR <Address error>")
	case EARG:
		fmt.Println("iberr EARG <Invalid argument>")
	case ESAC:
		fmt.Println("iberr ESAC <Not System Controller>")
	case EABO:
		fmt.Println("iberr EABO <Operation aborted>")
	case ENEB:
		fmt.Println("iberr ENEB <No GPIB board>")
	case EOIP:
		fmt.Println("iberr EOIP <Async I/O in progress>")
	case ECAP:
		fmt.Println("iberr ECAP <No capability>")
	case EFSO:
		fmt.Println("iberr EFSO <File system error>")
	case EBUS:
		fmt.Println("iberr EBUS <GPIB bus error>")
	case ESTB:
		fmt.Println("iberr ESTB <Status byte lost>")
	case ESRQ:
		fmt.Println("iberr ESRQ <SRQ stuck on>")
	case ETAB:
		fmt.Println("iberr ETAB <Table Overflow>")
	}

	fmt.Printf("ibcntl: %d\n", ThreadIbcntl())

	if exit {
		Ibonl(dev, 0)
		Ibonl(ud, 0)
		os.Exit(0)
	}
}

func main() {
	//	var priAddr int = 20 // instrument address
	//	var secAddr int = 96
	var buffer []byte = make([]byte, 101)
	GPIB0 := 0 // board index

	ud := Ibfind("GPIB0")
	if ud == -1 || (ThreadIbsta()&ERR) != 0 {
		GpibError(GPIB0, ud, "Ibfind", true)
	} else {
		fmt.Printf("GPIB0 found - %d \n", ud)
	}
	fmt.Println("-")

	// system controller check
	if err := Ibconfig(ud, IbcSC, 1); err != 0 {
		GpibError(GPIB0, ud, "Ibconfig IbcSC", false)
	} else {
		fmt.Printf("GPIB%d is the system controller \n", GPIB0)
	}
	fmt.Println("-")

	if err := Ibsic(ud); err != 0 {
		GpibError(GPIB0, ud, "Ibsic", false)
	} else {
		fmt.Println("Ibsic passed...")
	}
	fmt.Println("-")

	if err := Ibconfig(ud, IbcSRE, 1); err != 0 {
		GpibError(GPIB0, ud, "Ibconfig IbcSRE", false)
	} else {
		fmt.Println("IbcSRE passed...")
	}
	fmt.Println("-")

	if err := Ibconfig(ud, IbcTIMING, 1); err != 0 {
		GpibError(GPIB0, ud, "Ibconfig IbcTIMING", false)
	} else {
		fmt.Println("IbcTIMING passed...")
	}
	fmt.Println("-")

	v, err := Ibask(ud, IbaPAD)
	if err != 0 {
		GpibError(GPIB0, ud, "Ibask IbaPAD", false)
	}
	fmt.Printf("Ibask IbaPAD returned: %d \n", v)
	fmt.Println("-")

	v, err = Ibask(ud, IbaSAD)
	if err != 0 {
		GpibError(GPIB0, ud, "Ibask IbaSAD", false)
	}
	fmt.Printf("Ibask IbaSAD returned: %d \n", v)
	fmt.Println("-")

	padList := []int16{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xA,
		0xB, 0xC, 0xD, 0xE, 0xF, 0x10, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E}

	resultList := FindLstn(GPIB0, padList, 20)
	if (ThreadIbsta() & ERR) != 0 {
		GpibError(GPIB0, ud, "FindLstn Error", false)
	}
	fmt.Printf("FindLstn - cnt: %d list: %X\n", len(resultList), resultList)
	fmt.Println("-\n")

	for i, address := range resultList {
		if address == 0 {
			continue
		}
		p := GetPad(uint16(address))
		s := GetSad(uint16(address))
		fmt.Printf("PAD: %d, Instrument %d, SAD: %d \n", p, i, s)

		ud = Ibdev(GPIB0,
			p,
			s,
			T1s, //Timeout setting
			1,   // Assert EOI line at end of write
			0)   // EOS termination mode

		Ibclr(ud)

		cmd := "*IDN?"
		Ibwrt(ud, cmd)
		if (ThreadIbsta() & ERR) == 0 {
			fmt.Println("Ibwrt success...\n")
		} else {
			GpibError(GPIB0, ud, "Ibwrt", false)
		}

		Ibrd(ud, buffer)
		if (ThreadIbsta() & ERR) == 0 {
			fmt.Printf("buffer: %s", buffer)
		} else {
			GpibError(GPIB0, ud, "Ibrd", false)
		}

		Ibonl(ud, 0) // Take the interface offline
		if (ThreadIbsta() & ERR) == 0 {
			fmt.Println("Ibonl success, interface offline...\n")
		} else {
			GpibError(GPIB0, ud, "Ibonl", false)
		}
		fmt.Println("\n----------------------")
	}

	Ibonl(GPIB0, 0)
}
