package message

import (
	"reflect"
	"testing"
)

func TestNewSlackWords(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want SlackWords
	}{
		{
			"split by space",
			args{"this is test"},
			[]string{"this", "is", "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSlackWords(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSlackWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlackWords_HasMentionAtFirst(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name  string
		words SlackWords
		args  args
		want  bool
	}{
		{
			"with mention",
			[]string{"<@myBot>"},
			args{"myBot"},
			true,
		},
		{
			"without mention",
			[]string{"test", "test"},
			args{"myBot"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.words.HasMentionAtFirst(tt.args.userName); got != tt.want {
				t.Errorf("SlackWords.HasMentionAtFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlackWords_Command(t *testing.T) {
	tests := []struct {
		name  string
		words SlackWords
		want  string
	}{
		{
			"no command",
			[]string{"<@myBot>"},
			"help",
		},
		{
			"command",
			[]string{"<@myBot>", "someCom"},
			"someCom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.words.Command(); got != tt.want {
				t.Errorf("SlackWords.Command() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlackWords_CommandArgs(t *testing.T) {
	tests := []struct {
		name  string
		words SlackWords
		want  []string
	}{
		{
			"no command",
			[]string{"<@myBot>"},
			[]string{},
		},
		{
			"command no args",
			[]string{"<@myBot>", "someCom"},
			[]string{},
		},
		{
			"command with args",
			[]string{"<@myBot>", "someCom", "a", "b"},
			[]string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.words.CommandArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlackWords.CommandArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
