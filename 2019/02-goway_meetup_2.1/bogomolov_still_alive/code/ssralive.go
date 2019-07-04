func (c *CheckService) CheckAlive(ctx context.Context) AliveStatus {
	dbStatus := "OK"
	if err := c.DBPinger.PingContext(ctx); err != nil {
		dbStatus = "Error"
	}
	return &ssr.AliveStatus{
		Alive:        true,
		DB:           dbStatus,
		RuntimeData:  c.CollectRuntimeData(ctx),
		PkgBuildMeta: c.CollectPkgBuildMeta(ctx),
	}
}