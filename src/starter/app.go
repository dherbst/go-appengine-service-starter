package starter

func init() {
	conf := NewConfig()
	ctx := NewContextWithConfig(conf)
	SetupHandlers(ctx)
}
