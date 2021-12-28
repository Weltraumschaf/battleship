package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func retrieveId(r *http.Request) (uuid.UUID, error) {
	paramName := "id"
	literalId, err := retrievePathParameter(r, paramName)

	if err != nil {
		return uuid.UUID{}, err
	}

	id, err := uuid.Parse(literalId)

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("malformed UUID given for '%s'", paramName)
	}

	return id, nil
}

func retrievePathParameter(r *http.Request, name string) (string, error) {
	vars := mux.Vars(r)
	literal, ok := vars[name]

	if ok {
		return literal, nil
	}

	return "", fmt.Errorf("path parameter '%s' missing", name)
}
