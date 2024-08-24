package main

import (
	"testing"
)

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		length, width int
		expected      string
	}{
		{2, 3, "odd rectangle"},  // area is 6
		{4, 5, "even rectangle"}, // area is 20
		{1, 1, "odd rectangle"},  // area is 1
	}

	for _, tt := range tests {
		t.Run("RectangleArea", func(t *testing.T) {
			result := RectangleArea(tt.length, tt.width)
			if result != tt.expected {
				t.Errorf("RectangleArea(%d, %d) = %s; want %s", tt.length, tt.width, result, tt.expected)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		length, width int
		expected      bool
	}{
		{2, 2, false}, // perimeter is 8
		{5, 5, true},  // perimeter is 20
		{1, 3, false}, // perimeter is 8
		{10, 5, true}, // perimeter is 30
	}

	for _, tt := range tests {
		t.Run("RectanglePerimeter", func(t *testing.T) {
			result := RectanglePerimeter(tt.length, tt.width)
			if result != tt.expected {
				t.Errorf("RectanglePerimeter(%d, %d) = %v; want %v", tt.length, tt.width, result, tt.expected)
			}
		})
	}
}

func TestSquareArea(t *testing.T) {
	tests := []struct {
		side     int
		expected string
	}{
		{2, "even square"}, // area is 4
		{3, "odd square"},  // area is 9
		{4, "even square"}, // area is 16
		{5, "odd square"},  // area is 25
	}

	for _, tt := range tests {
		t.Run("SquareArea", func(t *testing.T) {
			result := SquareArea(tt.side)
			if result != tt.expected {
				t.Errorf("SquareArea(%d) = %s; want %s", tt.side, result, tt.expected)
			}
		})
	}
}

func TestSquarePerimeter(t *testing.T) {
	tests := []struct {
		side     int
		expected bool
	}{
		{5, false}, // perimeter is 20
		{10, true}, // perimeter is 40
		{11, true}, // perimeter is 44
		{8, false}, // perimeter is 32
	}

	for _, tt := range tests {
		t.Run("SquarePerimeter", func(t *testing.T) {
			result := SquarePerimeter(tt.side)
			if result != tt.expected {
				t.Errorf("SquarePerimeter(%d) = %v; want %v", tt.side, result, tt.expected)
			}
		})
	}
}
