package todo

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestNew(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want *Resource
	}{
		{
			name: "success",
			args: args{
				db: nil,
			},
			want: &Resource{
				db: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
