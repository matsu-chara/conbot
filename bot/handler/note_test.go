package handler

import (
	"reflect"
	"testing"
)

func Test_parseNoteArg(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want noteArg
	}{
		{
			"parse single",
			args{[]string{"service", "tag"}},
			noteArg{[]string{"service"}, []string{"tag"}},
		},
		{
			"parse multi",
			args{[]string{"s1,s2", "t1,t2"}},
			noteArg{[]string{"s1", "s2"}, []string{"t1", "t2"}},
		},
		{
			"nothing",
			args{[]string{}},
			noteArg{[]string{}, []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseNoteArg(tt.args.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNoteArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
