package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	MEM_SIZE = 30000
)

func main() {
	if len(os.Args) != 2 {
		panic("ファイル名を一つだけ指指してください。")
	}

	filename := os.Args[1]
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	s, ret := brainfuck(string(bytes))
	fmt.Println(ret)
	fmt.Println(s, "steps")
}

func brainfuck(program string) (int, string) {
	steps := 0
	result := ""

	mem := make([]byte, MEM_SIZE)
	ptr := 0

	n := len(program)
	for pc := 0; pc < n; pc++ {
		steps++
		c := program[pc]

		switch c {
		case '>':
			ptr = (ptr + 1) % MEM_SIZE
		case '<':
			ptr = (ptr - 1 + MEM_SIZE) % MEM_SIZE
		case '+':
			mem[ptr]++
		case '-':
			mem[ptr]--
		case '.':
			result += string(mem[ptr])
		// case ',':
		case '[':
			if mem[ptr] == 0 {
				dep := 1
				for dep > 0 {
					pc++
					if program[pc] == '[' {
						dep++
					} else if program[pc] == ']' {
						dep--
					}
				}
			}
		case ']':
			if mem[ptr] != 0 {
				dep := 1
				for dep > 0 {
					pc--
					if program[pc] == ']' {
						dep++
					} else if program[pc] == '[' {
						dep--
					}
				}
			}
		}

		// fmt.Printf("%v %v %v %v %v\n", steps, pc, string(c), ptr, mem[ptr])
		// if steps == 300 {
		// 	break
		// }
	}

	return steps, result
}
