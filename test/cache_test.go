package cache

import (
	"log"
	"testing"

	cp "github.com/UedaTakeyuki/compare"
	"local.packages/cache"
)

// AddOrReplace, Get
func Test_01(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	c, err := cache.NewCache(3, true)
	cp.Compare(t, err, nil)
	{
		val := c.AddOrReplace(1, "a")
		cp.Compare(t, val, "a")
	}
	c.AddOrReplace(2, "b")
	c.AddOrReplace(3, "c")
	{
		val, exist := c.Get(1)
		cp.Compare(t, val, "a")
		cp.Compare(t, exist, true)
	}
	c.AddOrReplace(4, "d")
	{
		_, exist := c.Get(1)
		// should be exist
		cp.Compare(t, exist, true)
		// should not be exist, already deleted.
		_, exist = c.Get(2)
		cp.Compare(t, exist, false)
		val, _ := c.Get(4)
		cp.Compare(t, val, "d")
	}
	c.AddOrReplace(5, "e")
}

// getNextFunc
func Test_02(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	c, err := cache.NewCache(3, true)
	cp.Compare(t, err, nil)
	{
		val := c.AddOrReplace(1, "a")
		cp.Compare(t, val, "a")
	}
	{
		val := c.AddOrReplace(2, "b")
		cp.Compare(t, val, "b")
	}
	{
		val := c.AddOrReplace(3, "c")
		cp.Compare(t, val, "c")
	}
	getNext := c.GetNextFunc()
	cp.Compare(t, getNext(), "a")
	cp.Compare(t, getNext(), "b")
	cp.Compare(t, getNext(), "c")
	cp.Compare(t, getNext(), nil)
}
