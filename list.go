package dll

type Dll struct {
	head *node
	tail *node
	len int
}

//region Exported functions

func MakeEmptyDLL() *Dll {
	sentinel := new(node)

	sentinel.data = "sentinel"
	sentinel.next = nil
	sentinel.prev = nil

	return &Dll{
		head: sentinel,
		tail: sentinel,
		len:  0,
	}
}

func (self *Dll) IsEmpty() bool{
	return self.len == 0
}

//creates new node to append to the list
func (self *Dll) Append(str string){
	newNode := new(node)
	newNode.data = str
	self.pushBack(newNode)
}

//creates a new node to enqueue to the list
func (self *Dll) AppendFront(str string){
	newNode := new(node)
	newNode.data = str
	self.enqueue(newNode)
}

func (self *Dll) Pop() string{

	if self.len == 0 {
		return ""
	}

	return self.popBack().data

}

func (self *Dll) PopFront() string {
	return self.dequeue().data
}

func (self *Dll) Len() int{
	return self.len
}

func (self *Dll) Peek() string{
	if self.len == 0{
		return ""
	}

	return self.head.next.data
}

//endregion

//region Private

//push node to the back of the list
func (self *Dll) pushBack(n *node){
	//special empty list case
	if self.len == 0 {
		n.setNext(self.tail)
		n.setPrev(self.head)
		self.tail.setPrev(n)
		self.head.setNext(n)
	}
	//general case
	if self.len > 0 {
		n.setNext(self.tail)
		n.setPrev(self.tail.prev)
		self.tail.prev.setNext(n)
		self.tail.setPrev(n)
	}
	self.len++

}

//remove and return node at the back of the list
func (self *Dll) popBack() *node{

	ret := self.tail.prev
	temp := ret.prev
	temp.setNext(self.tail)
	self.tail.setPrev(temp)

	ret.setPrev(nil)
	ret.setNext(nil)

	self.len -= 1

	return ret
}

//insert node n at the front of the list
func (self *Dll) enqueue(n *node){
	//general case
	if self.len > 0 {
		//set the new node's pointers
		n.setPrev(self.head)
		n.setNext(self.head.getNext())
		//captures the node after the head and sets its prev to the new node
		temp := self.getHead().getNext()
		temp.setPrev(n)
		//points the head to the new node
		self.getHead().setNext(n)
	}
	//special case: Empty List
	if self.len == 0{
		//set the new node's pointers
		n.setNext(self.tail)
		n.setPrev(self.head)
		//point the sentinel to the new node
		self.head.setNext(n)
		self.tail.setPrev(n)

	}
	//increase the length of the list
	self.len++
}

//remove and return front node of the list
func (self *Dll)dequeue() *node{
	ret := self.head.getNext()

	//point the head and the node after the return node to each other
	temp := self.head.next.next
	temp.setPrev(self.head)
	self.head.setNext(temp)

	//remove dangling pointers from the return node
	ret.next = nil
	ret.prev = nil

	self.len -= 1

	return ret

}

//return pointer to the sentinel of the list
//pragma: does exactly the same as getTail() but is used for code clarity in functions pertaining to the "head"
func (self *Dll) getHead() *node  {
	return self.head
}

//returns a pointer to the sentinel of the list
//pragma: does exactly the same of getHead() but is used for code clarity in functions pertaining to the "tail"
func (self *Dll) getTail() *node  {
	return self.tail
}

//endregion
