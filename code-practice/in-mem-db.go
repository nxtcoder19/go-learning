package main

//
//import (
//	"fmt"
//	"sort"
//)
//
//Based on the images you provided, your task is to implement an in-memory database with the following functionalities for Level 1:
//
//*SetOrInc(key string, field string, value int) int
//
//Inserts a field-value pair into the record associated with key.
//If the field already exists, it increments the existing value.
//If the record or field does not exist, it creates a new one.
//Returns the updated value.
//*Get(key string, field string) int
//
//Returns the value of the field in the record associated with key.
//If the record or field does not exist, return nil.
//Delete(key string, field string) bool
//
//Removes a field from the record associated with key.
//If the field was deleted, return true; otherwise, return false.
//If all fields in a record are deleted, remove the record.
//
//func (db *InMemoryDBImpl) SetOrInc(key string, field string, value int) *int {
//	db.mu.Lock()
//	defer db.mu.Unlock() // Ensure unlock is always called
//
//	// Ensure key exists
//	if _, exists := db.store[key]; !exists {
//		db.store[key] = make(map[string]int)
//	}
//
//	// Increment value
//	db.store[key][field] += value
//	result := db.store[key][field]
//	return &result
//}
//
//func (db *InMemoryDBImpl) Get(key string, field string) *int {
//	db.mu.RLock()
//	defer db.mu.RUnlock()
//
//	if record, exists := db.store[key]; exists {
//		if value, found := record[field]; found {
//			return &db.store[key][field] // ✅ Safe return
//		}
//	}
//	return nil // ✅ Return nil safely
//}
//
//func (db *InMemoryDBImpl) Delete(key string, field string) bool {
//	db.mu.Lock()
//	defer db.mu.Unlock()
//
//	if record, exists := db.store[key]; exists {
//		if _, found := record[field]; found {
//			delete(record, field)
//
//			// If the record is empty, remove the entire key
//			if len(record) == 0 {
//				delete(db.store, key)
//			}
//			return true
//		}
//	}
//
//	return false
//}
//
//Implementation Plan
//Ensure Basic Operations Work Properly:
//
//SetOrInc(key, field, value int): Insert or update values.
//Get(key, field): Retrieve values.
//Delete(key, field): Remove values and clear records if needed.
//Track Modification Count:
//
//Maintain a map that stores the number of updates per key.
//Whenever a key is modified (added, updated, or deleted), increment its counter.
//If a key is fully deleted, remove it from the modification counter.
//Implement TopNKeys(n int) []string:
//
//Sort records in descending order of modification count.
//If two records have the same modification count, sort them lexicographically.
//Format output as ["<key1>(<count1>)", "<key2>(<count2>)", ...].
//
//func (a *AbstractInMemoryDB) TopNKeys(n int) []string {
//	if len(a.modificationCount) == 0 {
//		return []string{}
//	}
//
//	type KeyMod struct {
//		Key        string
//		ModCount   int
//	}
//
//	var sortedKeys []KeyMod
//	for key, count := range a.modificationCount {
//		sortedKeys = append(sortedKeys, KeyMod{Key: key, ModCount: count})
//	}
//
//	// Sort by modification count (descending), then by lexicographical order
//	sort.Slice(sortedKeys, func(i, j int) bool {
//		if sortedKeys[i].ModCount == sortedKeys[j].ModCount {
//			return sortedKeys[i].Key < sortedKeys[j].Key
//		}
//		return sortedKeys[i].ModCount > sortedKeys[j].ModCount
//	})
//
//	// Prepare result list
//	result := []string{}
//	for i := 0; i < n && i < len(sortedKeys); i++ {
//		result = append(result, fmt.Sprintf("%s(%d)", sortedKeys[i].Key, sortedKeys[i].ModCount))
//	}
//
//	return result
//}
