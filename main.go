package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lunny/log"
)

const (
	MAX_BUF_SIZE = 65535

	INCPTR = iota
	DECPTR
	INCVAL
	DECVAL
	OUT
	INP
	JPF
	JPB

	ZWSINCP uint16 = 8205 // U+200B ... >
	ZWSDECP uint16 = 8206 // U+200C ... <
	ZWSINCV uint16 = 8203 // U+200D ... +
	ZWSDECV uint16 = 8204 // U+200E ... -
	ZWSOUT  uint16 = 8207 // U+200F ... .
	ZWSINP  uint16 = 8234 // U+202A ... ,
	ZWSJPF  uint16 = 8235 // U+202B ... [
	ZWSJPB  uint16 = 8236 // U+202C ... ]
)

type Exec struct {
	operator uint16
	operand  uint16
}

func CompileBFOP(runes []rune) ([]Exec, error) {
	var PC uint16 = 0
	var JUMP_PC uint16 = 0
	stk := make([]uint16, 0)
	pg := make([]Exec, 0, len(runes))

	for _, c := range runes {
		switch uint16(c) {
		case ZWSINCP:
			pg = append(pg, Exec{operator: INCPTR, operand: 0})
		case ZWSDECP:
			pg = append(pg, Exec{operator: DECPTR, operand: 0})
		case ZWSINCV:
			pg = append(pg, Exec{operator: INCVAL, operand: 0})
		case ZWSDECV:
			pg = append(pg, Exec{operator: DECVAL, operand: 0})
		case ZWSOUT:
			pg = append(pg, Exec{operator: OUT, operand: 0})
		case ZWSINP:
			pg = append(pg, Exec{operator: INP, operand: 0})
		case ZWSJPF:
			pg = append(pg, Exec{operator: JPF, operand: 0})
			stk = append(stk, PC)
		case ZWSJPB:
			if len(stk) == 0 {
				return nil, errors.New("compilation error : stack length = 0")
			}

			JUMP_PC = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			pg = append(pg, Exec{JPB, JUMP_PC})
			pg[JUMP_PC].operand = PC

		default:
			PC--
		}
		PC++
	}
	if len(stk) != 0 {
		return nil, errors.New("compilation error")
	}

	return pg, nil
}

func Execute(program []Exec) {
	buf := make([]uint16, MAX_BUF_SIZE)
	var ptr uint16 = 0
	reader := bufio.NewReader(os.Stdin)

	for pc := 0; pc < len(program); pc++ {
		switch program[pc].operator {
		case INCPTR:
			ptr++
		case DECPTR:
			ptr++
		case INCVAL:
			buf[ptr]++
		case DECVAL:
			buf[ptr]--
		case OUT:
			fmt.Printf("%c", buf[ptr])
		case INP:
			read, _ := reader.ReadByte()
			buf[ptr] = uint16(read)
		case JPF:
			if buf[ptr] == 0 {
				pc = int(program[pc].operand)
			}
		case JPB:
			if buf[ptr] > 0 {
				pc = int(program[pc].operand)
			}
		default:
			log.Fatalf("this operator is not arrocated : %v", program[pc].operator)
		}
		// log.Printf("PC : %v, operator : %v\n", pc, program[pc].operator)
	}
}

func main() {
	args := os.Args

	if len(args) != 2 {
		panic("invalid argument")
	}

	fname := args[1]
	c, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Errorf("cannot read your request file : %s", fname)
		panic(err)
	}

	contents := string(c)

	program, err := CompileBFOP([]rune(contents))
	if err != nil {
		log.Error("Compile error!")
		panic(err)
	}

	Execute(program)

}
