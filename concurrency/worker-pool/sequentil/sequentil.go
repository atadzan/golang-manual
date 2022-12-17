package main

import (
	"fmt"
	"time"
)

type JobInput struct {
	startTime time.Time
	endTime   time.Time
}

type JobResponse struct {
}

type Response struct {
	finalOutput string
}

func main() {
	fmt.Println(businessFunctionality().finalOutput)
}

func businessFunctionality() Response {
	var jobInputs []JobInput //Create job input for each job
	var jobResponses []JobResponse
	for _, jobInput := range jobInputs {
		jobResponses = append(jobResponses, businessFunctionalityJob(jobInput))
	}
	return combineResponses(jobResponses)
}

func combineResponses(jobResponses []JobResponse) Response {
	return Response{finalOutput: "Well done!"}
}

func businessFunctionalityJob(jobInput JobInput) JobResponse {
	return JobResponse{}
}
