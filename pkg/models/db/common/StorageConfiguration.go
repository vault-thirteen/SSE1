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
// File: github.com/vault-thirteen/SSE1/pkg/models/db/common/StorageConfiguration.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package common

import (
	"fmt"
	"strings"
)

// Connection String Syntax Settings.
const (
	ConnectionParameterDelimiter         = "&"
	ConnectionParameterKeyValueDelimiter = "="
)

// Errors.
const (
	ErrfSyntaxErrorAt = "Syntax Error at: '%v'"
	ErrfDuplicateKey  = "Duplicate Key: '%v'"
)

// Configuration of a generic Storage.
// Each Storage implements the 'IStorage' Interface.
type StorageConfiguration struct {
	Address              string
	ConnectionParameters map[string]string
	Database             string
	Password             string
	User                 string
}

// Parses the Storage Connection String into an associative Array (Map) of
// Settings.
func ParseConnectionParameters(
	parametersString string,
) (parameters map[string]string, err error) {
	parameters = make(map[string]string)
	if len(parametersString) == 0 {
		return
	}
	var kvPairs = strings.Split(parametersString, ConnectionParameterDelimiter)
	for _, kvPair := range kvPairs {
		parts := strings.Split(kvPair, ConnectionParameterKeyValueDelimiter)
		if len(parts) != 2 {
			err = fmt.Errorf(ErrfSyntaxErrorAt, kvPair)
			return
		}
		_, duplicateKey := parameters[parts[0]]
		if duplicateKey {
			err = fmt.Errorf(ErrfDuplicateKey, parts[0])
			return
		}
		parameters[parts[0]] = parts[1]
	}
	return
}
