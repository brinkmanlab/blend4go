package workflows

import (
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/histories"
	"github.com/brinkmanlab/blend4go/jobs"
)

type WorkflowInvocationStep struct {
	Id                   blend4go.GalaxyID                                `json:"id"`
	WorkflowStepUuid     string                                           `json:"workflow_step_uuid"`
	UpdateTime           string                                           `json:"update_time"`
	Jobs                 []*jobs.Job                                      `json:"jobs"`
	JobId                string                                           `json:"job_id"`
	Outputs              []*histories.HistoryDatasetAssociation           `json:"outputs"`
	OrderIndex           uint                                             `json:"order_index"`
	OutputCollections    []*histories.HistoryDatasetCollectionAssociation `json:"output_collections"`
	WorkflowStepLabel    string                                           `json:"workflow_step_label"`
	State                string                                           `json:"state"`
	Action               string                                           `json:"action"`
	ModelClass           string                                           `json:"model_class"`
	WorkflowStepId       blend4go.GalaxyID                                `json:"workflow_step_id"`
	WorkflowInvocationId blend4go.GalaxyID                                `json:"workflow_invocation_id"`
}
