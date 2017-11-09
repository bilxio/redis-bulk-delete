# @Author: billxiong
# @Date:   2017-11-08 18:42:32
# @Last Modified by:   billxiong
# @Last Modified time: 2017-11-08 18:45:15

cat << EOF | xargs -L1 -t redis-cli -c -p 16380
INCR key1
INCR key2
INCR key3
INCR key4
INCR key5
EOF
