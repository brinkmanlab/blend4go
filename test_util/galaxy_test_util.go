package test_util

import (
	"github.com/brinkmanlab/blend4go"
	"os"
)

var Host = os.Getenv("GALAXY_HOST")
var ApiKey = os.Getenv("GALAXY_API_KEY")
var User = os.Getenv("GALAXY_USER")
var Pass = os.Getenv("GALAXY_PASSWORD")

func NewTestInstance() *blend4go.GalaxyInstance {
	if Host == "" {
		panic("GALAXY_HOST environment variable not found")
	}
	if ApiKey == "" {
		panic("GALAXY_API_KEY environment variable not found")
	}
	return blend4go.NewGalaxyInstance(Host, ApiKey)
}
