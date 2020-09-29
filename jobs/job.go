package jobs

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/histories"
	"path"
)

type Job struct {
	galaxyInstance *blend4go.GalaxyInstance
	Id             blend4go.GalaxyID `json:"id,omitempty"`
	ToolId         string            `json:"tool_id,omitempty"`
	UpdateTime     string            `json:"update_time,omitempty"`
	HistoryId      string            `json:"history_id,omitempty"`
	ExitCode       uint              `json:"exit_code,omitempty"`
	State          string            `json:"state,omitempty"`
	CreateTime     string            `json:"create_time,omitempty"`
	ModelClass     string            `json:"model_class,omitempty"`
	Inputs         interface{}       `json:"inputs,omitempty"`
	Outputs        interface{}       `json:"outputs,omitempty"`
	Params         interface{}       `json:"params,omitempty"`
}

func (j *Job) GetBasePath() string {
	return BasePath
}

func (j *Job) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	j.galaxyInstance = g
}

func (j *Job) GetID() blend4go.GalaxyID {
	return j.Id
}

func (j *Job) SetID(id blend4go.GalaxyID) {
	j.Id = id
}

type invocationResponse struct {
	Jobs                []*Job                                           `json:"jobs"`
	Outputs             []*histories.HistoryDatasetAssociation           `json:"outputs"`
	OutputCollections   []*histories.HistoryDatasetCollectionAssociation `json:"output_collections"`
	ImplicitCollections []*histories.HistoryDatasetCollectionAssociation `json:"implicit_collections"`
}

// Execute tool with a given parameter payload
func NewJob(ctx context.Context, g *blend4go.GalaxyInstance, payload map[string]interface{}) ([]*Job, []*histories.HistoryDatasetAssociation, []*histories.HistoryDatasetCollectionAssociation, []*histories.HistoryDatasetCollectionAssociation, error) {
	//POST /api/tools
	if res, err := g.R(ctx).SetBody(payload).SetResult(&invocationResponse{}).Post("/api/tools"); err == nil {
		r := res.Result().(invocationResponse)
		return r.Jobs, r.Outputs, r.OutputCollections, r.ImplicitCollections, err
	} else {
		return nil, nil, nil, nil, err
	}
}

// Delete or stop a job
func (j *Job) Delete(ctx context.Context) error {
	//Delete /api/jobs/{id}
	return j.galaxyInstance.Delete(ctx, j)
}

// Resume paused job
func (j *Job) Resume(ctx context.Context) error {
	//PUT /api/jobs/{id}/resume
	_, err := j.galaxyInstance.R(ctx).Put(path.Join(j.GetBasePath(), j.GetID(), "resume"))
	return err
}

//GET /api/jobs/{id}/common_problems
//GET /api/jobs/{id}/inputs
//GET /api/jobs/{id}/outputs
//GET /api/jobs/{job_id}/metrics
//GET /api/jobs/{job_id}/destination_params
//GET /api/jobs/{job_id}/parameters_display
//GET /api/jobs/{id}/build_for_rerun
//POST /api/jobs/{id}/error
