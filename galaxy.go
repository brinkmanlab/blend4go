package blend4go

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"path"
	"runtime"
	"runtime/debug"
	"strings"
)

// https://blog.golang.org/publishing-go-modules
// https://github.com/go-resty/resty
// https://pkg.go.dev/github.com/go-resty/resty
type GalaxyID = string
type GalaxyRequest = *resty.Request

type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	URL      string
	Method   string
	Message1 string `json:"message"`
	Code1    string `json:"code"`
	Message  string `json:"err_msg"`
	Code     int    `json:"err_code"`
}

func (e *ErrorResponse) String() string {
	return fmt.Sprintf("%v %v %v%v: %v%v", e.Method, e.URL, e.Code, e.Code1, e.Message, e.Message1)
}

func (e *ErrorResponse) Error() string {
	return e.String()
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
	if res, err := r.R().SetError(&ErrorResponse{}).SetContext(ctx).Get("/api/authenticate/baseauth"); err == nil {
		if result, err := HandleResponse(res); err == nil {
			return result.(map[string]interface{})["api_key"].(string), nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

// Handle responses from Galaxy API, checking for errors
func HandleResponse(response *resty.Response) (interface{}, error) {
	if response.IsError() {
		err := response.Error().(*ErrorResponse)
		err.Method = response.Request.Method
		err.URL = response.Request.URL
		if err.Message == "" {
			err.Message = err.Message1
		}
		if err.Message == "" {
			return nil, fmt.Errorf("%v %v: %v", response.Request.Method, response.Request.URL, string(response.Body()))
		}
		return nil, err
	}
	if response.IsSuccess() {
		return response.Result(), nil
	}
	return nil, fmt.Errorf("unhandled response: %v", response.Status())
}

// Create a new connection handle to an instance of Galaxy
func NewGalaxyInstance(host, apiKey string) (g *GalaxyInstance) {
	// Automatically attach caller package name to agent
	pc, _, _, _ := runtime.Caller(1)
	components := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	if len(components) > 1 {
		components = components[:len(components)-1]
	}
	agent := strings.Join(components, ".") + " - blend4go"
	if info, ok := debug.ReadBuildInfo(); ok {
		agent = agent + " " + info.Main.Version + " " + info.Main.Sum
	}
	r := resty.New()
	r.SetHostURL(host)
	r.SetHeader("X-API-KEY", apiKey)
	r.SetHeader("Accept", "application/json")
	r.SetHeader("Content-Type", "application/json")
	r.SetHeader("User-Agent", agent)
	return &GalaxyInstance{Client: r}
}

// Helper to make generic requests against Galaxy API that return lists of objects
// params is a map of query parameters to add to the request
func (g *GalaxyInstance) List(ctx context.Context, path string, models interface{}, params *map[string]string) (interface{}, error) {
	r := g.R(ctx)
	if params != nil {
		r.SetQueryParams(*params)
	}
	if res, err := r.SetResult(models).Get(path); err == nil {
		if results, err := HandleResponse(res); err == nil {
			if r, ok := results.([]GalaxyModel); ok {
				for _, m := range r {
					m.SetGalaxyInstance(g)
				}
			}
			return results, err
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// Helper to make generic requests against Galaxy API that returns single object
// params is a map of query parameters to add to the request
func (g *GalaxyInstance) Get(ctx context.Context, id GalaxyID, model GalaxyModel, params *map[string]string) (GalaxyModel, error) {
	r := g.R(ctx)
	if params != nil {
		r.SetQueryParams(*params)
	}
	if res, err := r.SetResult(model).Get(path.Join(model.GetBasePath(), id)); err == nil {
		if result, err := HandleResponse(res); err == nil {
			m := result.(GalaxyModel)
			m.SetGalaxyInstance(g)
			return m, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// Helper to make generic PUT requests to Galaxy API to update objects
// params is a map of query parameters to add to the request
func (g *GalaxyInstance) Put(ctx context.Context, model GalaxyModel, params *map[string]string) (GalaxyModel, error) {
	r := g.R(ctx)
	if params != nil {
		r.SetQueryParams(*params)
	}
	if res, err := r.SetResult(model).SetBody(model).Put(path.Join(model.GetBasePath(), model.GetID())); err == nil {
		if result, err := HandleResponse(res); err == nil {
			m := result.(GalaxyModel)
			m.SetGalaxyInstance(g)
			return m, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// Helper to make generic DELETE requests to delete single objects
// params is a map of query parameters to add to the request
func (g *GalaxyInstance) Delete(ctx context.Context, model GalaxyModel, params *map[string]string) error {
	r := g.R(ctx)
	if params != nil {
		r.SetQueryParams(*params)
	}
	if res, err := r.Delete(path.Join(model.GetBasePath(), model.GetID())); err == nil {
		_, err := HandleResponse(res)
		return err
	} else {
		return err
	}
}

// Helper to create a new request to the Galaxy API
// Use this if one of the other functions in this package are not appropriate for the request
func (g *GalaxyInstance) R(ctx context.Context) GalaxyRequest {
	return g.Client.R().SetContext(ctx).SetError(&ErrorResponse{})
}

type ToolShed struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// Get a list of toolsheds configured in the galaxy instance
func (g *GalaxyInstance) ToolSheds(ctx context.Context) ([]*ToolShed, error) {
	//GET /api/tool_shed Interact with the Toolshed registry of this instance.
	if res, err := g.R(ctx).SetResult([]*ToolShed{}).Get("/api/tool_shed"); err == nil {
		return res.Result().([]*ToolShed), nil
	} else {
		return nil, err
	}
}

//GET /api/tool_shed/request

//GET /api/whoami Return information about the current authenticated user.
//GET /api/configuration Return an object containing exposable configuration settings.
//GET /api/version Return a description of the major version of Galaxy (e.g. 15.03).
//PUT /api/configuration/toolbox Reload the Galaxy toolbox (but not individual tools).
