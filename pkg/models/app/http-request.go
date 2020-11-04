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
// File: github.com/vault-thirteen/SSE1/pkg/models/app/http-request.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpHelper "github.com/vault-thirteen/SSE1/pkg/helper/http"
	netHelper "github.com/vault-thirteen/SSE1/pkg/helper/net"
	"github.com/vault-thirteen/SSE1/pkg/models/http/request"
)

// Reads the HTTP Request for the Handler: httpHandlerApiUserRegister.
// Returns 'true' on Success.
func (app *Application) getHttpRequest_ApiUserRegister(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) (requestObject *request.UserRegistrationRequest, ok bool) {

	// Read the raw HTTP Request.
	var requestBody []byte
	var err error
	requestBody, err = httpHelper.GetHttpRequestBody(r)
	if err != nil {
		app.handleCriticalError(w, err, SenderGetHttpRequestBody)
		return
	}
	// Decode and validate the Request.
	requestObject, err = request.NewUserRegistrationRequest(requestBody)
	if err != nil {
		app.handleBadRequestError(w, err, SenderNewUserRegistrationRequest)
		return
	}
	ok = true
	return
}

// Reads the HTTP Request for the Handler: httpHandlerApiUserDisable.
// Returns 'true' on Success.
func (app *Application) getHttpRequest_ApiUserDisable(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) (requestObject *request.UserDisablingRequest, ok bool) {

	// Read the raw HTTP Request.
	var requestBody []byte
	var err error
	requestBody, err = httpHelper.GetHttpRequestBody(r)
	if err != nil {
		app.handleCriticalError(w, err, SenderGetHttpRequestBody)
		return
	}
	// Decode and validate the Request.
	requestObject, err = request.NewUserDisablingRequest(requestBody)
	if err != nil {
		app.handleBadRequestError(w, err, SenderNewUserDisablingRequest)
		return
	}
	ok = true
	return
}

// Reads the HTTP Request for the Handler: httpHandlerApiUserLogIn.
// Returns 'true' on Success.
func (app *Application) getHttpRequest_ApiUserLogIn(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) (requestObject *request.UserLogInRequest, ok bool) {

	// Read the raw HTTP Request.
	var requestBody []byte
	var err error
	requestBody, err = httpHelper.GetHttpRequestBody(r)
	if err != nil {
		app.handleCriticalError(w, err, SenderGetHttpRequestBody)
		return
	}
	// Decode and validate the Request.
	requestObject, err = request.NewUserLogInRequest(requestBody)
	if err != nil {
		app.handleBadRequestError(w, err, SenderNewUserLogInRequest)
		return
	}

	// Add the Machine Settings.
	requestObject.Machine = request.UserLogRequestMachine{
		Host: netHelper.GetAddressHost(r.RemoteAddr),
		BrowserUserAgent: request.UserLogRequestMachineBrowserUserAgent{
			Name: r.UserAgent(),
			//Id:   -1,
		},
	}
	err = request.ValidateMachineHost(requestObject.Machine.Host)
	if err != nil {
		app.handleForbiddenError(w, err, SenderValidateMachineHost)
		return
	}
	err = request.ValidateMachineBrowserUserAgent(requestObject.Machine.BrowserUserAgent.Name)
	if err != nil {
		app.handleForbiddenError(w, err, SenderValidateMachineBrowserUserAgent)
		return
	}
	requestObject.Machine.BrowserUserAgent.Id, err =
		app.buam.GetBrowserUserAgentId(requestObject.Machine.BrowserUserAgent.Name)
	if err != nil {
		app.handleCriticalError(w, err, SenderGetBrowserUserAgentId)
		return
	}

	ok = true
	return
}

// Reads the HTTP Request for the Handler: httpHandlerApiUserLogOut.
// Returns 'true' on Success.
func (app *Application) getHttpRequest_ApiUserLogOut(
	w http.ResponseWriter,
	r *http.Request,
	ps httprouter.Params,
) (requestObject *request.UserLogOutRequest, ok bool) {

	// Read the raw HTTP Request.
	var requestBody []byte
	var err error
	requestBody, err = httpHelper.GetHttpRequestBody(r)
	if err != nil {
		app.handleCriticalError(w, err, SenderGetHttpRequestBody)
		return
	}
	// Decode and validate the Request.
	requestObject, err = request.NewUserLogOutRequest(requestBody)
	if err != nil {
		app.handleBadRequestError(w, err, SenderNewUserLogOutRequest)
		return
	}
	ok = true
	return
}
