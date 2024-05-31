import ctypes

example6 = ctypes.CDLL("./example6.so")

l = example6.hello("World!".encode("utf-8"))
print(l)
