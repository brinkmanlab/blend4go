package libraries

// GET /api/libraries/datasets/{encoded_dataset_id} Show the details of a library dataset.

// GET /api/libraries/datasets/{encoded_dataset_id}/versions/{encoded_ldda_id} Display a specific version of a library dataset (i.e. ldda).

// GET /api/libraries/datasets/{encoded_dataset_id}/permissions Display information about current or available roles for a given dataset permission.

// PATCH /api/libraries/datasets/{encoded_dataset_id} Update the given library dataset (the latest linked ldda).

// POST /api/libraries/datasets/{encoded_dataset_id}/permissions Set permissions of the given library dataset to the given role ids.

// DELETE /api/libraries/datasets/{encoded_dataset_id} Mark the dataset deleted or undeleted.

// GET /api/libraries/datasets/download/{archive_format} POST /api/libraries/datasets/download/{archive_format}