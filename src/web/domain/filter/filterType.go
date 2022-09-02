package filter

type FilterType int

const (
	Latest FilterType = iota
	Best
)

func (f *FilterType) contains(filterType FilterType) bool {
	return filterType == Latest ||
		filterType == Best
}
