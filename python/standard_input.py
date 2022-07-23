import sys


def r(): return lambda: sys.stdin.readline().strip()


read = r()
line = read()
x, y = line.split()
a = int(x)
b = int(y)
print(a + b)
