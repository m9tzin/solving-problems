package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMiniMaxSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int32
		expected string
	}{
		{
			name:     "positive numbers",
			input:    []int32{1, 3, 5, 7, 9},
			expected: "16 24",
		},
		{
			name:     "negative numbers",
			input:    []int32{-1, -3, -5, -7, -9},
			expected: "-24 -16",
		},
		{
			name:     "mixed numbers",
			input:    []int32{-1, 2, -3, 4, -5},
			expected: "-7 2",
		},
		{
			name:     "all same numbers",
			input:    []int32{5, 5, 5, 5, 5},
			expected: "20 20",
		},
		{
			name:     "large numbers",
			input:    []int32{396285104, 573261094, 759641832, 819230764, 364801279},
			expected: "2093989309 2548418794",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Run the function
			miniMaxSum(tt.input)

			// Restore stdout and read captured output
			w.Close()
			os.Stdout = old
			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			// Remove trailing newline for comparison
			output = output[:len(output)-1]

			if output != tt.expected {
				t.Errorf("miniMaxSum(%v) = %q, want %q", tt.input, output, tt.expected)
			}
		})
	}
}

func BenchmarkMiniMaxSum(b *testing.B) {
	arr := []int32{1, 3, 5, 7, 9}
	
	// Redirect stdout to discard output during benchmark
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		miniMaxSum(arr)
	}
}
