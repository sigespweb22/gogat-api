var server *Server = nil

func NewServer() *Server {
	if server == nil {
		r := chi.NewRouter()
		server = &Server{
			router: r,
			httpServer: http.Server{
				Addr:    fmt.Sprintf(":%d", configs.GetConfig.Service.APP_PORT),
				Handler: r,
			},
		}
	}

	return server
}