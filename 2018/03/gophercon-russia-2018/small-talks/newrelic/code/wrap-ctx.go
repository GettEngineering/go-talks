func WrapHandlerCtx(app newrelic.Application, name string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txn := app.StartTransaction(name, w, r)
		defer txn.End()

		ctx := r.Context() // HL
		ctx = context.WithValue(ctx, ContextKeyNewRelicTxn, txn) // HL

		handler.ServeHTTP(txn, r.WithContext(ctx)) // HL
	})
}
