package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Node is a Single Node in a Binary Tree
type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

type BST struct {
	Head *Node
}

func insertNode(b *BST, n *Node) {
	var old *Node = nil
	var current *Node = b.Head
	for current != nil {
		// set trailing pointer
		old = current
		// reassign the current position to a child
		if n.Value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	if old == nil {
		b.Head = n
	} else if n.Value < old.Value {
		old.Left = n
	} else {
		old.Right = n
	}
	n.Parent = old
}

func (b *BST) addNumber(val int) {
	newNode := Node{
		Value: val,
	}
	insertNode(b, &newNode)
}

// flatten outputs the nodes to a list, breadth first
func (b *BST) flatten() []*Node {
	Queue := make([]*Node, 1, 16)
	Queue[0] = b.Head
	index := 0
	for Queue[index] != nil {
		if Queue[index].Left != nil {
			Queue = append(Queue, Queue[index].Left)
		}
		if Queue[index].Right != nil {
			Queue = append(Queue, Queue[index].Right)
		}
		index++
		if index >= len(Queue) {
			break
		}
	}
	return Queue
}

func (b *BST) checkDepth() (depth int) {
	depth = getD(b.Head)
	return
}

func getD(n *Node, d ...int) int {
	if len(d) == 0 {
		d = append(d, 0)
	}
	if n != nil {
		getD(n.Left, d[0]+1)
		getD(n.Right, d[0]+1)
	}
	return d[0]
}

func main() {
	bst := BST{}
	bst.inOrderTW()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Binary Tree Add")

	for {
		fmt.Print("Insert: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		re := regexp.MustCompile(`\d+`)
		stripped := string(re.Find([]byte(text)))
		num, err := strconv.Atoi(stripped)
		if err != nil {
			log.Fatal(err)
		}
		bst.addNumber(num)
		bst.inOrderTW()
	}
}

func (b *BST) pprint() {
	for _, val := range b.flatten() {
		fmt.Println(val.Value)
	}
}

func (b *BST) inOrderTW() {
	tw(b.Head)
}

func tw(n *Node) {
	if n != nil {
		tw(n.Left)
		println(n.Value)
		tw(n.Right)
	}
}
