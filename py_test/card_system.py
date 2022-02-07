def system_show():
    print('---------------------')
    print('    名片管理系统')
    print('1:添加名片')
    print('2:删除名片')
    print('3:修改名片')
    print('4:查询名片')
    print('5:显示所有名片')
    print('0:exit')
    print('---------------------')
    global num
    num = input('请输入要进行得操作:')
    if num.isdigit():
        num = int(num)


def operation_add():
    global cards
    card = {}
    name = input('请输入name:')
    for item in cards:
        if item['name'] == name:
            print('已存在{}名片'.format(item['name']))
            print('请重新输入')
            print()
            break
    else:
        card['name'] = name
        card['age'] = int(input('请输入age:'))
        card['tel'] = input('请输入tel:')
        cards.append(card)


def operation_del():
    global cards
    name = input('请输入你要删除的名片拥有者姓名:')
    for item in cards:
        if item['name'] == name:
            cards.remove(item)
            print('已经删除{}'.format(name))
            break
    else:
        print('要删除的名字不存在!')
        print()


def operation_mod():
    global cards
    card = {}
    for item in cards:
        if item.get('name') == input('请输入要修改名片的名字:'):
            cards.remove(item)
            card['name'] = input('请输入新名字:')
            card['age'] = int(input('请输入新年龄:'))
            card['tel'] = input('请输入新电话:')
            cards.append(card)
            break
    else:
        print('要修改的名字不在名片列表中!')
        print()


def operation_query():
    global cards
    name = input('请输入要搜索名片的名字:')
    for item in cards:
        if item.get('name') == name:
            print('名字为{}的名片信息:{}'.format(name, item))
            break
    else:
        print('要搜索的名字不存在!')
        print()



def operation_list():
    global cards
    print('list如下:')
    for item in cards:
        print(item)


def operation_exit():
    # global flag
    if 'Yes' == input('确定退出系统吗？Yes(exit) or others(cancel)\n'):
        flag = False
    else:
        print('已经取消')
        print()


def main():
    flag = True
    while flag:
        system_show()
        if num == 0:
            flag = False if 'Yes' == input(
                '确定退出系统吗？Yes(exit) or others(cancel)\n') else True
        elif num == 1:
            operation_add()
        elif num == 2:
            operation_del()
        elif num == 3:
            operation_mod()
        elif num == 4:
            operation_query()
        elif num == 5:
            operation_list()
        else:
            print('请正确输入！')

    print('正在退出系统...')
    cards.clear()
    print('已退出')


if __name__ == "__main__":
    num = 0
    cards = []
    main()
