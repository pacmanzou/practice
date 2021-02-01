import threading
import time

ticket = 20
lock = threading.Lock()


def sell_ticket():
    global ticket
    while True:
        lock.acquire()
        if ticket > 0:
            ticket -= 1
            print('{}购买成功,剩余{}张'.format(threading.currentThread().name,
                                        ticket))
            lock.release()
        else:
            lock.release()
            print('全部卖完')
            break


t3 = threading.Thread(target=sell_ticket, name='333')
t4 = threading.Thread(target=sell_ticket, name='444')

if __name__ == "__main__":
    t3.start()
    t4.start()
    # sell_ticket()
