package output

type State map[int]int

func (s *State) Keys() (keys []int) {
	for k := range *s {
		keys = append(keys, k)
	}
	return
}

type Output interface {
	Set(State) error
}
