package main

import (
	graphformatter "github.com/HappyR0b0t/graph-formatting/pkg"
)

func main() {
	interval := "HOUR"
	structs := []graphformatter.Transaction{}

	graph := map[int]int64{ 
		1: 1616026248,
		2: 1616019048,
		3: 1616022648,
		4: 1615889448,
		5: 1615871448,
		6: 1234545757, 
		7: 1613672577,
		8: 1615493354,
		9: 1614849048,
		10: 1613639545,
		11: 1610961145,
		12: 1615453945,
		13: 1615972345,
		14: 1615885945,
		15: 1615799545,
		16: 1615626745,
		17: 1616015448,
	}

	structs = graphformatter.SliceSorter(graphformatter.SliceFiller(structs, graph))

	switch interval {
	case "MONTH":
		graphformatter.TimestampToUnixTime(
			graphformatter.TimeDifferenceMonth(structs))
	case "WEEK":
		graphformatter.TimestampToUnixTime(
			graphformatter.TimeDifferenceWeek(structs))
	case "DAY":
		graphformatter.TimestampToUnixTime(
			graphformatter.TimeDifferenceDay(structs))
	case "HOUR":
		graphformatter.TimestampToUnixTime(
			graphformatter.TimeDifferenceHour(structs))
	}

}