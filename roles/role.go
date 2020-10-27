package roles

import "github.com/brinkmanlab/blend4go"

type Role struct {
	galaxyInstance *blend4go.GalaxyInstance
	Id             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	ModelClass     string `json:"model_class,omitempty"`
	Url            string `json:"url,omitempty"`
	Description    string `json:"description,omitempty"`
	Type           string `json:"type,omitempty"`
}
