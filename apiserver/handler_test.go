package apiserver

import "net/http"

func ExampleHandler() {
	type CreateUserRequest struct {
		Username string `json:"username"`
	}

	type CreateBlogRequest struct {
		Name string `json:"name"`
	}

	NewHandler().
		Handle("POST /user", func(w http.ResponseWriter, r *http.Request, data CreateUserRequest) {}).
		Handle("POST /blog", func(w http.ResponseWriter, r *http.Request, data CreateBlogRequest) {})
}
