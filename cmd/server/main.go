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
// File: github.com/vault-thirteen/SSE1/cmd/server/main.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"log"
	"os"

	"github.com/vault-thirteen/SSE1/pkg/models/app"
	"github.com/vault-thirteen/SSE1/pkg/models/configuration"
)

// Messages.
const (
	MsgCommandLineArguments     = "Command Line Arguments"
	MsgApplicationConfiguration = "Application Configuration"
	MsgApplication              = "Application"
)

// OS Exit Codes.
const (
	OsExitCodeError = 1
)

func main() {
	var err error
	var cla CommandLineArguments
	cla, err = getCommandLineArguments()
	mustBeNoError(MsgCommandLineArguments, err)
	var appConfiguration *configuration.AppConfiguration
	appConfiguration, err = configuration.NewAppConfiguration(cla.PathToConfigurationFile)
	mustBeNoError(MsgApplicationConfiguration, err)
	var application *app.Application
	application, err = app.NewApplication(appConfiguration)
	mustBeNoError(MsgApplication, err)
	err = application.Run()
	checkError(err)
}

func getCommandLineArguments() (cla CommandLineArguments, err error) {
	flag.StringVar(
		&cla.PathToConfigurationFile,
		PathToConfigurationFileArgName,
		PathToConfigurationFileDefaultValue,
		PathToConfigurationFileUsageHint,
	)
	flag.Parse()
	return
}

func mustBeNoError(
	errorPrefix string,
	err error,
) {
	const SymbolColon = ":"
	if err != nil {
		log.Println(errorPrefix+SymbolColon, err)
		os.Exit(OsExitCodeError)
	}
}

func checkError(
	err error,
) {
	if err != nil {
		log.Println(err)
	}
}
