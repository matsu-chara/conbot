package query

import (
	"reflect"
	"testing"
)

func TestParseServicePrefixes(t *testing.T) {
	type args struct {
		arg   []string
		index int
	}
	tests := []struct {
		name string
		args args
		want ServicePrefixes
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
			if got := ParseServicePrefixes(tt.args.arg, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseServicePrefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServicePrefixes_Matches(t *testing.T) {
	type args struct {
		service string
	}
	tests := []struct {
		name     string
		prefixes ServicePrefixes
		args     args
		want     bool
	}{
		{
			"match all when no condition",
			[]string{},
			args{"test"},
			true,
		},
		{
			"match when prefix is matched",
			[]string{"te", "a"},
			args{"test"},
			true,
		},
		{
			"not match when no condition matched",
			[]string{"a"},
			args{"test"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.prefixes.Matches(tt.args.service); got != tt.want {
				t.Errorf("ServicePrefixes.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}
