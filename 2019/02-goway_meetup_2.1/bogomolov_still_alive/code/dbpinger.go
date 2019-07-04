package code

import "context"

// Interface of database access object acceptable by factory
type DBPinger interface {
	PingContext(ctx context.Context) error
}

// Build database check function from db connection.
// START OMIT
func MakeDbPinger(db DBPinger, dbName string) Pinger {
	return func(ctx context.Context) (map[string]interface{}, error) {
		prefix := "db"
		if dbName != "" {
			prefix = "db_" + dbName
		}
		response := map[string]interface{}{
			prefix: "OK",
		}
		err := db.PingContext(ctx) // HL
		if err != nil {
			response[prefix] = err.Error()
		}
		return response, err
		// END OMIT
	}
}


