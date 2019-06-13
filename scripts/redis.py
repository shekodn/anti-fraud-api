import socket
import time
import os
import sys
import subprocess


# port = int(os.environ["CACHE_PORT"])
# print("ENV CACHE_PORT:", port)
port = 6379

while True:
    try:
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        print("Connection attempt")
        s.connect(('localhost', port))
        print("SUCCESS")
        s.close()
        break
    except socket.error as ex:
        print("ex:", ex)
        time.sleep(1)
