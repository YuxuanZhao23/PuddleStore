package main

import "testing"

func TestExample(t *testing.T){
	if toTest("", 0) != 0 {
		t.Errorf("Error when str is nil")
	}
	if toTest("distributed", 0) != 1380 {
		t.Errorf("Error when str is distributed")
	}
	if toTest("golang", 0) != 5 {
		t.Errorf("Error when str is golang and num is smaller than 10")
	}
	if toTest("golang", 20) != 1 {
		t.Errorf("Error when str is golang and num is larger or equal to than 10")
	}
	if toTest("other", 0) != 10 {
		t.Errorf("Error when str is not nil and not distributed and golang")
	}
}

func BenchmarkExample(b *testing.B){
	for i := 0; i < b.N; i++ {
		toTest("", i)
		toTest("distributed", i)
		toTest("golang", i)
		toTest("other", i)
	}
}