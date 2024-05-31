import ctypes

example3 = ctypes.CDLL("./example3.so")

a = 1
b = 2

c = example3.add(a, b)
print(c)
