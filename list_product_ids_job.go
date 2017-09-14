package main

import (
	"encoding/json"
	"net/http"
)

type ListProductIDsJob struct {
	writer http.ResponseWriter
	reader *http.Request
	ProductProvider
}

func (job *ListProductIDsJob) Execute() {
	ids := job.ProductProvider.ProductIDs()
	job.writeResponse(ids)
}

func (job *ListProductIDsJob) writeResponse(ids []string) {
	job.writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(job.writer).Encode(ids)
}
