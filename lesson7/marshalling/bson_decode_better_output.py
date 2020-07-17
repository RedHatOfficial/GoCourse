import bson

def printTree(node, level):
    if node is not None:
        format = level * "       "
        format += "---[ "
        level += 1
        printTree(node["left"], level)
        print(format+str(node["value"]))
        printTree(node["right"], level)
    pass

with open("tree.bson", "rb") as fin:
    content = fin.read()
    binary_tree = bson.loads(content)
    printTree(binary_tree["root"], 0)
