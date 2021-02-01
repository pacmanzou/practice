import sys
import uuid
import requests
import hashlib
import time
import json
from importlib import reload


reload(sys)

YOUDAO_URL = "https://openapi.youdao.com/api"
APP_KEY = "0d2c703fe28a5b36"
APP_SECRET = "qVPudMYrrJmeoTGpIyzzVliIfeacQ3uE"


def encrypt(signStr):
    hash_algorithm = hashlib.sha256()
    hash_algorithm.update(signStr.encode("utf-8"))
    return hash_algorithm.hexdigest()


def truncate(q):
    size = len(q)
    return q if size <= 20 else q[0:10] + str(size) + q[size - 10 : size]


def do_request(data):
    headers = {"Content-Type": "application/x-www-form-urlencoded"}
    return requests.post(YOUDAO_URL, data=data, headers=headers)


def connect(q):

    data = {}
    data["to"] = "zh-CHS"
    data["from"] = "en"
    data["signType"] = "v3"
    curtime = str(int(time.time()))
    data["curtime"] = curtime
    salt = str(uuid.uuid1())
    signStr = APP_KEY + truncate(q) + salt + curtime + APP_SECRET
    sign = encrypt(signStr)
    data["appKey"] = APP_KEY
    data["q"] = q
    data["salt"] = salt
    data["sign"] = sign
    data["vocabId"] = "1"

    response = do_request(data)
    contentType = response.headers["Content-Type"]
    if contentType == "audio/mp3":
        millis = int(round(time.time() * 1000))
        filePath = "合成的音频存储路径" + str(millis) + ".mp3"
        fo = open(filePath, "wb")
        fo.write(response.content)
        fo.close()
    else:
        a_dict = json.loads(response.content)
        a_list = a_dict["web"]
        for item in a_list:
            print(item["value"])


if __name__ == "__main__":
    # q = input("请输入要翻译的词汇:")
    # print()
    connect(q="你好")
