// Copyright 2017 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package proxy

import (
	"github.com/dgraph-io/grpc-proxy/proxy/codec"
)

// ResponseModifier modifies the response before returning it to the client
type ResponseModifier func(*codec.Frame)
