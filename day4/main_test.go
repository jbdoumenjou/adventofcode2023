package main

import "testing"

func Test_score(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "no winning number",
			s:    "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			want: 0,
		},
		{
			name: "has one winning number",
			s:    "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			want: 1,
		},
		{
			name: "has two winning numbers",
			s:    "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			want: 2,
		},
		{
			name: "has 4 winning numbers",
			s:    "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := score(tt.s); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}
