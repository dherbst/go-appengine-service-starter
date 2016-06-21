package starter

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
)

// SetupHandlers creates the route to handler mapping.
func SetupHandlers(ctx context.Context) {
	http.Handle("/version", VersionHandler(ctx))
}

// VersionHandler returns the version of this service.
func VersionHandler(ctx context.Context) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodOptions {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		config, ok := ConfigFromContext(ctx)
		if !ok {
			http.Error(w, "No Config", http.StatusInternalServerError)
			return
		}

		val, err := json.Marshal(config)
		if err != nil {
			http.Error(w, "No Config", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, string(val))
	}
	return http.HandlerFunc(fn)
}
