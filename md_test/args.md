# args 可变非关键参数

Sed is a stream editors. A stream editor is used to perform basic text transforma‐
tions on an input stream (a file or input from a pipeline). While in some ways sim‐
ilar to an editor which permits scripted edits (such as ed), sed works by making
only one pass over the input(s), and is consequently more efficient. But it is
sed's ability to filter text in a pipeline which particularly distinguishes it from
other types of editors

```python
def add(a, b):
    return a + b




print(add(2, 3))
"""
5
"""
```

```python
def add(*args):
    return sum(args) # args是一个元组(),用sum()把里面的数字相加


print(add(2, 3))
print(add(2, 3, 4))
print(add(2, 3, 4, 5))
"""
5
9
14
"""
```
