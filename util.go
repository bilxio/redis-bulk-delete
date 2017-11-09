/*
* Copyright 2017 bilxio
*
* @File: util.go
* @Author: billxiong
* @Date:   2017-11-08 16:39:49
* @Last Modified by:   Bill Xiong
* @Last Modified time: 2017-11-09 19:23:23
*/

package main

import (
	"fmt"
	"strconv"
)

const (
	kClusterSlots	= 16384
)

func key(arg interface{}) (string, error) {
	switch arg := arg.(type) {
	case int:
		return strconv.Itoa(arg), nil
	case int64:
		return strconv.Itoa(int(arg)), nil
	case float64:
		return strconv.FormatFloat(arg, 'g', -1, 64), nil
	case string:
		return arg, nil
	case []byte:
		return string(arg), nil
	default:
		return "", fmt.Errorf("key: unknown type %T", arg)
	}
}

func hash(key string) uint16 {
	var s, e int
	for s = 0; s < len(key); s++ {
		if key[s] == '{' {
			break
		}
	}

	if s == len(key) {
		return crc16(key) & (kClusterSlots-1)
	}

	for e = s+1; e < len(key); e++ {
		if key[e] == '}' {
			break
		}
	}

	if e == len(key) || e == s+1 {
		return crc16(key) & (kClusterSlots-1)
	}

	return crc16(key[s+1:e]) & (kClusterSlots-1)
}
