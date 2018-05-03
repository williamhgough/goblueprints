package meander_test

import (
	"testing"

	"github.com/cheekybits/is"
	"github.com/williamhgough/goblueprints/meander"
)

func TestParseCost(t *testing.T) {
	is := is.New(t)
	is.Equal(meander.Cost1, meander.ParseCost("$"))
	is.Equal(meander.Cost2, meander.ParseCost("$$"))
	is.Equal(meander.Cost3, meander.ParseCost("$$$"))
	is.Equal(meander.Cost4, meander.ParseCost("$$$$"))
	is.Equal(meander.Cost5, meander.ParseCost("$$$$$"))
}

func TestParseCostRange(t *testing.T) {
	is := is.New(t)

	var l *meander.CostRange
	l = meander.ParseCostRange("$$...$$$")
	is.Equal(l.From, meander.Cost2)
	is.Equal(l.To, meander.Cost3)

	l = meander.ParseCostRange("$...$$$$")
	is.Equal(l.From, meander.Cost1)
	is.Equal(l.To, meander.Cost4)
}

func TestCostRangeString(t *testing.T) {
	is := is.New(t)
	is.Equal("$$...$$$$", (&meander.CostRange{
		From: meander.Cost2,
		To:   meander.Cost4,
	}).String())
}
