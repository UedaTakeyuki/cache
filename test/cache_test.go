package cache

import (
	"log"
	"testing"

	cp "github.com/UedaTakeyuki/compare"
	"local.packages/cache"
)

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
