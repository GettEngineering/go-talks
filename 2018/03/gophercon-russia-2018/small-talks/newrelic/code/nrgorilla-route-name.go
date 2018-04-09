
func RouteName(route *mux.Route) string {
	...
	// r.Handle("/api/v1/{env}/callers/{phone}", h).Methods("GET")
	// Out: "GET /api/v1/{env}/callers/{phone}"
	if n, _ := route.GetPathTemplate(); n != "" {
		ms, _ := route.GetMethods() // HL
		return strings.Join(ms, "/") + " " + n // HL
	}
	...
	return n
}
