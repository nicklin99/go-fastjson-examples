package gofastjsonexamples

type TreeNode interface {
	GetParentID() string
	GetID() string
}

type Node[T TreeNode] struct {
	Data     T          `json:"data"`
	ID       string     `json:"id"`
	ParentID string     `json:"parentId"`
	Parent   *Node[T]   `json:"-"` // 忽略json，json报错
	Children []*Node[T] `json:"children"`
}

func (n *Node[T]) GetChildren() []*T {
	var children []*T
	if len(n.Children) > 0 {
		for _, node := range n.Children {
			children = append(children, &node.Data)
		}
	}
	return children
}
func (n *Node[T]) GetDescendants() []T {
	var children []T
	if n.Children != nil {
		for _, node := range n.Children {
			children = append(children, node.Data)
			children = append(children, node.GetDescendants()...)
		}
	}
	return children
}

type TreeManager[T TreeNode] struct {
	nodes map[string]*Node[T]
}

func NewTreeManager[T TreeNode](data []T) TreeManager[T] {
	var treeMap TreeManager[T]
	count := len(data)
	treeMap.nodes = make(map[string]*Node[T], count)
	for i := 0; i < count; i++ {
		nodeId := data[i].GetID()
		parentId := data[i].GetParentID()
		node := Node[T]{Data: data[i]}
		node.ID = nodeId
		if _, ok := treeMap.nodes[nodeId]; !ok {
			treeMap.nodes[nodeId] = &node
		}
		if parentId != "" {
			node.ParentID = parentId
		}
	}
	return treeMap
}
func (t *TreeManager[T]) GetTree() []*Node[T] {
	var root []*Node[T]
	for _, node := range t.nodes {
		nodeId := node.ParentID
		if len(nodeId) > 0 && nodeId != "0" {
			parent, ok := t.nodes[nodeId]
			if ok {
				parent.Children = append(parent.Children, node)
				node.Parent = parent
			}
		} else {
			root = append(root, node)
		}
	}
	return root
}
func (t *TreeManager[T]) GetNode(id string) *Node[T] {
	return t.nodes[id]
}

func MapNode[T TreeNode, R []*T](nodes []*Node[T], f func(i int, node *Node[T])) (children R) {
	for i, v := range nodes {
		children = append(children, &v.Data)
		f(i, v)
	}
	return children
}