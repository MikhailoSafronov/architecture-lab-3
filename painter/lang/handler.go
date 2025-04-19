package lang

import (
    "net/http"
    "github.com/roman-mazur/architecture-lab-3/painter"
)

type Handler struct {
    loop   *painter.Loop
    parser *Parser
    srv    *http.Server
}

// NewHandler створює новий обробник HTTP запитів
func NewHandler(loop *painter.Loop) *Handler {
    return &Handler{
        loop:   loop,
        parser: &Parser{},
    }
}

// StartServer запускає HTTP сервер
func (h *Handler) StartServer() {
    mux := http.NewServeMux()
    mux.Handle("/api/", HttpHandler(h.loop, h.parser))
    
    h.srv = &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }
    
    go h.srv.ListenAndServe()
}
