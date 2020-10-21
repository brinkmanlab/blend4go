package jobs

import (
	"context"
	"github.com/brinkmanlab/blend4go"
)

const BasePath = "/api/jobs"

func List(ctx context.Context, g *blend4go.GalaxyInstance) ([]*Job, error) {
	//GET /api/jobs
	var jobs []*Job
	_, err := g.List(ctx, BasePath, &jobs, nil)
	return jobs, err
}

func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID) (*Job, error) {
	//GET /api/jobs/{id}
	if res, err := g.Get(ctx, id, &Job{}, nil); err == nil {
		return res.(*Job), nil
	} else {
		return nil, err
	}
}

//POST /api/jobs/search
//GET /api/job_lock
//PUT /api/job_lock
