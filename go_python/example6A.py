import ctypes

example6 = ctypes.CDLL("./example6.so")

l = example6.hello("World!")
print(l)
