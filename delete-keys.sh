# @Author: billxiong
# @Date:   2017-11-08 18:45:48
# @Last Modified by:   billxiong
# @Last Modified time: 2017-11-08 18:46:06

cat << EOF | xargs -L1 -t redis-cli -c -p 16380
DEL key1
DEL key2
DEL key3
DEL key4
DEL key5
EOF
