import ctypes

example3 = ctypes.CDLL("./example3.so")

a = 1
b = 10000000000000000

c = example3.add(a, b)
print(c)
