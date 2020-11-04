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
// File: github.com/vault-thirteen/SSE1/pkg/helper/random/random.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package random

import (
	"encoding/hex"
	"strings"

	"github.com/google/uuid"
	"github.com/vault-thirteen/random"
)

// Generates a unique random Marker using a mixed Technique including:
//	- An UUID of the fourth Version,
//	- Some random Bytes received from the random Numbers' Generator of the O.S.
func CreateUniqueMarker() (marker string, err error) {
	const (
		Separator = "-"
		Void      = ""
	)
	var uid uuid.UUID
	uid, err = uuid.NewRandom()
	if err != nil {
		return
	}
	marker = uid.String()
	marker = strings.Replace(marker, Separator, Void, -1)
	for i := 1; i <= 16; i++ {
		var rnd uint
		rnd, err = random.Uint(0, 255)
		if err != nil {
			return
		}
		marker = marker + hex.EncodeToString([]byte{byte(rnd)})
	}
	marker = strings.ToUpper(marker)
	return
}
