// time 2019/06/11
package main

import (
	"bytes"
	"fmt"
	"log"

	toml2 "github.com/BurntSushi/toml"
	toml1 "github.com/naoina/toml"
	toml0 "github.com/pelletier/go-toml"
	toml3 "github.com/shafreeck/toml"
)

type A struct {
	Str         string
	B           B
	B_nil_point *B     // naoina unsupport nil point
	in          *B     // pelletier over write
	in_str      string // pelletier over write
}

type B struct {
	Func func() `toml:"-"` // BurntSushi & shafreeck will not skip
}

func main() {
	pelletier()
	fmt.Println("toml0")
	shafreeck()
	fmt.Println("toml3")
	naoina()
	fmt.Println("toml1")
	burntsushi()
	fmt.Println("toml2")
}

func make_struct() *A {
	a := &A{}
	a.in = &B{}
	a.in_str = "adasds"
	return a
}

func check_struct(a *A) {
	if a.in == nil {
		log.Fatalln("in is nil")
	}
	if a.in_str != "adasds" {
		log.Fatalln("in_str is changed")
	}
}

// override even if not export filed
func pelletier() {
	a := make_struct()
	b, err := toml0.Marshal(a)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))

	err = toml0.Unmarshal(b, a)
	if err != nil {
		log.Fatalln(err)
	}
	check_struct(a)
}

// unsupport nil point
func naoina() {
	a := make_struct()
	b, err := toml1.Marshal(a)
	if err != nil {
		log.Fatalln(err)
	}
	err = toml1.Unmarshal(b, a)
	if err != nil {
		log.Fatalln(err)
	}
	check_struct(a)
}

// unsupport nest struct tag
func burntsushi() {
	a := make_struct()
	buffer := bytes.NewBuffer(nil)
	err := toml2.NewEncoder(buffer).Encode(a)
	if err != nil {
		log.Fatalln(err)
	}
	err = toml2.Unmarshal(buffer.Bytes(), a)
	if err != nil {
		log.Fatalln(err)
	}
	check_struct(a)
}

func shafreeck() {
	a := make_struct()
	b, err := toml3.Marshal(a)
	if err != nil {
		log.Fatalln(err)
	}
	err = toml3.Unmarshal(b, a)
	if err != nil {
		log.Fatalln(err)
	}
	check_struct(a)
}
