package jobs

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"reflect"
	"testing"
)

func TestJob_Delete(t *testing.T) {
	type fields struct {
		galaxyInstance *blend4go.GalaxyInstance
		Id             blend4go.GalaxyID
		ToolId         string
		UpdateTime     string
		HistoryId      string
		ExitCode       uint
		State          string
		CreateTime     string
		ModelClass     string
		Inputs         interface{}
		Outputs        interface{}
		Params         interface{}
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Job{
				galaxyInstance: tt.fields.galaxyInstance,
				Id:             tt.fields.Id,
				ToolId:         tt.fields.ToolId,
				UpdateTime:     tt.fields.UpdateTime,
				HistoryId:      tt.fields.HistoryId,
				ExitCode:       tt.fields.ExitCode,
				State:          tt.fields.State,
				CreateTime:     tt.fields.CreateTime,
				ModelClass:     tt.fields.ModelClass,
				Inputs:         tt.fields.Inputs,
				Outputs:        tt.fields.Outputs,
				Params:         tt.fields.Params,
			}
			if err := j.Delete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJob_GetBasePath(t *testing.T) {
	type fields struct {
		galaxyInstance *blend4go.GalaxyInstance
		Id             blend4go.GalaxyID
		ToolId         string
		UpdateTime     string
		HistoryId      string
		ExitCode       uint
		State          string
		CreateTime     string
		ModelClass     string
		Inputs         interface{}
		Outputs        interface{}
		Params         interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Job{
				galaxyInstance: tt.fields.galaxyInstance,
				Id:             tt.fields.Id,
				ToolId:         tt.fields.ToolId,
				UpdateTime:     tt.fields.UpdateTime,
				HistoryId:      tt.fields.HistoryId,
				ExitCode:       tt.fields.ExitCode,
				State:          tt.fields.State,
				CreateTime:     tt.fields.CreateTime,
				ModelClass:     tt.fields.ModelClass,
				Inputs:         tt.fields.Inputs,
				Outputs:        tt.fields.Outputs,
				Params:         tt.fields.Params,
			}
			if got := j.GetBasePath(); got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJob_GetID(t *testing.T) {
	type fields struct {
		galaxyInstance *blend4go.GalaxyInstance
		Id             blend4go.GalaxyID
		ToolId         string
		UpdateTime     string
		HistoryId      string
		ExitCode       uint
		State          string
		CreateTime     string
		ModelClass     string
		Inputs         interface{}
		Outputs        interface{}
		Params         interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   blend4go.GalaxyID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Job{
				galaxyInstance: tt.fields.galaxyInstance,
				Id:             tt.fields.Id,
				ToolId:         tt.fields.ToolId,
				UpdateTime:     tt.fields.UpdateTime,
				HistoryId:      tt.fields.HistoryId,
				ExitCode:       tt.fields.ExitCode,
				State:          tt.fields.State,
				CreateTime:     tt.fields.CreateTime,
				ModelClass:     tt.fields.ModelClass,
				Inputs:         tt.fields.Inputs,
				Outputs:        tt.fields.Outputs,
				Params:         tt.fields.Params,
			}
			if got := j.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJob_Resume(t *testing.T) {
	type fields struct {
		galaxyInstance *blend4go.GalaxyInstance
		Id             blend4go.GalaxyID
		ToolId         string
		UpdateTime     string
		HistoryId      string
		ExitCode       uint
		State          string
		CreateTime     string
		ModelClass     string
		Inputs         interface{}
		Outputs        interface{}
		Params         interface{}
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Job{
				galaxyInstance: tt.fields.galaxyInstance,
				Id:             tt.fields.Id,
				ToolId:         tt.fields.ToolId,
				UpdateTime:     tt.fields.UpdateTime,
				HistoryId:      tt.fields.HistoryId,
				ExitCode:       tt.fields.ExitCode,
				State:          tt.fields.State,
				CreateTime:     tt.fields.CreateTime,
				ModelClass:     tt.fields.ModelClass,
				Inputs:         tt.fields.Inputs,
				Outputs:        tt.fields.Outputs,
				Params:         tt.fields.Params,
			}
			if err := j.Resume(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Resume() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJob_SetGalaxyInstance(t *testing.T) {
	type fields struct {
		galaxyInstance *blend4go.GalaxyInstance
		Id             blend4go.GalaxyID
		ToolId         string
		UpdateTime     string
		HistoryId      string
		ExitCode       uint
		State          string
		CreateTime     string
		ModelClass     string
		Inputs         interface{}
		Outputs        interface{}
		Params         interface{}
	}
	type args struct {
		g *blend4go.GalaxyInstance
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Job{
				galaxyInstance: tt.fields.galaxyInstance,
				Id:             tt.fields.Id,
				ToolId:         tt.fields.ToolId,
				UpdateTime:     tt.fields.UpdateTime,
				HistoryId:      tt.fields.HistoryId,
				ExitCode:       tt.fields.ExitCode,
				State:          tt.fields.State,
				CreateTime:     tt.fields.CreateTime,
				ModelClass:     tt.fields.ModelClass,
				Inputs:         tt.fields.Inputs,
				Outputs:        tt.fields.Outputs,
				Params:         tt.fields.Params,
			}
		})
	}
}

func TestJob_SetID(t *testing.T) {
	type fields struct {
		galaxyInstance *blend4go.GalaxyInstance
		Id             blend4go.GalaxyID
		ToolId         string
		UpdateTime     string
		HistoryId      string
		ExitCode       uint
		State          string
		CreateTime     string
		ModelClass     string
		Inputs         interface{}
		Outputs        interface{}
		Params         interface{}
	}
	type args struct {
		id blend4go.GalaxyID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Job{
				galaxyInstance: tt.fields.galaxyInstance,
				Id:             tt.fields.Id,
				ToolId:         tt.fields.ToolId,
				UpdateTime:     tt.fields.UpdateTime,
				HistoryId:      tt.fields.HistoryId,
				ExitCode:       tt.fields.ExitCode,
				State:          tt.fields.State,
				CreateTime:     tt.fields.CreateTime,
				ModelClass:     tt.fields.ModelClass,
				Inputs:         tt.fields.Inputs,
				Outputs:        tt.fields.Outputs,
				Params:         tt.fields.Params,
			}
		})
	}
}

func TestNewJob(t *testing.T) {
	type args struct {
		ctx     context.Context
		g       *blend4go.GalaxyInstance
		payload map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []*Job
		want1   []*histories.HistoryDatasetAssociation
		want2   []*histories.HistoryDatasetCollectionAssociation
		want3   []*histories.HistoryDatasetCollectionAssociation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, err := NewJob(tt.args.ctx, tt.args.g, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJob() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewJob() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("NewJob() got2 = %v, want %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(got3, tt.want3) {
				t.Errorf("NewJob() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
