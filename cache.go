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
				log.Println("cache.fifo", cache.fifo)
				// get rid of cache.fifo[i]
				cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
				log.Println("cache.fifo", cache.fifo)
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
	log.Println("cache.fifo", cache.fifo)
	cache.fifo = append(cache.fifo, key)
	log.Println("cache.fifo", cache.fifo)

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
				log.Println("cache.fifo", cache.fifo)
				cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
				log.Println("cache.fifo", cache.fifo)
				break
			}
		}
		// add bottom CacheOrder
		log.Println("cache.fifo", cache.fifo)
		cache.fifo = append(cache.fifo, key)
		log.Println("cache.fifo", cache.fifo)
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
			log.Println("cache.fifo", cache.fifo)
			// get rid of cache.fifo[i]
			cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
			log.Println("cache.fifo", cache.fifo)
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
