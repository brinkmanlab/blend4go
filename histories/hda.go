package histories

type HistoryDatasetAssociation struct {
	Id          string `json:"id,omitempty"`
	Accessible  bool   `json:"accessible,omitempty"`
	TypeId      string `json:"type_id,omitempty"`
	FileName    string `json:"file_name,omitempty"`
	Resubmitted bool   `json:"resubmitted,omitempty"`
	CreateTime  string `json:"create_time,omitempty"`
	CreatingJob string `json:"creating_job,omitempty"`
	DatasetId   string `json:"dataset_id,omitempty"`
	FileSize    uint   `json:"file_size,omitempty"`
	FileExt     string `json:"file_ext,omitempty"`
	MiscInfo    string `json:"misc_info,omitempty"`
	HdaLdda     string `json:"hda_ldda,omitempty"`
	DownloadUrl string `json:"download_url,omitempty"`
	State       string `json:"state,omitempty"`
	//DisplayTypes         attr    `json:"display_types,omitempty"`
	//DisplayApps          attr    `json:"display_apps,omitempty"`
	MetadataDbkey string   `json:"metadata_dbkey,omitempty"`
	Type          string   `json:"type,omitempty"`
	MiscBlurb     string   `json:"misc_blurb,omitempty"`
	Peek          string   `json:"peek,omitempty"`
	UpdateTime    string   `json:"update_time,omitempty"`
	DataType      string   `json:"data_type,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Deleted       bool     `json:"deleted,omitempty"`
	HistoryId     string   `json:"history_id,omitempty"`
	//MetaFiles          attr    `json:"meta_files,omitempty"`
	GenomeBuild       string `json:"genome_build,omitempty"`
	MetadataSequences uint   `json:"metadata_sequences,omitempty"`
	Hid               uint   `json:"hid,omitempty"`
	ModelClass        string `json:"model_class,omitempty"`
	MetadataDataLines uint   `json:"metadata_data_lines,omitempty"`
	Annotation        string `json:"annotation,omitempty"`
	//Permissions        attr    `json:"permissions,omitempty"`
	HistoryContentType string `json:"history_content_type,omitempty"`
	Name               string `json:"name,omitempty"`
	Extension          string `json:"extension,omitempty"`
	Visible            bool   `json:"visible,omitempty"`
	Url                string `json:"url,omitempty"`
	Uuid               string `json:"uuid,omitempty"`
	//Visualizations     attr    `json:"visualizations,omitempty"`
	Rerunnable bool   `json:"rerunnable,omitempty"`
	Purged     bool   `json:"purged,omitempty"`
	ApiType    string `json:"api_type,omitempty"`
}
