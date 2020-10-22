package users

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"path"
	"regexp"
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
	if u.Deleted {
		return path.Join(BasePath, "deleted")
	}
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
	if res, err := g.R(ctx).SetResult(&User{galaxyInstance: g}).SetBody(map[string]string{
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

// returns an API key for authenticated user
// password is optional
// if password == "" the user associated with the Galaxy connection must be an admin
func (u *User) GetAPIKey(ctx context.Context, password string) (string, error) {
	if password == "" {
		// GET /api/users/{id}/api_key/inputs
		if res, err := u.galaxyInstance.R(ctx).Get(path.Join(u.GetBasePath(), u.Id, "api_key", "inputs")); err == nil {
			if _, err := blend4go.HandleResponse(res); err == nil {
				if re, err := regexp.Compile(`"value": "([^"]+)"`); err == nil {
					if match := re.FindStringSubmatch(string(res.Body())); match != nil {
						key := match[1]
						if key == "Not available." {
							if res, err := u.galaxyInstance.R(ctx).SetResult("").Post(path.Join(u.GetBasePath(), u.Id, "api_key")); err == nil {
								if result, err := blend4go.HandleResponse(res); err == nil {
									return *(result.(*string)), nil
								} else {
									return "", err
								}
							} else {
								return "", err
							}
						}
						return key, nil
					} else {
						return "", err
					}
				} else {
					return "", err
				}
			} else {
				return "", err
			}
		} else {
			return "", err
		}
	}
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

// Update the user with the state of the model
func (u *User) Update(ctx context.Context) error {
	// PUT /api/users/{id}
	_, err := u.galaxyInstance.Put(ctx, u, nil)
	return err
}

// Set the users password
func (u *User) SetPassword(ctx context.Context, current, new string) error {
	if res, err := u.galaxyInstance.R(ctx).SetBody(map[string]string{
		"id":       u.Id,
		"password": new,
		"confirm":  new,
		"current":  current,
	}).Put(path.Join(u.GetBasePath(), u.Id, "password", "inputs")); err == nil {
		if _, err := blend4go.HandleResponse(res); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

// Delete the user with the given id
func (u *User) Delete(ctx context.Context, purge bool) error {
	// DELETE /api/users/{id}
	params := map[string]string{}
	if purge {
		params["purge"] = "true"
		// Must delete before purge request
		if err := u.galaxyInstance.Delete(ctx, u, nil); err != nil {
			return err
		}
	}
	return u.galaxyInstance.Delete(ctx, u, &params)
}

// Undelete the user
func (u *User) Undelete(ctx context.Context) error {
	// POST /api/users/deleted/{id}/undelete
	if res, err := u.galaxyInstance.R(ctx).SetResult(u).Post(path.Join(u.GetBasePath(), u.Id, "undelete")); err == nil {
		if _, err := blend4go.HandleResponse(res); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

// GET /api/users/{id}/information/inputs Return user details such as username, email, addresses etc.

// PUT /api/users/{id}/information/inputs Save a user’s email, username, addresses etc.

// Add the object to user’s favorites PUT /api/users/{id}/favorites/{object_type}

// Remove the object from user’s favorites DELETE /api/users/{id}/favorites/{object_type}/{object_id:.*?}

// GET /api/users/{id}/custom_builds Returns collection of custom builds.

// PUT /api/users/{id}/custom_builds/{key} Add new custom build.

// DELETE /api/users/{id}/custom_builds/{key} Delete a custom build.
