import ctypes

example3 = ctypes.CDLL("./example3.so")

a = 1.2
b = 3.4

c = example3.add(a, b)
print(c)
