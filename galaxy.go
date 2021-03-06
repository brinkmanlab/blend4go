// blend4go Galaxy API client library for Golang
//
// For most use cases, NewGalaxyInstance and GetAPIKey are the only functions in the root package that should be used.
// Subpackages provide their own implementations of the remaining functions and those should be preferred.
// Subpackages are organised by subject. See workflows, users, tools, roles, repositories, libraries, jobs, histories, groups, or datatypes subpackages for more information.
//
// A number of subpackages are currently unimplemented, they will be implemented with interest in this project or as needed
package blend4go

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"log"
	"path"
	"reflect"
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

type LogLevel int

const (
	NONE LogLevel = iota
	ERROR
	WARN
	INFO
	DEBUG
)

type GalaxyInstance struct {
	Client   *resty.Client
	logErr   *log.Logger
	logInfo  *log.Logger
	logDebug *log.Logger
	logLevel LogLevel
}

type GalaxyModel interface {
	GetBasePath() string
	SetGalaxyInstance(*GalaxyInstance)
	GetID() GalaxyID
	SetID(GalaxyID)
}

// returns an API key for authenticated user based on BaseAuth headers
// Username is the users email address until https://github.com/galaxyproject/galaxy/pull/10521
func GetAPIKey(ctx context.Context, host, username, password string) (string, error) {
	r := resty.New()
	r.SetHostURL(host)
	log.Printf("[DEBUG] trying to get api key %v", host)
	r.SetHeader("Accept", "application/json")
	r.SetBasicAuth(username, password)
	if res, err := r.R().SetError(&ErrorResponse{}).SetContext(ctx).SetResult(map[string]interface{}{}).Get("/api/authenticate/baseauth"); err == nil {
		if result, err := HandleResponse(res); err == nil {
			return (*result.(*map[string]interface{}))["api_key"].(string), nil
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
func NewGalaxyInstanceLogger(host, apiKey string, logWriter io.Writer, logLevel LogLevel) (g *GalaxyInstance) {
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
	return &GalaxyInstance{
		Client:   r,
		logErr:   log.New(logWriter, "[Error]", log.Ldate|log.Ltime|log.Lshortfile),
		logInfo:  log.New(logWriter, "[Info]", log.Ldate|log.Ltime|log.Lshortfile),
		logDebug: log.New(logWriter, "[Debug]", log.Ldate|log.Ltime|log.Lshortfile),
		logLevel: logLevel,
	}
}

func NewGalaxyInstance(host, apiKey string) (g *GalaxyInstance) {
	return NewGalaxyInstanceLogger(host, apiKey, log.Writer(), ERROR)
}

func (g *GalaxyInstance) Error(v ...interface{}) {
	if g.logLevel >= ERROR {
		g.logErr.Print(v...)
	}
}

func (g *GalaxyInstance) Errorf(format string, v ...interface{}) {
	if g.logLevel >= ERROR {
		g.logErr.Printf(format, v...)
	}
}

func (g *GalaxyInstance) Info(v ...interface{}) {
	if g.logLevel >= INFO {
		g.logErr.Print(v...)
	}
}

func (g *GalaxyInstance) Infof(format string, v ...interface{}) {
	if g.logLevel >= INFO {
		g.logErr.Printf(format, v...)
	}
}

func (g *GalaxyInstance) Debug(v ...interface{}) {
	if g.logLevel >= DEBUG {
		g.logErr.Print(v...)
	}
}

func (g *GalaxyInstance) Debugf(format string, v ...interface{}) {
	if g.logLevel >= DEBUG {
		g.logErr.Printf(format, v...)
	}
}

func (g *GalaxyInstance) HandleResponse(response *resty.Response) (interface{}, error) {
	g.Infof("%v %v", response.Request.Method, response.Request.URL)
	g.Debugf("Request: %+v\nResponse: %+v", response.Request.RawRequest, response.RawResponse)
	g.Infof("%v %v", response.StatusCode(), response.Status())
	res, err := HandleResponse(response)
	if err != nil {
		g.Debug(err)
	}
	return res, err
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
			v := reflect.Indirect(reflect.ValueOf(results))
			if v.Kind() == reflect.Slice {
				for i := 0; i < v.Len(); i++ {
					if m, ok := v.Index(i).Interface().(GalaxyModel); ok {
						m.SetGalaxyInstance(g)
					} else {
						return nil, errors.New("models param element does not implement GalaxyModel")
					}
				}
				return results, nil
			} else {
				return nil, errors.New("models param was not of type slice")
			}
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

// Helper to make generic PATCH requests to Galaxy API to update objects
// params is a map of query parameters to add to the request
func (g *GalaxyInstance) Patch(ctx context.Context, model GalaxyModel, params *map[string]string) (GalaxyModel, error) {
	r := g.R(ctx)
	if params != nil {
		r.SetQueryParams(*params)
	}
	if res, err := r.SetResult(model).SetBody(model).Patch(path.Join(model.GetBasePath(), model.GetID())); err == nil {
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

// Return a description of the major version of Galaxy (e.g. 15.03).
func (g *GalaxyInstance) Version(ctx context.Context) (string, error) {
	//GET /api/version
	if res, err := g.R(ctx).SetResult(map[string]interface{}{}).Get("/api/version"); err == nil {
		if result, err := HandleResponse(res); err == nil {
			return (*result.(*map[string]interface{}))["version_major"].(string), nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

//PUT /api/configuration/toolbox Reload the Galaxy toolbox (but not individual tools).
