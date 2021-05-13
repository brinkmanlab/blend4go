package quotas

import (
	"context"
	"errors"
	"path"
	"strconv"

	"github.com/brinkmanlab/blend4go"
)

type defaultQuotaAssociation string
type quotaOperation string

const (
	UnregisteredUsers defaultQuotaAssociation = "unregistered"
	RegisteredUsers   defaultQuotaAssociation = "registered"
	NotDefault        defaultQuotaAssociation = "no"
	IncreaseBy        quotaOperation          = "+"
	DecreaseBy        quotaOperation          = "-"
	SetTo             quotaOperation          = "="
	Unlimited                                 = "unlimited"
)

type Quota struct {
	galaxyInstance *blend4go.GalaxyInstance
	Id             blend4go.GalaxyID       `json:"id,omitempty"`
	Name           string                  `json:"name,omitempty"`
	Bytes          uint64                  `json:"bytes,omitempty"`
	Operation      quotaOperation          `json:"operation,omitempty"`
	Default        defaultQuotaAssociation `json:"-"`
	Description    string                  `json:"description,omitempty"`
	DisplayAmount  string                  `json:"display_amount,omitempty"`
	Users          []blend4go.GalaxyID     `json:"users,omitempty"`
	Groups         []blend4go.GalaxyID     `json:"groups,omitempty"`
	Deleted        bool                    `json:"-"`
	RawDefaults    []map[string]string     `json:"default,omitempty"`
}

// NewQuota creates a new quota for groups or users
// operation - may be forced to SetTo depending on other parameters
func NewQuota(ctx context.Context, g *blend4go.GalaxyInstance, name, amount, description string, operation quotaOperation, users, groups []string, default_for defaultQuotaAssociation) (*Quota, error) {
	if name == "" || description == "" {
		return nil, errors.New("name and description required")
	}
	if amount == "" {
		return nil, errors.New("invalid amount")
	}
	if default_for != NotDefault || amount == "unlimited" || amount == "none" || amount == "no limit" {
		operation = SetTo
	}
	//POST /api/quotas
	if res, err := g.R(ctx).SetResult(&Quota{galaxyInstance: g, Default: NotDefault}).SetBody(map[string]interface{}{
		"name":        name,
		"amount":      amount,
		"operation":   operation,
		"description": description,
		"in_users":    users,
		"in_groups":   groups,
		"default":     default_for,
	}).Post(BasePath); err == nil {
		if result, err := blend4go.HandleResponse(res); err == nil {
			quota := result.(*Quota)
			quota.populateDefault()
			return quota, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (q *Quota) populateDefault() {
	q.Default = NotDefault
	for _, d := range q.RawDefaults {
		q.Default = defaultQuotaAssociation(d["type"])
	}
}

func (q *Quota) GetBasePath() string {
	if q.Deleted {
		return path.Join(BasePath, "deleted")
	}
	return BasePath
}

func (q *Quota) SetGalaxyInstance(instance *blend4go.GalaxyInstance) {
	q.galaxyInstance = instance
}

func (q *Quota) GetID() blend4go.GalaxyID {
	return q.Id
}

func (q *Quota) SetID(id blend4go.GalaxyID) {
	q.Id = id
}

// Update changes, sending to server
// amount - optional amount same as NewQuota, Quota.Bytes used if left empty
func (q *Quota) Update(ctx context.Context, amount string) error {
	// PUT /api/quotas/{encoded_quota_id}
	if amount == "" {
		amount = strconv.FormatUint(q.Bytes, 10)
	}
	_, err := q.galaxyInstance.R(ctx).SetBody(map[string]interface{}{
		"name":        q.Name,
		"amount":      amount,
		"operation":   q.Operation,
		"description": q.Description,
		"in_users":    q.Users,
		"in_groups":   q.Groups,
		"default":     q.Default,
	}).Put(path.Join(q.GetBasePath(), q.Id))
	if err == nil {
		_, err = get(ctx, q.galaxyInstance, q)
	}
	return err
}

// Delete quota
func (q *Quota) Delete(ctx context.Context, purge bool) error {
	// DELETE /api/quotas/{encoded_quota_id}
	if q.Default != NotDefault {
		// unset default first
		q.Default = NotDefault
		if err := q.Update(ctx, ""); err != nil {
			return err
		}
	}
	params := map[string]string{}
	if purge && false { // TODO https://github.com/galaxyproject/galaxy/issues/11975
		params["purge"] = "true"
		// Must delete before purge request
		if err := q.galaxyInstance.Delete(ctx, q, nil); err != nil {
			return err
		}
		q.Deleted = true
	}
	err := q.galaxyInstance.Delete(ctx, q, &params)
	q.Deleted = true
	return err
}

// Undelete quota
func (q *Quota) Undelete(ctx context.Context) error {
	// POST /api/quotas/deleted/{encoded_quota_id}/undelete
	// TODO https://github.com/galaxyproject/galaxy/issues/11971
	if res, err := q.galaxyInstance.R(ctx).Post(path.Join(q.GetBasePath(), q.Id, "undelete")); err == nil {
		if _, err := blend4go.HandleResponse(res); err == nil {
			q.Deleted = false
			_, err = get(ctx, q.galaxyInstance, q)
			return err
		} else {
			return err
		}
	} else {
		return err
	}
}
