import ctypes

example6 = ctypes.CDLL("./example6.so")

l = example6.hello("ěščř ЩжΛλ".encode("utf-8"))
print(l)
