package runtime

type Frame struct {
	hash map[string]interface{}
}

func NewFrame() Frame {
	return Frame{make(map[string]interface{})}
}

func (f *Frame) getFromHash(key string) interface{} {
	return f.hash[key]
}

func (f *Frame) Get(key string) interface{} {
	owner := f.Resolve(key)

	if owner == nil {
		return nil
	}

	return owner.getFromHash(key)
}

func (f *Frame) Set(key string, value interface{}) {
	f.hash[key] = value
}

// Search through frame pointers until you find the original owner
func (f *Frame) Resolve(key string) *Frame {
	fr, ok := f.getFromHash(key).(*Frame)

	if ok {
		return fr.Resolve(key)
	}

	val := f.getFromHash(key)

	if val == nil {
		return nil
	}

	return f
}
