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

func TestAppendNonEmptyList(t *testing.T){
	list := dll.MakeEmptyDLL()
	list.Append("stephanie")
	list.Append("nikos")

	if list.Len() != 2 {
		t.Errorf("Appended to populated list. Expected len: 2, Actual: %d", list.Len())
	}
}

func TestAppendFrontEmptyList(t *testing.T){
	dllist := dll.MakeEmptyDLL()
	dllist.AppendFront("stephanie")
	if dllist.IsEmpty(){
		t.Errorf("Appended to empty list. Expected len = 1. List returns empty")
	}

	if dllist.Peek() != "stephanie" {
		t.Errorf("Unexpected data node at front of list. Expected: 'stephanie', Recieved: %q", dllist.Peek())
	}
}

func TestAppendFrontNonEmptyList(t *testing.T){
	list := dll.MakeEmptyDLL()
	list.AppendFront("stephanie")
	list.AppendFront("nikos")

	if list.Len() != 2 {
		t.Errorf("Appended to populated list. Expected len: 2, Actual: %d", list.Len())
	}

	if list.Peek() != "nikos" {
		t.Errorf("Unexpected data node at front of list. Expected: 'nikos', Recieved: %q", list.Peek())
	}
}

func TestPopFromEmptyList(t *testing.T){

}

func TestPopFromList(t *testing.T){

}

func TestPopFrontFromEmptyList(t *testing.T){

}

func TestPopFrontFromList(t *testing.T){

}