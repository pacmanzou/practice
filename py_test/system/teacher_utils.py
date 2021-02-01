import sys
import modul
import file_utils
import student_utils


def register():
    teacher_name = input("请输入你的帐号姓名(3~6位):")
    teacher_pwd = input("请输入你的帐号密码(6~12位):")
    teacher_dict = file_utils.read_json("teacher_data.json", {})
    t = modul.Teacher(teacher_name, teacher_pwd)
    if teacher_name not in teacher_dict:
        teacher_dict[t.name] = t.pwd
        file_utils.write_json(teacher_dict, "teacher_data.json")
    else:
        print("已存在的用户，请重新输入:")


def login():
    if file_utils.read_file("teacher_data.json"):
        teacher_name = input("请输入你的帐号姓名:")
        teacher_pwd = input("请输入你的帐号密码:")
        teacher_dict = file_utils.read_json("teacher_data.json", {})
        if teacher_dict.get(teacher_name) == teacher_pwd:
            student_utils.manager_stu()
        else:
            print("账户名字或者密码不正确!")
    else:
        print("当前teacher账户为空，请先注册")


def exit():
    while True:
        content = input("确认要退出吗?(y or n)")
        if content == "y":
            print("已退出")
            sys.exit(0)
        elif content == "n":
            break
        else:
            print("输入错误")
