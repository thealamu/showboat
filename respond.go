package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//Error responds with an error message in JSON format
func (s *Server) Error(w http.ResponseWriter, err error, statusCode int) {
	errResp := toErrorResponse(err)
	log.Println(errResp)
	s.JSON(w, errResp, statusCode)
}

//JSON responds with data in JSON format
func (s *Server) JSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		s.logger.Error(err)
	}
}

type errResponse struct {
	Message interface{} `json:"message"`
}

func toErrorResponse(err error) errResponse {
	if err != nil {
		return errResponse{err.Error()}
	}
	return errResponse{nil}
}
