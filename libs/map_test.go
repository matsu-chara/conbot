package libs

import (
	"reflect"
	"testing"
)

func TestFilterByKey(t *testing.T) {
	type args struct {
		m map[string][]string
		f func(string) bool
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			"filter",
			args{
				map[string][]string{"a": {"1", "2"}, "b": {"2", "3"}},
				func(strs string) bool { return strs == "a" },
			},
			map[string][]string{"a": {"1", "2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterByKey(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterByValues(t *testing.T) {
	type args struct {
		m map[string][]string
		f func([]string) ([]string, bool)
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			"filter",
			args{
				map[string][]string{"a": {"1", "2"}, "b": {"2", "3"}},
				func(strs []string) ([]string, bool) { return []string{"2"}, Contains(strs, "2") },
			},
			map[string][]string{"a": {"2"}, "b": {"2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterByValues(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterByValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromSlice(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{
			"from",
			args{[]string{"1", "2"}},
			map[string]struct{}{"1": {}, "2": {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromSlice(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
