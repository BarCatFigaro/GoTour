package main

import "golang.org/x/tour/tree"
import "fmt"

//Walk with brevity
/*
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var Walker(t *tree.Tree)
	func Walker(t *tree.Tree) {
		if (t == nil) { return }
		Walker(t.Left)
		ch <- t.Value
		Walker(t.Right)
	}
	Walker(t)
}

func Same(t1, t2 *tree.Tree) bool {
    ch1, ch2 := make(chan int), make(chan int)

    go Walk(t1, ch1)
    go Walk(t2, ch2)

    for {
        v1,ok1 := <- ch1
        v2,ok2 := <- ch2

        if v1 != v2 || ok1 != ok2 {
            return false
        }

        if !ok1 {
            break
        }
    }

    return true
}
*/

//helper recursive function for walking
func NaturalWalk(t *tree.Tree, ch chan int) {
	right, left := t.Right, t.Left
	if (left != nil) {
		NaturalWalk(left, ch)
	}
	ch <- t.Value
	if (right != nil) {
		NaturalWalk(right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	NaturalWalk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)			
		
	for i := range ch1 {
		select {
			case v := <-ch2: 
				if (i == v) {
					return true
				}
				
			default:
				break
		}
	}
	return false
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
