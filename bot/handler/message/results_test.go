package message

import (
	"reflect"
	"testing"
)

func TestResultMessage_ResultOrNotFound(t *testing.T) {
	tests := []struct {
		name   string
		result ResultMessage
		want   []string
	}{
		{
			"result",
			[]string{"my", "result"},
			[]string{"my", "result"},
		},
		{
			"not found",
			[]string{},
			[]string{"not found."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.result.ResultOrNotFound(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResultMessage.ResultOrNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
