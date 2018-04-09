func InstrumentRoutes(r *mux.Router, app newrelic.Application) *mux.Router {
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		h := newrelicutil.WrapHandlerCtx(app, RouteName(route), route.GetHandler())
		route.Handler(h)

		return nil
	})
	if nil != r.NotFoundHandler {
		r.NotFoundHandler = newrelicutil.WrapHandlerCtx(app, "NotFoundHandler", r.NotFoundHandler)
	}
	return r
}

