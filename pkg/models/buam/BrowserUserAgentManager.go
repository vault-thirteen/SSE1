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
// File: github.com/vault-thirteen/SSE1/pkg/models/buam/BrowserUserAgentManager.go.
// Last Modification Time: 2020-11-04.
//
////////////////////////////////////////////////////////////////////////////////

package buam

import (
	"errors"

	"github.com/vault-thirteen/FixedSizeBubbleCache"
	"github.com/vault-thirteen/SSE1/pkg/interfaces/storage"
)

// Errors.
const (
	ErrStorageIsNotSet = "Storage is not set"
	ErrTypeCasting     = "Type Casting Failure"
)

// A Manager working with Browsers' 'User Agent' Fields.
type BrowserUserAgentManager struct {
	storage storage.IStorage
	cache   *fsbcache.FixedSizeBubbleCache
}

// Manager's Constructor.
func NewBrowserUserAgentManager(
	storage storage.IStorage,
) (buam *BrowserUserAgentManager, err error) {
	if storage == nil {
		err = errors.New(ErrStorageIsNotSet)
		return
	}
	buam = &BrowserUserAgentManager{
		storage: storage,
		cache:   fsbcache.NewFixedSizeBubbleCache(100, 1*60),
	}
	return
}

// Returns an Id of textual Representation of a 'User Agent' Setting.
// Uses the in-Memory Cache and the external Storage.
func (buam *BrowserUserAgentManager) GetBrowserUserAgentId(
	browserUserAgentName string,
) (id uint, err error) {

	// Try to get an Id from the Cache.
	var idAsInterface interface{}
	idAsInterface, err = buam.cache.GetActualRecordDataByUID(browserUserAgentName)
	if err == nil {
		var ok bool
		id, ok = idAsInterface.(uint)
		if !ok {
			err = errors.New(ErrTypeCasting)
			return
		}
		return
	}

	// Search in the Database and update the Cache.
	id, err = buam.storage.GetBrowserUserAgentId(browserUserAgentName)
	if err != nil {
		return
	}
	err = buam.cache.AddRecord(&fsbcache.FixedSizeBubbleCacheRecord{
		UID:  browserUserAgentName,
		Data: id,
	})
	if err != nil {
		return
	}
	return
}
