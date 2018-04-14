package runtime

type Frame struct {
	hash map[string]interface{}
}

func NewFrame() Frame {
	return Frame{make(map[string]interface{})}
}

func (f *Frame) Get(key string) interface{} {
	return f.hash[key]
}

func (f *Frame) Set(key string, value interface{}) {
	f.hash[key] = value
}
