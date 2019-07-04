package code

import (
	"net/http"
	"context"

	"github.com/gtforge/go-healthcheck"
	"database/sql"
)

func main() {
	db := sql.OpenDB()
	// START OMIT
	// type ICache interface {
	//  IsLoaded(context.Context) Bool
	//  IsHealthy(context.Context) Bool
	//}
	cache := cache.InitCache()
	hc := healthcheck.NewHealthCheck(
		healthcheck.MakeDbPinger(db, "master"),
	)

	handler := healthcheck.MakeHealthcheckHandler(hc)
	// END OMIT
	mux := http.NewServeMux()
	mux.Handle("/alive", handler)
	http.ListenAndServe(":3000", mux)
}
