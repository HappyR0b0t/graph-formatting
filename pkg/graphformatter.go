package graphformatter

import (
	"sort"
	"time"
)

type Transaction struct {
	Value		int
	Timestamp 	time.Time
}

func NewTransaction(Value int, Timestamp int64) *Transaction {
	t := &Transaction{Value: Value, Timestamp: time.Unix(Timestamp, 0).UTC()}
	return t
}

func TimeDifferenceMonth(structs []Transaction) []Transaction {
	result := []Transaction{}
	for i := 0; i < len(structs); i++ {
		for j := i + 1; j < len(structs); j++ {
			t1 := structs[i].Timestamp
			t2 := structs[j].Timestamp
			oneMonthLater := t1.AddDate(0, -1, 0)
			if i == 0 && 
			oneMonthLater.Year() == t2.Year() && 
			oneMonthLater.Month() == t2.Month() && 
			oneMonthLater.Day() == t2.Day() {
				structs[i].Timestamp = roundToMidnight(structs[i].Timestamp)
				structs[j].Timestamp = roundToMidnight(structs[j].Timestamp)
				result = append(result, structs[i], structs[j])
				i = j-1
				break
			} else if i != 0 && 
			oneMonthLater.Year() == t2.Year() && 
			oneMonthLater.Month() == t2.Month() && 
			oneMonthLater.Day() == t2.Day() {
				structs[j].Timestamp = roundToMidnight(structs[j].Timestamp)
				result = append(result, structs[j])
				i = j-1
				break				
			}
		}
	}
	return result
}

func TimeDifferenceWeek(structs []Transaction) []Transaction {
	result := []Transaction{}
	for i := 0; i < len(structs); i++ {
		for j := i + 1; j < len(structs); j++ {
			_, w1 := structs[i].Timestamp.ISOWeek()
			_, w2 := structs[j].Timestamp.ISOWeek()
			weekDifference := w1 - w2
			if weekDifference == 1 && i == 0 {
				structs[i].Timestamp = roundToMidnight(structs[i].Timestamp)
				structs[j].Timestamp = roundToMidnight(structs[j].Timestamp)
				result = append(result, structs[i], structs[j])
				i = j-1
				break
			} else if  weekDifference == 1 && i > 0 {
				structs[j].Timestamp = roundToMidnight(structs[j].Timestamp)
				result = append(result, structs[j])
				i = j-1
				break
			}		
		}
	}
	return result
}

func TimeDifferenceDay(structs []Transaction) []Transaction {
	result := []Transaction{}
	for i := 0; i < len(structs); i++ {
		for j := i + 1; j < len(structs); j++ {
			_, _, d1 := structs[i].Timestamp.Date()
			_, _, d2 := structs[j].Timestamp.Date()
			dayDifference := d1 - d2
			if dayDifference == 1 && i == 0 {
				structs[i].Timestamp = roundToMidnight(structs[i].Timestamp)
				structs[j].Timestamp = roundToMidnight(structs[j].Timestamp)
				result = append(result, structs[i], structs[j])
				i = j-1
				break
			} else if dayDifference == 1 && i > 0 {
				structs[j].Timestamp = roundToMidnight(structs[j].Timestamp)
				result = append(result, structs[j])
				i = j-1
				break
			}	
		}
	}
	return result
}

func TimeDifferenceHour(structs []Transaction) []Transaction {
	result := []Transaction{}
	for i := 0; i < len(structs); i++ {
		for j := i + 1; j < len(structs); j++ {
			t1 := structs[i].Timestamp
			t2 := structs[j].Timestamp
			duration := t1.Sub(t2)
			hour := time.Hour
			if i == 0 && j == 1 && duration == hour {
				structs[i].Timestamp = roundToNearestHour(structs[i].Timestamp)
				structs[j].Timestamp = roundToNearestHour(structs[j].Timestamp)
				result = append(result, structs[i], structs[j])
				i = j-1
				break
			} else if i == 0 && j == 1 && duration != hour{
				structs[i].Timestamp = roundToNearestHour(structs[i].Timestamp)
				result = append(result, structs[i])
			} 

			if duration == hour{
				structs[j].Timestamp = roundToNearestHour(structs[j].Timestamp)
				result = append(result, structs[j])
				i = j-1
				break
			}	
		}
	}
	return result
}

func roundToNearestHour(t time.Time) time.Time{
	return t.Truncate(time.Hour).Add(time.Hour).UTC()
}

func roundToMidnight(t time.Time) time.Time{
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	t = t.In(location)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, location)
}

func TimestampToUnixTime(structs []Transaction) []map[int]int64 {
	result := make([]map[int]int64, 0)
	if len(structs) == 0 {
		return result
	}
	for i := range structs{
		map1 := make(map[int]int64)
		map1[structs[i].Value] = structs[i].Timestamp.Unix()
		result = append(result, map1)
	}
	return result
}

func SliceSorter(structs []Transaction) []Transaction {
	sort.Slice(structs, func(i, j int) bool {
		return structs[i].Timestamp.After(structs[j].Timestamp)
	})
	return structs
}

func SliceFiller(structs []Transaction, graph map[int]int64) []Transaction {
	for key, value := range graph{
		t := NewTransaction(key, value)
		structs = append(structs, *t)
	}
	return structs
}