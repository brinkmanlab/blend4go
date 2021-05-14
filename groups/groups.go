// groups models represent and manipulate groups within a Galaxy instance
// Relevant api endpoints are: `/api/groups`
package groups

import (
	"context"
	"path"

	"github.com/brinkmanlab/blend4go"
)

const BasePath = "/api/groups"

// get full group, populating model
func get(ctx context.Context, g *blend4go.GalaxyInstance, model *Group) (*Group, error) {
	// GET /api/groups/{encoded_id} GET /api/groups/deleted/{encoded_id}
	if res, err := g.Get(ctx, model.GetID(), model, nil); err == nil {
		group := res.(*Group)
		group.populateAssoc()
		return group, nil
	} else {
		return nil, err
	}
}

// List groups
// deleted - If true, show deleted groups
func List(ctx context.Context, g *blend4go.GalaxyInstance, deleted bool) ([]*Group, error) {
	fullpath := BasePath
	if deleted {
		fullpath = path.Join(BasePath, "deleted")
	}
	// GET /api/groups GET /api/groups/deleted
	var groups []*Group
	_, err := g.List(ctx, fullpath, &groups, nil)
	for _, group := range groups {
		group.Deleted = deleted
		group, err = get(ctx, g, group)
		if err != nil {
			return nil, err
		}
		group.populateAssoc()
	}
	return groups, err
}

// Get group
func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID, deleted bool) (*Group, error) {
	return get(ctx, g, &Group{Id: id, Deleted: deleted})
}

// GetName get group by name
func GetName(ctx context.Context, g *blend4go.GalaxyInstance, name string) (*Group, error) {
	for _, deleted := range []bool{true, false} {
		if groups, err := List(ctx, g, deleted); err == nil {
			for _, group := range groups {
				if group.Name == name {
					return group, nil
				}
			}
		}
	}
	return nil, nil
}
