package config

import (
	"regexp"
	"net/http"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	routes []*route
}

func (hdl *RegexpHandler) Handler(
	pattern *regexp.Regexp,
	handler http.Handler) {
	hdl.routes = append(hdl.routes, &route{pattern: pattern, handler: handler})
}

func (hdl *RegexpHandler) HandleFunc(
	strRegexp string,
	handler func(http.ResponseWriter, *http.Request)) {
	pattern, _ := regexp.Compile(strRegexp)
	hdl.routes = append(hdl.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (hdl *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range hdl.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}
