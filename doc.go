/*
* Copyright 2017 bilxio
*
* @File: doc.go
* @Author: billxiong
* @Date:   2017-11-09 18:53:33
* @Last Modified by:   Bill Xiong
* @Last Modified time: 2017-11-09 19:23:29
*/

/*
Bulk delete keys from a redis cluster. It depends on a special redis lib `bilxio/redis-go-cluster`.


Bulk Delete

	cat keys | xargs -P6 -L100 ./redis-bulk-delete -redis-nodes=127.0.0.1:16380 MDEL


Bulk Delete and write error to file

	cat keys | xargs -P6 -L100 ./redis-bulk-delete -redis-nodes=127.0.0.1:16380 --error-log=foo.err MDEL
*/
package main
