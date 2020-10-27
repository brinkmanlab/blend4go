package libraries

import "github.com/brinkmanlab/blend4go"

type LibraryItem struct {
	libraryId blend4go.GalaxyID
	Id        blend4go.GalaxyID `json:"id"`
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Url       string            `json:"url"`
}
