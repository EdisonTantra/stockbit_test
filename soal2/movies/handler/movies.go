package handler

import (
	"context"
	"strconv"

	log "github.com/micro/micro/v3/service/logger"

	movies "movies/proto"
	usecase "movies/usecase"
	"movies/utils/pagination"
)

type Movies struct{}

func (e *Movies) Call(ctx context.Context, req *movies.Request, rsp *movies.Response) error {
	log.Info("Received Movies.Call request")
	rsp.Msg = "Hello"
	return nil
}

func (e *Movies) Stream(ctx context.Context, req *movies.StreamingRequest, stream movies.Movies_StreamStream) error {
	log.Infof("Received Movies.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&movies.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (e *Movies) PingPong(ctx context.Context, stream movies.Movies_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&movies.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func (e *Movies) SearchMovie(ctx context.Context, req *movies.SearchRequest, rsp *movies.SearchResponse) error {
	const PerPage = 20
	jsonRsp, err := usecase.CreateImdbSearchRequest(req.Query, int(req.CurrentPage))
	if err != nil {
		log.Fatal(err)
		return err
	}

	count, _ := strconv.Atoi(jsonRsp.TotalResults)

	list, err := pagination.NewPagination(ctx, &pagination.PaginationRequest{
		CurrentPage: int(req.CurrentPage),
		PerPage:     PerPage,
		Count:       count,
		GetCount: func(ctx context.Context, dataLen int) (i int, e error) {
			return len(jsonRsp.Data), nil
		},
		GetData: func(ctx context.Context) (i []interface{}, e error) {
			var resp []interface{}
			for _, mv := range jsonRsp.Data {
				mo := &movies.MovieData{}
				usecase.MovieSerializer(mv, mo)
				resp = append(resp, mo)
			}
			return resp, nil
		},
	})
	if err != nil {
		log.Fatal(err)
		return err
	}

	list.Cast(rsp)
	return nil
}

func (e *Movies) GetMovieByID(ctx context.Context, req *movies.GetMovieByRequest, rsp *movies.MovieData) error {
	jsonRsp, err := usecase.CreateImdbDetailRequest(req.Value)
	if err != nil {
		log.Fatal(err)
		return err
	}

	usecase.MovieSerializer(jsonRsp, rsp)
	return nil
}
