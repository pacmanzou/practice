import file_utils


def manager_stu():
    while True:
        print(file_utils.read_file("student_ui.txt"))
        num = input("请选择(1-5):")
        student_list = file_utils.read_json("student_data.json", [])
        if num == "1":
            add_stu(student_list, "student_data.json")
        elif num == "2":
            search_stu(student_list)
        elif num == "3":
            mod_stu(student_list)
        elif num == "4":
            del_stu(student_list)
        elif num == "5":
            break
        else:
            print("输入错误")


def search_stu(data: list):
    while True:
        print(file_utils.read_file("student_search_ui.txt"))
        num = input("请选择(1-5):")
        if num == "1":
            search_byname(data)
        elif num == "2":
            search_bysex(data)
        elif num == "3":
            search_byage(data)
        elif num == "4":
            search_all(data)
        elif num == "5":
            break
        else:
            print("输入错误")


def add_stu(data: list, fp):
    while True:
        while True:
            name = input("请输入学生姓名:")
            for items in data:
                if items.get("name") == name:
                    print("已存在的学生姓名")
                    break
            else:
                break
        sex = input("请输入学生性别:")
        age = input("请输入学生年龄:")
        import modul

        s = modul.Student(name, sex, age)
        data.append(s.__dict__)
        while True:
            num = input("是否继续添加(y or n)")
            if num == "y":
                break
            elif num == "n":
                file_utils.write_json(data, "student_data.json")
                return
            else:
                print("输入错误")


def search_byname(data: list):
    name = input("请输入学生姓名:")
    for i in data:
        if i.get("name") == name:
            print(i)
            break
    else:
        print("没有name={}学生的信息".format(name))


def search_bysex(data: list):
    sex = input("请输入学生性别:")
    for i in data:
        if i.get("sex") == sex:
            print(i)
            break
    else:
        print("没有sex={}学生的信息".format(sex))


def search_byage(data: list):
    age = input("请输入学生年龄:")
    for i in data:
        if i.get("age") == age:
            print("i")
            break
    else:
        print("没有age={}学生的信息".format(age))


def search_all(data: list):
    for i in data:
        print(i)


def mod_stu(data: list):
    name = input("请输入要修改学生姓名:")
    for items in data:
        if items.get("name") == name:
            items["name"] = input("请输入新名字:")
            items["sex"] = input("请输入新性别:")
            items["age"] = input("请输入新年龄:")
            file_utils.write_json(data, "student_data.json")
            print("修改成功")
            break
    else:
        print("修改失败\n没有name={}学生的信息".format(name))


def del_stu(data: list):
    name = input("请输入要删除学生姓名:")
    for items in data:
        if items.get("name") == name:
            data.remove(items)
            file_utils.write_json(data, "student_data.json")
            print("删除成功")
            break
    else:
        print("删除失败\n没有name={}学生的信息".format(name))
    pass
