package pagination

import (
	"context"
	"testing"

	"gotest.tools/assert"
)

type Author struct {
	Name string `json:"name"`
}

type Book struct {
	Name      string  `json:"name"`
	Publisher string  `json:"publisher"`
	Author    *Author `json:"author"`
}

type Protobuf struct {
	Count       int     `json:"count"`
	Data        []*Book `json:"data"`
	CurrentPage int     `json:"currentPage"`
	PerPage     int     `json:"perPage"`
}

func TestNewPagination(t *testing.T) {

	pl, err := NewPagination(context.Background(), &PaginationRequest{
		CurrentPage: 1,
		PerPage:     10,
		GetCount: func(ctx context.Context, dataLen int) (int, error) {
			return dataLen, nil
		},
		GetData: func(ctx context.Context) ([]interface{}, error) {
			return []interface {
			}{
				&Book{
					Name:      "Book1",
					Publisher: "Publisher1",
					Author: &Author{
						Name: "Author1",
					},
				},
				&Book{
					Name:      "Book2",
					Publisher: "Publisher2",
				},
			}, nil
		},
	})

	out := &Protobuf{}
	pl.Cast(out)

	assert.NilError(t, err)
	assert.Equal(t, out.Data[0].Author.Name, "Author1")
	assert.Equal(t, out.Count, 2)
}
