package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
)

func TestGiveTypes(t *testing.T) {
	var compl complex64
	expected := fmt.Sprintf("%T", compl)
	if output := GiveTypes(compl); output.String() != expected {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
}

func TestStringReturn(t *testing.T) {
	int1 := 1
	int2 := 0x2A
	pi := 3.14
	expected := strconv.Itoa(int1) + strconv.Itoa(int2) + fmt.Sprintf("%.2f", pi)
	if output := StringReturn(int1) + StringReturn(int2) + StringReturn(pi); output != expected {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
}

func TestFourth(t *testing.T) {
	str := "it is all runes"
	strRunes := Fourth(str)
	if strRunes == nil {
		t.Error("Line is nill")
	}
}

func TestFifths(t *testing.T) {
	str := []rune("inTheMiddle")
	hasher := sha256.New()
	hasher.Write([]byte("inThego-2024Middle"))
	expected := hex.EncodeToString(hasher.Sum(nil))
	if output := Fifths(str); output != expected {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
}
