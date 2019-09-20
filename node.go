package dll

type node struct {
	prev *node
	next *node
	data string
}

func (self *node) getNext() *node  {
	return self.next
}

func (self *node) getPrev() *node  {
	return self.next
}

func (self *node) setNext(n *node){
	self.next = n
}


func (self *node) setPrev(n *node){
	self.prev = n
}
