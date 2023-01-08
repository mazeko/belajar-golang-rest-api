package exception

import (
	"belajar-goalng-rest-api/helper"
	"belajar-goalng-rest-api/model/web"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationError(writer, request, err) {
		return
	}

	if notFoundError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, isErr := err.(validator.ValidationErrors)
	if isErr {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		var message []string
		var errMessage string
		for _, msg := range exception {
			errMessage = fmt.Sprintf("The %s is %s", msg.Field(), msg.ActualTag())
			message = append(message, errMessage)
		}

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   message,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, isErr := err.(NotFoundError)
	if isErr {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
