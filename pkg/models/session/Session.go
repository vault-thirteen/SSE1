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
// File: github.com/vault-thirteen/SSE1/pkg/models/session/Session.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package session

import (
	"database/sql"
	"time"
)

// Session.
type Session struct {

	// Basic Attributes...

	// Session Id.
	Id uint

	// Session User.
	User SessionUser

	// Time Points...

	// Time of the Session's Start.
	StartTime time.Time

	// Time of the last Activity on the Session.
	LastAccessTime time.Time

	// Time when the Session was closed (finished).
	EndTime sql.NullTime

	// A unique Marker of a Session and its Hash Sum.
	Marker     string
	MarkerHash string

	// Access Token Key which was created at the Moment of Token Generation.
	TokenKey string
}
