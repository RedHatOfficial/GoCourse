import ctypes

so3 = ctypes.CDLL("./so3.so")

a = 1.2
b = 3.4

c = so3.add(a, b)
print(c)
