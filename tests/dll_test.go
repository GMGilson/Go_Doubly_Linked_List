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
	list := dll.MakeEmptyDLL()
	if list.Peek() != ""{
		t.Errorf("Pop from empty list. Expected empty string, Recieved: %q", list.Pop())
	}
}

func TestPopFromList(t *testing.T){
	list := dll.MakeEmptyDLL()

	list.Append("Stephanie")
	list.Append("Nikos")

	temp := list.Pop()
	if temp != "Nikos"{
		t.Errorf("Pop from list. Expected: 'Nikos, Recieved: %q", temp)
	}
}

func TestPopFrontFromEmptyList(t *testing.T){
	list := dll.MakeEmptyDLL()
	if list.Peek() != ""{
		t.Errorf("Pop from empty list. Expected empty string, Recieved: %q", list.PopFront())
	}
}

func TestPopFrontFromList(t *testing.T){
	list := dll.MakeEmptyDLL()

	list.AppendFront("Stephanie")
	list.AppendFront("Nikos")

	temp := list.PopFront()
	if temp != "Nikos"{
		t.Errorf("Pop from list. Expected: 'Nikos, Recieved: %q", temp)
	}
}

func TestIncrement(t *testing.T) {
	list := dll.MakeEmptyDLL()

	list.AppendFront("Stephanie")
	if list.Len() != 1{
		t.Errorf("AppendFront to empty list did not increment length")
	}

	list.AppendFront("Nikos")
	if list.Len() != 2 {
		t.Errorf("AppendFront to list did not increment length")
	}

	list = dll.MakeEmptyDLL()

	list.Append("Will")
	if list.Len() != 1 {
		t.Errorf("Append to empty list did not increment length")
	}

	list.Append("Danny")
	if list.Len() != 2{
		t.Errorf("Append to list did not increment length")
	}


}

func TestDecrement(t *testing.T){
	list := dll.MakeEmptyDLL()

	list.AppendFront("Stephanie")
	list.AppendFront("Nikos")
	list.Pop()
	if list.Len() != 1{
		t.Errorf("Failed to decrement after pop")
	}
	list.Pop()
	if list.Len() != 0{
		t.Errorf("Unexpected Change to list length after pop from empty list. Expected: 0, Recieved: %q", list.Len())
	}
	if list.Len() < 0 {
		t.Errorf("Decremented too far after pop from empty list")
	}

	list.AppendFront("Will")
	list.AppendFront("Danny")
	list.PopFront()

	if list.Len() != 1 {
		t.Errorf("Failed to decrement after PopFront")
	}
	list.PopFront()
	if list.Len() != 0 {
		t.Errorf("Unexpected Change to list length after popFront from empty list. Expected: 0, Recieved: %q", list.Len())
	}
	if list.Len() < 0 {
		t.Errorf("Decremented too far after popFront from empty list")
	}
}

func TestIncDec(t *testing.T)  {
	list := dll.MakeEmptyDLL()
	for i := 1; i < 10; i++{
		list.AppendFront(string(i))
		if list.Len() != i{
			t.Errorf("List Length desync with loop iteration. Test: AppendFront. Expected: %d, Recieved: %d", i, list.Len())
		}
	}
	i := list.Len()
	for !list.IsEmpty(){
		i--
		list.PopFront()
		if list.Len() != i {
			t.Errorf("List Length desync with loop iteration. Test: PopFront. Expected: %d, Recieved: %d", i, list.Len())
		}
	}

	for i := 1; i < 10; i++{
		list.Append(string(i))
		if list.Len() != i{
			t.Errorf("List Length desync with loop iteration. Test: Append. Expected: %d, Recieved: %d", i, list.Len())
		}
	}
	i = list.Len()
	for !list.IsEmpty(){
		i--
		list.Pop()
		if list.Len() != i {
			t.Errorf("List Length desync with loop iteration. Test: Pop. Expected: %d, Recieved: %d", i, list.Len())
		}
	}
}