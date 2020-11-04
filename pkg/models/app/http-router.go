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
// File: github.com/vault-thirteen/SSE1/pkg/models/app/http-router.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package app

// A Configurator of an HTTP Server's Router.
func (app *Application) initHttpRouter() (err error) {

	// Actions with a single User.
	app.httpRouter.POST(
		"/api/user/register",
		app.httpProtocolCheck(
			app.httpAuthentication(
				app.httpHandlerApiUserRegister, false,
			),
		),
	)

	app.httpRouter.POST(
		"/api/user/disable",
		app.httpProtocolCheck(
			app.httpAuthentication(
				app.httpHandlerApiUserDisable, true,
			),
		),
	)

	app.httpRouter.POST(
		"/api/user/log-in",
		app.httpProtocolCheck(
			app.httpAuthentication(
				app.httpHandlerApiUserLogIn, false,
			),
		),
	)

	app.httpRouter.POST(
		"/api/user/log-out",
		app.httpProtocolCheck(
			app.httpAuthentication(
				app.httpHandlerApiUserLogOut, true,
			),
		),
	)

	// Actions with multiple Users.
	app.httpRouter.POST(
		"/api/users/list",
		app.httpProtocolCheck(
			app.httpAuthentication(
				app.httpHandlerApiUsersList, true,
			),
		),
	)

	return
}
