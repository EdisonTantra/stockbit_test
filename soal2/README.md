

using go-micro v3


gomicro - https://github.com/micro/micro


How to run this service:
1. go get github.com/micro/micro/v3
2. micro server
3. micro login 
    - username: admin
    - password: micro

4. micro services
5. micro run .

micro new movies
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/micro/v3/cmd/protoc-gen-micro



micro server &
micro run .
## kalo update proto 
make proto
## trus build lagi
make build
## liat status
micro status
## liat endpoints
micro registry getService --service=movies
## get service
micro get service movies

## liat logs service
micro logs movies
## call grpc
micro call movies  Movies.SearchMovie '{"query": "harry", "currentPage": 1}'
micro call movies  Movies.GetMovieByID '{"value": "tt7134194"}'

################################### PERLU API GATEWAY PAKE ECHO
// TODO custom default page value
// req.currentPage -> default 1
// req.PerPage -> default 20


################################### PERLU API GATEWAY PAKE ECHO

## perlu set ini di awal
micro config set movies.imdb_api_key faf7e5bb
micro config set movies.imdb_base_url https://www.omdbapi.com/