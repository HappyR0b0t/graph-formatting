package main

import (
	"sort"

	graphformatter "github.com/HappyR0b0t/graph-formatting/pkg"
)

func main() {
	interval := "month"
	structs := []graphformatter.Transaction{}
	result := []graphformatter.Transaction{}

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
	}

	for key, value := range graph{
		t := graphformatter.NewTransaction(key, value)
		structs = append(structs, *t)
	}

	sort.Slice(structs, func(i, j int) bool {
		return structs[i].Timestamp.After(structs[j].Timestamp)
	})

	if interval == "month"{
		graphformatter.TimeDifferenceMonth(structs)
	}

	if interval == "week"{
		graphformatter.TimeDifferenceWeek(structs)
	}
	
	if interval == "day"{
		graphformatter.TimeDifferenceDay(structs)
	}
	
	if interval == "hour"{
		graphformatter.TimeDifferenceHour(structs)
	}
}