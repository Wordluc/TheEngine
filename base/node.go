package base

type Node struct {
	objects []Object
	nodes   []Node
	pos     *Vec[int32]
}

func NewNode(p *Vec[int32]) Node {
	return Node{
		pos: p,
	}
}

func (n *Node) MoveTo(v Vec[int32]) {
	xNode, yNode := v.Get()
	xDest, yDest := v.Get()
	diffX, diffY := xDest-xNode, yDest-yNode
	n.ForEachObjects(func(o Object) error {
		o.GetPos().Add(NewVec(diffX, diffY))
		return nil
	})
}

func (n *Node) MoveBy(v Vec[int32]) {
	n.ForEachObjects(func(o Object) error {
		o.GetPos().Add(v)
		return nil
	})
	n.pos.Add(v)
}

func (n *Node) ForEachObjects(c func(o Object) error) (err error) {
	for i := range n.objects {
		err = c(n.objects[i])
		if err != nil {
			return err
		}
	}
	for i := range n.nodes {
		err = n.nodes[i].ForEachObjects(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) AddObject(o Object) {
	n.objects = append(n.objects, o)
}

func (n *Node) AddNode(a Node) {
	n.nodes = append(n.nodes, a)
}
