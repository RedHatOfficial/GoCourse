def testList():
    l1 = []
    l1.append(10)
    l1.append("Start")
    l1.append(None)
    print(l1)

testList()

import dis
dis.dis(testList)
