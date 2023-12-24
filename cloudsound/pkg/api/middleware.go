package api

import (
	"fmt"
	"net/http"
	"github.com/go-http-utils/headers"
)

func (api *api) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		methods := fmt.Sprintf(
			"%v, %v, %v, %v",
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		);
		origin := "*"
		w.Header().Set(headers.AccessControlAllowOrigin,origin)
		w.Header().Set(headers.AccessControlAllowMethods, methods)
		api.logger.Info(r.RemoteAddr)
		api.logger.Info(fmt.Sprintf("Method: %v", r.Method));
		next.ServeHTTP(w, r);
	})
}