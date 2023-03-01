package todo

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		todoUsecase todoUsecase
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "success",
			args: args{
				todoUsecase: nil,
			},
			want: &Handler{
				todoUsecase: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.todoUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
