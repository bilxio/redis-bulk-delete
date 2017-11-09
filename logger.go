/*
* Copyright 2017 bilxio
*
* @File: logger.go
* @Author: billxiong
* @Date:   2017-11-09 17:34:39
* @Last Modified by:   Bill Xiong
* @Last Modified time: 2017-11-09 19:23:32
*/

package main

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})

	Printf(format string, v ...interface{})
	Println(v ...interface{})

	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})

	Close()
}

type closeHandler func()

type logger struct {
	l *log.Logger // default logger
	e *log.Logger // error logger

	// 关闭处理
	closeFuncs []closeHandler
}

func newLogger(logFile string, errFile string) Logger {
	log1 := log.New(os.Stdout, "LOG ", log.Ldate | log.Ltime | log.Lshortfile)
	log2 := log.New(os.Stderr, "ERROR ", log.Ldate | log.Ltime | log.Lshortfile)

	log_ := &logger{
		l: log1,
		e: log2,
		closeFuncs: make([]closeHandler, 0),
	}

	// write logs to File
	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("file open error : %v", err)
		}
		log_.l.SetOutput(f)
		log_.closeFuncs = append(log_.closeFuncs, func() {
			f.Close()
		})
	}

	// write errs to File
	if errFile != "" {
		f, err := os.OpenFile(errFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("file open error : %v", err)
		}
		log_.e.SetOutput(f)
		log_.closeFuncs = append(log_.closeFuncs, func() {
			f.Close()
		})
	}

	return log_
}

// 关闭处理
func (l *logger) Close() {
	for i := range l.closeFuncs {
		if l.closeFuncs[i] != nil {
			l.closeFuncs[i]()
		}
	}
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	s := fmt.Sprintf(format , v...)
	l.e.Output(2, s)
	os.Exit(1)
}

func (l *logger) Fatalln(v ...interface{}) {
	l.e.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	s := fmt.Sprintf(format , v...)
	l.e.Output(2, s)
}

func (l *logger) Errorln(v ...interface{}) {
	l.e.Output(2, fmt.Sprint(v...))
}

func (l *logger) Printf(format string, v ...interface{}) {
	s := fmt.Sprintf(format , v...)
	l.l.Output(2, s)
}

func (l *logger) Println(v ...interface{}) {
	l.l.Output(2, fmt.Sprint(v...))
}
