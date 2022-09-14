/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start

package isAnagram

import (
    "fmt"
)

func isAnagram(s string, t string) bool {
    
    smap := make(map[string]int)
    tmap := make(map[string]int)
    
    schar := []rune(s)
    tchar := []rune(t)
    
    if len(schar)!= len(tchar) {
        return false
    }
    
    for _,v := range schar {
        smap[string(v)]+=1
    }
    
    for _,v := range tchar {
        tmap[string(v)]+=1
    }
    
    for k,_ := range smap {
        if smap[k] != tmap[k]  {
            return false
        }
    }
    return true
}