/*
 * ICOS Shell
 *
 * This is the ICOS Shell based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package shellbackend

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}