// Copyright 2017 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package proxy

// ResponseModifier modifies the response before returning it to the client
type ResponseModifier func()
