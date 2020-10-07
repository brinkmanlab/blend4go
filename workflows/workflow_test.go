package workflows

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"reflect"
	"testing"
)

func TestNewStoredWorkflow(t *testing.T) {
	type args struct {
		ctx  context.Context
		g    *blend4go.GalaxyInstance
		json string
	}
	tests := []struct {
		name    string
		args    args
		want    *StoredWorkflow
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStoredWorkflow(tt.args.ctx, tt.args.g, tt.args.json)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStoredWorkflow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStoredWorkflow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoredWorkflow_Delete(t *testing.T) {
	type fields struct {
		galaxyInstance     *blend4go.GalaxyInstance
		Id                 blend4go.GalaxyID
		Name               string
		Tags               []string
		Deleted            bool
		LatestWorkflowUuid string
		ShowInToolPanel    bool
		Url                string
		NumberOfSteps      uint
		Published          bool
		Owner              string
		ModelClass         string
		Inputs             []*StoredWorkflowInput
		Annotation         string
		Version            uint
		Steps              []*WorkflowInvocationStep
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
			w := &StoredWorkflow{
				galaxyInstance:     tt.fields.galaxyInstance,
				Id:                 tt.fields.Id,
				Name:               tt.fields.Name,
				Tags:               tt.fields.Tags,
				Deleted:            tt.fields.Deleted,
				LatestWorkflowUuid: tt.fields.LatestWorkflowUuid,
				ShowInToolPanel:    tt.fields.ShowInToolPanel,
				Url:                tt.fields.Url,
				NumberOfSteps:      tt.fields.NumberOfSteps,
				Published:          tt.fields.Published,
				Owner:              tt.fields.Owner,
				ModelClass:         tt.fields.ModelClass,
				Inputs:             tt.fields.Inputs,
				Annotation:         tt.fields.Annotation,
				Version:            tt.fields.Version,
				Steps:              tt.fields.Steps,
			}
			if err := w.Delete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStoredWorkflow_Download(t *testing.T) {
	type fields struct {
		galaxyInstance     *blend4go.GalaxyInstance
		Id                 blend4go.GalaxyID
		Name               string
		Tags               []string
		Deleted            bool
		LatestWorkflowUuid string
		ShowInToolPanel    bool
		Url                string
		NumberOfSteps      uint
		Published          bool
		Owner              string
		ModelClass         string
		Inputs             []*StoredWorkflowInput
		Annotation         string
		Version            uint
		Steps              []*WorkflowInvocationStep
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &StoredWorkflow{
				galaxyInstance:     tt.fields.galaxyInstance,
				Id:                 tt.fields.Id,
				Name:               tt.fields.Name,
				Tags:               tt.fields.Tags,
				Deleted:            tt.fields.Deleted,
				LatestWorkflowUuid: tt.fields.LatestWorkflowUuid,
				ShowInToolPanel:    tt.fields.ShowInToolPanel,
				Url:                tt.fields.Url,
				NumberOfSteps:      tt.fields.NumberOfSteps,
				Published:          tt.fields.Published,
				Owner:              tt.fields.Owner,
				ModelClass:         tt.fields.ModelClass,
				Inputs:             tt.fields.Inputs,
				Annotation:         tt.fields.Annotation,
				Version:            tt.fields.Version,
				Steps:              tt.fields.Steps,
			}
			got, err := w.Download(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Download() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoredWorkflow_GetBasePath(t *testing.T) {
	type fields struct {
		galaxyInstance     *blend4go.GalaxyInstance
		Id                 blend4go.GalaxyID
		Name               string
		Tags               []string
		Deleted            bool
		LatestWorkflowUuid string
		ShowInToolPanel    bool
		Url                string
		NumberOfSteps      uint
		Published          bool
		Owner              string
		ModelClass         string
		Inputs             []*StoredWorkflowInput
		Annotation         string
		Version            uint
		Steps              []*WorkflowInvocationStep
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
			w := &StoredWorkflow{
				galaxyInstance:     tt.fields.galaxyInstance,
				Id:                 tt.fields.Id,
				Name:               tt.fields.Name,
				Tags:               tt.fields.Tags,
				Deleted:            tt.fields.Deleted,
				LatestWorkflowUuid: tt.fields.LatestWorkflowUuid,
				ShowInToolPanel:    tt.fields.ShowInToolPanel,
				Url:                tt.fields.Url,
				NumberOfSteps:      tt.fields.NumberOfSteps,
				Published:          tt.fields.Published,
				Owner:              tt.fields.Owner,
				ModelClass:         tt.fields.ModelClass,
				Inputs:             tt.fields.Inputs,
				Annotation:         tt.fields.Annotation,
				Version:            tt.fields.Version,
				Steps:              tt.fields.Steps,
			}
			if got := w.GetBasePath(); got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoredWorkflow_GetID(t *testing.T) {
	type fields struct {
		galaxyInstance     *blend4go.GalaxyInstance
		Id                 blend4go.GalaxyID
		Name               string
		Tags               []string
		Deleted            bool
		LatestWorkflowUuid string
		ShowInToolPanel    bool
		Url                string
		NumberOfSteps      uint
		Published          bool
		Owner              string
		ModelClass         string
		Inputs             []*StoredWorkflowInput
		Annotation         string
		Version            uint
		Steps              []*WorkflowInvocationStep
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
			w := &StoredWorkflow{
				galaxyInstance:     tt.fields.galaxyInstance,
				Id:                 tt.fields.Id,
				Name:               tt.fields.Name,
				Tags:               tt.fields.Tags,
				Deleted:            tt.fields.Deleted,
				LatestWorkflowUuid: tt.fields.LatestWorkflowUuid,
				ShowInToolPanel:    tt.fields.ShowInToolPanel,
				Url:                tt.fields.Url,
				NumberOfSteps:      tt.fields.NumberOfSteps,
				Published:          tt.fields.Published,
				Owner:              tt.fields.Owner,
				ModelClass:         tt.fields.ModelClass,
				Inputs:             tt.fields.Inputs,
				Annotation:         tt.fields.Annotation,
				Version:            tt.fields.Version,
				Steps:              tt.fields.Steps,
			}
			if got := w.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoredWorkflow_Invoke(t *testing.T) {
	type fields struct {
		galaxyInstance     *blend4go.GalaxyInstance
		Id                 blend4go.GalaxyID
		Name               string
		Tags               []string
		Deleted            bool
		LatestWorkflowUuid string
		ShowInToolPanel    bool
		Url                string
		NumberOfSteps      uint
		Published          bool
		Owner              string
		ModelClass         string
		Inputs             []*StoredWorkflowInput
		Annotation         string
		Version            uint
		Steps              []*WorkflowInvocationStep
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
			w := &StoredWorkflow{
				galaxyInstance:     tt.fields.galaxyInstance,
				Id:                 tt.fields.Id,
				Name:               tt.fields.Name,
				Tags:               tt.fields.Tags,
				Deleted:            tt.fields.Deleted,
				LatestWorkflowUuid: tt.fields.LatestWorkflowUuid,
				ShowInToolPanel:    tt.fields.ShowInToolPanel,
				Url:                tt.fields.Url,
				NumberOfSteps:      tt.fields.NumberOfSteps,
				Published:          tt.fields.Published,
				Owner:              tt.fields.Owner,
				ModelClass:         tt.fields.ModelClass,
				Inputs:             tt.fields.Inputs,
				Annotation:         tt.fields.Annotation,
				Version:            tt.fields.Version,
				Steps:              tt.fields.Steps,
			}
			if err := w.Invoke(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Invoke() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStoredWorkflow_Repositories(t *testing.T) {
	type fields struct {
		galaxyInstance     *blend4go.GalaxyInstance
		Id                 blend4go.GalaxyID
		Name               string
		Tags               []string
		Deleted            bool
		LatestWorkflowUuid string
		ShowInToolPanel    bool
		Url                string
		NumberOfSteps      uint
		Published          bool
		Owner              string
		ModelClass         string
		Inputs             []*StoredWorkflowInput
		Annotation         string
		Version            uint
		Steps              []*WorkflowInvocationStep
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*repositories.Repository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &StoredWorkflow{
				galaxyInstance:     tt.fields.galaxyInstance,
				Id:                 tt.fields.Id,
				Name:               tt.fields.Name,
				Tags:               tt.fields.Tags,
				Deleted:            tt.fields.Deleted,
				LatestWorkflowUuid: tt.fields.LatestWorkflowUuid,
				ShowInToolPanel:    tt.fields.ShowInToolPanel,
				Url:                tt.fields.Url,
				NumberOfSteps:      tt.fields.NumberOfSteps,
				Published:          tt.fields.Published,
				Owner:              tt.fields.Owner,
				ModelClass:         tt.fields.ModelClass,
				Inputs:             tt.fields.Inputs,
				Annotation:         tt.fields.Annotation,
				Version:            tt.fields.Version,
				Steps:              tt.fields.Steps,
			}
			got, err := w.Repositories(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repositories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repositories() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoredWorkflow_SetGalaxyInstance(t *testing.T) {
	type fields struct {
		galaxyInstance     *blend4go.GalaxyInstance
		Id                 blend4go.GalaxyID
		Name               string
		Tags               []string
		Deleted            bool
		LatestWorkflowUuid string
		ShowInToolPanel    bool
		Url                string
		NumberOfSteps      uint
		Published          bool
		Owner              string
		ModelClass         string
		Inputs             []*StoredWorkflowInput
		Annotation         string
		Version            uint
		Steps              []*WorkflowInvocationStep
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
			w := &StoredWorkflow{
				galaxyInstance:     tt.fields.galaxyInstance,
				Id:                 tt.fields.Id,
				Name:               tt.fields.Name,
				Tags:               tt.fields.Tags,
				Deleted:            tt.fields.Deleted,
				LatestWorkflowUuid: tt.fields.LatestWorkflowUuid,
				ShowInToolPanel:    tt.fields.ShowInToolPanel,
				Url:                tt.fields.Url,
				NumberOfSteps:      tt.fields.NumberOfSteps,
				Published:          tt.fields.Published,
				Owner:              tt.fields.Owner,
				ModelClass:         tt.fields.ModelClass,
				Inputs:             tt.fields.Inputs,
				Annotation:         tt.fields.Annotation,
				Version:            tt.fields.Version,
				Steps:              tt.fields.Steps,
			}
		})
	}
}

func TestStoredWorkflow_SetID(t *testing.T) {
	type fields struct {
		galaxyInstance     *blend4go.GalaxyInstance
		Id                 blend4go.GalaxyID
		Name               string
		Tags               []string
		Deleted            bool
		LatestWorkflowUuid string
		ShowInToolPanel    bool
		Url                string
		NumberOfSteps      uint
		Published          bool
		Owner              string
		ModelClass         string
		Inputs             []*StoredWorkflowInput
		Annotation         string
		Version            uint
		Steps              []*WorkflowInvocationStep
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
			w := &StoredWorkflow{
				galaxyInstance:     tt.fields.galaxyInstance,
				Id:                 tt.fields.Id,
				Name:               tt.fields.Name,
				Tags:               tt.fields.Tags,
				Deleted:            tt.fields.Deleted,
				LatestWorkflowUuid: tt.fields.LatestWorkflowUuid,
				ShowInToolPanel:    tt.fields.ShowInToolPanel,
				Url:                tt.fields.Url,
				NumberOfSteps:      tt.fields.NumberOfSteps,
				Published:          tt.fields.Published,
				Owner:              tt.fields.Owner,
				ModelClass:         tt.fields.ModelClass,
				Inputs:             tt.fields.Inputs,
				Annotation:         tt.fields.Annotation,
				Version:            tt.fields.Version,
				Steps:              tt.fields.Steps,
			}
		})
	}
}

func TestStoredWorkflow_Update(t *testing.T) {
	type fields struct {
		galaxyInstance     *blend4go.GalaxyInstance
		Id                 blend4go.GalaxyID
		Name               string
		Tags               []string
		Deleted            bool
		LatestWorkflowUuid string
		ShowInToolPanel    bool
		Url                string
		NumberOfSteps      uint
		Published          bool
		Owner              string
		ModelClass         string
		Inputs             []*StoredWorkflowInput
		Annotation         string
		Version            uint
		Steps              []*WorkflowInvocationStep
	}
	type args struct {
		ctx  context.Context
		json string
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
			w := &StoredWorkflow{
				galaxyInstance:     tt.fields.galaxyInstance,
				Id:                 tt.fields.Id,
				Name:               tt.fields.Name,
				Tags:               tt.fields.Tags,
				Deleted:            tt.fields.Deleted,
				LatestWorkflowUuid: tt.fields.LatestWorkflowUuid,
				ShowInToolPanel:    tt.fields.ShowInToolPanel,
				Url:                tt.fields.Url,
				NumberOfSteps:      tt.fields.NumberOfSteps,
				Published:          tt.fields.Published,
				Owner:              tt.fields.Owner,
				ModelClass:         tt.fields.ModelClass,
				Inputs:             tt.fields.Inputs,
				Annotation:         tt.fields.Annotation,
				Version:            tt.fields.Version,
				Steps:              tt.fields.Steps,
			}
			if err := w.Update(tt.args.ctx, tt.args.json); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
