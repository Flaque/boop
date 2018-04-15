package runtime

type Tree struct {
	value  *Frame
	parent *Tree
}

func NewTree(parent *Tree) Tree {
	f := NewFrame()

	return Tree{
		&f,
		parent,
	}
}

func (t *Tree) Get(label string) (interface{}, error) {
	return "", nil
}

func (t *Tree) Set(label string, value interface{}) {
	// Do something
}
