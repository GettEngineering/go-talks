func TxnFromContext(ctx context.Context) newrelic.Transaction {
	v := ctx.Value(ContextKeyNewRelicTxn)
	if v == nil {
		return nil
	}

	if txn, ok := v.(newrelic.Transaction); ok {
		return txn
	}
	return nil
}
