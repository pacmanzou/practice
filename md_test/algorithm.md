# 简单的算法

<!-- vim-markdown-toc GFM -->

- [正整数的反转](#正整数的反转)
- [百钱百鸡](#百钱百鸡)
- [生成斐波那契数列的前 20 个数](#生成斐波那契数列的前-20-个数)
- [10000 以内的完美数](#10000-以内的完美数)
- [素数](#素数)
- [打印各种三角形图案](#打印各种三角形图案)

<!-- vim-markdown-toc -->

```python
"""
水仙花数
"""
for i in range(100, 1000):
    if (i % 10) ** 3 + (i // 10 % 10) ** 3 + (i // 100) ** 3 == i:
        print(i)
"""
153
370
371
407
"""
```

## 正整数的反转

```python
# (1)
num = 1234
# num = 12340 # 4321
print(int(str(num)[::-1]))
# (2)
num = int(input('num = '))
reversed_num = 0
while num > 0:
    reversed_num = reversed_num * 10 + num % 10
    num //= 10
print(reversed_num)
"""
4321
num = 4321
1234
"""
```

## 百钱百鸡

```python
# b:5 m:3 s:1/3
for b in range(100 // 5 + 1):
    for m in range(100 // 3 + 1):
        if (100 - b - m) * 1 / 3 + 5 * b + 3 * m == 100:
            print(b, m, 100 - b - m)
"""
0 25 75
4 18 78
8 11 81
12 4 84
"""

```

## 生成斐波那契数列的前 20 个数

```python
x = 1
y = 1
list1 = [1]
for _ in range(20):
    list1.append(y)
    x, y = y, y + x # python 独有方法
print(list1)
"""
[1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144,
233, 377,610, 987, 1597, 2584, 4181, 6765, 10946]
"""

```

## 10000 以内的完美数

```python
for num in range(1, 10000):
    result = 1
    for factor in range(2, int(num ** 0.5) + 1):
        if num % factor == 0:
            result += factor
            if factor > 1 and num // factor != factor:
                result += num // factor
    if result == num and result != 1:
        print(num)
"""
6
28
496
8128
"""
```

## 素数

```python
n = int(input('请输入要判断的数:'))
n = 4 if n == 1 else n # n为1时输出NO
for i in range(2, int(n ** 0.5) + 1):
    if n % i == 0:
        print('NO')
        break
else:
    print('YES')
"""
请输入要判断的数:103
YES
请输入要判断的数:104
NO
"""

```

## 打印各种三角形图案

```python
row = int(input('请输入要打印的行数:'))

for i in range(row):
    for _ in range(i + 1):
        print('*', end='')
    print(' ')

for i in range(row):
    for _ in range(row - i - 1):
        print(' ', end='')
    for _ in range(i + 1):
        print('*', end='')
    print()

for i in range(row):
    for _ in range(row - i - 1):
        print(' ', end='')
    for _ in range(2 * i + 1):
        print('*', end='')
    print()

# 这个不太好理解
for i in range(row):
    for j in range(row * 2 - 1):
        if j < row - 1 - i:
            print(' ', end='')
        elif j <= row - 1 + i:
            print('*', end='')
        else:
            print(' ', end='')
    print()
"""
请输入要打印的行数:5
*
**
***
****
*****
    *
   **
  ***
 ****
*****
    *
   ***
  *****
 *******
*********
    *
   ***
  *****
 *******
*********
"""

```
