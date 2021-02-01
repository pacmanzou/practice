class House(object):
    def __init__(self, shape, area_all, fur_list: list = None):
        if fur_list is None:
            fur_list = []
        self.shape = shape
        self.area_all = area_all
        self.fur_list = fur_list
        self.free = self.area_all
        if len(fur_list) != 0:
            for fur in fur_list:
                self.free -= fur.area

    def add(self, fur):
        if fur.area > self.free:
            print('There is not enough space left.')
        else:
            self.fur_list.append(fur)
            self.free -= fur.area

    def __str__(self):
        fur_list = []
        for fur in self.fur_list:
            fur_list.append(fur.name)

        return 'shape={},area_all={},free={},fur_list={}'.format(
            self.shape, self.area_all, self.free, fur_list)


class Furniture(object):
    def __init__(self, name, area):
        self.name = name
        self.area = area


f1 = Furniture('bed', 4)
f2 = Furniture('chest', 2)
f3 = Furniture('table', 1.5)
h1 = House('舒适型', 20, [f1, f2, f3])
print(h1)
f4 = Furniture('test', 2)
h1.add(f4)
print(h1)
h1.add(f4)
print(h1)
h2 = House('经济型', 15)
print(h2)
