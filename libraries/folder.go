package libraries

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/jobs"
	"path"
)

type LibraryFolder struct {
	galaxyInstance *blend4go.GalaxyInstance
	Id        blend4go.GalaxyID `json:"id"`
	Name      string            `json:"name"`
	ParentID  blend4go.GalaxyID `json:"parent_id"`
	Description string `json:"description"`
	ItemCount int `json:"item_count"`
	GenomeBuild string `json:"genome_build"`
	UpdateTime string `json:"update_time"`
	Deleted bool `json:"deleted"`
	LibraryPath []string `json:"library_path"`
	ParentLibraryID blend4go.GalaxyID `json:"parent_library_id"`
}

func (f *LibraryFolder) GetBasePath() string {
	return path.Join(BasePath, f.ParentLibraryID, "contents", f.Id)
}

func (f *LibraryFolder) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	f.galaxyInstance = g
}

func (f *LibraryFolder) GetID() blend4go.GalaxyID {
	return f.Id
}

func (f *LibraryFolder) SetID(id blend4go.GalaxyID) {
	f.Id = id
}

// Import dataset(s) from the given source into the folder.
func (f *LibraryFolder) Import(ctx context.Context, source string, link, preserve, tagUsingFileName bool, fileType, db string) (*jobs.Job, error) {
	if fileType == "" {
		fileType = "auto"
	}
	if db == "" {
		db = "?"
	}
	payload := struct {
		FolderID    blend4go.GalaxyID   `json:"encoded_folder_id"`
		Source string `json:"source"`
		LinkData    bool `json:"link_data"`
		PreserveDirs bool `json:"preserve_dirs"`
		FileType string `json:"file_type"`
		DBKey string `json:"dbkey"`
		TagUsingFileNames bool `json:"tag_using_filenames"`
	}{
		FolderID: f.Id,
		Source: source,
		LinkData: link,
		PreserveDirs: preserve,
		FileType: fileType,
		DBKey: db,
		TagUsingFileNames: tagUsingFileName,
	}
	// POST /api/libraries/datasets
	if res, err := f.galaxyInstance.R(ctx).SetBody(payload).SetResult(&jobs.Job{}).Post(path.Join(BasePath, "datasets")); err == nil {
		result, err := blend4go.HandleResponse(res)
		job := result.(*jobs.Job)
		job.SetGalaxyInstance(f.galaxyInstance)
		return job, err
	} else {
		return nil, err
	}
}