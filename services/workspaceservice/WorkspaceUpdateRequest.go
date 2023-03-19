package workspaceservice

type WorkspaceUpdateRequest struct {
	Id      string
	Request WorkspaceUpdateApiRequest
}

type WorkspaceUpdateApiRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
