package wl2

import (
	"bytes"
	"fmt"
	"net"
)

type LEDStripController struct {
	Hostname   string
	Port       string
	Connection net.Conn
}

func NewStripController(hostname string) *LEDStripController {
	return &LEDStripController{
		Hostname: hostname,
		Port:     "5577", // Default
	}
}

func (lsc *LEDStripController) fullHost() string {
	return lsc.Hostname + ":" + lsc.Port
}

func (lsc *LEDStripController) Connect() error {
	if lsc.Connection == nil {
		conn, e := net.Dial("tcp", lsc.fullHost())
		if e != nil {
			fmt.Println(e.Error())
			return e
		}
		lsc.Connection = conn

		return nil
	}
	fmt.Printf("DEBUG::: Already connected: %+v\n", lsc.Connection)
	return nil
}

func (lsc *LEDStripController) Disconnect() error {
	if lsc.Connection != nil {
		return lsc.Connection.Close()
	}
	// error on nil close?
	return nil
}

func (lsc *LEDStripController) SetColor(red, green, blue byte) error {
	buffer := &bytes.Buffer{}

	// header
	buffer.WriteByte(byte(0x31))

	buffer.WriteByte(red)
	buffer.WriteByte(green)
	buffer.WriteByte(blue)

	// warm white level
	// ww * level / 100
	buffer.WriteByte(byte(0xFF))

	// cool white level
	// cw * level / 100
	buffer.WriteByte(byte(0xFF))

	// footer
	buffer.WriteByte(byte(0x0F))

	// calculate checksum and append
	var checkSum int
	for _, b := range buffer.Bytes() {
		checkSum += int(b)
	}
	ckSum := checkSum & 0xFF
	buffer.WriteByte(byte(ckSum))

	_, err := lsc.Connection.Write(buffer.Bytes())

	return err
}

// [name:"Soft White",	r: 255, g: 241, b: 224	],
// [name:"Daylight", 	r: 255, g: 255, b: 251	],
// [name:"Warm White", r: 255, g: 244, b: 229	],

// [name:"Red", 		r: 255, g: 0,	b: 0	],
// 	[name:"Green", 		r: 0, 	g: 255,	b: 0	],
// [name:"Blue", 		r: 0, 	g: 0,	b: 255	],

// 	[name:"Cyan", 		r: 0, 	g: 255,	b: 255	],
// [name:"Magenta", 	r: 255, g: 0,	b: 33	],
// [name:"Orange", 	r: 255, g: 102, b: 0	],

// [name:"Purple", 	r: 170, g: 0,	b: 255	],
// 	[name:"Yellow", 	r: 255, g: 255, b: 0	],
// 	[name:"Black", 		r: 0, 	g: 0, 	b: 0	],
// [name:"White", 		r: 255, g: 255, b: 255	]
// ]

func (lsc *LEDStripController) SetSoftWhite() error {
	r := byte(255)
	g := byte(241)
	b := byte(224)
	return lsc.SetColor(r, g, b)
}

func (lsc *LEDStripController) SetDaylight() error {
	r := byte(255)
	g := byte(255)
	b := byte(251)
	return lsc.SetColor(r, g, b)
}

func (lsc *LEDStripController) SetWarmWhite() error {
	r := byte(255)
	g := byte(244)
	b := byte(229)
	return lsc.SetColor(r, g, b)
}

func (lsc *LEDStripController) SetCyan() error {
	r := byte(0)
	g := byte(255)
	b := byte(255)
	return lsc.SetColor(r, g, b)
}

func (lsc *LEDStripController) SetMagenta() error {
	r := byte(255)
	g := byte(0)
	b := byte(33)
	return lsc.SetColor(r, g, b)
}

func (lsc *LEDStripController) SetOrange() error {
	r := byte(255)
	g := byte(102)
	b := byte(0)
	return lsc.SetColor(r, g, b)
}

func (lsc *LEDStripController) SetPurple() error {
	r := byte(170)
	g := byte(0)
	b := byte(255)
	return lsc.SetColor(r, g, b)
}

func (lsc *LEDStripController) SetYellow() error {
	r := byte(255)
	g := byte(255)
	b := byte(0)
	return lsc.SetColor(r, g, b)
}

func (lsc *LEDStripController) TurnOff() error {
	buffer := &bytes.Buffer{}

	// Turn Off
	//[0x71, 0x24, 0x0F, 0xA4]
	buffer.WriteByte(byte(0x71))
	buffer.WriteByte(byte(0x24))
	buffer.WriteByte(byte(0x0F))
	buffer.WriteByte(byte(0xA4))

	_, err := lsc.Connection.Write(buffer.Bytes())

	return err
}

func (lsc *LEDStripController) TurnOn() error {
	buffer := &bytes.Buffer{}

	// Turn On
	//[0x71, 0x23, 0x0F, 0xA3]
	buffer.WriteByte(byte(0x71))
	buffer.WriteByte(byte(0x23))
	buffer.WriteByte(byte(0x0F))
	buffer.WriteByte(byte(0xA3))

	_, err := lsc.Connection.Write(buffer.Bytes())

	return err
}
