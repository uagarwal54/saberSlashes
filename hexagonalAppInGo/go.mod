module hexagonalAppInGo

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/mux v1.8.0
)

replace rest v0.0.0 => ./pkg/http/rest
