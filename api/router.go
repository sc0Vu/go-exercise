package main

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() Router {
	return Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) notFound() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		res := []byte("not found")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
	}
}

func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		panic("empty handler")
	}
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			notFound := r.notFound()
			notFound(w, req)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		panic("empty handler")
	}
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			notFound := r.notFound()
			notFound(w, req)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Put(path string, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		panic("empty handler")
	}
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "PUT" {
			notFound := r.notFound()
			notFound(w, req)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Delete(path string, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		panic("empty handler")
	}
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "DELETE" {
			notFound := r.notFound()
			notFound(w, req)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Option(path string, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		panic("empty handler")
	}
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "OPTION" {
			notFound := r.notFound()
			notFound(w, req)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Patch(path string, handler func(http.ResponseWriter, *http.Request)) {
	if handler == nil {
		panic("empty handler")
	}
	r.mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "PATCH" {
			notFound := r.notFound()
			notFound(w, req)
			return
		}
		handler(w, req)
	})
}

func (r *Router) Handler() *http.ServeMux {
	return r.mux
}
