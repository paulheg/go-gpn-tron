package strategies

type Quadtree struct {
	Arena
}

type Node struct {
	Parent   *Node
	Children []Node
}

// func (q *Quadtree) Generate() *Node {
// 	return
// }
