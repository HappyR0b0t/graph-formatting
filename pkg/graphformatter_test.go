package graphformatter

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTransaction(t *testing.T) {
	// Test cases
	tests := []struct {
		name         string
		value        int
		timestamp    int64
		expected     *Transaction
		expectedTime time.Time
	}{
		{
			name:         "Valid transaction",
			value:        100,
			timestamp:    1672531200, // Unix timestamp for 2023-01-01 00:00:00 UTC
			expected:     &Transaction{Value: 100, Timestamp: time.Unix(1672531200, 0)},
			expectedTime: time.Unix(1672531200, 0),
		},
		{
			name:         "Negative value",
			value:        -100,
			timestamp:    1672534800,
			expected:     &Transaction{Value: -100, Timestamp: time.Unix(1672534800, 0)},
			expectedTime: time.Unix(1672534800, 0),
		},
		{
			name:         "Zero timestamp",
			value:        200,
			timestamp:    0,
			expected:     &Transaction{Value: 200, Timestamp: time.Unix(0, 0)},
			expectedTime: time.Unix(0, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewTransaction(tt.value, tt.timestamp)
			
			// Check value
			if result.Value != tt.expected.Value {
				t.Errorf("NewTransaction().Value = %d, want %d", result.Value, tt.expected.Value)
			}
			
			// Check timestamp
			if !result.Timestamp.Equal(tt.expectedTime) {
				t.Errorf("NewTransaction().Timestamp = %v, want %v", result.Timestamp, tt.expectedTime)
			}
		})
	}
}
func TestTimeDifferenceMonth(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    []Transaction
		expected []Transaction
	}{
		{
			name: "Single pair one month apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 3, 1, 12, 0, 0, 0, time.UTC)}, // January 1, 2023
				{Value: 200, Timestamp: time.Date(2023, 2, 1, 13, 0, 0, 0, time.UTC)}, // February 1, 2023
				{Value: 300, Timestamp: time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC)}, // March 1, 2023
			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)},  // Rounded to January 1, 2023, midnight
				{Value: 200, Timestamp: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)},  // Rounded to February 1, 2023, midnight
				{Value: 300, Timestamp: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},  // Rounded to March 1, 2023, midnight
			},
		},
		{
			name: "No pairs one month apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 3, 1, 12, 0, 0, 0, time.UTC)},
			},
			expected: []Transaction{},
		},
		{
			name: "Multiple pairs one month apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 3, 1, 12, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 2, 1, 12, 0, 0, 0, time.UTC)},
				{Value: 300, Timestamp: time.Date(2023, 2, 15, 12, 0, 0, 0, time.UTC)},
				{Value: 400, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)},  // Rounded to January 1, 2023, midnight
				{Value: 200, Timestamp: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)},  // Rounded to February 1, 2023, midnight
				{Value: 400, Timestamp: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},  // Rounded to March 1, 2023, midnight
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeDifferenceMonth(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TimeDifferenceMonth() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTimeDifferenceWeek(t *testing.T) {
	tests := []struct {
		name     string
		input    []Transaction
		expected []Transaction
	}{
		{
			name: "Single pair one week apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 15, 12, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 1, 8, 12, 0, 0, 0, time.UTC)},
				{Value: 300, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "Multiple pairs one week apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 2, 8, 12, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 2, 1, 12, 0, 0, 0, time.UTC)},
				{Value: 300, Timestamp: time.Date(2023, 1, 26, 12, 0, 0, 0, time.UTC)},
				{Value: 400, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 2, 8, 0, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)},
				{Value: 300, Timestamp: time.Date(2023, 1, 26, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "No pairs one week apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},   // January 1, 2023
				{Value: 200, Timestamp: time.Date(2023, 1, 10, 12, 0, 0, 0, time.UTC)},  // January 10, 2023
				{Value: 300, Timestamp: time.Date(2023, 1, 15, 12, 0, 0, 0, time.UTC)},  // January 15, 2023
			},
			expected: []Transaction{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeDifferenceWeek(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TimeDifferenceWeek() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTimeDifferenceDay(t *testing.T) {
	tests := []struct {
		name     string
		input    []Transaction
		expected []Transaction
	}{
		{
			name: "Single pair one day apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 4, 12, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 1, 3, 12, 0, 0, 0, time.UTC)},
				{Value: 300, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "Multiple pairs one day apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 7, 12, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 1, 6, 12, 0, 0, 0, time.UTC)},
				{Value: 300, Timestamp: time.Date(2023, 1, 5, 12, 0, 0, 0, time.UTC)},
				{Value: 400, Timestamp: time.Date(2023, 1, 3, 22, 0, 0, 0, time.UTC)},

			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 7, 0, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 1, 6, 0, 0, 0, 0, time.UTC)},
				{Value: 300, Timestamp: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "No pairs one day apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 5, 12, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 1, 3, 12, 0, 0, 0, time.UTC)},
				{Value: 300, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
			},
			expected: []Transaction{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeDifferenceDay(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TimeDifferenceDay() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTimeDifferenceHour(t *testing.T) {
	tests := []struct {
		name     string
		input    []Transaction
		expected []Transaction
	}{
		{
			name: "Single pair one hour apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},   // January 1, 2023 12:00 PM
				{Value: 200, Timestamp: time.Date(2023, 1, 1, 11, 0, 0, 0, time.UTC)},   // January 1, 2023 11:00 AM
			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 1, 13, 0, 0, 0, time.UTC)},   // Rounded to January 1, 2023 1:00 PM
				{Value: 200, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},   // Rounded to January 1, 2023 12:00 PM
			},
		},
		{
			name: "Multiple pairs one hour apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 1, 13, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
				{Value: 300, Timestamp: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC)},
			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 1, 14, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 1, 1, 13, 0, 0, 0, time.UTC)},
			},
		},
		{
			name: "No pairs one hour apart",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 1, 1, 10, 30, 0, 0, time.UTC)},
			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 1, 13, 0, 0, 0, time.UTC)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeDifferenceHour(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TimeDifferenceHour() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestTimestampToUnixTime(t *testing.T) {
	tests := []struct {
		name     string
		input    []Transaction
		expected []map[int]int64
	}{
		{
			name: "Single transaction",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
			},
			expected: []map[int]int64{
				{100: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC).Unix()},
			},
		},
		{
			name: "Multiple transactions",
			input: []Transaction{
				{Value: 100, Timestamp: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)},
				{Value: 200, Timestamp: time.Date(2023, 2, 1, 12, 0, 0, 0, time.UTC)},
				{Value: 300, Timestamp: time.Date(2023, 3, 1, 12, 0, 0, 0, time.UTC)},
			},
			expected: []map[int]int64{
				{100: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC).Unix()},
				{200: time.Date(2023, 2, 1, 12, 0, 0, 0, time.UTC).Unix()},
				{300: time.Date(2023, 3, 1, 12, 0, 0, 0, time.UTC).Unix()},
			},
		},
		{
			name:     "Empty input",
			input:    []Transaction{},
			expected: []map[int]int64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimestampToUnixTime(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TimestampToUnixTime() = %v, want %v", result, tt.expected)
			}
		})
	}
}
func TestSliceFiller(t *testing.T) {
	tests := []struct {
		name     string
		structs  []Transaction
		graph    map[int]int64
		expected []Transaction
	}{
		{
			name:     "Empty structs and graph",
			structs:  []Transaction{},
			graph:    map[int]int64{},
			expected: []Transaction{},
		},
		{
			name:    "Non-empty structs and empty graph",
			structs: []Transaction{{Value: 100, Timestamp: time.Unix(1672531200, 0)}},
			graph:   map[int]int64{},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Unix(1672531200, 0)},
			},
		},
		{
			name:    "Empty structs and non-empty graph",
			structs: []Transaction{},
			graph: map[int]int64{
				200: 1672534800,
				300: 1672538400,
			},
			expected: []Transaction{
				{Value: 200, Timestamp: time.Unix(1672534800, 0).UTC()},
				{Value: 300, Timestamp: time.Unix(1672538400, 0).UTC()},
			},
		},
		{
			name:    "Non-empty structs and non-empty graph",
			structs: []Transaction{{Value: 100, Timestamp: time.Unix(1672531200, 0).UTC()}},
			graph: map[int]int64{
				200: 1672534800,
				300: 1672538400,
			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Unix(1672531200, 0).UTC()},
				{Value: 200, Timestamp: time.Unix(1672534800, 0).UTC()},
				{Value: 300, Timestamp: time.Unix(1672538400, 0).UTC()},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SliceFiller(tt.structs, tt.graph)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SliceFiller() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSliceSorter(t *testing.T) {
	tests := []struct {
		name     string
		input    []Transaction
		expected []Transaction
	}{
		{
			name:     "Empty slice",
			input:    []Transaction{},
			expected: []Transaction{},
		},
		{
			name: "Single element slice",
			input: []Transaction{
				{Value: 100, Timestamp: time.Unix(1672531200, 0)},
			},
			expected: []Transaction{
				{Value: 100, Timestamp: time.Unix(1672531200, 0)},
			},
		},
		{
			name: "Multiple elements slice",
			input: []Transaction{
				{Value: 100, Timestamp: time.Unix(1672531200, 0)},
				{Value: 200, Timestamp: time.Unix(1672534800, 0)},
				{Value: 300, Timestamp: time.Unix(1672527600, 0)},
			},
			expected: []Transaction{
				{Value: 200, Timestamp: time.Unix(1672534800, 0)},
				{Value: 100, Timestamp: time.Unix(1672531200, 0)},
				{Value: 300, Timestamp: time.Unix(1672527600, 0)},
			},
		},
		{
			name: "Already sorted slice",
			input: []Transaction{
				{Value: 200, Timestamp: time.Unix(1672534800, 0)},
				{Value: 100, Timestamp: time.Unix(1672531200, 0)},
				{Value: 300, Timestamp: time.Unix(1672527600, 0)},
			},
			expected: []Transaction{
				{Value: 200, Timestamp: time.Unix(1672534800, 0)},
				{Value: 100, Timestamp: time.Unix(1672531200, 0)},
				{Value: 300, Timestamp: time.Unix(1672527600, 0)},
			},
		},
		{
			name: "Reverse sorted slice",
			input: []Transaction{
				{Value: 300, Timestamp: time.Unix(1672527600, 0)},
				{Value: 100, Timestamp: time.Unix(1672531200, 0)},
				{Value: 200, Timestamp: time.Unix(1672534800, 0)},
			},
			expected: []Transaction{
				{Value: 200, Timestamp: time.Unix(1672534800, 0)},
				{Value: 100, Timestamp: time.Unix(1672531200, 0)},
				{Value: 300, Timestamp: time.Unix(1672527600, 0)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SliceSorter(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SliceSorter() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRoundToMidnight(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Already midnight",
			input:    time.Date(2023, 6, 13, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2023, 6, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Middle of the day",
			input:    time.Date(2023, 6, 13, 15, 30, 45, 123456789, time.UTC),
			expected: time.Date(2023, 6, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "End of the day",
			input:    time.Date(2023, 6, 13, 23, 59, 59, 999999999, time.UTC),
			expected: time.Date(2023, 6, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Beginning of the day",
			input:    time.Date(2023, 6, 13, 0, 0, 0, 1, time.UTC),
			expected: time.Date(2023, 6, 13, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := roundToMidnight(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("roundToMidnight(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRoundToNearestHour(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected time.Time
	}{
		{
			name:     "Already on the hour",
			input:    time.Date(2023, 6, 13, 14, 0, 0, 0, time.UTC),
			expected: time.Date(2023, 6, 13, 15, 0, 0, 0, time.UTC),
		},
		{
			name:     "Half past the hour",
			input:    time.Date(2023, 6, 13, 14, 30, 0, 0, time.UTC),
			expected: time.Date(2023, 6, 13, 15, 0, 0, 0, time.UTC),
		},
		{
			name:     "Just before the next hour",
			input:    time.Date(2023, 6, 13, 14, 59, 59, 999999999, time.UTC),
			expected: time.Date(2023, 6, 13, 15, 0, 0, 0, time.UTC),
		},
		{
			name:     "Just after the hour",
			input:    time.Date(2023, 6, 13, 14, 0, 1, 0, time.UTC),
			expected: time.Date(2023, 6, 13, 15, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := roundToNearestHour(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("roundToNearestHour(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}