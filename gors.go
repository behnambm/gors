package gors

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Config is the plugin configuration.
type Config struct {
	AllowedOrigins  []string
	AllowedHeaders  []string
	AllowedMethods  []string
	PreflightMaxAge int
	// Disabled is used to disable all functionalities of plugin, when it's disabled it just passes the request through
	Disabled bool
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		AllowedOrigins:  []string{},
		AllowedHeaders:  []string{},
		AllowedMethods:  []string{http.MethodOptions},
		PreflightMaxAge: 86400, // in seconds
		Disabled:        false,
	}
}

// Gors is the plugin struct
type Gors struct {
	next   http.Handler
	name   string
	Config *Config
}

// New creates a new plugin when Traefik gets started.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Gors{
		next:   next,
		name:   name,
		Config: config,
	}, nil
}

func (g *Gors) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if g.Config.Disabled {
		g.next.ServeHTTP(rw, req)

		return
	}
	fmt.Println("here ", req.Method)
	if contains(g.Config.AllowedOrigins, "*") {
		rw.Header().Add("Access-Control-Allow-Origin", "*")
	} else {
		rw.Header().Add("Vary", "Origin")
		rw.Header().Add("Access-Control-Allow-Origin", strings.Join(g.Config.AllowedOrigins, ", "))
	}

	rw.Header().Add("Access-Control-Allow-Methods", strings.Join(g.Config.AllowedMethods, ", "))
	rw.Header().Add("Access-Control-Allow-Headers", strings.Join(g.Config.AllowedHeaders, ", "))
	rw.Header().Add("Access-Control-Max-Age", strconv.Itoa(g.Config.PreflightMaxAge))

	if req.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusNoContent)

		return
	}

	g.next.ServeHTTP(rw, req)
}

// Utility functions

func contains(list []string, elem string) bool {
	for _, value := range list {
		if value == elem {
			return true
		}
	}
	return false
}

//f = () => {console.log("sf")}; const xhr = new XMLHttpRequest(); xhr.open("POST", "http://localhost:3030"); xhr.setRequestHeader("X-PINGOTHER", "pingpong"); xhr.setRequestHeader("Content-Type", "text/xml"); xhr.onreadystatechange = f; xhr.send("<person><name>Arun</name></person>");
