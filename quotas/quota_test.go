package quotas

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"path"
	"reflect"
	"testing"
)

func TestNewQuota(t *testing.T) {
	t.SkipNow() // Quotas cant actually be deleted, a fresh server would need to be created for each test
	type args struct {
		ctx         context.Context
		g           *blend4go.GalaxyInstance
		name        string
		amount      string
		description string
		operation   quotaOperation
		users       []string
		groups      []string
		default_for defaultQuotaAssociation
	}
	tests := []struct {
		name    string
		args    args
		success func(args, *Quota) bool
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				ctx:         context.Background(),
				g:           galaxyInstance,
				name:        "test",
				amount:      "0",
				description: "test",
				operation:   SetTo,
				users:       nil,
				groups:      nil,
				default_for: NotDefault,
			},
			success: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQuota(tt.args.ctx, tt.args.g, tt.args.name, tt.args.amount, tt.args.description, tt.args.operation, tt.args.users, tt.args.groups, tt.args.default_for)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQuota() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Cleanup(func() {
				err = got.Delete(tt.args.ctx, true)
				if err != nil {
					t.Errorf("Failed to clean up created quota: %v", got)
				}
			})
			if !tt.success(tt.args, got) {
				t.Errorf("NewQuota() got = %v, want %v", got, tt.args)
			}
		})
	}
}

func TestQuota_Delete(t *testing.T) {
	test_quota, err := createTestQuota()
	if test_quota == nil {
		t.Fatalf("test quota could not be set up: %v", err)
	}
	defer test_quota.Delete(context.Background(), true)

	t.Run("basic", func(t *testing.T) {
		if err := test_quota.Delete(context.Background(), false); (err != nil) != false {
			t.Errorf("Delete() error = %v, wantErr %v", err, false)
		}
	})
}

func TestQuota_GetBasePath(t *testing.T) {
	type fields struct {
		galaxyInstance *blend4go.GalaxyInstance
		Id             blend4go.GalaxyID
		Name           string
		Bytes          uint64
		Operation      quotaOperation
		Default        defaultQuotaAssociation
		Description    string
		DisplayAmount  string
		Users          []blend4go.GalaxyID
		Groups         []blend4go.GalaxyID
		Deleted        bool
		RawDefaults    []map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "basic",
			fields: fields{Deleted: false},
			want:   BasePath,
		}, {
			name:   "deleted",
			fields: fields{Deleted: true},
			want:   path.Join(BasePath, "deleted"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quota{
				galaxyInstance: tt.fields.galaxyInstance,
				Id:             tt.fields.Id,
				Name:           tt.fields.Name,
				Bytes:          tt.fields.Bytes,
				Operation:      tt.fields.Operation,
				Default:        tt.fields.Default,
				Description:    tt.fields.Description,
				DisplayAmount:  tt.fields.DisplayAmount,
				Users:          tt.fields.Users,
				Groups:         tt.fields.Groups,
				Deleted:        tt.fields.Deleted,
				RawDefaults:    tt.fields.RawDefaults,
			}
			if got := q.GetBasePath(); got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuota_GetID(t *testing.T) {
	type fields struct {
		galaxyInstance *blend4go.GalaxyInstance
		Id             blend4go.GalaxyID
		Name           string
		Bytes          uint64
		Operation      quotaOperation
		Default        defaultQuotaAssociation
		Description    string
		DisplayAmount  string
		Users          []blend4go.GalaxyID
		Groups         []blend4go.GalaxyID
		Deleted        bool
		RawDefaults    []map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   blend4go.GalaxyID
	}{
		{
			name:   "basic",
			fields: fields{Id: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quota{
				galaxyInstance: tt.fields.galaxyInstance,
				Id:             tt.fields.Id,
				Name:           tt.fields.Name,
				Bytes:          tt.fields.Bytes,
				Operation:      tt.fields.Operation,
				Default:        tt.fields.Default,
				Description:    tt.fields.Description,
				DisplayAmount:  tt.fields.DisplayAmount,
				Users:          tt.fields.Users,
				Groups:         tt.fields.Groups,
				Deleted:        tt.fields.Deleted,
				RawDefaults:    tt.fields.RawDefaults,
			}
			if got := q.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuota_SetGalaxyInstance(t *testing.T) {
	type fields struct {
		galaxyInstance *blend4go.GalaxyInstance
		Id             blend4go.GalaxyID
		Name           string
		Bytes          uint64
		Operation      quotaOperation
		Default        defaultQuotaAssociation
		Description    string
		DisplayAmount  string
		Users          []blend4go.GalaxyID
		Groups         []blend4go.GalaxyID
		Deleted        bool
		RawDefaults    []map[string]string
	}
	type args struct {
		instance *blend4go.GalaxyInstance
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *blend4go.GalaxyInstance
	}{
		{
			name:   "basic",
			fields: fields{},
			args: args{
				instance: galaxyInstance,
			},
			want: galaxyInstance,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quota{
				galaxyInstance: tt.fields.galaxyInstance,
				Id:             tt.fields.Id,
				Name:           tt.fields.Name,
				Bytes:          tt.fields.Bytes,
				Operation:      tt.fields.Operation,
				Default:        tt.fields.Default,
				Description:    tt.fields.Description,
				DisplayAmount:  tt.fields.DisplayAmount,
				Users:          tt.fields.Users,
				Groups:         tt.fields.Groups,
				Deleted:        tt.fields.Deleted,
				RawDefaults:    tt.fields.RawDefaults,
			}
			q.SetGalaxyInstance(tt.args.instance)
			if q.galaxyInstance != tt.want {
				t.Errorf("GetID() = %v, want %v", q.galaxyInstance, tt.want)
			}
		})
	}
}

func TestQuota_SetID(t *testing.T) {
	type fields struct {
		galaxyInstance *blend4go.GalaxyInstance
		Id             blend4go.GalaxyID
		Name           string
		Bytes          uint64
		Operation      quotaOperation
		Default        defaultQuotaAssociation
		Description    string
		DisplayAmount  string
		Users          []blend4go.GalaxyID
		Groups         []blend4go.GalaxyID
		Deleted        bool
		RawDefaults    []map[string]string
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
			name:   "basic",
			fields: fields{},
			args:   args{id: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quota{
				galaxyInstance: tt.fields.galaxyInstance,
				Id:             tt.fields.Id,
				Name:           tt.fields.Name,
				Bytes:          tt.fields.Bytes,
				Operation:      tt.fields.Operation,
				Default:        tt.fields.Default,
				Description:    tt.fields.Description,
				DisplayAmount:  tt.fields.DisplayAmount,
				Users:          tt.fields.Users,
				Groups:         tt.fields.Groups,
				Deleted:        tt.fields.Deleted,
				RawDefaults:    tt.fields.RawDefaults,
			}
			q.SetID(tt.args.id)
			if q.Id != tt.want {
				t.Errorf("GetID() = %v, want %v", q.Id, tt.want)
			}
		})
	}
}

func TestQuota_Undelete(t *testing.T) {
	test_quota, err := createTestQuota()
	if test_quota == nil {
		t.Fatalf("test quota could not be set up: %v", err)
	}
	defer test_quota.Delete(context.Background(), true)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "basic",
			args:    args{ctx: context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test_quota.Delete(tt.args.ctx, false)
			if err := test_quota.Undelete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Undelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQuota_Update(t *testing.T) {
	test_quota, err := createTestQuota()
	if test_quota == nil {
		t.Fatalf("test quota could not be set up: %v", err)
	}
	defer test_quota.Delete(context.Background(), true)
	type fields struct {
		Bytes       uint64
		Operation   quotaOperation
		Default     defaultQuotaAssociation
		Description string
		Users       []blend4go.GalaxyID
		Groups      []blend4go.GalaxyID
	}
	type args struct {
		ctx    context.Context
		amount string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "basic",
			fields: fields{
				Bytes:       100,
				Operation:   IncreaseBy,
				Default:     NotDefault,
				Description: "test_basic",
				Users:       []string{},
				Groups:      []string{},
			},
			args: args{
				ctx:    context.Background(),
				amount: "1G",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quota{
				galaxyInstance: test_quota.galaxyInstance,
				Id:             test_quota.Id,
				Name:           test_quota.Name,
				Bytes:          tt.fields.Bytes,
				Operation:      tt.fields.Operation,
				Default:        tt.fields.Default,
				Description:    tt.fields.Description,
				DisplayAmount:  test_quota.DisplayAmount,
				Users:          tt.fields.Users,
				Groups:         tt.fields.Groups,
				Deleted:        test_quota.Deleted,
				RawDefaults:    test_quota.RawDefaults,
			}
			test_quota.Bytes = tt.fields.Bytes
			test_quota.Operation = tt.fields.Operation
			test_quota.Default = tt.fields.Default
			test_quota.Description = tt.fields.Description
			test_quota.Users = tt.fields.Users
			test_quota.Groups = tt.fields.Groups
			if err := test_quota.Update(tt.args.ctx, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.amount != "" {
				q.Bytes = test_quota.Bytes
			}
			q.DisplayAmount = test_quota.DisplayAmount
			if !reflect.DeepEqual(test_quota, q) {
				t.Errorf("Update()\ngot = %#v\nwant= %#v", test_quota, q)
			}
		})
	}
}

func TestQuota_populateDefault(t *testing.T) {
	type fields struct {
		Default     defaultQuotaAssociation
		RawDefaults []map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   defaultQuotaAssociation
	}{
		{
			name: "empty",
			fields: fields{
				Default:     "",
				RawDefaults: []map[string]string{{"type": string(NotDefault)}},
			},
			want: NotDefault,
		}, {
			name: "overwrite",
			fields: fields{
				Default:     RegisteredUsers,
				RawDefaults: []map[string]string{{"type": string(NotDefault)}},
			},
			want: NotDefault,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quota{
				Default:     tt.fields.Default,
				RawDefaults: tt.fields.RawDefaults,
			}
			q.populateDefault()
			if q.Default != tt.want {
				t.Errorf("populateDefault() got = %v, want %v", q.Default, tt.want)
			}
		})
	}
}

func Test_get(t *testing.T) {
	test_quota, err := createTestQuota()
	if test_quota == nil {
		t.Fatalf("test quota could not be set up: %v", err)
	}
	defer test_quota.Delete(context.Background(), true)
	type args struct {
		ctx   context.Context
		g     *blend4go.GalaxyInstance
		model *Quota
	}
	tests := []struct {
		name    string
		args    args
		want    *Quota
		wantErr bool
	}{
		{
			name: "unpopulated",
			args: args{
				ctx:   context.Background(),
				g:     galaxyInstance,
				model: &Quota{Id: test_quota.Id, Deleted: test_quota.Deleted},
			},
			want:    test_quota,
			wantErr: false,
		}, {
			name: "populated",
			args: args{
				ctx:   context.Background(),
				g:     galaxyInstance,
				model: test_quota,
			},
			want:    test_quota,
			wantErr: false,
		}, {
			name: "not_exist",
			args: args{
				ctx:   context.Background(),
				g:     galaxyInstance,
				model: &Quota{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := get(tt.args.ctx, tt.args.g, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() got = %+v, want %+v", got, tt.want)
			}
		})
	}
}
