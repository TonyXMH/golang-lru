package unsafelru

type LRUCache interface {
	Add(key, val interface{}) bool
	Get(key interface{}) (val interface{}, ok bool)
	Contains(key interface{}) (ok bool)
	Peek(key interface{}) (val interface{}, ok bool)
	Remove(key interface{}) bool
	RemoveOldest() (key, val interface{}, ok bool)
	GetOldest() (key, val interface{}, ok bool)
	Keys() []interface{}
	Len() int
	Purge()
	Resize(int) int
}
