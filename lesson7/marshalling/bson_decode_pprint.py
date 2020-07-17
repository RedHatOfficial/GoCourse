import pprint
import bson

with open("tree.bson", "rb") as fin:
    content = fin.read()
    binary_tree = bson.loads(content)
    pp = pprint.PrettyPrinter(indent=4)
    pp.pprint(binary_tree)
