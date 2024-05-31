import ctypes

example4 = ctypes.CDLL("./example4.so")

a = 2**31-1
b = 1

c = example4.add(a, b)
print(c)
