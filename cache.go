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
	"time"
)

type fifoElmType struct {
	lastUpdated int64
	id          interface{}
}

type Cache struct {
	maxSize int
	body    map[interface{}]interface{}
	fifo    []fifoElmType
	//	fifo    []interface{}
	debug bool
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
		for i, fifoElm := range cache.fifo {
			if fifoElm.id == key {
				// get rid of cache.fifo[i]
				cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
				break
			}
		}
	} else if len(cache.body) >= cache.maxSize {
		// delete oldest
		delete(cache.body, cache.fifo[0].id)
		cache.fifo = cache.fifo[1:]
	}
	// add (or replace) new one
	cache.body[key] = entity
	cache.fifo = append(cache.fifo, makeFifoElm(key))

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
		for i, fifoElm := range cache.fifo {
			if fifoElm.id == key {
				// get rid of cache[i]
				cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
				break
			}
		}
		// add bottom CacheOrder
		cache.fifo = append(cache.fifo, makeFifoElm(key))
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
	for i, fifoElm := range cache.fifo {
		if fifoElm.id == key {
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

// destructive, only use for read, DONT WRITE
/*
func (cache *Cache) PullValueMap() (valueMap map[interface{}]interface{}) {
	return cache.body
}
*/

// https://zenn.dev/toriwasa/articles/c7428879d624cd
func (cache *Cache) GetNextFunc() func() interface{} {
	i := -1

	return func() interface{} {
		i++
		if i < len(cache.body) {
			value, _ := cache.body[cache.fifo[i].id]
			log.Println("value", value)
			return value
		} else {
			return nil
		}
	}
}

/*
 * mskr fifoElm
 */
func makeFifoElm(key interface{}) fifoElmType {
	return fifoElmType{id: key, lastUpdated: time.Now().Unix()}
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
