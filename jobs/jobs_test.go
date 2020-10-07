package jobs

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
		id  blend4go.GalaxyID
	}
	tests := []struct {
		name    string
		args    args
		want    *Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.ctx, tt.args.g, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
	}
	tests := []struct {
		name    string
		args    args
		want    []*Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.ctx, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}
