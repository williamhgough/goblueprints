package meander

import "strings"

// ParseCost allows us to access string representation
// of cost value.
func ParseCost(s string) Cost {
	return costStrings[s]
}

// CostRange holds the range of costs for places
type CostRange struct {
	From Cost
	To   Cost
}

func (r CostRange) String() string {
	return r.From.String() + "..." + r.To.String()
}

// ParseCostRange builds a new CostRange from string
func ParseCostRange(s string) *CostRange {
	segs := strings.Split(s, "...")
	return &CostRange{
		From: ParseCost(segs[0]),
		To:   ParseCost(segs[1]),
	}
}
