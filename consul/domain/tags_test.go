package domain

import "testing"

func TestTags_String(t *testing.T) {
	tests := []struct {
		name string
		ts   Tags
		want string
	}{
		{
			"join",
			[]string{"t1", "t2"},
			"t1, t2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.String(); got != tt.want {
				t.Errorf("Tags.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
