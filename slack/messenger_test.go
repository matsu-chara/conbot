package slack

import (
	"reflect"
	"testing"

	"github.com/nlopes/slack"
)

type testRtm struct {
	sentTexts []string
}

func (rtm *testRtm) SendMessage(msg *slack.OutgoingMessage) {
	rtm.sentTexts = append(rtm.sentTexts, msg.Text)
}
func (rtm *testRtm) NewOutgoingMessage(text string, channel string) *slack.OutgoingMessage {
	return &slack.OutgoingMessage{
		Text:    text,
		Channel: channel,
	}
}

func TestMessenger_SendChecking(t *testing.T) {
	type fields struct {
		rtm slackRtm
	}
	type args struct {
		ev *slack.MessageEvent
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"should send checking",
			fields{&testRtm{}},
			args{&slack.MessageEvent{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rtm := tt.fields.rtm.(*testRtm)
			messenger := &Messenger{
				rtm: rtm,
			}
			messenger.SendChecking(tt.args.ev)
			if rtm.sentTexts[0] != "checking..." {
				t.Error("invalid sent text:", rtm.sentTexts)
			}

		})
	}
}

func TestMessenger_SendMessages(t *testing.T) {
	case1 := []string{"msg1", "msg2"}
	case2 := []string{}
	type fields struct {
		rtm slackRtm
	}
	type args struct {
		msgs []string
		ev   *slack.MessageEvent
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"should send message",
			fields{&testRtm{[]string{}}},
			args{case1, &slack.MessageEvent{}},
		},
		{
			"should not send when empty",
			fields{&testRtm{[]string{}}},
			args{case2, &slack.MessageEvent{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messenger := &Messenger{
				rtm: tt.fields.rtm,
			}
			args := tt.args.msgs
			messenger.SendMessages(args, tt.args.ev)
			result := messenger.rtm.(*testRtm).sentTexts
			if !reflect.DeepEqual(args, result) {
				t.Errorf("not equal %#v, %#v", args, result)
			}
		})
	}
}
