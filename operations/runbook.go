package operations

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetRunbookRequest struct {
	ID string `path:"runbookId" doc:"The ID of the runbook"`
}

type GetRunbookResponse struct {
}

type ListRunbooksResponse struct {
	Body Runbook
}

type RunbooksHandler struct {}

func (s *RunbooksHandler) RegisterGetRunbook(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-runbook",
		Method: http.MethodGet,
		Path: "/runbooks/{runbookId}",
	}, func(ctx context.Context, input *GetRunbookRequest) (*GetRunbookResponse, error) {
		resp := &GetRunbookResponse{}
		return resp, nil
	})
}

func (s *RunbooksHandler) RegisterListRunbooks(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "list-runbooks",
		Method: http.MethodGet,
		Path: "/runbooks",
	}, func(ctx context.Context, input *struct{}) (*ListRunbooksResponse, error) {
		resp := &ListRunbooksResponse{}
		return resp, nil
	})
}
