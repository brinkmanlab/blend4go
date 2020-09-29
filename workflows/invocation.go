package workflows

import "github.com/brinkmanlab/blend4go"

type WorkflowInvocation struct {
	galaxyInstance *blend4go.GalaxyInstance
	Id             blend4go.GalaxyID `json:"id,omitempty"`
	UpdateTime     string            `json:"update_time,omitempty"`
	Uuid           string            `json:"uuid,omitempty"`
	// Outputs ? `json:"outputs,omitempty"`
	// Output_collections ? `json:"output_collections,omitempty"`
	HistoryId  blend4go.GalaxyID `json:"history_id,omitempty"`
	WorkflowId blend4go.GalaxyID `json:"workflow_id,omitempty"`
	State      string            `json:"state,omitempty"`
	ModelClass string            `json:"model_class,omitempty"`
	// Inputs ? `json:"inputs,omitempty"`
	Steps []*WorkflowInvocationStep `json:"steps,omitempty"`
}

func (w *WorkflowInvocation) GetBasePath() string {
	return BasePath
}

func (w *WorkflowInvocation) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	w.galaxyInstance = g
}

func (w *WorkflowInvocation) GetID() blend4go.GalaxyID {
	return w.Id
}

func (w *WorkflowInvocation) SetID(id blend4go.GalaxyID) {
	w.Id = id
}

// DELETE /api/workflows/{workflow_id}/invocations/{invocation_id} DELETE /api/invocations/{invocation_id} Cancel the specified workflow invocation.

// GET /api/workflows/{workflow_id}/invocations/{invocation_id}/report GET /api/invocations/{invocation_id}/report Get JSON summarizing invocation for reporting.

// GET /api/workflows/{workflow_id}/invocations/{invocation_id}/report.pdf GET /api/invocations/{invocation_id}/report.pdf Get JSON summarizing invocation for reporting.

// GET /api/invocations/{invocations_id}/biocompute Return a BioCompute Object for the workflow invocation.

// GET /api/invocations/{invocations_id}/biocompute/download Returns a selected BioCompute Object as a file for download (HTTP headers configured with filename and such).

// GET /api/workflows/{workflow_id}/invocations/{invocation_id}/steps/{step_id} GET /api/invocations/{invocation_id}/steps/{step_id}

// GET /api/workflows/{workflow_id}/invocations/{invocation_id}/step_jobs_summary GET /api/invocations/{invocation_id}/step_jobs_summary return job state summary info aggregated across per step of the workflow invocation

// GET /api/workflows/{workflow_id}/invocations/{invocation_id}/jobs_summary GET /api/invocations/{invocation_id}/jobs_summary return job state summary info aggregated across all current jobs of workflow invocation

// PUT /api/workflows/{workflow_id}/invocations/{invocation_id}/steps/{step_id} PUT /api/invocations/{invocation_id}/steps/{step_id} Update state of running workflow step invocation - still very nebulous but this would be for stuff like confirming paused steps can proceed etc….
