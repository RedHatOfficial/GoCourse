  2           0 BUILD_LIST               0
              2 STORE_FAST               0 (l1)

  3           4 LOAD_FAST                0 (l1)
              6 LOAD_ATTR                0 (append)
              8 LOAD_CONST               1 (10)
             10 CALL_FUNCTION            1
             12 POP_TOP

  4          14 LOAD_FAST                0 (l1)
             16 LOAD_ATTR                0 (append)
             18 LOAD_CONST               2 ('Start')
             20 CALL_FUNCTION            1
             22 POP_TOP

  5          24 LOAD_FAST                0 (l1)
             26 LOAD_ATTR                0 (append)
             28 LOAD_CONST               0 (None)
             30 CALL_FUNCTION            1
             32 POP_TOP

  6          34 LOAD_GLOBAL              1 (print)
             36 LOAD_FAST                0 (l1)
             38 CALL_FUNCTION            1
             40 POP_TOP
             42 LOAD_CONST               0 (None)
             44 RETURN_VALUE
