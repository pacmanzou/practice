#!/usr/bin/env bash
#
# Description: demo.sh test
# Time: 08:42:35 2020-10-23
# a=$((1+2))
# echo $a
# echo "hello $a"
# echo "$(echo "scale=2; 1*1.5*2-1.25" | bc)%"
# echo $((100/3))
# if [ "$1" -eq "$2" ]; then
#     echo "$1 = $2"
# else
#     echo "$1 != $2"
# fi
# if [[ "$1" -eq "$2" ]]; then
#     echo "$1 = $2"
# fi
# Countdown
# echo -n "Countdown: "
# for (( i = 9; i >= 0; i-- )); do
#     echo -e -n "$i\b"
#     sleep 1
# done
# echo ""
# read
# echo -n -e "username: "
# read -r name
# echo -n -e "password: "
# read -r -s password
# echo ""
# echo "$name $password"
# read -r -p "username: " name
# read -r -s -p "password: " password
# echo
# echo "$name $password"
# echo -e "System: $(uname --help)\n$(uname -a)"
# ARRAY1=('a' 'b' 'c' 'd')
# echo "${ARRAY1[1]}"
# ARRAY1[5]='e'
# echo "${ARRAY1[@]}"
# if ping -c 1 -w 1 "$1" &> /dev/null; then
#     echo "success"
# else
#     echo "failure"
# fi
# my_num=0
# while true; do
#     echo "${my_num}"
#     [[ $((my_num++)) == 10 ]] && break
# done
# while true; do
#     read -r -p "请输入1-4:" NUM
# case "${NUM}" in
#     1)
#         echo 'a'
#         ;;
#     2)
#         echo 'b'
#         ;;
#     3)
#         echo 'c'
#         ;;
#     4)
#         echo 'd'
#         ;;
#     *)
#         break
#         ;;
# esac
# done
# do_something() {
#     echo "hello world"
# }
# my_var=""
# if [[ -z "${my_var}" ]]; then
#     do_something
# fi
# set -e
# if ! mv abc def; then
#     echo "error"
# fi

# sed
# sed '2,4a\*********************' demo.txt
# sed '2i\*********************' demo.txt
# sed '1i\# This is a comment' demo.txt
# sed '$a\# This is a comment' demo.txt
# sed '2,4d' demo.txt
# sed '2,$d' demo.txt
# sed '/^$/d' demo.txt
# sed '2c\*********' demo.txt
# sed 's/^.*hello.*$/substitute/g' demo.txt
# sed -n '2,3p' demo.txt
# sed -n '/hello/p' demo.txt
# sed 's/\.$/\!/g' demo.txt
# sed 'G' demo.txt

# cut
# cut -d ' ' -f 2,3 demo1.txt
# cut -d ' ' -f 2- demo1.txt
# cut -d ' ' -f -2 demo1.txt
# cut -c -2 demo1.txt
# cut -c 2- demo1.txt
# grep "shell" demo1.txt | cut -d " " -f 2

# awk
# default space
# echo "abc 123 456" | awk '{print $1"&"$2"&"$3}'
# awk '/root/{print $0}' passwd
# awk  -F: '/root/{print $0}' passwd
# awk -F: '{print "filename:"FILENAME",row:"NR",column:"NF",content:"$0" "}' passwd
# awk -F: '{printf("filename:%s,row:%s,column:%s,content:%s\n",FILENAME,NR,NF,$0)}' passwd
# 空行行号
# awk '/^$/{print NR}' practice.sh
# show ip
# ip addr show dev wlp2s0 \
#     | grep 'inet '  \
#     | awk '{print $2}' \
#     | awk -F "/" '{print $1}'
# ip addr show dev wlp2s0 \
#     | grep 'link/ether' -A 1 \
#     | tail -n 1 \
#     | cut -f 1 -d '/' \
#     | awk '{print $2}'
# ip addr \
#     | grep 'link/ether' -A 1 \
#     | awk 'NR==2{print $2}' \
#     | awk -F '/' '{print $1}'
# ifconfig wlp2s0 | grep -E "inet " | awk '{print $2}'

# sort
# sort -t " " -nrk 2 demo2.txt
# sort demo2.txt

sed -n '/\[file-build-and-run\]/,/\[.*\]/p' tasks.ini | head -n -1

# awk 脚本
# /\[file-build-and-run\]/{
# # index($0, "[file-build-and-run]"){
#     print
#     while( (getline var) > 0 ){
#         if(var ~ /\[.*\]/){
#             exit
#         }
#         print var
#     }
# }

# getline返回值:
# >0: 表示读取到数据
# =0: 表示遇到结尾EOF,没有读到数据
# <0: 表示读取报错

# sum
awk '{arr[$0]++}END{for(i in arr){print i,arr[i]}}' repeat.txt
