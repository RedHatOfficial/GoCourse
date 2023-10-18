import ctypes

so4 = ctypes.CDLL("./so4.so")

a = 2**31-1
b = 1

c = so4.add(a, b)
print(c)
