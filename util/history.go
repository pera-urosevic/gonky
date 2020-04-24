package util

// Segment //
type Segment struct {
	items []float64
}

// ItemFunction //
type ItemFunction func(item float64) bool

// Items //
func (s Segment) Items() []float64 {
	return s.items
}

// ItemsLength //
func (s Segment) ItemsLength() int {
	return len(s.items)
}

// Count //
func (s Segment) Count(itemFunction ItemFunction) int {
	counter := 0
	for _, item := range s.items {
		if itemFunction(item) {
			counter++
		}
	}
	return counter
}

// Contains //
func (s Segment) Contains(itemFunction ItemFunction) bool {
	for _, item := range s.items {
		if itemFunction(item) {
			return true
		}
	}
	return false
}

// History //
type History struct {
	id          string
	lastItem    float64
	lastSegment int
	segments    []Segment
}

// HistoryCreate //
func HistoryCreate(size int, id string) *History {
	segments := make([]Segment, size)
	for _, segment := range segments {
		segment.items = []float64{}
	}
	h := History{
		id:          id,
		lastSegment: size - 1,
		segments:    segments,
	}
	return &h
}

func (h *History) recycle(id string) {
	items := []float64{}
	segment := Segment{
		items: items,
	}
	h.segments = h.segments[1:]
	h.segments = append(h.segments, segment)
	h.id = id
}

// Add //
func (h *History) Add(id string, item float64) {
	if h.id != id {
		h.recycle(id)
	}
	h.lastItem = item
	h.segments[h.lastSegment].items = append(h.segments[h.lastSegment].items, item)
}

// Segments //
func (h *History) Segments() []Segment {
	return h.segments
}

// Last //
func (h *History) Last() float64 {
	return h.lastItem
}
