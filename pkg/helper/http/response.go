////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
// Web Site:  'https://github.com/vault-thirteen'.
// Author:     McArcher.
// Web Site Address is an Address in the global Computer Internet Network.
//
// File: github.com/vault-thirteen/SSE1/pkg/helper/http/response.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package http

import (
	"net/http"

	"github.com/vault-thirteen/SSE1/pkg/interfaces/response"
	"github.com/vault-thirteen/header"
	"github.com/vault-thirteen/mime"
)

// Writes an Error to the HTTP Response Stream.
func SendHttpResponseError(
	w http.ResponseWriter,
	httpStatusCode int,
	errorText error,
) (err error) {
	w.WriteHeader(httpStatusCode)
	_, err = w.Write([]byte(errorText.Error()))
	if err != nil {
		return
	}
	return
}

// Writes an 'Internal Server' Error to the HTTP Response Stream.
func SendHttpResponseInternalServerError(
	w http.ResponseWriter,
	errorText error,
) (err error) {
	return SendHttpResponseError(w, http.StatusInternalServerError, errorText)
}

// Writes a 'Bad Request' Error to the HTTP Response Stream.
func SendHttpResponseBadRequestError(
	w http.ResponseWriter,
	errorText error,
) (err error) {
	return SendHttpResponseError(w, http.StatusBadRequest, errorText)
}

// Writes a 'Forbidden' Error to the HTTP Response Stream.
func SendHttpResponseForbiddenError(
	w http.ResponseWriter,
	errorText error,
) (err error) {
	return SendHttpResponseError(w, http.StatusForbidden, errorText)
}

// Sets an HTTP Header of the HTTP Response.
func SetHttpResponseHeader(
	w http.ResponseWriter,
	headerName string,
	headerValue string,
) {
	w.Header().Set(headerName, headerValue)
}

// Sets a secure HTTP Cookie of the HTTP Response.
func SetHttpSecureCookie(
	w http.ResponseWriter,
	cookieName string,
	cookieValue string,
	cookiePath string,
) {
	var cookie = http.Cookie{
		Name:     cookieName,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     cookiePath,
		Value:    cookieValue,
	}
	w.Header().Add(header.HttpHeaderSetCookie, cookie.String())
}

// Writes an Object in JSON Format to the HTTP Response Stream.
func RespondWithJsonObject(
	w http.ResponseWriter,
	responseObject response.IResponseObject,
) (err error) {
	var buffer []byte
	buffer, err = responseObject.MarshalJSON()
	if err != nil {
		return
	}
	w.Header().Set(
		header.HttpHeaderContentType,
		mime.TypeApplicationJson,
	)
	_, err = w.Write(buffer)
	if err != nil {
		return
	}
	return
}
