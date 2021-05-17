// quotas models represent and manipulate quotas within a Galaxy instance
// Relevant api endpoints are: `/api/quotas`
package quotas

import (
	"context"
	"path"

	"github.com/brinkmanlab/blend4go"
)

const BasePath = "/api/quotas"

// get full quota, populating model
func get(ctx context.Context, g *blend4go.GalaxyInstance, model *Quota) (*Quota, error) {
	// GET /api/quotas/{encoded_id} GET /api/quotas/deleted/{encoded_id}
	if res, err := g.Get(ctx, model.GetID(), model, nil); err == nil {
		quota := res.(*Quota)
		quota.populateIndirect()
		return quota, nil
	} else {
		return nil, err
	}
}

// List quotas
// deleted - If true, show deleted quotas
func List(ctx context.Context, g *blend4go.GalaxyInstance, deleted bool) ([]*Quota, error) {
	fullpath := BasePath
	if deleted {
		fullpath = path.Join(BasePath, "deleted")
	}
	// GET /api/quotas GET /api/quotas/deleted
	var quotas []*Quota
	_, err := g.List(ctx, fullpath, &quotas, nil)
	for _, quota := range quotas {
		quota.Deleted = deleted
		quota, err = get(ctx, g, quota)
		if err != nil {
			return nil, err
		}
		quota.populateIndirect()
	}
	return quotas, err
}

// Get quota
func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID, deleted bool) (*Quota, error) {
	return get(ctx, g, &Quota{Id: id, Deleted: deleted})
}

// GetName get quota by name
func GetName(ctx context.Context, g *blend4go.GalaxyInstance, name string) (*Quota, error) {
	for _, deleted := range []bool{true, false} {
		if quotas, err := List(ctx, g, deleted); err == nil {
			for _, quota := range quotas {
				if quota.Name == name {
					return quota, nil
				}
			}
		}
	}
	return nil, nil
}
