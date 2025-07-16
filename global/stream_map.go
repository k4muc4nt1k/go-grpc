package global

import (
	pb "mbooke/grpc-server/proto/message"
	"sync"
)

// StreamServerMap is a global map for pb.StreamService_FetchResponseServer
var StreamServerMap = struct {
	M map[string]pb.StreamService_FetchResponseServer
	sync.RWMutex
}{
	M: make(map[string]pb.StreamService_FetchResponseServer),
}
