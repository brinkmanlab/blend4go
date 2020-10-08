package users_test

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/users"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		ctx      context.Context
		g        *blend4go.GalaxyInstance
		username string
		password string
		email    string
	}
	tests := []struct {
		name    string
		args    args
		success func(args, *users.User) bool
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				ctx:      context.Background(),
				g:        galaxyInstance,
				username: "test",
				password: "test",
				email:    "test@example.com",
			},
			success: func(a args, user *users.User) bool {
				return user != nil &&
					!user.Deleted &&
					!user.IsAdmin &&
					!user.Purged &&
					user.Id != "" &&
					user.Username == a.username &&
					user.Email == a.email
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := users.NewUser(tt.args.ctx, tt.args.g, tt.args.username, tt.args.password, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Cleanup(func() {
				err = got.Delete(tt.args.ctx)
				if err != nil {
					t.Errorf("Failed to clean up created user: %v", got)
				}
			})
			if !tt.success(tt.args, got) {
				t.Errorf("NewUser() got = %v, want %v", got, tt.args)
			}
		})
	}
}

func TestUser_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Basic",
			args:    args{ctx: context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := users.NewUser(tt.args.ctx, galaxyInstance, "test", "test", "test@example.com")
			if err != nil {
				t.Errorf("Failed to prepare user for deletion")
				return
			}
			if err := u.Delete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUser_GetBasePath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Basic",
			want: users.BasePath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &users.User{}
			if got := u.GetBasePath(); got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetID(t *testing.T) {
	type fields struct {
		Id blend4go.GalaxyID
	}
	tests := []struct {
		name   string
		fields fields
		want   blend4go.GalaxyID
	}{
		{
			name:   "Basic",
			fields: fields{Id: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &users.User{
				Id: tt.fields.Id,
			}
			if got := u.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SetGalaxyInstance(t *testing.T) {
	t.SkipNow()
}

func TestUser_SetID(t *testing.T) {
	type fields struct {
		Id blend4go.GalaxyID
	}
	type args struct {
		id blend4go.GalaxyID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   blend4go.GalaxyID
	}{
		{
			name:   "Basic",
			fields: fields{Id: "oldid"},
			args:   args{id: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &users.User{
				Id: tt.fields.Id,
			}
			u.SetID(tt.args.id)
			if u.Id != tt.want {
				t.Errorf("SetID() = %v, wanted %v", u.Id, tt.want)
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	t.SkipNow()
	return
	type fields struct {
		Email    string
		TagsUsed []string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		success func(*users.User) bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := users.NewUser(tt.args.ctx, galaxyInstance, "test", "test", "test@example.com")
			if err != nil {
				t.Errorf("Failed to prepare user for deletion")
				return
			}
			if err := u.Update(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
