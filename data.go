package main

import (
	"encoding/json"
)

type postUserID struct {
	UserID int64 `json:"user_id"`
}

type resp struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Err     string      `json:"err,omitempty"`
	Message string      `json:"message,omitempty"`
}

func (r resp) JSON() []byte {
	b, _ := json.Marshal(r)
	return b
}

func newSuccessResp(data interface{}, msg string) resp {
	return resp{
		Success: true,
		Data:    data,
		Message: msg,
	}
}

func successUsersResp(data interface{}) resp {
	return resp{
		Success: true,
		Data:    data,
	}
}

func newBadRequestResp(msg string) resp {
	return resp{
		Success: false,
		Err:     "ERR_BAD_REQUEST",
		Message: msg,
	}
}

func newErrInternalResp(msg string) resp {
	return resp{
		Success: false,
		Err:     "ERR_INTERNAL_ERROR",
		Message: msg,
	}
}
