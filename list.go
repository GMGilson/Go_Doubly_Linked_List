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

func (self *Dll) Append(str string){
	newNode := new(node)
	newNode.data = str
	self.pushBack(newNode)
}

func (self *Dll) AppendFront(str string){
	newNode := new(node)
	newNode.data = str
	self.enqueue(newNode)
}

func (self *Dll) Pop() string{
	return self.popBack().data

}

func (self *Dll) PopFront() string {
	return self.dequeue().data
}

func (self *Dll) Len() int{
	return self.len
}

func (self *Dll) Peek() string{
	return self.head.next.data
}

//endregion

//region Private
func (self *Dll) pushBack(n *node){

	if self.len == 0 {
		n.setNext(self.tail)
		n.setPrev(self.head)
		self.tail.setPrev(n)
		self.head.setNext(n)
	}

	if self.len > 0 {
		n.setNext(self.tail)
		n.setPrev(self.tail.prev)
		self.tail.prev.setNext(n)
		self.tail.setPrev(n)
	}
	self.len++

}

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

//pushes n to the front of the list
func (self *Dll) enqueue(n *node){
	//creates a new node and points it to the head and the item after the head
	n.setPrev(self.head)
	n.setNext(self.head.getNext())
	//captures the node after the head and sets its prev to the new node
	temp := self.getHead().getNext()
	temp.setPrev(n)
	//points the head to the new node
	self.getHead().setNext(n)
	//increase the length of the list
	self.len++
}

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

func (self *Dll) getHead() *node  {
	return self.head
}

func (self *Dll) getTail() *node  {
	return self.tail
}

//endregion
