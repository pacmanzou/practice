# 装饰器的简单使用

```python
"""
计算代码的运行时长(普通函数)
"""
import time


def calc_time(fn):
    start = time.time()
    print(fn())
    end = time.time()
    print(end - start)


def demo():
    sum = 0
    for i in range(10000000):
        sum += i
    print(sum)


calc_time(demo)
"""
49999995000000
0.6397993564605713
"""
```

```python
"""
计算代码的运行时长(decorator)
"""
import time


def calc_time_decorator(fn): # fn相当于普通的demo
    def inner():
        start = time.time()
        fn()
        end = time.time()
        print(end - start, '(decorator)')

    return inner


@calc_time_decorator
def demo():
    sum = 0
    for i in range(10000000):
        sum += i
    print(sum)


demo() # 此时demo相当于装饰器里面的inner,不再是普通的demo
"""
49999995000000
0.7188940048217773 (decorator)
"""
```
