package pagination

import (
	"context"
	"movies/utils/caster"
)

type PaginationList struct {
	Count       int           `json:"count"`
	Data        []interface{} `json:"data"`
	CurrentPage int           `json:"currentPage"`
	PerPage     int           `json:"perPage"`
}

type PaginationRequest struct {
	CurrentPage int
	PerPage     int
	Count       int
	GetCount    func(ctx context.Context, dataLen int) (int, error)
	GetData     func(ctx context.Context) ([]interface{}, error)
}

func (c *PaginationList) Cast(out interface{}) error {
	return caster.Cast(c, out)
}

func NewPagination(ctx context.Context, req *PaginationRequest) (*PaginationList, error) {

	data, err := req.GetData(ctx)
	if err != nil {
		return nil, err
	}

	count, err := req.GetCount(ctx, len(data))
	if err != nil {
		return nil, err
	}

	if req.PerPage == 0 {
		req.PerPage = count
	}

	if req.CurrentPage == 0 {
		req.CurrentPage = 1
	}

	return &PaginationList{
		Count:       req.Count,
		Data:        data,
		PerPage:     req.PerPage,
		CurrentPage: req.CurrentPage,
	}, nil
}
