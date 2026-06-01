package apiserver

import "net/http"

func ExampleHandler() {
	type CreateUserRequest struct {
		Username string `json:"username"`
	}

	type CreateBlogRequest struct {
		Name string `json:"name"`
	}

	app := NewHandler()
	app.Handle("POST /user", func(w http.ResponseWriter, r *http.Request, data CreateUserRequest) {})
	app.Handle("POST /blog", func(w http.ResponseWriter, r *http.Request, data CreateBlogRequest) {})
}
