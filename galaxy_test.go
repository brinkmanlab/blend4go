package blend4go_test

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/test_util"
	"github.com/go-resty/resty/v2"

	//"github.com/jarcoal/httpmock"
	"reflect"
	"testing"
)

//func init() {
//	httpmock.Activate()
//}

func TestGalaxyInstance_Delete(t *testing.T) {
	type fields struct {
		Client *resty.Client
	}
	type args struct {
		ctx   context.Context
		model blend4go.GalaxyModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := test_util.NewTestInstance()
			if err := g.Delete(tt.args.ctx, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGalaxyInstance_Get(t *testing.T) {
	type args struct {
		ctx   context.Context
		id    blend4go.GalaxyID
		model blend4go.GalaxyModel
	}
	tests := []struct {
		name    string
		args    args
		want    blend4go.GalaxyModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := test_util.NewTestInstance()
			got, err := g.Get(tt.args.ctx, tt.args.id, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGalaxyInstance_List(t *testing.T) {
	type args struct {
		ctx    context.Context
		path   string
		models interface{}
		params *map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := test_util.NewTestInstance()
			got, err := g.List(tt.args.ctx, tt.args.path, tt.args.models, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGalaxyInstance_Put(t *testing.T) {
	type fields struct {
	}
	type args struct {
		ctx   context.Context
		model blend4go.GalaxyModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    blend4go.GalaxyModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := test_util.NewTestInstance()
			got, err := g.Put(tt.args.ctx, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGalaxyInstance_R(t *testing.T) {
	type fields struct {
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		success func(blend4go.GalaxyRequest) bool
	}{
		{
			name:   "Basic",
			fields: fields{},
			args: args{
				ctx: context.Background(),
			},
			success: func(request blend4go.GalaxyRequest) bool {
				return request.Context() == context.Background()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := test_util.NewTestInstance()
			if got := g.R(tt.args.ctx); !tt.success(got) {
				t.Errorf("R() = %v", got)
			}
		})
	}
}

func TestGalaxyInstance_ToolSheds(t *testing.T) {
	type fields struct {
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		success func([]*blend4go.ToolShed) bool
	}{
		{
			name:   "Basic",
			fields: fields{},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
			success: func(sheds []*blend4go.ToolShed) bool {
				return len(sheds) > 0 && sheds[0].Name != "" && sheds[0].Url != ""
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := test_util.NewTestInstance()
			got, err := g.ToolSheds(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToolSheds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.success(got) {
				t.Errorf("ToolSheds() got = %v", got)
			}
		})
	}
}

func TestGetAPIKey(t *testing.T) {
	type args struct {
		ctx      context.Context
		host     string
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				ctx:      context.Background(),
				host:     test_util.Host,
				username: test_util.User,
				password: test_util.Pass,
			},
			want:    test_util.ApiKey,
			wantErr: false,
		}, {
			name: "Empty args",
			args: args{
				ctx:      context.Background(),
				host:     test_util.Host,
				username: "",
				password: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := blend4go.GetAPIKey(tt.args.ctx, tt.args.host, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGalaxyInstance(t *testing.T) {
	type args struct {
		host   string
		apiKey string
	}
	tests := []struct {
		name    string
		args    args
		success func(*blend4go.GalaxyInstance) bool
	}{
		{
			name: "Basic",
			args: args{
				host:   test_util.Host,
				apiKey: test_util.ApiKey,
			},
			success: func(g *blend4go.GalaxyInstance) bool {
				return g != nil &&
					g.Client != nil &&
					g.Client.HostURL == test_util.Host &&
					g.Client.Header.Get("X-AUTH-KEY") == test_util.ApiKey &&
					g.Client.Header.Get("Accept") == "application/json" &&
					g.Client.Header.Get("Content-Type") == "application/json" &&
					g.Client.Header.Get("User-Agent") != ""
			},
		}, {
			name: "Empty args", // Should still return 'valid' instance
			args: args{
				host:   "",
				apiKey: "",
			},
			success: func(g *blend4go.GalaxyInstance) bool {
				return g != nil &&
					g.Client != nil &&
					g.Client.HostURL == "" &&
					g.Client.Header.Get("X-AUTH-KEY") == ""
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotG := blend4go.NewGalaxyInstance(tt.args.host, tt.args.apiKey); !tt.success(gotG) {
				t.Errorf("NewGalaxyInstance() = %v", gotG)
			}
		})
	}
}
