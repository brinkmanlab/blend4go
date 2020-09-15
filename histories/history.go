package histories

// GET /api/histories/{id} return the history with id

// GET /api/histories/deleted/{id} return the deleted history with id

// GET /api/histories/most_recently_used return the most recently used history
// GET /api/histories/published return all histories that are published
// GET /api/histories/shared_with_me return all histories that are shared with the current user

// GET /api/histories/{id}/citations Return all the citations for the tools used to produce the datasets in the history.

// DELETE /api/histories/{id} delete the history with the given id

// POST /api/histories/deleted/{id}/undelete undelete history (that hasn’t been purged) with the given id

// PUT /api/histories/{id} updates the values for the history with the given id

// PUT /api/histories/{id}/exports start job (if needed) to create history export for corresponding history.

// GET /api/histories/{id}/exports/{jeha_id} If ready and available, return raw contents of exported history.

// GET /api/histories/{id}/custom_builds_metadata Returns meta data for custom builds.
