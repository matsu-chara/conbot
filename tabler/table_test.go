package tabler

import "testing"

func TestToSlackTable(t *testing.T) {
	type args struct {
		titles []string
		data   [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"generate table",
			args{
				titles: []string{"t1", "t2"},
				data:   [][]string{{"v1", "v2"}, {"v3", "v4"}},
			},
			"```\n" + `+----+----+
| T1 | T2 |
+----+----+
| v1 | v2 |
| v3 | v4 |
+----+----+` + "\n```",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSlackTable(tt.args.titles, tt.args.data); got != tt.want {
				t.Errorf("ToSlackTable() = \n%v\n%v\n", got, tt.want)
			}
		})
	}
}
