# redis-bulk-delete

Bulk Delete Keys from Redis Cluster by concurrent goroutines.

## build

```
go build -o redis-bulk-delete
```

or for Linux x64

```
GOOS=linux GOARCH=amd64 go build -o redis-bulk-delete-amd64
```

## run

Basic usage

```
$ cat keys | xargs -P6 -L100 ./redis-bulk-delete -redis-nodes=127.0.0.1:16380 MDEL
```

Append errors to a file

```
$ cat keys | xargs -P6 -L100 ./redis-bulk-delete -redis-nodes=127.0.0.1:16380 --error-log=foo.err MDEL
```

## thanks

- `xargs`
	Thanks to `xargs` cmdline tool, it's great.

- `redis-go-cluster`
	Thanks to [redis-go-cluster](http://github.com/chasex/redis-go-cluster). It's a nice redis client to handle cluster operations.
	I forked it (see [redis-go-cluster](http://github.com/bilxio/redis-go-cluster)), and wrote a `MDEL` command to adapter the cluster situation.
