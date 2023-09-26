// HandleServerError handles separate server errors and sends appropriate responses
func HandleServerError(writer http.ResponseWriter, err error) {
	log.Error().Err(err).Msg("handleServerError()")

	var respErr error

	switch err := err.(type) {
	case *RouterMissingParamError, *RouterParsingError, *json.SyntaxError, *NoBodyError, *ValidationError:
		respErr = responses.SendBadRequest(writer, err.Error())
	case *json.UnmarshalTypeError:
		respErr = responses.SendBadRequest(writer, "bad type in json data")
	case *ItemNotFoundError:
		respErr = responses.SendNotFound(writer, err.Error())
	case *UnauthorizedError:
		respErr = responses.SendUnauthorized(writer, err.Error())
	case *ForbiddenError:
		respErr = responses.SendForbidden(writer, err.Error())
	default:
		respErr = responses.SendInternalServerError(writer, "Internal Server Error")
	}

	if respErr != nil {
		log.Error().Err(respErr).Msg(responseDataError)
	}
}
