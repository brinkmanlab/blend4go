// roles models represent and manipulate roles within a Galaxy instance
// Relevant api endpoints are: `/api/roles`
package roles

import (
	"context"
	"github.com/brinkmanlab/blend4go"
)

const BasePath = "/api/roles"

/*
Displays a collection (list) of roles.
deleted - If true, show deleted roles
*/
func List(ctx context.Context, g *blend4go.GalaxyInstance, deleted bool) ([]*Role, error) {
	q := make(map[string]string)
	if deleted {
		q["deleted"] = "true"
	}
	// GET /api/roles GET /api/roles/deleted
	var roles []*Role
	_, err := g.List(ctx, BasePath, &roles, &q)
	for _, role := range roles {
		role.Deleted = deleted
	}
	return roles, err
}

// Displays information about a role.
func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID, deleted bool) (*Role, error) {
	// GET /api/roles/{encoded_id} GET /api/roles/deleted/{encoded_id}
	if res, err := g.Get(ctx, id, &Role{Deleted: deleted}, nil); err == nil {
		return res.(*Role), err
	} else {
		return nil, err
	}
}

// GetName get role by name
func GetName(ctx context.Context, g *blend4go.GalaxyInstance, name string) (*Role, error) {
	for _, deleted := range []bool{true, false} {
		if roles, err := List(ctx, g, deleted); err == nil {
			for _, role := range roles {
				if role.Name == name {
					return role, nil
				}
			}
		}
	}
	return nil, nil
}
