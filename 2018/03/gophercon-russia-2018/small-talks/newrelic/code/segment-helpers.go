func SgmFromContext(ctx context.Context) (extsgm newrelic.ExternalSegment) {
	v := ctx.Value(ContextKeyNewRelicSgm)
	if v == nil {
		return
	}
	if sgm, ok := v.(newrelic.ExternalSegment); ok {
		extsgm = sgm
	}
	return
}

func StartExternalSegmentCtx(ctx context.Context, r *http.Request) context.Context {
	txn := TxnFromContext(ctx)
	if txn == nil {
		return ctx
	}
	s := newrelic.StartExternalSegment(txn, r)
	return context.WithValue(ctx, ContextKeyNewRelicSgm, s)
}

func EndExternalSegmentCtx(ctx context.Context, resp *http.Response) context.Context {
	sgm := SgmFromContext(ctx)
	sgm.Response = resp
	sgm.End()

	return context.WithValue(ctx, ContextKeyNewRelicSgm, nil)
}
