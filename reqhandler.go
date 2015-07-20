// Req handler is used to handle a route
// The "handle" method calls 4 different funtions based on the request type (get, post, put, delete)
// It calls a function of type func(res http.ResponseWriter, req *http.Request) for every request type
package httputils

import (
	"net/http"
	"fmt"
)

type ReqHandler struct {
    Path string
}

func (r *ReqHandler) Handle(
        get    func(res http.ResponseWriter, req *http.Request),
        post   func(res http.ResponseWriter, req *http.Request),
        put    func(res http.ResponseWriter, req *http.Request),
        delete func(res http.ResponseWriter, req *http.Request)) {
    http.HandleFunc(r.Path, func(res http.ResponseWriter, req *http.Request) {
        m := req.Method
        var status int
        status = http.StatusOK
        fmt.Printf("%s %s ", m, r.Path)
        var call func(res http.ResponseWriter, req *http.Request)
        switch m {
            case "GET": if get != nil {call = get}else{status = http.StatusNotFound}
            case "POST": if post != nil {call = post}else{status = http.StatusNotFound}
            case "PUT": if put != nil {call = put}else{status = http.StatusNotFound}
            case "DELETE": if delete != nil {call = delete}else{status = http.StatusNotFound}
        }
        fmt.Printf("(%d)\n", status)
        if status == http.StatusNotFound {
            http.Error(res, "Not found", http.StatusNotFound)
        }else{
            call(res, req)
        }
    })
}
