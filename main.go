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

	ZWSINCP uint16 = 8203 // U+200B ... +
	ZWSDECP uint16 = 8204 // U+200C ... -
	ZWSINCV uint16 = 8205 // U+200D ... >
	ZWSDECV uint16 = 8206 // U+200E ... <
	ZWSOUT  uint16 = 8207 // U+200F ... ,
	ZWSINP  uint16 = 8232 // U+2028 ... .
	ZWSJPF  uint16 = 8233 // U+2029 ... [
	ZWSJPB  uint16 = 8234 // U+202A ... ]
)

type Exec struct {
	operator uint16
	operand  uint16
}

func CompileBFOP(runes []rune) ([]Exec, error) {
	var PC uint16 = 0
	var JUMP_PC uint16 = 0
	stk := make([]uint16, 0)
	pg := make([]Exec, len(runes))

	for _, c := range runes {
		switch uint16(c) {
		case ZWSINCP:
			pg = append(pg, Exec{INCPTR, 0})
		case ZWSDECP:
			pg = append(pg, Exec{DECPTR, 0})
		case ZWSINCV:
			pg = append(pg, Exec{INCVAL, 0})
		case ZWSDECV:
			pg = append(pg, Exec{DECVAL, 0})
		case ZWSOUT:
			pg = append(pg, Exec{OUT, 0})
		case ZWSINP:
			pg = append(pg, Exec{INP, 0})
		case ZWSJPF:
			pg = append(pg, Exec{JPF, 0})
			stk = append(stk, PC)
		case ZWSJPB:
			if len(stk) == 0 {
				return nil, errors.New("compilation error")
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
			rval, _ := reader.ReadByte()
			buf[ptr] = uint16(rval)
		case JPF:
			if buf[ptr] == 0 {
				pc = int(program[pc].operand)
			}
		case JPB:
			if buf[ptr] > 0 {
				pc = int(program[pc].operand)
			}
		default:
			// panic("this operator is not arrocated : " + strconv.Itoa(program[pc].operand))
		}
	}
	fmt.Println()
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

	log.Println(program)

}
