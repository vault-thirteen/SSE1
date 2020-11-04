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
// File: github.com/vault-thirteen/SSE1/pkg/models/app/messages.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package app

import "fmt"

// Messages.
const (
	MsgAppStarted           = "Application has been started."
	MsgAppStopped           = "Application has been stopped."
	MsgSignalSigterm        = "'SIGTERM' Signal is received."
	MsgSignalSigint         = "'SIGINT' Signal is received."
	MsgHttpServerStop       = "HTTP Server has been stopped"
	MsgStorageConnection    = "Storage has been connected"
	MsgStorageCheck         = "Storage has been checked"
	MsgStorageDisconnection = "Storage has been disconnected"
	MsgIsEnabled            = "is enabled"
	MsgIsDisabled           = "is disabled"
)

// Message Formats.
const (
	MsgHttpServerStart = "HTTP Server Start (TLS %v)"
)

// Composes a Message for the Start of an HTTP Server.
func makeMsgHttpServerStart(
	isTlsEnabled bool,
) (msg string) {
	if isTlsEnabled {
		msg = fmt.Sprintf(MsgHttpServerStart, MsgIsEnabled)
	} else {
		msg = fmt.Sprintf(MsgHttpServerStart, MsgIsDisabled)
	}
	return
}
