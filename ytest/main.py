import socket
import random


def test():
    sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    try:
        while True:
            port = random.randrange(50000, 59000)
            result = sock.connect_ex(('127.0.0.1', port))
            if result != 0:
                print("Port is avaliable.")
                return port
    except Exception as e:
        print("Failed Connect sock, info: ", e)
        return False


print(test())
