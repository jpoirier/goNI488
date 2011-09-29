package main

import (
	"os"
	"fmt"
	. "github.com/jpoirier/ni488"
)

func GpibError(dev, brdIndx int, msg string) {
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
	fmt.Println(" >\n")

	iberr := ThreadIberr()
	fmt.Printf("iberr: %d", iberr)
	switch iberr {
	default:
		fmt.Println(" ???? <Unknown Error>\n")
	case EDVR:
		fmt.Println(" EDVR <System Error>\n")
	case ECIC:
		fmt.Println(" ECIC <Not Controller-In-Charge>\n")
	case ENOL:
		fmt.Println(" ENOL <No Listener>\n")
	case EADR:
		fmt.Println(" EADR <Address error>\n")
	case EARG:
		fmt.Println(" EARG <Invalid argument>\n")
	case ESAC:
		fmt.Println(" ESAC <Not System Controller>\n")
	case EABO:
		fmt.Println(" EABO <Operation aborted>\n")
	case ENEB:
		fmt.Println(" ENEB <No GPIB board>\n")
	case EOIP:
		fmt.Println(" EOIP <Async I/O in progress>\n")
	case ECAP:
		fmt.Println(" ECAP <No capability>\n")
	case EFSO:
		fmt.Println(" EFSO <File system error>\n")
	case EBUS:
		fmt.Println(" EBUS <GPIB bus error>\n")
	case ESTB:
		fmt.Println(" ESTB <Status byte lost>\n")
	case ESRQ:
		fmt.Println(" ESRQ <SRQ stuck on>\n")
	case ETAB:
		fmt.Println(" ETAB <Table Overflow>\n")
	}

	fmt.Printf("\nibcntl: %d\n", ThreadIbcntl())
	Ibonl(dev, 0)
	Ibonl(brdIndx, 0)
	os.Exit(0)
}

func main() {
	var device int = 0
	var brdIndx int = 0
	var priAddr int = 20 // instrument address
	var secAddr int = 96
	var buffer []byte = make([]byte, 101)

	//--- Init
	device = Ibdev(brdIndx,
		priAddr,
		secAddr,
		T10s, //Timeout setting
		1,    // Assert EOI line at end of write
		0)    // EOS termination mode

	if (ThreadIbsta() & ERR) == 0 {
		fmt.Printf("Device 0x%X:%d \n", device, device)
		fmt.Println("Ibdev success...\n")
	} else {
		GpibError(device, brdIndx, "Ibdev Error")
	}

	Ibclr(device)
	if (ThreadIbsta() & ERR) == 0 {
		fmt.Println("Ibclr success...\n")
	} else {
		GpibError(device, brdIndx, "Ibclr Error")
	}

	DevClear(brdIndx, priAddr)
	if (ThreadIbsta() & ERR) == 0 {
		fmt.Println("DevClear success...\n")
	} else {
		GpibError(device, brdIndx, "Ibclr Error")
	}

	//--- App body
	cmd := "*IDN?"
	Ibwrt(device, cmd) // Send the identification query command
	if (ThreadIbsta() & ERR) == 0 {
		fmt.Println("Ibwrt success...\n")
	} else {
		GpibError(device, brdIndx, "Ibwrt Error")
	}

	Ibrd(device, buffer) // Read up to 100 bytes from the device
	if (ThreadIbsta() & ERR) == 0 {
		fmt.Println("Ibrd success...\n")
		fmt.Printf("buffer: %s", buffer)
		fmt.Println()
	} else {
		GpibError(device, brdIndx, "Ibrd Error")
	}

	//--- Uninit
	Ibonl(device, 0) // Take the device offline
	if (ThreadIbsta() & ERR) == 0 {
		fmt.Println("Ibonl success, device offline...\n")
	} else {
		GpibError(device, brdIndx, "Ibonl Error")
	}

	Ibonl(brdIndx, 0) // Take the interface offline
	if (ThreadIbsta() & ERR) == 0 {
		fmt.Println("Ibonl success, interface offline...\n")
	} else {
		GpibError(device, brdIndx, "Ibonl Error")
	}
}
