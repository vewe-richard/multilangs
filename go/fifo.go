package main

import "fmt"

type Node struct{
	val int
}

func (n *Node) String() string {
	return fmt.Sprint(n.val)
}

type ififo interface {
	init(int)
	push(Node)
	pop() *Node
}

type myfifo struct {
	queue [] Node
	len int
	count int
	head int
	tail int
}

func (f *myfifo) init(len int) {
	f.len = len
	f.count = 0
	f.head = 0
	f.tail = 0
	f.queue = make([] Node, len, len)
}

func (f *myfifo) push(n Node){

	if f.count >= f.len {
		fmt.Println("queue is full")
		return
	}

	if f.head >= f.len {
		f.head = 0
	}

	f.queue[f.head] = n
	f.count += 1
	f.head += 1
}

func (f *myfifo) pop() *Node {
	if f.count == 0 {
		return nil
	}

	if f.tail >= f.len {
		f.tail = 0
	}

	k := f.tail

	f.tail += 1
	f.count -= 1

	return &(f.queue[k])
}

func main(){
	var f myfifo

	f.init(5)

	fmt.Println(f)
	f.push(Node{1})
	f.push(Node{4})
	f.push(Node{41})
	f.push(Node{101})
	f.push(Node{339})

	fmt.Println(f.pop(), f.pop())
	fmt.Println(f)


}
