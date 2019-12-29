package main

import (
	clr "github.com/logrusorgru/aurora"
)

var defaultLineColor = clr.CyanFg
var defaultSpaceColor = clr.MagentaFg

var defaultNULLEOFColor = clr.RedFg

var defaultPrintableASCIIColor = clr.BlueFg

// Default colors for bytes 0-255
var defaultCharacterColors = map[uint8]clr.Color{
	0:   defaultNULLEOFColor,
	10:  defaultLineColor,
	13:  defaultLineColor,
	32:  defaultSpaceColor,
	33:  defaultPrintableASCIIColor,
	34:  defaultPrintableASCIIColor,
	35:  defaultPrintableASCIIColor,
	36:  defaultPrintableASCIIColor,
	37:  defaultPrintableASCIIColor,
	38:  defaultPrintableASCIIColor,
	39:  defaultPrintableASCIIColor,
	40:  defaultPrintableASCIIColor,
	41:  defaultPrintableASCIIColor,
	42:  defaultPrintableASCIIColor,
	43:  defaultPrintableASCIIColor,
	44:  defaultPrintableASCIIColor,
	45:  defaultPrintableASCIIColor,
	46:  defaultPrintableASCIIColor,
	47:  defaultPrintableASCIIColor,
	48:  defaultPrintableASCIIColor,
	49:  defaultPrintableASCIIColor,
	50:  defaultPrintableASCIIColor,
	51:  defaultPrintableASCIIColor,
	52:  defaultPrintableASCIIColor,
	53:  defaultPrintableASCIIColor,
	54:  defaultPrintableASCIIColor,
	55:  defaultPrintableASCIIColor,
	56:  defaultPrintableASCIIColor,
	57:  defaultPrintableASCIIColor,
	58:  defaultPrintableASCIIColor,
	59:  defaultPrintableASCIIColor,
	60:  defaultPrintableASCIIColor,
	61:  defaultPrintableASCIIColor,
	62:  defaultPrintableASCIIColor,
	63:  defaultPrintableASCIIColor,
	64:  defaultPrintableASCIIColor,
	65:  defaultPrintableASCIIColor,
	66:  defaultPrintableASCIIColor,
	67:  defaultPrintableASCIIColor,
	68:  defaultPrintableASCIIColor,
	69:  defaultPrintableASCIIColor,
	70:  defaultPrintableASCIIColor,
	71:  defaultPrintableASCIIColor,
	72:  defaultPrintableASCIIColor,
	73:  defaultPrintableASCIIColor,
	74:  defaultPrintableASCIIColor,
	75:  defaultPrintableASCIIColor,
	76:  defaultPrintableASCIIColor,
	77:  defaultPrintableASCIIColor,
	78:  defaultPrintableASCIIColor,
	79:  defaultPrintableASCIIColor,
	80:  defaultPrintableASCIIColor,
	81:  defaultPrintableASCIIColor,
	82:  defaultPrintableASCIIColor,
	83:  defaultPrintableASCIIColor,
	84:  defaultPrintableASCIIColor,
	85:  defaultPrintableASCIIColor,
	86:  defaultPrintableASCIIColor,
	87:  defaultPrintableASCIIColor,
	88:  defaultPrintableASCIIColor,
	89:  defaultPrintableASCIIColor,
	90:  defaultPrintableASCIIColor,
	91:  defaultPrintableASCIIColor,
	92:  defaultPrintableASCIIColor,
	93:  defaultPrintableASCIIColor,
	94:  defaultPrintableASCIIColor,
	95:  defaultPrintableASCIIColor,
	96:  defaultPrintableASCIIColor,
	97:  defaultPrintableASCIIColor,
	98:  defaultPrintableASCIIColor,
	99:  defaultPrintableASCIIColor,
	100: defaultPrintableASCIIColor,
	101: defaultPrintableASCIIColor,
	102: defaultPrintableASCIIColor,
	103: defaultPrintableASCIIColor,
	104: defaultPrintableASCIIColor,
	105: defaultPrintableASCIIColor,
	106: defaultPrintableASCIIColor,
	107: defaultPrintableASCIIColor,
	108: defaultPrintableASCIIColor,
	109: defaultPrintableASCIIColor,
	110: defaultPrintableASCIIColor,
	111: defaultPrintableASCIIColor,
	112: defaultPrintableASCIIColor,
	113: defaultPrintableASCIIColor,
	114: defaultPrintableASCIIColor,
	115: defaultPrintableASCIIColor,
	116: defaultPrintableASCIIColor,
	117: defaultPrintableASCIIColor,
	118: defaultPrintableASCIIColor,
	119: defaultPrintableASCIIColor,
	120: defaultPrintableASCIIColor,
	121: defaultPrintableASCIIColor,
	122: defaultPrintableASCIIColor,
	123: defaultPrintableASCIIColor,
	124: defaultPrintableASCIIColor,
	125: defaultPrintableASCIIColor,
	126: defaultPrintableASCIIColor,
	255: defaultNULLEOFColor,
}
