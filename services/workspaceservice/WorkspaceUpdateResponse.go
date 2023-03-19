package workspaceservice

type WorkspaceUpdateResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
