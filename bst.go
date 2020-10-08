package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Node is a Single Node in a Binary Tree
type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

// BST is just a container for the tree structure giving a target for methods
type BST struct {
	Head *Node
}

func (b *BST) insertNode(n *Node) {
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
	b.insertNode(&newNode)
}

// flatten outputs the nodes to a list, breadth first
func (b *BST) flatten() []*Node {
	Queue := make([]*Node, 1, 16)
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
	depth = 0
	// pass pointer to max to be used in recursive function
	getD(b.Head, &depth)
	return
}

func getD(n *Node, max *int, d ...int) int {
	if len(d) == 0 {
		d = append(d, 0)
	}
	if *max < d[0] {
		*max = d[0]
	}
	if n != nil {
		getD(n.Left, max, d[0]+1)
		getD(n.Right, max, d[0]+1)
	}
	return *max
}

func (b *BST) get(val int) (n *Node) {
	n = b.Head
	for n != nil && val != n.Value {
		if val < n.Value {
			n = n.Left
		} else if val > n.Value {
			n = n.Right
		}
	}
	return
}

func (b *BST) delete(n *Node) {
	if n.Left == nil {
		fmt.Println("n.Parent", n.Parent)
		*n.Parent = nil
		fmt.Println("n.Parent", n.Parent)

	} else if n.Right == nil {
		n.Parent = n.Left
	} else if n.Right == n.successor() {
		n.Right.Left = n.Left
		n = n.Right
	} else {
		s := n.successor()
		s.Parent.Left = s.Right
		s.Left = n.Left
		n = s
	}
}

func (b *BST) deleteNum(num int) {
	node := b.get(num)
	b.delete(&node)
}

func treeMin(n *Node) (min *Node) {
	min = n
	for min.Left != nil {
		min = min.Left
	}
	return
}

func treeMax(n *Node) (max *Node) {
	max = n
	for max.Right != nil {
		max = max.Right
	}
	return
}

func (n *Node) successor() (suc *Node) {
	suc = n
	if suc.Right != nil {
		suc = treeMin(suc.Right)
	} else {
		for suc == suc.Parent.Right {
			suc = suc.Parent
		}
		suc = suc.Parent
	}
	return
}

func main() {
	bst := BST{}
	bst.inOrderTW()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Binary Tree")

	for {
		re := regexp.MustCompile(`\d+`)
		fmt.Println("Options:")
		fmt.Println("(1) Add to tree")
		fmt.Println("(2) Delete from tree")
		fmt.Println("(3) Pre-populate tree")
		option, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		option = string(re.Find([]byte(option)))
		fmt.Print("Give number: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		stripped := string(re.Find([]byte(text)))
		num, err := strconv.Atoi(stripped)
		if err != nil {
			log.Fatal(err)
		}
		switch option {
		case "1":
			bst.addNumber(num)
		case "2":
			bst.deleteNum(num)
		case "3":
			for _, val := range []int{20, 15, 25, 10, 17, 27, 8, 12, 6, 11, 13} {
				bst.addNumber(val)
			}
		default:
			return
		}
		bst.printTree()
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

func (b *BST) printTree() {
	currLayerIndex := 0
	newLayer := make([]*Node, 0)
	newLayer = append(newLayer, b.Head)
	// starting at index 0 create slices, of size 2^index, initialised with nil and nodes inserted in place
	for currLayerIndex < b.checkDepth() {
		// assign a trailing pointer
		currLayer := newLayer
		newLayer = make([]*Node, int(math.Pow(2, float64(currLayerIndex+1))))
		// create padding to ensure numbers in correct position
		spaces := strings.Repeat(" ", int(math.Pow(2, float64(b.checkDepth()-currLayerIndex))))
		for ind, val := range currLayer {
			var currVal string
			if val != nil {
				currVal = fmt.Sprint(val.Value)
			} else {
				currVal = " "
			}
			fmt.Print(spaces + currVal + spaces)
			if val != nil {
				if val.Left != nil {
					newLayer[ind*2] = val.Left
				}
				if val.Right != nil {
					newLayer[ind*2+1] = val.Right
				}
			}
		}
		fmt.Println("")
		// TODO Print out the tree structure
		// fmt.Print(strings.Repeat(" ", len(spaces)))
		// for _, val := range newLayer {
		// 	if val != nil {
		// 		fmt.Print(strings.Repeat("_", len(spaces)))
		// 	} else {
		// 		fmt.Print(strings.Repeat(" ", len(spaces)))
		// 	}
		// }
		// fmt.Println("")
		currLayerIndex++
	}
}
