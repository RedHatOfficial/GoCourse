import ctypes

so4 = ctypes.CDLL("./so4.so")

a = 1
b = 2

c = so4.add(a, b)
print(c)
