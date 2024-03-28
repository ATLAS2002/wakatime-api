package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Error struct {
  Message string `json:"message"`
  Code    int    `json:"code"`
}

func (e *Error) Error() string {
  return e.Message
}

func DefaultError(e error) *Error {
  return &Error{
    Message: e.Error(),
    Code:    http.StatusInternalServerError,
  }
}

func SendResponse(w http.ResponseWriter, code int, res interface{}) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  json.NewEncoder(w).Encode(res)
}

func ParseLimit(r *http.Request) int {
  limitParam := r.URL.Query().Get("limit")

  limit, err := strconv.Atoi(limitParam)
  if err != nil {
    limit = -1
  }

  return limit
}