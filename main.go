// gio-cache instructs the browser to cache your wasm app until you update the version or significant time has passed
// in addition it shows a simple static css spinner while the app is being loaded and initialized
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	origin                 = "*"
	methods                = "GET"
	cacheControl           = "no-cache"
	wasmBinarVersionEnvVar = "WASM_BINARY_VERSION"
	port                   = "5173"
	staticDir              = "bin"
	compressedWASMFile     = "/main.wasm.br"
)

func main() {
	var (
		currentVersion = os.Getenv(wasmBinarVersionEnvVar)
		fileServer     = http.FileServer(http.Dir(staticDir))
	)
	if currentVersion == "" {
		panic(fmt.Errorf("You must set the env var: %s", wasmBinarVersionEnvVar))
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", methods)
		w.Header().Set("ETag", currentVersion)
		w.Header().Set("Cache-Control", cacheControl)

		var match = r.Header.Get("If-None-Match")
		cacheDiff := fmt.Sprintf("%s -> %s", match, currentVersion)
		if match == currentVersion {
			green("CACHED", fmt.Sprintf("%s %s", r.URL.Path, cacheDiff))
			w.WriteHeader(http.StatusNotModified)
			return
		}
		yellow("SERVED", fmt.Sprintf("%s %s", r.URL.Path, cacheDiff))

		if r.URL.Path == compressedWASMFile {
			w.Header().Set("Vary", "Accept-Encoding")
			w.Header().Set("Content-Encoding", "br")
			w.Header().Set("Content-Type", "application/wasm")
		}
		fileServer.ServeHTTP(w, r)
	})

	log.Println("Serving WASM app on port ", port)
	http.ListenAndServe(":"+port, nil)
}
