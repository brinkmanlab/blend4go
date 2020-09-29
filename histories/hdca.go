package histories

type HistoryDatasetCollectionAssociation struct {
	Id                    string        `json:"id,omitempty"`
	HistoryContentType    string        `json:"history_content_type,omitempty"`
	PopulatedStateMessage string        `json:"populated_state_message,omitempty"`
	Name                  string        `json:"name,omitempty"`
	Populated             bool          `json:"populated,omitempty"`
	Deleted               bool          `json:"deleted,omitempty"`
	Type                  string        `json:"type,omitempty"`
	HistoryId             string        `json:"history_id,omitempty"`
	Tags                  []string      `json:"tags,omitempty"`
	Visible               bool          `json:"visible,omitempty"`
	JobSourceId           string        `json:"job_source_id,omitempty"`
	JobSourceType         string        `json:"job_source_type,omitempty"`
	CollectionType        string        `json:"collection_type,omitempty"`
	Url                   string        `json:"url,omitempty"`
	ModelClass            string        `json:"model_class,omitempty"`
	Hid                   uint          `json:"hid,omitempty"`
	ElementCount          uint          `json:"element_count,omitempty"`
	PopulatedState        string        `json:"populated_state,omitempty"`
	Elements              []interface{} `json:"elements,omitempty"`
}
