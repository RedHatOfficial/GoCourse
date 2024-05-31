import ctypes

example4 = ctypes.CDLL("./example4.so")

a = 2**31-1
b = 1

example4.add.restype = ctypes.c_int64

c = example4.add(a, b)
print(c)
