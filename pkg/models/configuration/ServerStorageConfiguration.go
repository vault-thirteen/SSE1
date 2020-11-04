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
// File: github.com/vault-thirteen/SSE1/pkg/models/configuration/ServerStorageConfiguration.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package configuration

import (
	configuration "github.com/vault-thirteen/SSE1/pkg/models/configuration/xml"
	"github.com/vault-thirteen/SSE1/pkg/models/db/common"
)

type ServerStorageConfiguration struct {
	Type                  ServerStorageType
	CommonParameters      common.StorageConfiguration
	InitializationScripts ServerStorageIniScripts
	TableSettings         []common.TableSettings
	Time                  ServerStorageTimeConfiguration

	// Settings taken from Application's Settings.
	CoolDownPeriods       ServerStorageCoolDownPeriods
	IdleSessionTimeoutSec uint
	TokenLifeTimeSec      uint
}

func NewTableSettings(
	settings configuration.XmlServerStorageTableSettings,
) (result []common.TableSettings, err error) {
	result = make([]common.TableSettings, 0, len(settings.Table))
	for _, table := range settings.Table {
		columnNames := make([]string, 0, len(table.Column))
		for _, column := range table.Column {
			columnNames = append(columnNames, column.Name)
		}
		result = append(result, common.TableSettings{
			TableName:        table.Name,
			TableColumnNames: columnNames,
		})
	}
	return
}
