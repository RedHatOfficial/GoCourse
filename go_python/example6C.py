import ctypes

so6 = ctypes.CDLL("./so6.so")

l = so6.hello("ěščř ЩжΛλ".encode("utf-8"))
print(l)
