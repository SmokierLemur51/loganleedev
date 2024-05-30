package middleware

import (
  "net/http"
)

type Middleware func(http.Handler) http.Handler 


func CreateStack(xs ...Middleware) Middleware {
  return func(next http.Handler) http.Handler {
    for i := len(xs)-1; i >= 0; i-- {
      x := xs[i]
      next = x(next)
    }
    // This will return the topmost middleware which will nest all subsequent
    // middleware underneath.
    return next
  }
}


