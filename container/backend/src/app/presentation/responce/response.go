package responce

import (
	"net/http"
	"todo/app/application/schema"
	"todo/app/core"
)

func ConvertErrorCode(code core.ErrorCode) int {
	switch code {
	case core.BadRequestError:
		return http.StatusBadRequest
	case core.ValidationError:
		return http.StatusUnprocessableEntity
	case core.NotFoundError:
		return http.StatusNotFound
	case core.AlreadyExistsError:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

func NewDefaultRespoce(msg string) *schema.DefaultResponseModel {
	return &schema.DefaultResponseModel{
		Message: msg,
	}
}
