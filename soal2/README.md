How to run this service:
1. cd movies
2. micro server &
3. micro login 
    - username: admin
    - password: micro
4.  micro config set movies.imdb_api_key faf7e5bb
    micro config set movies.imdb_base_url https://www.omdbapi.com/
5. make init
6. micro server &
7. micro run .
8. open browser
9. untuk search dapat ke URL http://localhost:8080/movies/searchMovie/?query=<query>&currentPage=<page>
10. untuk detail dapat ke URL http://localhost:8080/movies/getMovieByID/?value=<imbdID>
