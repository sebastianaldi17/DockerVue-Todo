package todo

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		todoRes todoResource
	}
	tests := []struct {
		name string
		args args
		want *Usecase
	}{
		{
			name: "success init",
			args: args{
				todoRes: nil,
			},
			want: &Usecase{
				todoRes: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.todoRes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
