package miniMaxSum

import "testing"

//TestMiniMaxSum is a test function for miniMaxSum

func TestMiniMaxSum(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int32
		expected [2]int
	}{
		{
			name:     "Test case 0",
			arr:      []int32{1, 2, 3, 4, 5},
			expected: [2]int{10, 14},
		},
		{
			name:     "Test case 1",
			arr:      []int32{7, 69, 2, 221, 8974},
			expected: [2]int{299, 9271},
		},
		{
			name:     "Test case 2",
			arr:      []int32{256741038, 623958417, 467905213, 714532089, 938071625},
			expected: [2]int{2063136757, 2744467344},
		},
		{
			name:     "Test case 3",
			arr:      []int32{5, 5, 5, 5, 5},
			expected: [2]int{20, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := miniMaxSum(tt.arr)
			if actual != tt.expected {
				t.Errorf("Expected %v but got %v", tt.expected, actual)
			}
		})
	}

}
