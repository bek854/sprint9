package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantNil bool
	}{
		{"Positive size", 100, false},
		{"Zero size", 0, true},
		{"Negative size", -10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if tt.wantNil {
				if got != nil {
					t.Errorf("generateRandomElements(%d) = %v, want nil", tt.size, got)
				}
			} else {
				if len(got) != tt.size {
					t.Errorf("generateRandomElements(%d) returned slice with size %d", tt.size, len(got))
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{42}, 42},
		{"Multiple elements", []int{1, 3, 2, 5, 4}, 5},
		{"All equal", []int{7, 7, 7, 7}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.slice); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"Empty slice", []int{}, 0},
		{"Single element", []int{99}, 99},
		{"Less than chunks", []int{1, 3, 2, 5, 4}, 5},
		{"Large slice", make([]int, 1000), 999},
	}

	for i := range tests[3].slice {
		tests[3].slice[i] = i
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.slice); got != tt.want {
				t.Errorf("maxChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
