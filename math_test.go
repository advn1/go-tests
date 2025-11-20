package tests

import "testing"

func TestAdd_TableDriven(t *testing.T) {
	tests := [] struct {
		name string
		a,b int
		expected int
	} {
		{"positive numbers", 2,3,5},
		{"negative numbers", -2,-5,-7},
		{"mixed numbers", -2,3,1},
		{"zero numbers", 0,0,0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Add(tt.a, tt.b)
			expected := tt.expected

			if res != expected {
				t.Errorf("Add(%v, %v) = %v; Expected %v", tt.a, tt.b, res, expected)
			}
		})
	}
}