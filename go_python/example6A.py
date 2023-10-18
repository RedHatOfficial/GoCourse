import ctypes

so6 = ctypes.CDLL("./so6.so")

l = so6.hello("World!")
print(l)
