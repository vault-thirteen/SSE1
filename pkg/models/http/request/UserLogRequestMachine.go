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
// File: github.com/vault-thirteen/SSE1/pkg/models/http/request/UserLogRequestMachine.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package request

import (
	"errors"
	"fmt"
	"net"
)

// Settings.
const (
	MachineHostLengthMin             = 1
	MachineHostLengthMax             = 1024
	MachineBrowserUserAgentLengthMin = 3
	MachineBrowserUserAgentLengthMax = 4096
	NetworkIPType                    = "ip"
)

// Errors.
const (
	ErrfMachineIPAddress       = "Error in Machine.Host: '%v'."
	ErrMachineBrowserUserAgent = "Error in Machine.BrowserUserAgent."
)

type UserLogRequestMachine struct {
	Host             string
	BrowserUserAgent UserLogRequestMachineBrowserUserAgent
}

func ValidateMachineHost(
	host string,
) (err error) {
	if (len(host) < MachineHostLengthMin) ||
		(len(host) > MachineHostLengthMax) {
		err = fmt.Errorf(ErrfMachineIPAddress, host)
		return
	}
	_, err = net.ResolveIPAddr(NetworkIPType, host)
	if err != nil {
		return
	}
	return
}

func ValidateMachineBrowserUserAgent(
	bua string,
) (err error) {
	if (len(bua) < MachineBrowserUserAgentLengthMin) ||
		(len(bua) > MachineBrowserUserAgentLengthMax) {
		err = errors.New(ErrMachineBrowserUserAgent)
		return
	}
	return
}
