package tools

import (
	"github.com/brinkmanlab/blend4go"
	"testing"
)

func TestToolSection_GetID(t1 *testing.T) {
	type fields struct {
		Id         string
		Name       string
		Version    string
		ModelClass string
		Elems      []*Tool
	}
	tests := []struct {
		name   string
		fields fields
		want   blend4go.GalaxyID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &ToolSection{
				Id:         tt.fields.Id,
				Name:       tt.fields.Name,
				Version:    tt.fields.Version,
				ModelClass: tt.fields.ModelClass,
				Elems:      tt.fields.Elems,
			}
			if got := t.GetID(); got != tt.want {
				t1.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToolSection_SetID(t1 *testing.T) {
	type fields struct {
		Id         string
		Name       string
		Version    string
		ModelClass string
		Elems      []*Tool
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &ToolSection{
				Id:         tt.fields.Id,
				Name:       tt.fields.Name,
				Version:    tt.fields.Version,
				ModelClass: tt.fields.ModelClass,
				Elems:      tt.fields.Elems,
			}
		})
	}
}

func TestTool_GetBasePath(t1 *testing.T) {
	type fields struct {
		galaxyInstance   *blend4go.GalaxyInstance
		Id               string
		Name             string
		Version          string
		MinWidth         int
		Target           string
		Link             string
		PanelSectionId   string
		EdamTopics       []string
		FormStyle        string
		EdamOperations   []string
		Labels           []string
		Description      string
		ConfigFile       string
		Xrefs            []string
		PanelSectionName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tool{
				galaxyInstance:   tt.fields.galaxyInstance,
				Id:               tt.fields.Id,
				Name:             tt.fields.Name,
				Version:          tt.fields.Version,
				MinWidth:         tt.fields.MinWidth,
				Target:           tt.fields.Target,
				Link:             tt.fields.Link,
				PanelSectionId:   tt.fields.PanelSectionId,
				EdamTopics:       tt.fields.EdamTopics,
				FormStyle:        tt.fields.FormStyle,
				EdamOperations:   tt.fields.EdamOperations,
				Labels:           tt.fields.Labels,
				Description:      tt.fields.Description,
				ConfigFile:       tt.fields.ConfigFile,
				Xrefs:            tt.fields.Xrefs,
				PanelSectionName: tt.fields.PanelSectionName,
			}
			if got := t.GetBasePath(); got != tt.want {
				t1.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTool_GetID(t1 *testing.T) {
	type fields struct {
		galaxyInstance   *blend4go.GalaxyInstance
		Id               string
		Name             string
		Version          string
		MinWidth         int
		Target           string
		Link             string
		PanelSectionId   string
		EdamTopics       []string
		FormStyle        string
		EdamOperations   []string
		Labels           []string
		Description      string
		ConfigFile       string
		Xrefs            []string
		PanelSectionName string
	}
	tests := []struct {
		name   string
		fields fields
		want   blend4go.GalaxyID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tool{
				galaxyInstance:   tt.fields.galaxyInstance,
				Id:               tt.fields.Id,
				Name:             tt.fields.Name,
				Version:          tt.fields.Version,
				MinWidth:         tt.fields.MinWidth,
				Target:           tt.fields.Target,
				Link:             tt.fields.Link,
				PanelSectionId:   tt.fields.PanelSectionId,
				EdamTopics:       tt.fields.EdamTopics,
				FormStyle:        tt.fields.FormStyle,
				EdamOperations:   tt.fields.EdamOperations,
				Labels:           tt.fields.Labels,
				Description:      tt.fields.Description,
				ConfigFile:       tt.fields.ConfigFile,
				Xrefs:            tt.fields.Xrefs,
				PanelSectionName: tt.fields.PanelSectionName,
			}
			if got := t.GetID(); got != tt.want {
				t1.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTool_InstallDependencies(t1 *testing.T) {
	type fields struct {
		galaxyInstance   *blend4go.GalaxyInstance
		Id               string
		Name             string
		Version          string
		MinWidth         int
		Target           string
		Link             string
		PanelSectionId   string
		EdamTopics       []string
		FormStyle        string
		EdamOperations   []string
		Labels           []string
		Description      string
		ConfigFile       string
		Xrefs            []string
		PanelSectionName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tool{
				galaxyInstance:   tt.fields.galaxyInstance,
				Id:               tt.fields.Id,
				Name:             tt.fields.Name,
				Version:          tt.fields.Version,
				MinWidth:         tt.fields.MinWidth,
				Target:           tt.fields.Target,
				Link:             tt.fields.Link,
				PanelSectionId:   tt.fields.PanelSectionId,
				EdamTopics:       tt.fields.EdamTopics,
				FormStyle:        tt.fields.FormStyle,
				EdamOperations:   tt.fields.EdamOperations,
				Labels:           tt.fields.Labels,
				Description:      tt.fields.Description,
				ConfigFile:       tt.fields.ConfigFile,
				Xrefs:            tt.fields.Xrefs,
				PanelSectionName: tt.fields.PanelSectionName,
			}
			if err := t.InstallDependencies(); (err != nil) != tt.wantErr {
				t1.Errorf("InstallDependencies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTool_Reload(t1 *testing.T) {
	type fields struct {
		galaxyInstance   *blend4go.GalaxyInstance
		Id               string
		Name             string
		Version          string
		MinWidth         int
		Target           string
		Link             string
		PanelSectionId   string
		EdamTopics       []string
		FormStyle        string
		EdamOperations   []string
		Labels           []string
		Description      string
		ConfigFile       string
		Xrefs            []string
		PanelSectionName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tool{
				galaxyInstance:   tt.fields.galaxyInstance,
				Id:               tt.fields.Id,
				Name:             tt.fields.Name,
				Version:          tt.fields.Version,
				MinWidth:         tt.fields.MinWidth,
				Target:           tt.fields.Target,
				Link:             tt.fields.Link,
				PanelSectionId:   tt.fields.PanelSectionId,
				EdamTopics:       tt.fields.EdamTopics,
				FormStyle:        tt.fields.FormStyle,
				EdamOperations:   tt.fields.EdamOperations,
				Labels:           tt.fields.Labels,
				Description:      tt.fields.Description,
				ConfigFile:       tt.fields.ConfigFile,
				Xrefs:            tt.fields.Xrefs,
				PanelSectionName: tt.fields.PanelSectionName,
			}
			if err := t.Reload(); (err != nil) != tt.wantErr {
				t1.Errorf("Reload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTool_SetGalaxyInstance(t1 *testing.T) {
	type fields struct {
		galaxyInstance   *blend4go.GalaxyInstance
		Id               string
		Name             string
		Version          string
		MinWidth         int
		Target           string
		Link             string
		PanelSectionId   string
		EdamTopics       []string
		FormStyle        string
		EdamOperations   []string
		Labels           []string
		Description      string
		ConfigFile       string
		Xrefs            []string
		PanelSectionName string
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tool{
				galaxyInstance:   tt.fields.galaxyInstance,
				Id:               tt.fields.Id,
				Name:             tt.fields.Name,
				Version:          tt.fields.Version,
				MinWidth:         tt.fields.MinWidth,
				Target:           tt.fields.Target,
				Link:             tt.fields.Link,
				PanelSectionId:   tt.fields.PanelSectionId,
				EdamTopics:       tt.fields.EdamTopics,
				FormStyle:        tt.fields.FormStyle,
				EdamOperations:   tt.fields.EdamOperations,
				Labels:           tt.fields.Labels,
				Description:      tt.fields.Description,
				ConfigFile:       tt.fields.ConfigFile,
				Xrefs:            tt.fields.Xrefs,
				PanelSectionName: tt.fields.PanelSectionName,
			}
		})
	}
}

func TestTool_SetID(t1 *testing.T) {
	type fields struct {
		galaxyInstance   *blend4go.GalaxyInstance
		Id               string
		Name             string
		Version          string
		MinWidth         int
		Target           string
		Link             string
		PanelSectionId   string
		EdamTopics       []string
		FormStyle        string
		EdamOperations   []string
		Labels           []string
		Description      string
		ConfigFile       string
		Xrefs            []string
		PanelSectionName string
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
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tool{
				galaxyInstance:   tt.fields.galaxyInstance,
				Id:               tt.fields.Id,
				Name:             tt.fields.Name,
				Version:          tt.fields.Version,
				MinWidth:         tt.fields.MinWidth,
				Target:           tt.fields.Target,
				Link:             tt.fields.Link,
				PanelSectionId:   tt.fields.PanelSectionId,
				EdamTopics:       tt.fields.EdamTopics,
				FormStyle:        tt.fields.FormStyle,
				EdamOperations:   tt.fields.EdamOperations,
				Labels:           tt.fields.Labels,
				Description:      tt.fields.Description,
				ConfigFile:       tt.fields.ConfigFile,
				Xrefs:            tt.fields.Xrefs,
				PanelSectionName: tt.fields.PanelSectionName,
			}
		})
	}
}

func TestTool_UninstallDependencies(t1 *testing.T) {
	type fields struct {
		galaxyInstance   *blend4go.GalaxyInstance
		Id               string
		Name             string
		Version          string
		MinWidth         int
		Target           string
		Link             string
		PanelSectionId   string
		EdamTopics       []string
		FormStyle        string
		EdamOperations   []string
		Labels           []string
		Description      string
		ConfigFile       string
		Xrefs            []string
		PanelSectionName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tool{
				galaxyInstance:   tt.fields.galaxyInstance,
				Id:               tt.fields.Id,
				Name:             tt.fields.Name,
				Version:          tt.fields.Version,
				MinWidth:         tt.fields.MinWidth,
				Target:           tt.fields.Target,
				Link:             tt.fields.Link,
				PanelSectionId:   tt.fields.PanelSectionId,
				EdamTopics:       tt.fields.EdamTopics,
				FormStyle:        tt.fields.FormStyle,
				EdamOperations:   tt.fields.EdamOperations,
				Labels:           tt.fields.Labels,
				Description:      tt.fields.Description,
				ConfigFile:       tt.fields.ConfigFile,
				Xrefs:            tt.fields.Xrefs,
				PanelSectionName: tt.fields.PanelSectionName,
			}
			if err := t.UninstallDependencies(); (err != nil) != tt.wantErr {
				t1.Errorf("UninstallDependencies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
