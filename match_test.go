package main

import "testing"

func TestMatch(t *testing.T) {
	type test struct {
		name string
		file string
		want bool
	}

	tests := []test{
		{
			name: "a",
			file: "b",
			want: false,
		},
		{
			name: "a",
			file: "a",
			want: true,
		},
		{
			name: "20230627 - UNCUT: Michaela Kiersch - Wovenhand (8B/V13)",
			file: "20230627 - UNCUT： Michaela Kiersch - Wovenhand (8B⧸V13).webm",
			want: true,
		},
		{
			name: "20230628 - Michaela Kiersch - Wovenhand (8B/V13)",
			file: "20230627 - UNCUT： Michaela Kiersch - Wovenhand (8B⧸V13).webm",
			want: false,
		},
	}

	for _, tc := range tests {
		got := match(tc.name, tc.file)
		if got != tc.want {
			t.Fatalf(
				"name = %v\nfile = %v\nexpected: %v, got: %v",
				tc.name,
				tc.file,
				tc.want,
				got,
			)
		}
	}
}
