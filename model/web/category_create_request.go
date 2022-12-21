package web

type CategoryCreateRequest struct {
	Name string `validator:"required"`
}
