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
// File: github.com/vault-thirteen/SSE1/pkg/interfaces/storage/IStorage.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package storage

import (
	"github.com/vault-thirteen/SSE1/pkg/helper/jwt"
	"github.com/vault-thirteen/SSE1/pkg/models/http/request"
	"github.com/vault-thirteen/SSE1/pkg/models/session"
	"github.com/vault-thirteen/SSE1/pkg/models/user"
)

// Storage Interface.
type IStorage interface {

	// I. Public System Methods.

	// Reads the Data Source Name of the Storage.
	GetDsn() string

	// Connects the Storage.
	Connect() error

	// Checks the Integrity Storage.
	Check() error

	// Disconnects the Storage.
	Disconnect() error

	// II. User Methods.

	// Checks whether a User Name exists in the Database.
	// The State of a User (registered or disabled) is ignored.
	IsUserAuthenticationNameFree(string) (bool, error)

	// Check if a registered User with Id exists.
	RegisteredUserIdExists(uint) (bool, error)

	// Reads a User's Id by its Authentication Name.
	GetUserIdByAuthenticationName(string) (uint, error)

	// Registers a User.
	RegisterUser(*user.User) error

	// Disables a User.
	// When a User is disabled, it can not be registered again.
	// To register a User with the disabled Name,
	// the disabled User must be first deleted from the Database.
	DisableUser(*user.User) error

	// Logs a User in.
	LogUserIn(*user.User, *request.UserLogRequestMachine) (*session.Session, *jwt.TokenData, error)

	// Logs a User out.
	LogUserOut(*user.User, *session.Session) error

	// III. Users Methods.

	// Lists public Names of all registered Users.
	ListRegisteredUsersPublicNames() ([]string, error)

	// IV. Other Methods.

	// Returns an Id of a Browser's User Agent by its full Name.
	GetBrowserUserAgentId(string) (uint, error)

	// Returns a Key for a Token with a specified Marker Hash Sum.
	// Token's Session must be active.
	GetTokenKeyByMarkerHash(string) (interface{}, error)

	// Returns an active Session by its Id.
	GetActiveSessionById(uint) (*session.Session, error)

	// Updates the 'LastAccessTime' Parameter of an active Session.
	UpdateActiveSessionLastAccessTime(*session.Session) error
}
