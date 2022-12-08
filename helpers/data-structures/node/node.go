package node

type Node struct {
	Id       string
	Value    int
	Parent   *Node
	Children []*Node
}

func NewNode(id string, value int, parent *Node) *Node {
	return &Node{
		Id:     id,
		Value:  value,
		Parent: parent,
	}
}

func (n *Node) IsRoot() bool {
	return n.Parent == nil
}

func (n *Node) HasChildren() bool {
	return len(n.Children) > 0
}

func (n *Node) IsLastParent() bool {
	if n.IsRoot() || !n.HasChildren() {
		return false
	}

	for _, child := range n.Children {
		if child.HasChildren() {
			return false
		}
	}

	return true
}

func (n *Node) FindChildById(id string) *Node {
	if !n.HasChildren() {
		return nil
	}

	// shallow
	for _, child := range n.Children {
		if child.Id == id {
			return child
		}
	}

	return nil
}

func (n *Node) AddChild(child *Node) {
	if n.FindChildById(child.Id) != nil {
		return
	}

	n.Children = append(n.Children, child)
}

func (n *Node) RemoveChild(child *Node) {
	for i, c := range n.Children {
		if c == child {
			n.Children = append(n.Children[:i], n.Children[i+1:]...)
			break
		}
	}
}

func (n *Node) ComputeValue() int {
	if !n.HasChildren() {
		return n.Value
	}

	totalSize := 0
	for _, child := range n.Children {
		totalSize += child.ComputeValue()
	}

	return totalSize
}
