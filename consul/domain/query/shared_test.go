package query

import (
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		arg string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"split", args{"a,b,c"}, []string{"a", "b", "c"}},
		{"non empty", args{"a,,c"}, []string{"a", "c"}},
		{"trim", args{"a, b  ,c"}, []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
