package tools

import (
	"context"
	"github.com/brinkmanlab/blend4go"
)

const BasePath = "/api/tools"

// returns a list of tools defined by parameters
func List(ctx context.Context, g *blend4go.GalaxyInstance) ([]*ToolSection, error) {
	if res, err := g.List(ctx, BasePath, []*ToolSection{}, &map[string]string{}); err == nil {
		return res.([]*ToolSection), err
	} else {
		return nil, err
	}
}

func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID) (*Tool, error) {
	if res, err := g.Get(ctx, id, &Tool{}, nil); err == nil {
		return res.(*Tool), nil
	} else {
		return nil, err
	}
}

//GET /api/tools/tests_summary
//GET /api/tools/all_requirements Return list of unique requirements for all tools.
//GET /api/tools/error_stack Returns global tool error stack
