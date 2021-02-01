class Student(object):
    # __slots__ = ('name', 'age', 'height', 'city')

    # 类属性
    type = "person"
    __count = 0

    def __init__(self, name, age, height):
        Student.__count += 1
        self.name = name
        self.age = age
        self.__height = height

    def run(self):
        print("run")

    def eat(self):
        print("eat")

    def name_is(self):
        print(self.name)

    def __setitem__(self, key, value):
        self.__dict__[key] = value

    def __getitem__(self, key):
        return self.__dict__[key]

    def get_height(self):
        return self.__height

    @classmethod
    def get_count(cls):
        return cls.__count

    @classmethod
    def test(cls):
        print("test")
        print(cls.type)


s1 = Student("xiaoming", 18, 1.75)
s2 = Student("xiaohong", 17, 1.65)
print(Student.get_count())

# 类属性modify
# Student.type = 'modify'

# 类属性调用
# print(Student.type)
# print(s1.type)

# 类方法调用
# Student.test()
s1.test()

# 示例变量和方法
# s1.run()
# s2.eat()
# 类对象调用时要传入示例对象参数
# s1.name_is()
# Student.name_is(s1)
# print(s1.name)
# __变量名，__方法名表示私有
# 都可以访问私有属性,下面两行
print(s1._Student__height)
print(s1.get_height())

# __slots__ 有city的话才可已添加属性
# s1.city = 'shanghai'
# print(s1.city)

# 将对象转化为字典
# print(s1.__dict__)

# 不能把对象当作字典来使用 下面一行将报错
# s1['age'] = 20
# 对象[]语法会调用__setitem__方法
# print(s1.age)
# s1['age'] = 20
# print(s1.age)
# print(对象[])语法会调用__getitem__方法
# print(s1['age'])


# 魔术方法 __str__ __del__ __call__
class Test(object):
    def __init__(self, a, b):
        print("__init__执行")
        self.a = a
        self.b = b

    def __str__(self):
        return "__str__执行"

    def __del__(self):
        print("__del__执行")

    def __call__(self):
        print("__call__执行")


t1 = Test("abc", "def")
# t2 = Test('def', 'ghi')
print(t1)

# __call__
t1()


class Person(object):
    def __init__(self, name, age):
        self.name = name
        self.age = age

    # def __eq__(self):
    #     pass


# == 会调用对象的__eq__方法
p1 = Person("zhangsan", 18)
p2 = Person("zhangsan", 18)
print(p1 is p2)  # False
print(p1 == p2)  # False p1.__eq__(p2)
num1 = [1, 2, 3]
num2 = [1, 2, 3]
print(num1 is num2)  # False
print(num1 == num2)  # True
