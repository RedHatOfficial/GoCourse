import ctypes

so3 = ctypes.CDLL("./so3.so")

a = 1
b = 10000000000000000

c = so3.add(a, b)
print(c)
