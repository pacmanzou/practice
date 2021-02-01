# 3 4 12
# 12 16 4 48
"""
Least Common Multiple
最小公倍数
"""


def lcm(x, y):
    (x, y) = (y, x) if x > y else (x, y)
    for i in range(1, y + 1):
        if (i * x) % y == 0:
            return i * x


print(lcm(3, 4))
"""
Greatest Common Divisor
最大公因数
"""


def gcd(x, y):
    (x, y) = (y, x) if x > y else (x, y)
    for i in range(x, 1, -1):
        if (y % i) == (x % i) == 0:
            return i


print(gcd(12, 16))
"""
回文数
"""

# def is_palindrome(num):
#     temp = ''
#     for i in reversed(str(num)):
#         temp += i
#     return int(temp) == num


def is_palindrome(num):
    temp = num
    total = 0
    while temp > 0:
        total = total * 10 + temp % 10
        temp //= 10
    return total == num


print(is_palindrome(12321))


def foo(a):
    a = 200
    print(a) # 200


if __name__ == '__main__':
    a = 100
    foo(a)
    print(a) # 100

import re


def main():
    username = input('请输入用户名字:')
    qq = input('请输入qq号:')
    pattern = re.compile(r'^\w{6,20}$')
    username_validation = re.match(pattern, username)
    qq_validation = re.match(r'^[1-9]\d{4,11}$', qq)
    if username_validation and qq_validation:
        print('success')
    elif not username_validation and qq_validation:
        print('error 用户名字必须由字母数字或者下划线构成，且长度要在6-20个字符')
    elif not qq_validation and username_validation:
        print('error qq号是5-12个数字之间，第一个数字不能为0')
    else:
        print('error 用户名字必须由字母数字或者下划线构成，且长度要在6-20个字符')
        print('error qq号是5-12个数字之间，第一个数字不能为0')


if __name__ == "__main__":
    main()

import re


def main():
    sentence = '''
        重要的事情说8130123456789遍，我的手机号是13512346789这个靓号，
        不是15600998765，也是110或119，王大锤的手机号才是15600998765。
        '''
    pattern = re.compile(r'(?<=\D)1[345678]\d{9}(?=\D)')
    # pattern = re.compile(r'(?<!\d)1[345678]\d{9}(?!\d)')
    temp_list = re.findall(pattern, sentence)
    print(temp_list)
    temp_iter = re.finditer(pattern, sentence)
    for i in temp_iter:
        print(i.group())


if __name__ == "__main__":
    main()

import re


def main():
    sentence = '你丫是傻叉吗? 我操你大爷的. Fuck you.'
    purified = re.sub('[操艹肏]|fuck|shit|傻[逼叉子吊屌缺屄]|煞笔',
                      '*',
                      sentence,
                      flags=re.IGNORECASE)
    print(purified)


if __name__ == "__main__":
    main()

import re


def main():
    sentence = '床前明月光，疑是地上霜，举头望明月，低头思故乡。'
    sentence_list = re.split('[，。]', sentence)
    while '' in sentence_list:
        sentence_list.remove('')
    print(sentence_list)


if __name__ == "__main__":
    main()

# filter
newlist = filter(lambda items: items % 2 == 1, range(1, 11))
for items in newlist:
    print(items, end=' ')
# 解析式
print([i for i in range(1, 11) if i % 2 == 1])

# map
newlist = map(lambda items: items ** 2, range(1, 11))
for items in newlist:
    print(items, end=' ')

# 解析式
print([i ** 2 for i in range(1, 11)])

# reduce (python3需要导入)
from functools import reduce
newlist = reduce(lambda x, y: x + y, range(1, 11))
print(newlist)

# sum
print(sum(range(1, 11)))
my_dict = {'a': 3, 'b': 1, 'd': 4, 'c': 5}


# print(my_dict.items())
# print(sorted(my_dict.items()))
def index_1(items):
    """
    Return index[1]
    """
    return items[1]


print(sorted(my_dict.items(), key=index_1))
print(sorted(my_dict.items(), key=lambda x: x[1]))
# print({i: v for v in my_dict.values()})
# print(sorted(my_dict.values()))

r = open('test.txt', mode='r', encoding='utf8')
print(r.read())
r.close()

try:
    with open('test.txt', mode='r', encoding='utf8') as r:
        print(r.read())
        num = input('请输入')
except Exception as i:
    print(i)

w = open('test1.txt', mode='w', encoding='utf8')
w.write('test1')
w.close()

import json

# 序列化
# dumps:将数据转换为json字符串,不会将数据保存到文件里面
# dump:将数据转换为json字符串,同时将数据保存到文件里面
d = {'name': 'zou', 'age': 22}
l = ['name', 'zou', 'age', 22]
with open('test.json', mode='w', encoding='utf8') as w_file:
    # w.write(json.dumps(d))
    json.dump(l, w_file)

# 反序列化
# loads:将json字符串加载为python数据
# load:读取文件json文件，将文件的内容加载为python数据
d_str = '{"name": "zou", "age": 22}'
d_json = json.loads(d_str)
print(d_json['name'])

with open('test.json', mode='r', encoding='utf8') as r_file:
    p = json.load(r_file)
    print(p, type(p))

import json


class Person(object):
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def eat(self):
        print(self.name + 'is eating')


p = Person('zou', 22)
d = json.dumps(p.__dict__)
print(d)
j = json.loads(d)
print(j)

import re

string = 'set expandtab test'
string1 = 'set expandtabtest'
pattern = re.compile(r'\bex.*b\b')
print(re.search(pattern, string1))
print(re.search(pattern, string))
