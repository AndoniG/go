package main

import( 
	"net/http"
)

func NewServer(port string) *Server{
	return &Server{
		port: port,
		router: NewRouter(),
	}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exists := s.router.rules[path]
	if !exists {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

// ... INDICA QUE NO SABES LA CANTIDAD DE PARÁMETROS QUE VAS A RECIBIR DE ESE TIPO EN ESPECÍFICO. LOS GUARDA EN UN SLICE
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
	for _, middleware := range middlewares{
		f = middleware(f)
	}

	return f
}