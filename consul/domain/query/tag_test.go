package query

import (
	"reflect"
	"testing"
)

func TestParseTagExacts(t *testing.T) {
	type args struct {
		arg   []string
		index int
	}
	tests := []struct {
		name string
		args args
		want TagExacts
	}{
		{
			"single parse",
			args{[]string{"this", "is", "test"}, 1},
			[]string{"is"},
		},
		{
			"multi parse",
			args{[]string{"this,is,test", "end", "test"}, 0},
			[]string{"this", "is", "test"},
		},
		{
			"out of bound",
			args{[]string{"this"}, 1},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTagExacts(tt.args.arg, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTagExacts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagExacts_Matches(t *testing.T) {
	type args struct {
		tag []string
	}
	tests := []struct {
		name   string
		exacts TagExacts
		args   args
		want   []string
		want1  bool
	}{
		{
			"match all when no condition",
			[]string{},
			args{[]string{"test", "test2"}},
			[]string{"test", "test2"},
			true,
		},
		{
			"match when exact matched",
			[]string{"te", "a", "test"},
			args{[]string{"test", "test2"}},
			[]string{"test"},
			true,
		},

		{
			"not match when prefix is matched",
			[]string{"te", "a"},
			args{[]string{"test", "test2"}},
			[]string{},
			false,
		},
		{
			"not match when no condition matched",
			[]string{"a"},
			args{[]string{"test", "test2"}},
			[]string{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.exacts.Matches(tt.args.tag)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagExacts.Matches() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TagExacts.Matches() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTagExacts_MatchesOnly(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name   string
		exacts TagExacts
		args   args
		want   bool
	}{
		{
			"match all when no condition",
			[]string{},
			args{"test"},
			true,
		},
		{
			"match when exact matched",
			[]string{"te", "a", "test"},
			args{"test"},
			true,
		},

		{
			"not match when prefix is matched",
			[]string{"te", "a"},
			args{"test"},
			false,
		},
		{
			"match when no condition matched",
			[]string{"a"},
			args{"test"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exacts.MatchesOnly(tt.args.tag); got != tt.want {
				t.Errorf("TagExacts.MatchesOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}
