package workspaceservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mrkresnofatih/terraform-provider-xata/services/base"
	"github.com/mrkresnofatih/terraform-provider-xata/utilty"
	"io"
	"net/http"
	"time"
)

type IWorkspaceService interface {
	Create(request WorkspaceCreateRequest) (WorkspaceCreateResponse, error)
	Get(request WorkspaceGetRequest) (WorkspaceGetResponse, error)
	Update(request WorkspaceUpdateRequest) (WorkspaceUpdateResponse, error)
	Delete(request WorkspaceDeleteRequest) (WorkspaceDeleteResponse, error)
}

type WorkspaceService struct {
	XataApi base.XataApi
}

func (w WorkspaceService) Create(request WorkspaceCreateRequest) (WorkspaceCreateResponse, error) {
	utilty.LogInfo("Start Create-Workspace w. data: %s", request)
	serial, err := json.Marshal(request)
	if err != nil {
		utilty.LogError("failed to json-marshal %s", request)
		return utilty.ErrorReturner[WorkspaceCreateResponse]("failed to create workspace")
	}
	bodyReader := bytes.NewReader(serial)
	url := "https://api.xata.io/workspaces"
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		utilty.LogError("failed to build http request %s", request)
		return utilty.ErrorReturner[WorkspaceCreateResponse]("failed to create workspace")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", w.XataApi.ApiKey))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: time.Second * 30,
	}
	res, err := client.Do(req)
	if err != nil {
		utilty.LogError("failed http request failed")
		return utilty.ErrorReturner[WorkspaceCreateResponse]("http request failed")
	}
	defer res.Body.Close()
	utilty.LogInfo("status http request returned %s", res.Status)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		utilty.LogInfo("read all body failed")
		return utilty.ErrorReturner[WorkspaceCreateResponse]("read body failed")
	}
	utilty.LogInfo("response-body: %s", string(body))
	result := WorkspaceCreateResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		utilty.LogError("failed to unmarshal body %s", string(body))
		return utilty.ErrorReturner[WorkspaceCreateResponse]("failed to unmarshal")
	}
	utilty.LogInfo("success create workspace of req: %s", request)
	return result, nil
}

func (w WorkspaceService) Get(request WorkspaceGetRequest) (WorkspaceGetResponse, error) {
	utilty.LogInfo("Start Get-Workspace w. data: %s", request)
	url := fmt.Sprintf("https://api.xata.io/workspaces/%s", request.Id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		utilty.LogError("failed to build http request %s", url)
		return utilty.ErrorReturner[WorkspaceGetResponse]("failed to build http request")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", w.XataApi.ApiKey))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: time.Second * 30,
	}
	res, err := client.Do(req)
	if err != nil {
		utilty.LogError("http request failed %s", url)
		return utilty.ErrorReturner[WorkspaceGetResponse]("http req failed")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		utilty.LogInfo("read all body failed")
		return utilty.ErrorReturner[WorkspaceGetResponse]("read body failed")
	}
	utilty.LogInfo("response-body: %s", string(body))
	result := WorkspaceGetResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		utilty.LogError("failed to unmarshal body %s", string(body))
		return utilty.ErrorReturner[WorkspaceGetResponse]("failed to unmarshal")
	}
	utilty.LogInfo("success get workspace of req: %s", request)
	return result, nil
}

func (w WorkspaceService) Update(request WorkspaceUpdateRequest) (WorkspaceUpdateResponse, error) {
	utilty.LogInfo("start update-workspace w. data: %s", request)
	url := fmt.Sprintf("https://api.xata.io/workspaces/%s", request.Id)
	serial, err := json.Marshal(request.Request)
	if err != nil {
		utilty.LogError("failed to json-marshal %s", request)
		return utilty.ErrorReturner[WorkspaceUpdateResponse]("failed to update workspace")
	}
	bodyReader := bytes.NewReader(serial)
	req, err := http.NewRequest(http.MethodPut, url, bodyReader)
	if err != nil {
		utilty.LogError("failed to build http request %s", url)
		return utilty.ErrorReturner[WorkspaceUpdateResponse]("failed to build http request")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", w.XataApi.ApiKey))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: time.Second * 30,
	}
	res, err := client.Do(req)
	if err != nil {
		utilty.LogError("http request failed %s", url)
		return utilty.ErrorReturner[WorkspaceUpdateResponse]("http req failed")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		utilty.LogInfo("read all body failed")
		return utilty.ErrorReturner[WorkspaceUpdateResponse]("read body failed")
	}
	utilty.LogInfo("response-body: %s", string(body))
	result := WorkspaceUpdateResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		utilty.LogError("failed to unmarshal body %s", string(body))
		return utilty.ErrorReturner[WorkspaceUpdateResponse]("failed to unmarshal")
	}
	utilty.LogInfo("success get workspace of req: %s", request)
	return result, nil
}

func (w WorkspaceService) Delete(request WorkspaceDeleteRequest) (WorkspaceDeleteResponse, error) {
	utilty.LogInfo("start delete-workspace w. data: %s", request)
	url := fmt.Sprintf("https://api.xata.io/workspaces/%s", request.Id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		utilty.LogError("failed to build http request %s", url)
		return utilty.ErrorReturner[WorkspaceDeleteResponse]("failed to build http request")
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", w.XataApi.ApiKey))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: time.Second * 30,
	}
	_, err = client.Do(req)
	if err != nil {
		utilty.LogError("failed to execute api request for workspace")
		return utilty.ErrorReturner[WorkspaceDeleteResponse]("failed to execute api request workspace")
	}
	utilty.LogInfo("success delete workspace")
	return WorkspaceDeleteResponse{
		Id: request.Id,
	}, nil
}
