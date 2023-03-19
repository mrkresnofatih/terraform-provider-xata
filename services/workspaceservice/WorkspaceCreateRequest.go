package workspaceservice

type WorkspaceCreateRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
