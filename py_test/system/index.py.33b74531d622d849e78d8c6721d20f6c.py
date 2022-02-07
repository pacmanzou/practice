import file_utils


def main():
    while True:
        print(file_utils.read_file("teacher_ui.txt"))
        num = input("请选择(1-3):")
        if num == "1":
            teacher_utils.login()
        elif num == "2":
            teacher_utils.register()
        elif num == "3":
            teacher_utils.exit()
        else:
            print("输入错误")


if __name__ == "__main__":
    main()
