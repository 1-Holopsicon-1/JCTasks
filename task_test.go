package main

import (
	"reflect"
	"testing"
)

func TestStringIntMap_Add(t *testing.T) {
	var strIMap StringIntMap
	strIMap.NewStringIntMap()
	myMap := make(map[string]int)
	myMap["hello"] = 1
	strIMap.Add("hello", 1)
	if !reflect.DeepEqual(myMap, strIMap.mapa) {
		t.Errorf("Output %q not equal to expected %q", strIMap.mapa, myMap)
	}
}

func TestStringIntMap_Remove(t *testing.T) {
	var strIMap StringIntMap
	strIMap.NewStringIntMap()
	strIMap.Add("hello", 1)
	strIMap.Add("delete", 21312312)
	strIMap.Remove("delete")
	_, ok := strIMap.mapa["delete"]
	if ok {
		t.Errorf("Still exist %q", strIMap.mapa)
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	var strIMap StringIntMap
	strIMap.NewStringIntMap()
	myMap := make(map[string]int)
	myMap = strIMap.Copy()
	if reflect.DeepEqual(myMap, strIMap.mapa) {
		t.Errorf("Output %q not equal to expected %q", myMap, strIMap.mapa)
	}
}

func TestStringIntMap_Exists(t *testing.T) {
	var strIMap StringIntMap
	strIMap.NewStringIntMap()
	strIMap.Add("check", 123456)
	if !strIMap.Exists("check") {
		t.Error("Error cant find")

	}
}

func TestStringIntMap_Get(t *testing.T) {
	var strIMap StringIntMap
	strIMap.NewStringIntMap()
	strIMap.Add("check", 123456)
	if v, ok := strIMap.Get("check"); !ok || v != 123456 {
		t.Errorf("Output %q not equal to expected %q", v, 123456)

	}
}
