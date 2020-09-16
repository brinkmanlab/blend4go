package jobs

import "blend4go"

type Job struct {
	galaxyInstance *blend4go.GalaxyInstance
}

func (j *Job) GetBasePath() string {
	return BasePath
}

func (j *Job) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	j.galaxyInstance = g
}
