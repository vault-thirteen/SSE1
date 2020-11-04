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
// File: github.com/vault-thirteen/SSE1/pkg/helper/http/request.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package http

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vault-thirteen/errorz"
)

// Errors.
const (
	ErrfHeaderNotFound = "HTTP Header is not found: %v"
)

// Reads an HTTP Body of the HTTP Request.
func GetHttpRequestBody(
	r *http.Request,
) (reuestBody []byte, err error) {
	reuestBody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer func() {
		var derr = r.Body.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()
	return
}

// Reads an HTTP Cookie of the HTTP Request.
func GetHttpCookie(
	r *http.Request,
	cookieName string,
) (cookieValue string, err error) {
	var cookie *http.Cookie
	cookie, err = r.Cookie(cookieName)
	if err != nil {
		return
	}
	cookieValue = cookie.Value
	return
}

// Reads an HTTP Header of the HTTP Request.
func GetHttpRequestHeader(
	r *http.Request,
	headerName string,
) (headerValue string, err error) {
	var ok bool
	_, ok = r.Header[headerName]
	if !ok {
		err = fmt.Errorf(ErrfHeaderNotFound, headerName)
		return
	}
	headerValue = r.Header.Get(headerName)
	return
}
