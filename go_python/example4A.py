import ctypes

example4 = ctypes.CDLL("./example4.so")

a = 1
b = 2

c = example4.add(a, b)
print(c)
