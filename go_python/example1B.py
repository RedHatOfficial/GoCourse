import ctypes

example1 = ctypes.CDLL("./example1.so")

example1.hello()
