package exception

import (
	"net/http"
	"pznrestfulapi/helper"
	"pznrestfulapi/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	internalServerError(writer, request, err)
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
