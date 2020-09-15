package libraries

// GET /api/libraries/{encoded_id} returns detailed information about a library
// GET /api/libraries/deleted/{encoded_id} returns detailed information about a deleted library

// PATCH /api/libraries/{encoded_id} Updates the library defined by an encoded_id with the data in the payload.

// DELETE /api/libraries/{id} marks the library with the given id as deleted (or removes the deleted mark if the undelete param is true)

// GET /api/libraries/{id}/permissions Load all permissions for the given library id and return it.

// POST /api/libraries/{encoded_library_id}/permissions Set permissions of the given library to the given role ids.

// GET /api/libraries/{library_id}/contents Return a list of library files and folders.

// GET /api/libraries/{library_id}/contents/{id} Returns information about library file or folder.

// POST /api/libraries/{library_id}/contents Create a new library file or folder.

// PUT /api/libraries/{library_id}/contents/{id} Create an ImplicitlyConvertedDatasetAssociation.

// DELETE /api/libraries/{library_id}/contents/{id} Delete the LibraryDataset with the given id.

// GET /api/libraries/datasets/{encoded_dataset_id} Show the details of a library dataset.

// GET /api/libraries/datasets/{encoded_dataset_id}/versions/{encoded_ldda_id} Display a specific version of a library dataset (i.e. ldda).

// GET /api/libraries/datasets/{encoded_dataset_id}/permissions Display information about current or available roles for a given dataset permission.

// PATCH /api/libraries/datasets/{encoded_dataset_id} Update the given library dataset (the latest linked ldda).

// POST /api/libraries/datasets/{encoded_dataset_id}/permissions Set permissions of the given library dataset to the given role ids.

// DELETE /api/libraries/datasets/{encoded_dataset_id} Mark the dataset deleted or undeleted.

// POST /api/libraries/datasets Load dataset(s) from the given source into the library.

// GET /api/libraries/datasets/download/{archive_format} POST /api/libraries/datasets/download/{archive_format}
