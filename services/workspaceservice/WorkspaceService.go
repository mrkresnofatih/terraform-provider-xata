package workspaceservice

import "github.com/mrkresnofatih/terraform-provider-xata/services"

type IWorkspaceService interface {
	Create(request WorkspaceCreateRequest) (WorkspaceCreateResponse, error)
	Get(request WorkspaceGetRequest) (WorkspaceGetResponse, error)
	Update(request WorkspaceUpdateRequest) (WorkspaceUpdateResponse, error)
	Delete(request WorkspaceDeleteRequest) (WorkspaceDeleteResponse, error)
}

type WorkspaceService struct {
	XataApi services.XataApi
}

func (w WorkspaceService) Create(request WorkspaceCreateRequest) (WorkspaceCreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (w WorkspaceService) Get(request WorkspaceGetRequest) (WorkspaceGetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (w WorkspaceService) Update(request WorkspaceUpdateRequest) (WorkspaceUpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (w WorkspaceService) Delete(request WorkspaceDeleteRequest) (WorkspaceDeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}
