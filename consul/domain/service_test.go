package domain

import (
	"reflect"
	"testing"

	"github.com/matsu-chara/conbot/consul/domain/query"
)

func TestTagsByServiceMap_SortedServices(t *testing.T) {
	tests := []struct {
		name string
		tsm  TagsByServiceMap
		want Services
	}{
		{
			"sort",
			map[string][]string{"b": {"x"}, "a": {"y"}, "c": {"z"}},
			[]string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tsm.SortedServices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagsByServiceMap.SortedServices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagsByServiceMap_SortedTags(t *testing.T) {
	type args struct {
		service string
	}
	tests := []struct {
		name string
		tsm  TagsByServiceMap
		args args
		want Tags
	}{
		{
			"sort",
			map[string][]string{"b": {"2", "1"}, "a": {"y"}, "c": {"z"}},
			args{"b"},
			[]string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tsm.SortedTags(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagsByServiceMap.SortedTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagsByServiceMap_RemoveService(t *testing.T) {
	type args struct {
		service string
	}
	tests := []struct {
		name string
		tsm  TagsByServiceMap
		args args
	}{
		{
			"remove",
			map[string][]string{"b": {"x"}, "a": {"y"}, "c": {"z"}},
			args{"b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tsm.RemoveService(tt.args.service)
		})
		if _, ok := tt.tsm["b"]; ok {
			t.Error("remove but found")
		}
	}
}

func TestTagsByServiceMap_FilterByServicePrefixes(t *testing.T) {
	type args struct {
		servicePrefixes query.ServicePrefixes
	}
	tests := []struct {
		name string
		tsm  TagsByServiceMap
		args args
		want TagsByServiceMap
	}{
		{
			"filter by service prefix",
			map[string][]string{"apple": {"x"}, "appear": {"y"}, "banana": {"z"}},
			args{[]string{"a"}},
			map[string][]string{"apple": {"x"}, "appear": {"y"}},
		},
		{
			"no filter when service prefix is empty",
			map[string][]string{"apple": {"x"}, "appear": {"y"}, "banana": {"z"}},
			args{[]string{}},
			map[string][]string{"apple": {"x"}, "appear": {"y"}, "banana": {"z"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tsm.FilterByServicePrefixes(tt.args.servicePrefixes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagsByServiceMap.FilterByServicePrefixes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagsByServiceMap_FilterByTagExact(t *testing.T) {
	type args struct {
		tagExacts query.TagExacts
	}
	tests := []struct {
		name string
		tsm  TagsByServiceMap
		args args
		want TagsByServiceMap
	}{
		{
			"filter by service prefix",
			map[string][]string{"a": {"dog"}, "b": {"cat"}, "c": {"doge"}},
			args{[]string{"dog"}},
			map[string][]string{"a": {"dog"}},
		},
		{
			"return exact matched tags",
			map[string][]string{"a": {"dog", "doge"}, "b": {"cat", "dog"}, "c": {"doge"}},
			args{[]string{"dog"}},
			map[string][]string{"a": {"dog"}, "b": {"dog"}},
		},
		{
			"no filter when service prefix is empty",
			map[string][]string{"a": {"dog"}, "b": {"cat"}, "c": {"doge"}},
			args{[]string{}},
			map[string][]string{"a": {"dog"}, "b": {"cat"}, "c": {"doge"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tsm.FilterByTagExact(tt.args.tagExacts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagsByServiceMap.FilterByTagExact() = %v, want %v", got, tt.want)
			}
		})
	}
}
