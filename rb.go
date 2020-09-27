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

// RBNode is a Single Node in a Red Black Tree
type RBNode struct {
	Value  int
	Parent *RBNode
	Left   *RBNode
	Right  *RBNode
	Colour string
}

// RB is just a container for the tree structure giving a target for methods
type RB struct {
	Head *RBNode
}

func rbInsertNode(b *RB, n *RBNode) {
	var old *RBNode = nil
	var current *RBNode = b.Head
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

func (b *RB) addNumber(val int) {
	newNode := Node{
		Value: val,
	}
	insertNode(b, &newNode)
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
		// for _, n := range bst.flatten() {
		// 	fmt.Println(n.Value)
		// }
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
