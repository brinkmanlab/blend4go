package users

import (
	"context"
	"github.com/brinkmanlab/blend4go"
)

type User struct {
	galaxyInstance *blend4go.GalaxyInstance
	Id             blend4go.GalaxyID `json:"id,omitempty"`
	Username       string            `json:"username,omitempty"`
	QuotaPercent   uint              `json:"quota_percent,omitempty"`
	//Preferences ? `json:"preferences,omitempty"`
	TotalDiskUsage     float32  `json:"total_disk_usage,omitempty"`
	Deleted            bool     `json:"deleted,omitempty"`
	Purged             bool     `json:"purged,omitempty"`
	NiceTotalDiskUsage string   `json:"nice_total_disk_usage,omitempty"`
	Quota              string   `json:"quota,omitempty"`
	Email              string   `json:"email,omitempty"`
	IsAdmin            bool     `json:"is_admin,omitempty"`
	TagsUsed           []string `json:"tags_used,omitempty"`
}

func (u *User) GetBasePath() string {
	return BasePath
}

func (u *User) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	u.galaxyInstance = g
}

func (u *User) GetID() blend4go.GalaxyID {
	return u.Id
}

func (u *User) SetID(id blend4go.GalaxyID) {
	u.Id = id
}

// Creates a new Galaxy user.
func NewUser(ctx context.Context, g *blend4go.GalaxyInstance, username, password, email string) (*User, error) {
	//POST /api/users
	if res, err := g.R(ctx).SetResult(&User{galaxyInstance: g}).SetBody(map[string]string{ // TODO reuse User struct?
		"username": username,
		"password": password,
		"email":    email,
	}).Post(BasePath); err == nil {
		if result, err := blend4go.HandleResponse(res); err == nil {
			return result.(*User), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// returns an API key for authenticated user based on BaseAuth headers
func (u *User) GetAPIKey(ctx context.Context, password string) (string, error) {
	if res, err := u.galaxyInstance.R(ctx).SetBasicAuth(u.Username, password).Get("/api/authenticate/baseauth"); err == nil {
		if result, err := blend4go.HandleResponse(res); err == nil {
			return result.(map[string]interface{})["api_key"].(string), nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

// PUT /api/users/{id}
func (u *User) Update(ctx context.Context) error {
	_, err := u.galaxyInstance.Put(ctx, u, nil)
	return err
}

// delete the user with the given id
func (u *User) Delete(ctx context.Context) error {
	// DELETE /api/users/{id}
	return u.galaxyInstance.Delete(ctx, u, nil)
}

// POST /api/users/deleted/{id}/undelete Undelete the user with the given id

// GET /api/users/{id}/information/inputs Return user details such as username, email, addresses etc.

// PUT /api/users/{id}/information/inputs Save a user’s email, username, addresses etc.

// Add the object to user’s favorites PUT /api/users/{id}/favorites/{object_type}

// Remove the object from user’s favorites DELETE /api/users/{id}/favorites/{object_type}/{object_id:.*?}

// GET /api/users/{id}/custom_builds Returns collection of custom builds.

// PUT /api/users/{id}/custom_builds/{key} Add new custom build.

// DELETE /api/users/{id}/custom_builds/{key} Delete a custom build.
