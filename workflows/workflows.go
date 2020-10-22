// workflows models represent and manipulate workflows within a Galaxy instance
// Relevant api endpoints are: `/api/workflows`, `/api/invocations`
package workflows

import (
	"context"
	"encoding/json"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/repositories"
	"github.com/jmespath/go-jmespath"
	"path"
)

const BasePath = "/api/workflows"

/*
Get list of workflows
published – if True, show also published workflows
hidden – if True, show hidden workflows
deleted – if True, show deleted workflows
missingTools – if True, include a list of missing tools per workflow
*/
func List(ctx context.Context, g *blend4go.GalaxyInstance, published, hidden, deleted, missingTools bool) ([]*StoredWorkflow, error) {
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
	var workflows []*StoredWorkflow
	_, err := g.List(ctx, BasePath, &workflows, &q)
	return workflows, err
}

// Displays information needed to run a workflow.
func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID) (*StoredWorkflow, error) {
	// TODO instance (boolean) – true if fetch by Workflow ID instead of StoredWorkflow id, false by default.
	// GET /api/workflows/{encoded_workflow_id}
	if res, err := g.Get(ctx, id, &StoredWorkflow{}, nil); err == nil {
		return res.(*StoredWorkflow), err
	} else {
		return nil, err
	}
}

// Recursively search for all tool ids in workflow
func findToolIDs(data map[string]interface{}) ([]*repositories.Repository, error) {
	var res []*repositories.Repository

	// Search subworkflows
	// This can be replaced by a recursive query when added to JMESPath https://github.com/jmespath/jmespath.py/issues/110
	if subworkflows, err := jmespath.Search("steps.*.subworkflow", data); err == nil {
		for _, subworkflow := range subworkflows.([]interface{}) {
			if ids, err := findToolIDs(subworkflow.(map[string]interface{})); err == nil {
				res = append(res, ids...)
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	// Append repositories
	if repos, err := jmespath.Search("steps.*.tool_shed_repository", data); err == nil {
		for _, repo := range repos.([]interface{}) {
			r := repo.(map[string]interface{})
			res = append(res, &repositories.Repository{
				Name:              r["name"].(string),
				ToolShed:          r["tool_shed"].(string),
				Owner:             r["owner"].(string),
				ChangesetRevision: r["changeset_revision"].(string),
			})
		}
	} else {
		return nil, err
	}

	return res, nil
}

func Repositories(workflow string) ([]*repositories.Repository, error) {
	data := make(map[string]interface{})

	if err := json.Unmarshal([]byte(workflow), &data); err != nil {
		return nil, err
	}

	// Search for all tool ids in workflow json
	if res, err := findToolIDs(data); err == nil {
		// Reduce to unique values
		set := make(map[string]*repositories.Repository)
		for _, i := range res {
			i.SetID(path.Join(i.ToolShed, i.Owner, i.Name, i.ChangesetRevision))
			set[i.GetID()] = i
		}
		// Convert keys to list
		tools := make([]*repositories.Repository, 0, len(set))
		for _, i := range set {
			tools = append(tools, i)
		}
		return tools, nil
	} else {
		return nil, err
	}
}

// Get workflows present in the tools panel GET /api/workflows/menu
// Save workflow menu to be shown in the tool panel PUT /api/workflows/menu

// POST /api/workflows Run or create workflows from the api.

// POST /api/workflows/build_module Builds module models for the workflow editor.

// POST /api/workflows/get_tool_predictions Fetch predicted tools for a workflow

// POST /api/workflows/import Import a workflow shared by other users.
