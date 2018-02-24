package handler

import (
	"reflect"
	"testing"
)

func Test_parseCatalogArg(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want catalogArg
	}{
		{
			"parse single",
			args{[]string{"service", "tag"}},
			catalogArg{[]string{"service"}, []string{"tag"}},
		},
		{
			"parse multi",
			args{[]string{"s1,s2", "t1,t2"}},
			catalogArg{[]string{"s1", "s2"}, []string{"t1", "t2"}},
		},
		{
			"nothing",
			args{[]string{}},
			catalogArg{[]string{}, []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseCatalogArg(tt.args.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCatalogArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
