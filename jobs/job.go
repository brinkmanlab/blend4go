package jobs

import "github.com/brinkmanlab/blend4go"

type Job struct {
	galaxyInstance *blend4go.GalaxyInstance
	Id             blend4go.GalaxyID
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
