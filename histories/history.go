package histories

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"path"
)

type History struct {
	galaxyInstance  *blend4go.GalaxyInstance
	Id              blend4go.GalaxyID              `json:"id,omitempty"`
	Importable      bool                           `json:"importable,omitempty"`
	CreateTime      string                         `json:"create_time,omitempty"`
	ContentsUrl     string                         `json:"contents_url,omitempty"`
	Size            uint                           `json:"size,omitempty"`
	UserId          blend4go.GalaxyID              `json:"user_id,omitempty"`
	UsernameAndSlug string                         `json:"username_and_slug,omitempty"`
	Annotation      string                         `json:"annotation,omitempty"`
	StateDetails    map[string]uint                `json:"state_details,omitempty"`
	State           string                         `json:"state,omitempty"`
	Empty           bool                           `json:"empty,omitempty"`
	UpdateTime      string                         `json:"update_time,omitempty"`
	Tags            []string                       `json:"tags,omitempty"`
	Deleted         bool                           `json:"deleted,omitempty"`
	GenomeBuild     string                         `json:"genome_build,omitempty"`
	Slug            string                         `json:"slug,omitempty"`
	Name            string                         `json:"name,omitempty"`
	Url             string                         `json:"url,omitempty"`
	StateIds        map[string][]blend4go.GalaxyID `json:"state_ids,omitempty"`
	Published       bool                           `json:"published,omitempty"`
	ModelClass      string                         `json:"model_class,omitempty"`
	Purged          bool                           `json:"purged,omitempty"`
}

func (h *History) GetBasePath() string {
	return BasePath
}

func (h *History) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	h.galaxyInstance = g
}

func (h *History) GetID() blend4go.GalaxyID {
	return h.Id
}

func (h *History) SetID(id blend4go.GalaxyID) {
	h.Id = id
}

// Create a new history
func (h *History) NewHistory(ctx context.Context, g *blend4go.GalaxyInstance, name string) (*History, error) {
	// POST /api/histories
	if res, err := g.R(ctx).SetResult(&History{}).SetBody(map[string]string{"name": name}).Post(BasePath); err == nil {
		h := res.Result().(*History)
		h.SetGalaxyInstance(g)
		return h, err
	} else {
		return nil, err
	}
}

// GET /api/histories/{id}/citations Return all the citations for the tools used to produce the datasets in the history.

// Delete the history with the given id
func (h *History) Delete(ctx context.Context) error {
	// DELETE /api/histories/{id}
	return h.galaxyInstance.Delete(ctx, h)
}

// Undelete history (that hasn’t been purged) with the given id
func (h *History) Undelete(ctx context.Context) error {
	// POST /api/histories/deleted/{id}/undelete
	_, err := h.galaxyInstance.R(ctx).Post(path.Join(h.GetBasePath(), "deleted", h.GetID(), "undelete"))
	return err
}

// Update the values for the history
func (h *History) Update(ctx context.Context) error {
	// PUT /api/histories/{id}
	_, err := h.galaxyInstance.Put(ctx, h)
	return err
}

// PUT /api/histories/{id}/exports start job (if needed) to create history export for corresponding history.

// GET /api/histories/{id}/exports/{jeha_id} If ready and available, return raw contents of exported history.

// GET /api/histories/{id}/custom_builds_metadata Returns meta data for custom builds.
