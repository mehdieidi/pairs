package pairs

import (
	"reflect"
	"sort"
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected [][2]int
	}{
		{
			name:     "All Elements Same",
			numbers:  []int{1, 1, 1},
			target:   2,
			expected: [][2]int{{1, 1}, {1, 1}, {1, 1}},
		},
		{
			name:     "No Pairs",
			numbers:  []int{1, 2, 3},
			target:   7,
			expected: [][2]int{},
		},
		{
			name:     "Single Pair",
			numbers:  []int{1, 4, 5},
			target:   9,
			expected: [][2]int{{4, 5}},
		},
		{
			name:     "Multiple Pairs",
			numbers:  []int{1, 2, 3, 4, 5, 6},
			target:   7,
			expected: [][2]int{{1, 6}, {2, 5}, {3, 4}},
		},
		{
			name:     "Duplicates",
			numbers:  []int{1, 1, 2, 2, 3, 3},
			target:   4,
			expected: [][2]int{{1, 3}, {1, 3}, {2, 2}, {3, 1}, {3, 1}},
		},
		{
			name:     "Empty Input",
			numbers:  []int{},
			target:   4,
			expected: [][2]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Find(tt.numbers, tt.target)

			// Normalize and sort the pairs so they have deterministic order that makes comparison easy.
			normalizedResult := normalizePairs(result)
			normalizedExpected := normalizePairs(tt.expected)

			if !reflect.DeepEqual(normalizedResult, normalizedExpected) {
				t.Errorf("expected %v, got %v", normalizedExpected, normalizedResult)
			}
		})
	}
}

// normalizePairs sorts pairs and normalizes the order within pairs.
func normalizePairs(pairs [][2]int) [][2]int {
	normalized := make([][2]int, len(pairs))

	// Normalize each pair so that the smaller element comes first.
	for i, pair := range pairs {
		if pair[0] > pair[1] {
			normalized[i] = [2]int{pair[1], pair[0]}
		} else {
			normalized[i] = pair
		}
	}

	// Sort the pairs by first and second elements.
	sort.Slice(normalized, func(i, j int) bool {
		if normalized[i][0] != normalized[j][0] {
			return normalized[i][0] < normalized[j][0]
		}
		return normalized[i][1] < normalized[j][1]
	})

	return normalized
}
