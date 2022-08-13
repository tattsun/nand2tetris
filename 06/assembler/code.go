package main

import (
	"fmt"
	"strings"
)

type Code struct{}

func NewCode() *Code {
	return &Code{}
}

func (c *Code) Dest(mnemonic string) string {
	ret := 0b0

	if strings.Contains(mnemonic, "M") {
		ret = ret | 0b001
	}

	if strings.Contains(mnemonic, "D") {
		ret = ret | 0b010
	}

	if strings.Contains(mnemonic, "A") {
		ret = ret | 0b100
	}

	return fmt.Sprintf("%03b", ret)
}

func (c *Code) Comp(mnemonic string) string {
	switch mnemonic {
	case "0":
		return "0101010"
	case "1":
		return "0111111"
	case "-1":
		return "0111010"
	case "D":
		return "0001100"
	case "A":
		return "0110000"
	case "!D":
		return "0001101"
	case "!A":
		return "0110001"
	case "-D":
		return "0001111"
	case "-A":
		return "0110011"
	case "D+1":
		return "0011111"
	case "A+1":
		return "0110111"
	case "D-1":
		return "0001110"
	case "A-1":
		return "0110010"
	case "D+A":
		return "0000010"
	case "D-A":
		return "0010011"
	case "A-D":
		return "0000111"
	case "D&A":
		return "0000000"
	case "D|A":
		return "0010101"

	case "M":
		return "1110000"
	case "!M":
		return "1110001"
	case "-M":
		return "1110011"
	case "M+1":
		return "1110111"
	case "M-1":
		return "1110010"
	case "D+M":
		return "1000010"
	case "D-M":
		return "1010011"
	case "M-D":
		return "1000111"
	case "D&M":
		return "1000000"
	case "D|M":
		return "1010101"
	default:
		panic("unexpected mnemonic")
	}
}

func (c *Code) Jump(mnemonic string) string {
	switch mnemonic {
	case "JGT":
		return "001"
	case "JEQ":
		return "010"
	case "JGE":
		return "011"
	case "JLT":
		return "100"
	case "JNE":
		return "101"
	case "JLE":
		return "110"
	case "JMP":
		return "111"
	default:
		return "000"
	}
}
