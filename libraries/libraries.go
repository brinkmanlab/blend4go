// libraries models represent and manipulate libraries within a Galaxy instance
// Relevant api endpoints are: `/api/libraries`
package libraries

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"path"
)

const BasePath = "/api/libraries"

// Returns a list of summary data for all libraries
func List(ctx context.Context, g *blend4go.GalaxyInstance, deleted bool) ([]*Library, error) {
	p := BasePath
	if deleted {
		p = path.Join(BasePath, "deleted")
	}
	var libraries []*Library
	// GET /api/libraries
	_, err := g.List(ctx, p, &libraries, nil)
	return libraries, err
}

// returns detailed information about a library
func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID, deleted bool) (*Library, error) {
	// GET /api/libraries/{encoded_id}
	// GET /api/libraries/deleted/{encoded_id} returns detailed information about a deleted library
	if res, err := g.Get(ctx, id, &Library{}, nil); err == nil {
		return res.(*Library), err
	} else {
		return nil, err
	}
}
