package mysort

type IntSlice []int
type StringSlice []string

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// int slice methods
func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// string slice methods
func (s StringSlice) Len() int { return len(s) }
func (s StringSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s StringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i - i) {
			return false
		}
	}
	return true
}
