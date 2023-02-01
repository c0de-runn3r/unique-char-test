package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func assert(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestUniqueChar(t *testing.T) {
	i, ok := uniqChar("testingtest")
	assert(t, i, "i")
	assert(t, ok, true)

	r, ok := uniqChar("aabbccdd")
	assert(t, r, "")
	assert(t, ok, false)

}

func TestUniqCharsInArrStrings(t *testing.T) {
	str := "The Tao gave birth to machine language.  Machine language gave birth to the assembler."
	s := strings.Split(str, " ")
	fmt.Println(string(uniqCharsFromArr(s)))
}

func TestAll(t *testing.T) {
	str := `The Tao gave birth to machine language.  Machine language gave birth
	to the assembler.
	The assembler gave birth to the compiler.  Now there are ten thousand
	languages.
	Each language has its purpose, however humble.  Each language
	expresses the Yin and Yang of software.  Each language has its place within
	the Tao.
	But do not program in COBOL if you can avoid it.
			-- Geoffrey James, "The Tao of Programming"`

	answ, _ := UniqOfUniqsChar(str)
	assert(t, string(answ), "m")
	anotherString := `C makes it easy for you to shoot yourself in the foot. C++ makes that harder, but when you do, it blows away your whole leg. (с) Bjarne Stroustrup`
	answ2, _ := UniqOfUniqsChar(anotherString)
	println(string(answ2))
}
