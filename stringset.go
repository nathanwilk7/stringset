package stringset

type Stringset interface {
	Add (string)
	Contains (string) bool
	Remove (string)
	Size () int
	ToSlice () []string
}

func NewStringset () Stringset {
	return EmptyInterfaceStringset{}
}

type EmptyInterfaceStringset struct {
	stringMap map[string]interface{}
	empty interface{}
}

func (ss EmptyInterfaceStringset) Add (s string) {
	ss.stringMap[s] = ss.empty
}

func (ss EmptyInterfaceStringset) Contains (s string) bool {
	_, ok := ss.stringMap[s]
	return ok
}

func (ss EmptyInterfaceStringset) Remove (s string) {
	delete(ss.stringMap, s)
}

func (ss EmptyInterfaceStringset) Size () int {
	return len(ss.stringMap)
}

func (ss EmptyInterfaceStringset) ToSlice () []string {
	slice := make([]string, len(ss.stringMap))
	for s := range ss.stringMap {
		slice = append(slice, s)
	}
	return slice
}
