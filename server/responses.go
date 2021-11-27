package server

type response struct {
	success bool
	error   interface{}
	result  interface{}
}

func successResponse(result interface{}) *response {
	return &response{
		success: true,
		result:  result,
	}
}

func errorResponse(error interface{}) *response {
	return &response{
		success: true,
		error:   error,
	}
}
