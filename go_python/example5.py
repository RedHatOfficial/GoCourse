import ctypes

so5 = ctypes.CDLL("./so5.so")

l = so5.hello("World!")
print(l)
