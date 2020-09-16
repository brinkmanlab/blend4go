package workflows

import "github.com/brinkmanlab/blend4go"

const BasePath = "/api/workflows"

/*
Get list of workflows
published – if True, show also published workflows
hidden – if True, show hidden workflows
deleted – if True, show deleted workflows
missingTools – if True, include a list of missing tools per workflow
*/
func List(g *blend4go.GalaxyInstance, published, hidden, deleted, missingTools bool) ([]StoredWorkflow, error) {
	q := make(map[string]string)
	if published {
		q["show_published"] = "True"
	}
	if hidden {
		q["show_hidden"] = "True"
	}
	if deleted {
		q["show_deleted"] = "True"
	}
	if missingTools {
		q["missing_tools"] = "True"
	}
	// GET /api/workflows
	res, err := g.List(BasePath, []StoredWorkflow{}, &q)
	return res.([]StoredWorkflow), err
}

// Displays information needed to run a workflow.
func Get(g *blend4go.GalaxyInstance, id blend4go.GalaxyID) (*StoredWorkflow, error) {
	// TODO instance (boolean) – true if fetch by Workflow ID instead of StoredWorkflow id, false by default.
	// GET /api/workflows/{encoded_workflow_id}
	res, err := g.Get(id, &StoredWorkflow{})
	return res.(*StoredWorkflow), err
}

// Get workflows present in the tools panel GET /api/workflows/menu
// Save workflow menu to be shown in the tool panel PUT /api/workflows/menu

// POST /api/workflows Run or create workflows from the api.

// POST /api/workflows/build_module Builds module models for the workflow editor.

// POST /api/workflows/get_tool_predictions Fetch predicted tools for a workflow

// POST /api/workflows/import Import a workflow shared by other users.
