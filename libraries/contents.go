package libraries

import (
	"context"
	"github.com/brinkmanlab/blend4go"
)

type libraryItem struct {
	library	*Library
	Id        blend4go.GalaxyID `json:"id"`
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Url       string            `json:"url"`
}

func (i *libraryItem) GetModel(ctx context.Context) (*blend4go.GalaxyModel, error) {
	switch i.Type {
	case "folder":
		return i.library.
	}
}

// GET /api/libraries/{library_id}/contents/{id} Returns information about library file or folder.

// POST /api/libraries/{library_id}/contents Create a new library file or folder.

// PUT /api/libraries/{library_id}/contents/{id} Create an ImplicitlyConvertedDatasetAssociation.

// DELETE /api/libraries/{library_id}/contents/{id} Delete the LibraryDataset with the given id.
