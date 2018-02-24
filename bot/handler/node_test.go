package handler

import (
	"reflect"
	"testing"
)

func Test_parseNodeArg(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want nodeArg
	}{
		{
			"parse single",
			args{[]string{"service", "tag"}},
			nodeArg{[]string{"service"}, []string{"tag"}},
		},
		{
			"parse multi",
			args{[]string{"s1,s2", "t1,t2"}},
			nodeArg{[]string{"s1", "s2"}, []string{"t1", "t2"}},
		},
		{
			"nothing",
			args{[]string{}},
			nodeArg{[]string{}, []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseNodeArg(tt.args.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArg() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
