//
// Primitive Cach feature
//
// Last used last deleted [string]interface{}.
//
// Copyright 2020 Aterier UEDA
// Author: Takeyuki UEDA

package cache

import (
	"fmt"
	"log"
)

type Cache struct {
	maxSize int
	body    map[interface{}]interface{}
	fifo    []interface{}
	debug   bool
}

func NewCache(maxSize int, debug bool) (*Cache, error) {
	cache := Cache{} // initialize
	cache.maxSize = maxSize
	cache.body = map[interface{}]interface{}{}
	cache.debug = debug
	return &cache, nil
}

/*
 * AddOrReplace
 */
func (cache *Cache) AddOrReplace(key interface{}, entity interface{}) interface{} { // Add & Replace
	_, isExist := cache.body[key]
	if isExist {
		// remove ex CacheOrder
		for i, id := range cache.fifo {
			if id == key {
				// get rid of cache.fifo[i]
				cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
				break
			}
		}
	} else if len(cache.body) > cache.maxSize {
		// delete oldest
		delete(cache.body, cache.fifo[0])
		cache.fifo = cache.fifo[1:]
	}
	// add (or replace) new one
	cache.body[key] = entity
	cache.fifo = append(cache.fifo, key)

	if cache.debug {
		cache.DumpBody()
		cache.DumpFifo()
	}

	return entity
}

/*
 * Get
 */
func (cache *Cache) Get(key interface{}) (result interface{}, isExist bool) {
	result, isExist = cache.body[key]
	if isExist {
		fmt.Println("cache hit!")
		// remove ex CacheOrder
		for i, id := range cache.fifo {
			if id == key {
				// get rid of cache[i]
				cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
				break
			}
		}
		// add bottom CacheOrder
		cache.fifo = append(cache.fifo, key)
	}

	if cache.debug {
		cache.DumpBody()
		cache.DumpFifo()
	}

	return
}

/*
 * Delete
 */
func (cache *Cache) Delete(key interface{}) {
	// remove from CacheTable
	delete(cache.body, key)
	// remove from CacheOrder
	for i, id := range cache.fifo {
		if id == key {
			// get rid of cache.fifo[i]
			cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
			break
		}
	}

	if cache.debug {
		cache.DumpBody()
		cache.DumpFifo()
	}

	return
}

/*
 * DumpKeys
 */
func (cache *Cache) DumpKeys() {
	log.Println("*** Dump Cache Keys ***")
	for key, _ := range cache.body {
		log.Println(key)
	}
	log.Println("***********************")
}

/*
 * DumpBody
 */
func (cache *Cache) DumpBody() {
	log.Println("len(cache.body)", len(cache.body))
	log.Println("cache.body", cache.body)
}

/*
 * DumpFifo
 */
func (cache *Cache) DumpFifo() {
	log.Println("len(cache.fifo)", len(cache.fifo))
	log.Println("cache.fifo", cache.fifo)
}
