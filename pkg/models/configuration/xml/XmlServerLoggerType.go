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
// File: github.com/vault-thirteen/SSE1/pkg/models/configuration/xml/XmlServerLoggerType.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package configuration

import "strings"

const (
	ServerLoggerTypeUnknown = 0
	ServerLoggerTypeBuiltIn = 1
)

const (
	ServerLoggerTypeAliasBuiltIn = "built-in"
)

type ServerLoggerType uint8

func NewServerLoggerType(
	loggerType string,
) ServerLoggerType {
	loggerType = strings.ToLower(loggerType)
	switch loggerType {
	case ServerLoggerTypeAliasBuiltIn:
		return ServerLoggerTypeBuiltIn
	}
	return ServerLoggerTypeUnknown
}

func (slt ServerLoggerType) IsValid() bool {
	switch slt {
	case ServerLoggerTypeUnknown:
		return false
	case ServerLoggerTypeBuiltIn:
		return true
	}
	return false
}
