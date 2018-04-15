package runtime

type Tree struct {
	value  *Frame
	parent *Tree
}

func (t *Tree) Get(label string) (interface{}, error) {
	return "", nil
}

func (t *Tree) Set(label string, value interface{}) {
	// Do something
}
