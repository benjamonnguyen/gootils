package httputil

func Is2xx(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

func Is3xx(statusCode int) bool {
	return statusCode >= 300 && statusCode < 400
}

func Is4xx(statusCode int) bool {
	return statusCode >= 400 && statusCode < 500
}

func Is5xx(statusCode int) bool {
	return statusCode >= 500 && statusCode < 600
}
