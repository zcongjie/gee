package cache

import "github.com/ijunyu/gee/cache/cachepb"

type PeerGetter interface {
	Get(in *cachepb.Request, out *cachepb.Response) error
}

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}
