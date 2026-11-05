package calc

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2,3) = %d, want 5", got)
	}
}

func TestSub(t *testing.T) {
	if got := Sub(10, 3); got != 7 {
		t.Fatalf("Sub(10,3) = %d, want 7", got)
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 3, 4, 12},
		{"negative numbers", -2, 5, -10},
		{"zero", 0, 7, 0},
		{"both zero", 0, 0, 0},
		{"negative result", 3, -2, -6},
		{"large numbers", 100, 200, 20000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.a, tt.b); got != tt.expected {
				t.Errorf("Mul(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
