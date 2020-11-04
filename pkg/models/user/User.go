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
// File: github.com/vault-thirteen/SSE1/pkg/models/user/User.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package user

import "fmt"

// Settings.
const (
	UserPublicNameLengthMin = 3
	UserPublicNameLengthMax = 40
)

// Errors.
const (
	ErrUserPublicName = "Error in User.PublicName: '%v'."
)

// User.
type User struct {

	// 'true' means that a User is registered.
	// 'false' means that a User is disabled.
	IsEnabled bool

	// A unique Identifier in the List of all Users.
	Id uint

	// Authentication Data.
	Authentication UserAuthentication

	// Registration Data.
	Registration UserRegistration

	// Public Name displayed at the Website.
	PublicName string
}

func ValidateUserPublicName(
	publicName string,
) (err error) {
	if (len(publicName) < UserPublicNameLengthMin) ||
		(len(publicName) > UserPublicNameLengthMax) {
		err = fmt.Errorf(ErrUserPublicName, publicName)
		return
	}
	return
}
