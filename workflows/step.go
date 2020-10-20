package workflows

import (
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/histories"
	"github.com/brinkmanlab/blend4go/jobs"
)

type WorkflowInvocationStep struct {
	Id                   blend4go.GalaxyID                                `json:"id,omitempty"`
	WorkflowStepUuid     string                                           `json:"workflow_step_uuid,omitempty"`
	UpdateTime           string                                           `json:"update_time,omitempty"`
	Jobs                 []*jobs.Job                                      `json:"jobs,omitempty"`
	JobId                string                                           `json:"job_id,omitempty"`
	Outputs              []*histories.HistoryDatasetAssociation           `json:"outputs,omitempty"`
	OrderIndex           uint                                             `json:"order_index,omitempty"`
	OutputCollections    []*histories.HistoryDatasetCollectionAssociation `json:"output_collections,omitempty"`
	WorkflowStepLabel    string                                           `json:"workflow_step_label,omitempty"`
	State                string                                           `json:"state,omitempty"`
	Action               string                                           `json:"action,omitempty"`
	ModelClass           string                                           `json:"model_class,omitempty"`
	WorkflowStepId       blend4go.GalaxyID                                `json:"workflow_step_id,omitempty"`
	WorkflowInvocationId blend4go.GalaxyID                                `json:"workflow_invocation_id,omitempty"`
}
