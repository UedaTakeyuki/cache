//
// Primitive Cache wiht have 2 keys
//
// Assumption:
// Cached object shoul'd have 1st key, but 2nd key is optional.
//
// Usage:
// Last used last deleted [string]interface{}.
//
// Copyright 2021 Aterier UEDA
// Author: Dr. Takeyuki UEDA

package cache

import (
	"log"
)

type BiKeyCache struct {
	maxSize       int
	body          map[string]interface{}
  bikeys        map[string]string      // map[2ndkeystr] 1stkeystr
  bikeysreverse map[string]string      // map[1stkeystr] 2ndkeystr
	fifo          []string
}

func NewBiKeyCache(maxSize int) (*Cache, error) {
	cache := Cache{} // initialize
	cache.maxSize = maxSize
	cache.body = map[string]interface{}{}
	return &cache, nil
}

/*
 * AddOrReplace : without 2nd key
 */
func (cache Cache) AddOrReplace(key string, entity interface{}) interface{} { // Add & Replace
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
		delete(cache.body, cache.fifo[0])
		cache.fifo = cache.fifo[1:]
	}
	// add (or replace) new one
	cache.body[key] = entity
	cache.fifo = append(cache.fifo, key)

	return entity
}

/*
 * AddOrReplaceWith2ndKey
 */
func (cache Cache) AddOrReplaceWith2ndKey(key1st string, key2nd string, entity interface{}) interface{} {
  // Add or Replace body
  cache.AddOrReplace(key1st, entity)
  
  // update bikeys
  key1stPrev, isExist := cache.bikeys[key2nd]
	if isExist {
    if key1stPrev != key1st {
      // replace key1st
      cache.bikeys[key2nd] = key1st
    }
  } else {
    // add key1st
    cache.bikeys[key2nd] = key1st
	}
  
  // update bikeysreverse
  key2ndPrev, isExist := cache.bikeysreverse[key1st]
	if isExist {
    if key2ndPrev != key2nd {
      // replace key2nd
      cache.bikeysreverse[key1st] = key2nd
    }
  } else {
    // add key2nd
    cache.bikeysreverse[key1st] = key2nd
	}

  return entity
}

/*
 * Get : by 1st key
 */
func (cache Cache) Get(key string) (result interface{}, isExist bool) {
	result, isExist = cache.body[key]
	if isExist {
		fmt.Println("cache hit!")
		// remove ex CacheOrder
		for i, id := range cache.fifo {
			if id == key {
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
 * GetBy2ndKey
 */
func (cache Cache) GetBy2ndKey(key2nd string) (result interface{}, isExist bool) {
  key1st, isExist := cache.bikeys[key2nd]
	if isExist {
    result, isExist = cache.Get(key1st)
    if !isExist {
      log.Println("bikey system broken!!!")
    }
	}
	return
}

/*
 * Delete : by 1st key
 */
func (cache Cache) Delete(key1st string) {
	// remove from CacheTable
	delete(cache.body, key1st)
	// remove from CacheOrder
	for i, id := range cache.fifo {
		if id == key1st {
			cache.fifo = append(cache.fifo[:i], cache.fifo[i+1:]...)
			break
		}
	}
  // remove bikeys
  key2nd, isExist := cache.bikeysreverse[key1st]
  if isExist {
    delete(cache.bikeys, key2nd)
    delete(cache.bikeysreverse, key1st)
  }
}

/*
 * DeleteBy2ndKey
 */
func (cache Cache) DeleteBy2ndKey(key2nd string) {
  key1st, isExist := cache.bikeys[key2nd]
	if isExist {
    cache.Delete(key1st)
    delete(cache.bikeys, key2nd)
    delete(cache.bikeysreverse, key1st)
	}
}

/*
 * Dump1stKeys : 
 */
func (cache Cache) Dump1stKeys() {
	log.Println("*** Dump 1st Cache Keys ***")
	for key1st, _ := range cache.body {
		log.Println(key1st)
	}
	log.Println("***********************")
}

/*
 * Dump2ndKeys : 
 */
func (cache Cache) Dump2ndKeys() {
	log.Println("*** Dump 2nd Cache Keys ***")
	for key2nd, _ := range cache.bikeys {
		log.Println(key2nd)
	}
	log.Println("***********************")
}
