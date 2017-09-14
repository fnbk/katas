package main

import "net/http"

type NotFoundJob struct {
	reader *http.Request
	writer http.ResponseWriter
}

func (job *NotFoundJob) Execute() {
	http.NotFound(job.writer, job.reader)
}
