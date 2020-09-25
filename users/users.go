package users

import (
	"context"
	"github.com/brinkmanlab/blend4go"
)

const BasePath = "/api/users"

/*
Displays a collection (list) of users.
deleted - If true, show deleted users
filterEmail - (optional) An email address to filter on. (Non-admin users can only use this if expose_user_email is True in galaxy.ini)
filterName - (optional) A username to filter on. (Non-admin users can only use this if expose_user_name is True in galaxy.ini)
filterAny - (optional) Filter on username OR email. (Non-admin users can use this, the email filter and username filter will only be active if their corresponding expose_user_* is True in galaxy.ini)
*/
func List(ctx context.Context, g *blend4go.GalaxyInstance, deleted bool, filterEmail, filterName, filterAny string) ([]User, error) {
	q := make(map[string]string)
	if deleted {
		q["deleted"] = "true"
	}
	if filterEmail != "" {
		q["f_email"] = filterEmail
	}
	if filterName != "" {
		q["f_name"] = filterName
	}
	if filterAny != "" {
		q["f_any"] = filterAny
	}
	// GET /api/users GET /api/users/deleted
	res, err := g.List(ctx, BasePath, []User{}, &q)
	return res.([]User), err
}

// Displays information about a user.
func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID) (*User, error) {
	// GET /api/users/{encoded_id} GET /api/users/deleted/{encoded_id} GET /api/users/current
	if id == "" {
		id = "current"
	}
	res, err := g.Get(ctx, id, &User{})
	return res.(*User), err
}
