class Animals(object):
    def __init__(self, name, age):
        self.name = name
        self.age = age

    def sleep(self):
        print(self.name + ' is sleeping')

    def bark(self):
        print(self.name + ' is barking')


class Dog(Animals):
    pass


class Student(Animals):
    def __init__(self, name, age, sex):
        # 下面两行都可以调用父类的__init__方法(推荐第二种super)
        # Animals.__init__(self, name, age)
        super().__init__(name, age)
        self.sex = sex

    def study(self):
        print(self.name + ' is studying')


def newmethod330():
    dog1 = Animals('dog', 3)
    dog2 = Dog('dog', 3)
    student = Student('student', 22, 'man')
    # print(isinstance(dog1, Animals))
    print(isinstance(dog1, Dog))
    print(issubclass(Animals, Dog))
    print(student.sex)
    print(student.name)
    student.study()
    student.sleep()
    dog1.sleep()


newmethod330()
print(Dog.__mro__)
