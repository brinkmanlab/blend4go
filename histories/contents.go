package histories

type HistoryDatasetAssociation struct {}
type HistoryDatasetCollectionAssociation struct {

}

// GET /api/histories/{history_id}/contents



// GET /api/histories/{history_id}/contents/{id}
// GET /api/histories/{history_id}/contents/{type}/{id} return detailed information about an HDA or HDCA within a history

// GET /api/histories/{history_id}/jobs_summary return job state summary info for jobs, implicit groups jobs for collections or workflow invocations
// GET /api/histories/{history_id}/contents/{type}/{id}/jobs_summary return detailed information about an HDA or HDCAs jobs

// GET /api/histories/{history_id}/contents/{id}/download
// GET /api/dataset_collection/{id}/download Download the content of a HistoryDatasetCollection as a tgz archive while maintaining approximate collection structure.

// POST /api/histories/{history_id}/contents/{type}s
// POST /api/histories/{history_id}/contents create a new HDA or HDCA

// GET /api/histories/{history_id}/contents/datasets/{encoded_dataset_id}/permissions Display information about current or available roles for a given dataset permission.
// PUT /api/histories/{history_id}/contents/datasets/{encoded_dataset_id}/permissions Set permissions of the given library dataset to the given role ids.

// PUT /api/histories/{history_id}/contents
// PUT /api/histories/{history_id}/contents/{id} updates the values for the history content item with the given id

// PUT /api/histories/{history_id}/contents/{id}/validate updates the values for the history content item with the given id

// DELETE /api/histories/{history_id}/contents/{id}
// DELETE /api/histories/{history_id}/contents/{type}s/{id} delete the history content with the given id and specified type (defaults to dataset)

// GET /api/histories/{history_id}/contents/archive/{id}
// GET /api/histories/{history_id}/contents/archive/{filename}.{format} build and return a compressed archive of the selected history contents

