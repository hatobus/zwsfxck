package main

import (
	"errors"
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

	ZWSINCP = []rune("a")[0]
	ZWSDECP = []rune("b")[0]
	ZWSINCV = []rune("c")[0]
	ZWSDECV = []rune("d")[0]
	ZWSOUT  = []rune("e")[0]
	ZWSINP  = []rune("f")[0]
	ZWSJPF  = []rune("g")[0]
	ZWSJPB  = []rune("h")[0]
)

type Exec struct {
	operator uint16
	operand  uint16
}

func CompileBFOP(runes []rune) ([]Exec, error) {
	// panic("not impl")

	var PC uint16 = 0
	var JUMP_PC uint16 = 0
	stk := make([]uint16, 0)
	pg := make(Exec, len(runes))

	for _, c := range runes {
		switch c {
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
			PC = append(PC, Exec{JPB, JUMP_PC})
			PC[JUMP_PC].operand = PC

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

}
