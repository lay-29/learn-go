package No7

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"net/http"
	"time"
)

var user = 0

func Test() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))

	})
	r.Route("/api", func(r chi.Router) {
		//r.Use(func(next http.Handler) http.Handler {
		//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//		fmt.Println("中间件1")
		//		next.ServeHTTP(w, r)
		//		fmt.Println("中间件1: after")
		//	})
		//})
		//r.Use(func(next http.Handler) http.Handler {
		//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//		fmt.Println("中间件2")
		//		next.ServeHTTP(w, r)
		//		fmt.Println("中间件2: after")
		//	})
		//})

		r.Use(ApiContextMiddleware())
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			apiCtx := GetAPIContext(r)
			w.Write([]byte(fmt.Sprintf("Api Page: user：=> %v! start at :=> %v", apiCtx.UserID, apiCtx.StartAt)))
		})

		r.Get("/master", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Api Master Page"))
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				user++
				w.Write([]byte(fmt.Sprintf("User %v!", user)))
			})
		})
	})
	http.ListenAndServe(":8080", r)
}

// Api 调用时的上下文
type APIContext struct {
	RequestID string
	UserID    int64
	StartAt   time.Time
}
type apiContextKeyType struct{}

var apiContextKey = apiContextKeyType{}

func ApiContextMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiCtx := &APIContext{
				RequestID: uuid.New().String(),
				UserID:    int64(time.Now().Second()),
				StartAt:   time.Now(),
			}
			ctx := context.WithValue(r.Context(), apiContextKey, apiCtx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
func GetAPIContext(r *http.Request) *APIContext {
	if ctx, ok := r.Context().Value(apiContextKey).(*APIContext); ok {
		return ctx
	}
	return nil
}
