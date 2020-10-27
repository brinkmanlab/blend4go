package libraries

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/roles"
	"path"
)

type Library struct {
	galaxyInstance   *blend4go.GalaxyInstance
	Id               blend4go.GalaxyID `json:"id,omitempty"`
	Deleted          bool              `json:"deleted,omitempty"`
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	Synopsis         string            `json:"synopsis,omitempty"`
	RootFolderId     blend4go.GalaxyID `json:"root_folder_id,omitempty"`
	CreateTime       string            `json:"create_time,omitempty"`
	Public           bool              `json:"public,omitempty"`
	CreateTimePretty string            `json:"create_time_pretty,omitempty"`
	CanUserAdd       bool              `json:"can_user_add,omitempty"`
	CanUserModify    bool              `json:"can_user_modify,omitempty"`
	CanUserManage    bool              `json:"can_user_manage,omitempty"`
	ModelClass       string            `json:"model_class,omitempty"`
}

type LibraryPermissions struct {
	AccessLibraryRoleList  []*roles.Role
	ModifyLibraryRoleList  []*roles.Role
	ManageLibraryRoleList  []*roles.Role
	AddLibraryItemRoleList []*roles.Role
}

func (l *Library) GetBasePath() string {
	if l.Deleted {
		return path.Join(BasePath, "deleted")
	}
	return BasePath
}

func (l *Library) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	l.galaxyInstance = g
}

func (l *Library) GetID() blend4go.GalaxyID {
	return l.Id
}

func (l *Library) SetID(id blend4go.GalaxyID) {
	l.Id = id
}

// Creates a new library.
func NewLibrary(ctx context.Context, g *blend4go.GalaxyInstance, name, description, synopsis string) (*Library, error) {
	// POST /api/libraries
	if res, err := g.R(ctx).SetResult(&Library{galaxyInstance: g}).SetBody(map[string]string{
		"name":        name,
		"description": description,
		"synopsis":    synopsis,
	}).Post(BasePath); err == nil {
		if result, err := blend4go.HandleResponse(res); err == nil {
			return result.(*Library), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// Updates the library defined by an encoded_id with the data in the payload.
func (l *Library) Update(ctx context.Context) error {
	// PATCH /api/libraries/{encoded_id}
	_, err := l.galaxyInstance.Patch(ctx, l, nil)
	return err
}

func (l *Library) delete(ctx context.Context, undelete bool) error {
	params := map[string]string{}
	if undelete {
		params["undelete"] = "true"
	}
	// DELETE /api/libraries/{id}
	return l.galaxyInstance.Delete(ctx, l, &params)
}

// marks the library with the given id as deleted
func (l *Library) Delete(ctx context.Context) error {
	return l.delete(ctx, false)
}

// marks the library with the given id as not deleted
func (l *Library) Undelete(ctx context.Context) error {
	return l.delete(ctx, true)
}

// Load all permissions for the given library id and return it.
func (l *Library) Permissions(ctx context.Context) (*LibraryPermissions, error) {
	// GET /api/libraries/{id}/permissions
	if res, err := l.galaxyInstance.R(ctx).SetResult(map[string]interface{}{}).Get(path.Join(l.GetBasePath(), "permissions")); err == nil {
		if result, err := blend4go.HandleResponse(res); err == nil {
			// Handle https://github.com/galaxyproject/galaxy/issues/10562
			libPerms := &LibraryPermissions{}
			perms := *result.(*map[string]interface{})
			for _, role := range perms["access_library_role_list"].([]interface{}) {
				libPerms.AccessLibraryRoleList = append(libPerms.AccessLibraryRoleList, &roles.Role{Id: role.([]string)[1], Name: role.([]string)[0]})
			}
			for _, role := range perms["modify_library_role_list"].([]interface{}) {
				libPerms.ModifyLibraryRoleList = append(libPerms.ModifyLibraryRoleList, &roles.Role{Id: role.([]string)[1], Name: role.([]string)[0]})
			}
			for _, role := range perms["manage_library_role_list"].([]interface{}) {
				libPerms.ManageLibraryRoleList = append(libPerms.ManageLibraryRoleList, &roles.Role{Id: role.([]string)[1], Name: role.([]string)[0]})
			}
			for _, role := range perms["add_library_item_role_list"].([]interface{}) {
				libPerms.AddLibraryItemRoleList = append(libPerms.AddLibraryItemRoleList, &roles.Role{Id: role.([]string)[1], Name: role.([]string)[0]})
			}
			return libPerms, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// Set permissions of the given library to the given role ids.
func (l *Library) SetPermissions(ctx context.Context, permissions *LibraryPermissions) error {
	payload := struct {
		Action    string   `json:"action,omitempty"`
		AccessIds []string `json:"access_ids"`
		AddIds    []string `json:"add_ids"`
		ManageIds []string `json:"manage_ids"`
		ModifyIds []string `json:"modify_ids"`
	}{
		Action: "", // Omitted to trigger set_permissions_old
	}
	for _, perm := range permissions.AccessLibraryRoleList {
		payload.AccessIds = append(payload.AccessIds, perm.Id)
	}
	for _, perm := range permissions.AddLibraryItemRoleList {
		payload.AddIds = append(payload.AddIds, perm.Id)
	}
	for _, perm := range permissions.ManageLibraryRoleList {
		payload.ManageIds = append(payload.ManageIds, perm.Id)
	}
	for _, perm := range permissions.ModifyLibraryRoleList {
		payload.ModifyIds = append(payload.ModifyIds, perm.Id)
	}
	// POST /api/libraries/{encoded_library_id}/permissions
	if res, err := l.galaxyInstance.R(ctx).SetBody(payload).Post(path.Join(l.GetBasePath(), l.GetID(), "permissions")); err == nil {
		_, err := blend4go.HandleResponse(res)
		return err
	} else {
		return err
	}
}

// Return a list of library files and folders.
func (l *Library) Contents(ctx context.Context) ([]*LibraryItem, error) {
	var items []*LibraryItem
	// GET /api/libraries/{library_id}/contents
	_, err := l.galaxyInstance.List(ctx, path.Join(l.GetBasePath(), l.GetID(), "contents"), &items, nil)
	for _, item := range items {
		item.libraryId = l.GetID()
	}
	return items, err
}

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
