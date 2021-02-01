# 求1+2+3+...+n的值
def sum(num):
    nums = 0
    for i in range(1, num + 1):
        nums += i
    return nums


print(sum(5))


# 计算 staticmethod
class Calc(object):
    @staticmethod
    def add(a, b):
        return a + b

    @staticmethod
    def minus(a, b):
        return a - b

    @staticmethod
    def mul(a, b):
        return a * b

    @staticmethod
    def div(a, b):
        return a // b


# print(Calc.add(6, 2))
# print(Calc.minus(6, 2))
# print(Calc.mul(6, 2))
# print(Calc.div(6, 2))

try:
    x = Calc.div(5, 0)
    print("不会打印")
except Exception as e:
    print("程序error")
else:
    print(x)


# 计算
def Calc1(a, b, fn):
    return fn(a, b)


print(Calc1(6, 2, lambda x, y: x + y))
print(Calc1(6, 2, lambda x, y: x - y))
print(Calc1(6, 2, lambda x, y: x * y))
print(Calc1(6, 2, lambda x, y: x // y))


# 求多个数的最大值
def max_num(*args):
    # max = args[0]
    # for i in args:
    #     if i > max:
    #         max = i
    return max(args)


print(max_num(21, 243, 35, 7, 3, 46, 78, 8))
print(max_num("1", "a", "lsf"))
print(max_num("slfjdlf"))

# 打印n个骰子的点数之和
import random


def get_sum(num):
    sum = 0
    for i in range(1, num + 1):
        n = random.randint(1, 6)
        sum += n
        print(i, n)
    print("和是{}".format(sum))


get_sum(3)

# 交换字典的key和value，使用字典推导式
dict1 = {"name": "zou", "age": 22}
print({v: k for k, v in dict1.items()})


# 提取字符串中的所有字母
def alpah_extract(strings: str):
    string_tmp = ""
    for cha in strings:
        if cha.isalpha():
            string_tmp += cha
    return string_tmp


print(alpah_extract("lsssf$$^%..25k"))


# 多个数的平均值
def ave_get(*args: int):
    sum = 0
    for arg in args:
        sum += arg
    return sum // len(args)


print(ave_get(1, 2, 3, 4, 5, 6, 7, 8, 9))

# 求阶乘 导入math模块和自己写函数
import math


def fac_get(num: int):
    product = 1
    for n in range(1, num + 1):
        product *= n
    return product


print(math.factorial(5))
print(fac_get(5))


# 字符串首字母变大写
def cap_get(strings: str):
    if "z" >= strings[0] >= "a":
        s = strings[1:]
        return strings[0].upper() + s
    return strings


s1 = "klsjf"
print(s1.capitalize())
print(cap_get(s1))


# 判断字符串是否是以指定字符串结尾
def endwith_str(string: str, end: str):
    return string[-len(end) :] == end


string = "abcdefg"
end = "fg"

print(endwith_str(string, end))
print(string.endswith(end))


# 判断字符串是否为纯数字组成
def isdigit_str(string: str):
    for cha in string:
        if not ("9" >= cha >= "0"):
            return False
    return True


string = "24g43857"
print(isdigit_str(string))


# 字符串中所有小写字母变大写
def upper_str(string: str):
    string_tmp = ""
    for s in string:
        if "z" > s > "a":
            string_tmp += chr(ord(s) - 32)
        else:
            string_tmp += s
    return string_tmp


print(upper_str("hskfSG245j"))


# rjust
def my_rjust(string: str, width: int, fill):
    return (width - len(string)) * fill + string


print(my_rjust("abc", 6, "."))
print("abc".rjust(6, "."))


# index
def my_index(cha: str, lis: list):
    list_tmp = []
    for i, arg in enumerate(lis):
        if arg == cha:
            list_tmp.append(i)
    return list_tmp


list1 = ["l", "g", "k", "a", 2, 2, "l"]
print(my_index("l", list1))


# len
def my_len(sequence):
    count = 0
    for _ in sequence:
        count += 1
    return count


print(my_len([2, 3, 4, 5, 3, 2, "j", "h"]))
print(my_len((2, 3, 4, 5, 3, 2, "j", "h")))
print(my_len("slfjlsjf"))


# in
def my_in(seq, sequence):
    for i in sequence:
        if seq == i:
            return True
    return False


print(my_in("abc", ["lisi", "zhangsan", "abc"]))
print("abc" in ["lisi", "zhangsan", "abc"])


# replace
def my_replace(string: str, old: str, new: str):
    return new.join(string.split(old))


def my_replace1(string: str, old: str, new: str):
    i = 0
    string_tmp = ""
    while i < len(string):
        if old == string[i : i + len(old)]:
            string_tmp += new
            i += len(new)
        else:
            string_tmp += string[i]
        i += 1
    return string_tmp


string = "how are you? and you?"
print(string.replace("you", "me"))
print(my_replace1(string, "you", "me"))
print(my_replace(string, "you", "me"))


# 获取序列中元素的最大值，如果是字典，取value的最大值
def my_max(seq):
    if type(seq) == dict:
        seq = list(seq.values())
    max = seq[0]
    for k in seq:
        if k > max:
            max = k
    return max


list1 = [-7, -4, -1, 4]
tuple1 = (34, 5, 6, 7, 76)
str1 = "hljsfhgalz"
dict1 = {"a": 34, "b": 87, "c": 60}
print(my_max(dict1))
print(my_max(list1))
print(my_max(tuple1))
print(my_max(str1))


# 列表的交集
def jiaoji(list1: list, list2: list):
    list3 = []
    for i in list1:
        for j in list2:
            if i == j:
                list3.append(i)
                break
    return list3


list1 = ["abc", "zou", "li", "jj"]
list2 = ["bc", "zou", "li", "kk", "kk", "jj"]
print(jiaoji(list1, list2))


# 并集
def union(list1: list, list2: list):
    return list(set(list1 + list2))


def union1(list1: list, list2: list):
    list3 = list1
    for i in list2:
        if i not in list3:
            list3.append(i)
    return list(set(list3))


list1 = ["abc", "zou", "li", "jj"]
list2 = ["bc", "zou", "li", "kk", "kk", "jj"]
print(union1(list2, list1))
print(union(list1, list2))


# 差集
def sub(list1: list, list2: list):
    list3 = []
    for i in list1:
        if i not in list2:
            list3.append(i)
    return list(set(list3))


list1 = ["abc", "zou", "li", "jj"]
list2 = ["bc", "zou", "li", "kk", "kk", "jj"]
print(sub(list1, list2))
print(sub(list2, list1))


# 补集
def sep(list1: list, list2: list):
    list3 = []
    for i in list(set(list1 + list2)):
        if i not in list1:
            list3.append(i)
    return list3


list1 = ["abc", "zou", "li", "jj"]
list2 = ["bc", "zou", "li", "kk", "kk", "jj"]
print(sep(list1, list2))
"""
打印各种三角形图案

*
**
***
****
*****

    * 0
   ** 1
  *** 2
 **** 3
***** 4

    *
   ***
  *****
 *******
*********
"""

line = int(input("请输入要打印的行数:"))

for i in range(line):
    for _ in range(i + 1):
        print("*", end="")
    print(" ")

for i in range(line):
    for _ in range(line - i - 1):
        print(" ", end="")
    for _ in range(i + 1):
        print("*", end="")
    print()

for i in range(line):
    for _ in range(line - i - 1):
        print(" ", end="")
    for _ in range(2 * i + 1):
        print("*", end="")
    print()

# 这个不太好理解
for i in range(line):
    for j in range(line * 2 - 1):
        if j < line - 1 - i:
            print(" ", end="")
        elif j <= line - 1 + i:
            print("*", end="")
        else:
            print(" ", end="")
    print()

# 计算代码的运行时间
import time


def calc_time(fn):
    start = time.time()
    fn()
    end = time.time()
    print(end - start)


def test():
    sum = 0
    for i in range(10000000):
        sum += i
    print(sum)


calc_time(test)

# 计算代码的运行时间(装饰器)

import time


def calc_time_decorator(fn):
    def inner():
        start = time.time()
        fn()
        end = time.time()
        print(end - start, "decorator")

    return inner


@calc_time_decorator
def demo():
    sum = 0
    for i in range(10000000):
        sum += i
    print(sum)


demo()
