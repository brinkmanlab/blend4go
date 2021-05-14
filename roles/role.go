package roles

import (
	"context"
	"github.com/brinkmanlab/blend4go"
)

type Role struct {
	galaxyInstance *blend4go.GalaxyInstance
	Id             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Type           string `json:"type,omitempty"`
	Deleted        bool   `json:"-"`
}

func (r *Role) GetBasePath() string {
	return BasePath
}

func (r *Role) SetGalaxyInstance(instance *blend4go.GalaxyInstance) {
	r.galaxyInstance = instance
}

func (r *Role) GetID() blend4go.GalaxyID {
	return r.Id
}

func (r *Role) SetID(id blend4go.GalaxyID) {
	r.Id = id
}

func NewRole(ctx context.Context, name string) (*Role, error) {
	// POST /api/roles Creates a new role.
	panic("implement me")
}

func (r *Role) Update(ctx context.Context) error {
	panic("implement me")
}
