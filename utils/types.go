package utils

type ApiError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
