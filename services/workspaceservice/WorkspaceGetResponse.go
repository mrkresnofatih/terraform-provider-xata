package workspaceservice

type WorkspaceGetResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
