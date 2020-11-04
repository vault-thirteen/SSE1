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
// File: github.com/vault-thirteen/SSE1/pkg/models/app/errors.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package app

// Errors Messages.
const (
	ErrAuthenticationFailure      = "Authentication Failure"
	ErrNullPointer                = "Null Pointer"
	ErrTokenDataMismatch          = "Token Data Mismatch"
	ErrTypeCastFailure            = "Type Cast Failure"
	ErrUnsupportedLoggerType      = "Unsupported Logger Type"
	ErrUnsupportedStorageType     = "Unsupported Storage Type"
	ErrUserCanNotBeDisabled       = "User can not be disabled"
	ErrCanNotLogOutOtherUsers     = "Can not log out other Users"
	ErrRegisteredUserDoesNotExist = "Registered User does not exist"
)

// Errors Formats.
const (
	ErrfBadRequestError               = "Bad Request: %v"
	ErrfCriticalError                 = "Critical Error: %v"
	ErrfForbiddenError                = "Forbidden: %v"
	ErrfHttpResponseError             = "HTTP Response Error: %v"
	ErrfUserAuthenticationNameIsTaken = "User Authentication Name is not available: '%v'."
)

// Sender Names.
const (
	SenderDisableUser                       = "DisableUser"
	SenderGetBrowserUserAgentId             = "GetBrowserUserAgentId"
	SenderGetHttpRequestBody                = "GetHttpRequestBody"
	SenderGetUserIdByAuthenticationName     = "GetUserIdByAuthenticationName"
	SenderHttpAuthentication                = "HTTP Authentication"
	SenderHttpHandlerApiUserDisable         = "httpHandlerApiUserDisable"
	SenderHttpHandlerApiUserLogOut          = "HttpHandlerApiUserLogOut"
	SenderHttpProtocolCheck                 = "HTTP Protocol Check"
	SenderIsUserAuthenticationNameFree      = "IsUserAuthenticationNameFree"
	SenderListRegisteredUsersPublicNames    = "ListRegisteredUsersPublicNames"
	SenderLogUserIn                         = "LogUserIn"
	SenderLogUserOut                        = "LogUserOut"
	SenderNewUserDisablingRequest           = "NewUserDisablingRequest"
	SenderNewUserLogInRequest               = "NewUserLogInRequest"
	SenderNewUserLogOutRequest              = "NewUserLogOutRequest"
	SenderNewUserRegistrationRequest        = "NewUserRegistrationRequest"
	SenderRegisteredUserIdExists            = "RegisteredUserIdExists"
	SenderRegisterUser                      = "RegisterUser"
	SenderRespondWithJsonObject             = "RespondWithJsonObject"
	SenderUnpackAuthDataFromContext         = "unpackAuthDataFromContext"
	SenderUpdateActiveSessionLastAccessTime = "UpdateActiveSessionLastAccessTime"
	SenderValidateMachineBrowserUserAgent   = "ValidateMachineBrowserUserAgent"
	SenderValidateMachineHost               = "ValidateMachineHost"
)
