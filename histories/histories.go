package histories

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"path"
)

const BasePath = "/api/histories"

func list(ctx context.Context, g *blend4go.GalaxyInstance, category string) ([]History, error) {
	// GET /api/histories
	// GET /api/histories/deleted return deleted histories for the current user
	p := BasePath
	if category != "" {
		p = path.Join(BasePath, category)
	}
	if res, err := g.List(ctx, p, []History{}, nil); err == nil {
		return res.([]History), nil
	} else {
		return nil, err
	}
}

// List histories for the current user
func List(ctx context.Context, g *blend4go.GalaxyInstance) ([]History, error) {
	return list(ctx, g, "")
}

// List deleted histories for the current user
func ListDeleted(ctx context.Context, g *blend4go.GalaxyInstance) ([]History, error) {
	return list(ctx, g, "deleted")
}

// Get all histories that are published
func GetPublished(ctx context.Context, g *blend4go.GalaxyInstance) ([]History, error) {
	// GET /api/histories/published
	return list(ctx, g, "published")
}

// Get all histories that are shared with the current user
func GetSharedWithMe(ctx context.Context, g *blend4go.GalaxyInstance) ([]History, error) {
	// GET /api/histories/shared_with_me
	return list(ctx, g, "shared_with_me")
}

// Get the history with id
func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID) (*History, error) {
	//GET /api/histories/{id}
	if res, err := g.Get(ctx, id, &History{}); err == nil {
		return res.(*History), nil
	} else {
		return nil, err
	}
}

// Get the most recently used history
func GetMostRecent(ctx context.Context, g *blend4go.GalaxyInstance) (*History, error) {
	// GET /api/histories/most_recently_used
	return Get(ctx, g, "most_recently_used")
}
