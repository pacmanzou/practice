import json


def read_file(fp):
    try:
        with open(fp, mode="r", encoding="utf8") as fp:
            return fp.read()
    except Exception:
        print("file not found")


def write_json(data, fp):
    with open(fp, mode="w", encoding="utf8") as fp:
        json.dump(data, fp)


def read_json(fp, t):
    """
    If the file is empty ,return t
    t: [] or {}
    """
    try:
        json.loads(read_file(fp))
    except Exception as identifier:
        return t
    else:
        return json.loads(read_file(fp))
