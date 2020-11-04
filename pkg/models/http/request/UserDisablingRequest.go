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
// File: github.com/vault-thirteen/SSE1/pkg/models/http/request/UserDisablingRequest.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package request

import (
	_ "github.com/mailru/easyjson/gen"
	"github.com/vault-thirteen/SSE1/pkg/models/user"
)

type UserDisablingRequest struct {
	User UserDisablingRequestUser `json:"user"`
}

func NewUserDisablingRequest(
	httpRequestBody []byte,
) (request *UserDisablingRequest, err error) {
	request = new(UserDisablingRequest)
	err = request.UnmarshalJSON(httpRequestBody)
	if err != nil {
		return
	}
	err = request.Validate()
	if err != nil {
		return
	}
	return
}

func (urr *UserDisablingRequest) Validate() (err error) {
	err = user.ValidateUserAuthenticationName(urr.User.InternalName)
	if err != nil {
		return
	}
	err = user.ValidateUserAuthenticationPassword(urr.User.Password)
	if err != nil {
		return
	}
	err = user.ValidateUserRegistrationSecretCode(urr.User.SecretCode)
	if err != nil {
		return
	}
	return
}
