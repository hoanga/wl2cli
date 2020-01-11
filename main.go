package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hoanga/wl2cli/ledstripcontroller"
)

func valToHexByte(val string) byte {

	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return byte(0x00)
	}

	return byte(i)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <red-value> <green-value> <blue-value>\n", os.Args[0])
		os.Exit(1)
	}

	redVal := valToHexByte(os.Args[1])
	greenVal := valToHexByte(os.Args[2])
	blueVal := valToHexByte(os.Args[3])

	ledStrip1 := wl2.NewStripController("172.16.4.59")
	ledStrip2 := wl2.NewStripController("172.16.4.58")
	ledStrip1.Connect()
	ledStrip2.Connect()

	ledStrip1.SetColor(redVal, greenVal, blueVal)
	ledStrip2.SetColor(redVal, greenVal, blueVal)

	ledStrip1.Disconnect()
	ledStrip2.Disconnect()
}
