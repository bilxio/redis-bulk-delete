# xargs 命令使用 (仅适用 MacOS)

## 参数详解

- `-t` 执行之前打印命令

	cat keys|xargs -I @@ -t ./redis-bulk-delete -verbose f @@ r

- `-L` 指定一次命令执行使用参数个数

	$ cat keys|xargs -I @@ -t -L2 ./redis-bulk-delete -verbose f @@ r
 	./redis-bulk-delete -verbose f a b r

- 自定义占位符替换 `@@`

	cat keys|xargs -I @@ ./redis-bulk-delete -verbose f @@ r

- `-R` 指定替换次数

	$ cat keys|xargs -I @@ -R 1 -L2 ./redis-bulk-delete -verbose f @@ r @@
	./redis-bulk-delete -verbose f a b r @@

	$ cat keys|xargs -I @@ -R 2 -L2 ./redis-bulk-delete -verbose f @@ r @@
	./redis-bulk-delete -verbose f a b r a b

- `-P` 最大进程数

- `-s` 字符数

## Docker Run

	docker run --rm -v $(pwd):/app ubuntu:xenial bash -c "head /app/keys| xargs -t -i@@ -n2 /app/redis-bulk-delete-amd64 -verbose 00 @@ 99"
	docker run --rm -v $(pwd):/app ubuntu:xenial bash -c "head /app/keys| xargs -t -n2 /app/redis-bulk-delete-amd64 -verbose 00 99"
	docker run --rm -v $(pwd):/app ubuntu:xenial xargs --help
