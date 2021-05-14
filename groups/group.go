package groups

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/roles"
	"github.com/brinkmanlab/blend4go/users"
	"path"
)

type groupAssociations struct {
	group       *Group
	association string
}

func (ga *groupAssociations) Add(ctx context.Context, id blend4go.GalaxyID) error {
	// PUT /api/groups/{encoded_group_id}/users/{encoded_user_id} Adds a user to a group
	// PUT /api/groups/{encoded_group_id}/roles/{encoded_role_id} Adds a role to a group
	panic("implement me")
	if res, err := ga.group.galaxyInstance.R(ctx).Put(path.Join(BasePath, ga.group.Id, ga.association, id)); err == nil {
		_, err := blend4go.HandleResponse(res)
		return err
	} else {
		return err
	}
}

func (ga *groupAssociations) Remove(ctx context.Context, id blend4go.GalaxyID) error {
	// DELETE /api/groups/{encoded_group_id}/roles/{encoded_role_id} Removes a role from a group
	// DELETE /api/groups/{encoded_group_id}/users/{encoded_user_id} Removes a user from a group
	panic("implement me")
	if res, err := ga.group.galaxyInstance.R(ctx).Delete(path.Join(BasePath, ga.group.Id, ga.association, id)); err == nil {
		_, err := blend4go.HandleResponse(res)
		return err
	} else {
		return err
	}
}

type userGroupAssociations groupAssociations

func (u *userGroupAssociations) List(ctx context.Context) ([]*users.User, error) {
	// GET /api/groups/{encoded_group_id}/users Displays a collection (list) of groups.
	// GET /api/groups/{encoded_group_id}/users/{encoded_user_id} Displays information about a group user.
	panic("implement me")
}

type roleGroupAssociations groupAssociations

func (r *roleGroupAssociations) List(ctx context.Context) ([]*roles.Role, error) {
	// GET /api/groups/{encoded_group_id}/roles Displays a collection (list) of roles.
	// GET /api/groups/{encoded_group_id}/roles/{encoded_role_id} Displays information about a group role.
	panic("implement me")
}

type Group struct {
	galaxyInstance *blend4go.GalaxyInstance
	Id             blend4go.GalaxyID `json:"id,omitempty"`
	Name           string            `json:"name,omitempty"`
	Users          userGroupAssociations
	Roles          roleGroupAssociations
	Deleted        bool
}

func (g *Group) populateAssoc() {
	g.Users.association = "users"
	g.Roles.association = "roles"
	g.Users.group = g
	g.Roles.group = g
}

func NewGroup(ctx context.Context, name string) (*Group, error) {
	g := &Group{
		Name: name,
	}
	g.populateAssoc()
	// POST /api/groups Creates a new group.
	panic("implement me")
}

func (g *Group) GetBasePath() string {
	return BasePath
}

func (g *Group) SetGalaxyInstance(instance *blend4go.GalaxyInstance) {
	g.galaxyInstance = instance
}

func (g *Group) GetID() blend4go.GalaxyID {
	return g.Id
}

func (g *Group) SetID(id blend4go.GalaxyID) {
	g.Id = id
}

func (g *Group) Update(ctx context.Context) error {
	// PUT /api/groups/{encoded_group_id} Modifies a group.
	panic("implement me")
}

func (g *Group) Delete(ctx context.Context) error {
	panic("implement me")
}
