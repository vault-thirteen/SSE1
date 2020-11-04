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
// File: github.com/vault-thirteen/SSE1/pkg/models/user/UserAuthentication.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package user

import (
	"fmt"

	"github.com/vault-thirteen/unicode"
)

// Settings.
const (
	UserAuthenticationNameLengthMin     = 3
	UserAuthenticationNameLengthMax     = 40
	UserAuthenticationPasswordLengthMin = 8
	UserAuthenticationPasswordLengthMax = 64
)

// Errors.
const (
	ErrUserAuthenticationName     = "Error in UserAuthentication.Name: '%v'."
	ErrUserAuthenticationPassword = "Error in UserAuthentication.Password: '%v'."
)

// User Authentication Data.
type UserAuthentication struct {

	// Internal Name used for Authentication.
	Name string

	// Password used for Authentication.
	Password string
}

func ValidateUserAuthenticationName(
	name string,
) (err error) {
	if (len(name) < UserAuthenticationNameLengthMin) ||
		(len(name) > UserAuthenticationNameLengthMax) {
		err = fmt.Errorf(ErrUserAuthenticationName, name)
		return
	}
	var letters = []rune(name)
	if !unicode.SymbolIsLatLetter(letters[0]) {
		err = fmt.Errorf(ErrUserAuthenticationName, name)
		return
	}
	for i := 1; i < len(name); i++ {
		if (!unicode.SymbolIsLatLetter(letters[i])) &&
			(!unicode.SymbolIsNumber(letters[i])) {
			err = fmt.Errorf(ErrUserAuthenticationName, name)
			return
		}
	}
	return
}

func ValidateUserAuthenticationPassword(
	password string,
) (err error) {
	if (len(password) < UserAuthenticationPasswordLengthMin) ||
		(len(password) > UserAuthenticationPasswordLengthMax) {
		err = fmt.Errorf(ErrUserAuthenticationPassword, password)
		return
	}
	return
}
