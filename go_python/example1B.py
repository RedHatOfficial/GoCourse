import ctypes

so1 = ctypes.CDLL("./so1.so")

so1.hello()
