package dll

import (
	"github.com/GMGilson/dll"
	"testing"
)


func TestDllConstrutor(t *testing.T){
	dllist := dll.MakeEmptyDLL()
	if !dllist.IsEmpty() {
		t.Errorf("Empty List constructor returns non empty list")
	}
}

func TestAppendEmptyList(t *testing.T){
	dllist := dll.MakeEmptyDLL()
	dllist.Append("stephanie")
	if dllist.IsEmpty(){
		t.Errorf("Appended to empty list. Expected len = 1. List returns empty")
	}

	if dllist.Peek() != "stephanie" {
		t.Fail()
	}
}