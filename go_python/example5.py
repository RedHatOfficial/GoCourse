import ctypes

example5 = ctypes.CDLL("./example5.so")

l = example5.hello("World!")
print(l)
