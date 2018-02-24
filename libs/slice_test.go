package libs

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		vs []string
		f  func(string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"filter",
			args{[]string{"1", "2", "1"}, func(s string) bool { return s == "1" }},
			[]string{"1", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.vs, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		vs []string
		f  func(string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"map",
			args{[]string{"1", "2"}, func(s string) string { return s + "a" }},
			[]string{"1a", "2a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.vs, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"found", args{[]string{"1", "2"}, "1"}, true},
		{"not found", args{[]string{"1", "2"}, "3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	arg := []string{"1", "2", "1", "3"}
	type args struct {
		xs *[]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"remove duplicates",
			args{&arg},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveDuplicates(tt.args.xs)
		})
		want := []string{"1", "2", "3"}
		if !reflect.DeepEqual(arg, want) {
			t.Errorf("RemoveDuplicates() = %v, want %v", arg, want)
		}
	}
}

func TestGrouped(t *testing.T) {
	type args struct {
		xs []string
		n  int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"grouped",
			args{[]string{"1", "2", "3", "4"}, 2},
			[][]string{
				{"1", "2"},
				{"3", "4"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Grouped(tt.args.xs, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grouped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameLength(t *testing.T) {
	type args struct {
		xss [][]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"same",
			args{[][]string{
				{"a", "c"},
				{"b", "d"},
			}},
			true,
		},
		{
			"not same",
			args{[][]string{
				{"a"},
				{"b"},
				{"b", "d"},
			}},
			false,
		},
		{
			"not same empty",
			args{[][]string{
				{"a"},
				{},
			}},
			false,
		},
		{
			"empty",
			args{[][]string{}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SameLength(tt.args.xss); got != tt.want {
				t.Errorf("SameLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	type args struct {
		xss [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"flatten",
			args{[][]string{
				{"a"},
				{"b"},
				{"b", "d"},
			}},
			[]string{"a", "b", "b", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten(tt.args.xss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
