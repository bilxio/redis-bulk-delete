/*
* Copyright 2017 bilxio
*
* @File: main.go
* @Author: billxiong
* @Date:   2017-11-08 11:55:12
* @Last Modified by:   Bill Xiong
* @Last Modified time: 2017-11-09 19:23:24
*/

package main

import (
	"flag"
	"strings"
	"time"

	"github.com/bilxio/redis-go-cluster"
)

var (
	f_redis_nodes = flag.String("redis-nodes", "", "redis nodes' address and port, separeted by comma. <host>:<port>")
	f_verbose = flag.Bool("verbose", false, "verbose log")
	f_print_slot = flag.Bool("print-slot", false, "print slot for keys")

	f_error_log = flag.String("error-log", "", "error log file")
	f_log = flag.String("log", "", "log file")
)

func main() {
	flag.Parse()

	mylog := newLogger(*f_log, *f_error_log)
	defer mylog.Close()

	mylog.Printf("hahaha, go ! \n")

	if *f_redis_nodes == "" {
		mylog.Fatalln("redis-nodes required.")
	}

	nodes := strings.Split(*f_redis_nodes, ",")

	if *f_verbose {
		mylog.Println("try to connect cluster ", nodes)
	}

	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes: nodes,
			ConnTimeout: 5 * 1000 * time.Millisecond,
			ReadTimeout:  60 * 1000 * time.Millisecond,
			WriteTimeout: 60 * 1000 * time.Millisecond,
			KeepAlive: 16,
			AliveTime: 300 * time.Second,
			})

	if err != nil {
		mylog.Fatalln(err)
	}

	defer cluster.Close()

	if len(flag.Args()) < 1 {
		mylog.Fatalln("args length must be greater then 1")
	}

	cmd := flag.Args()[:1]
	args := flag.Args()[1:]

	if *f_verbose {
		mylog.Println("cmd = ", cmd, " , len(args) = ", len(args) ," , args = ", args)
		if *f_print_slot {
			slots := make([]uint16, len(args))
			args_lens := make([]int, len(args))
			for i := range args {
				slots[i] = hash(args[i])
				args_lens[i] = len(args[i])
			}
			mylog.Println("cmd = ", cmd, " , lens = ", args_lens,  " , slots = ", slots)
		}
	}

	// 复制args
	cmd_args := make([]interface{}, len(args))
	for i := range args {
		cmd_args[i] = strings.Trim(args[i], " \r\n")
	}

	reply, err := cluster.Do(cmd[0], cmd_args...) // 数组打散，深坑

	if err != nil {
		mylog.Errorln("ERROR @@ ", cmd, " @@ ", strings.Join(args, "][")) // 记录错误的 keys
		mylog.Fatalln(err)
	}

	mylog.Println(reply)
}
