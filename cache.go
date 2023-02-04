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
}

func NewCache(maxSize int) (*Cache, error) {
	cache := Cache{} // initialize
	cache.maxSize = maxSize
	cache.body = map[interface{}]interface{}{}
	cache.fifo = make([]interface{}, maxSize)
	return &cache, nil
}

/*
 * AddOrReplace
 */
func (cache Cache) AddOrReplace(key interface{}, entity interface{}) interface{} { // Add & Replace
	_, isExist := cache.body[key]
	if isExist {
		// remove ex CacheOrder
		for i, id := range cache.fifo {
			if id == key {
				cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
				break
			}
		}
	} else if len(cache.body) > cache.maxSize {
		// delete oldest
		log.Println("len(cache.fifo)", cache.fifo)
		delete(cache.body, cache.fifo[0])
		cache.fifo = cache.fifo[1:]
	}
	// add (or replace) new one
	cache.body[key] = entity
	cache.fifo = append(cache.fifo, key)
	log.Println("len(cache.body)", cache.body)
	log.Println("len(cache.fifo)", cache.fifo)

	return entity
}

/*
 * Get
 */
func (cache Cache) Get(key interface{}) (result interface{}, isExist bool) {
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
	return
}

/*
 * Delete
 */
func (cache Cache) Delete(key interface{}) {
	// remove from CacheTable
	delete(cache.body, key)
	// remove from CacheOrder
	for i, id := range cache.fifo {
		if id == key {
			cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
			break
		}
	}
}

/*
 * DumpKeys
 */
func (cache Cache) DumpKeys() {
	fmt.Println("*** Dump Cache Keys ***")
	for key, _ := range cache.body {
		fmt.Println(key)
	}
	fmt.Println("***********************")
}
