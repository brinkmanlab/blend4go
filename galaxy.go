package blend4go

import (
	"context"
	"github.com/go-resty/resty/v2"
	"path"
	"runtime/debug"
)

// https://blog.golang.org/publishing-go-modules
// https://github.com/go-resty/resty
// https://pkg.go.dev/github.com/go-resty/resty
type GalaxyID = string
type GalaxyRequest = *resty.Request

type StatusResponse struct {
	Status  string
	Message string
}

type GalaxyInstance struct {
	Client *resty.Client
}

type GalaxyModel interface {
	GetBasePath() string
	SetGalaxyInstance(*GalaxyInstance)
	GetID() GalaxyID
	SetID(GalaxyID)
}

// returns an API key for authenticated user based on BaseAuth headers
func GetAPIKey(ctx context.Context, host, username, password string) (string, error) {
	r := resty.New()
	r.SetHostURL(host)
	r.SetHeader("Accept", "application/json")
	r.SetBasicAuth(username, password)
	if res, err := r.R().SetContext(ctx).Get("/api/authenticate/baseauth"); err == nil {
		return res.Result().(map[string]string)["api_key"], nil
	} else {
		return "", err
	}
}

func NewGalaxyInstance(host, apiKey string) (g *GalaxyInstance) {
	agent := "blend4go"
	if info, ok := debug.ReadBuildInfo(); ok {
		agent = agent + " " + info.Main.Version + " " + info.Main.Sum
	}
	r := resty.New()
	r.SetHostURL(host)
	r.SetHeader("X-AUTH-KEY", apiKey)
	r.SetHeader("Accept", "application/json")
	r.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   agent,
	})
	return &GalaxyInstance{Client: r}
}

func (g *GalaxyInstance) List(ctx context.Context, path string, models interface{}, params *map[string]string) (interface{}, error) {
	r := g.R(ctx)
	if params != nil {
		r.SetQueryParams(*params)
	}
	if res, err := r.SetResult(models).Get(path); err == nil {
		results := res.Result()
		if r, ok := results.([]GalaxyModel); ok {
			for _, m := range r {
				m.SetGalaxyInstance(g)
			}
		}
		return results, err
	} else {
		return nil, err
	}
}

func (g *GalaxyInstance) Get(ctx context.Context, id GalaxyID, model GalaxyModel) (GalaxyModel, error) {
	if res, err := g.R(ctx).SetResult(model).Get(path.Join(model.GetBasePath(), id)); err == nil {
		m := res.Result().(GalaxyModel)
		m.SetGalaxyInstance(g)
		return m, nil
	} else {
		return nil, err
	}
}

func (g *GalaxyInstance) Put(ctx context.Context, model GalaxyModel) (GalaxyModel, error) {
	if res, err := g.R(ctx).SetResult(model).SetBody(model).Put(path.Join(model.GetBasePath(), model.GetID())); err == nil {
		return res.Result().(GalaxyModel), nil
	} else {
		return nil, err
	}
}

func (g *GalaxyInstance) Delete(ctx context.Context, model GalaxyModel) error {
	if _, err := g.R(ctx).Delete(path.Join(model.GetBasePath(), model.GetID())); err == nil {
		return err // TODO handle result. Status message?
	} else {
		return err
	}
}

func (g *GalaxyInstance) R(ctx context.Context) GalaxyRequest {
	return g.Client.R().SetContext(ctx)
}

type ToolShed struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (g *GalaxyInstance) ToolSheds() []ToolShed {
	//GET /api/tool_shed Interact with the Toolshed registry of this instance.
	return nil
}

//GET /api/tool_shed/request

//GET /api/whoami Return information about the current authenticated user.
//GET /api/configuration Return an object containing exposable configuration settings.
//GET /api/version Return a description of the major version of Galaxy (e.g. 15.03).
//PUT /api/configuration/toolbox Reload the Galaxy toolbox (but not individual tools).
