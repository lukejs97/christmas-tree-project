import sys

def checkArgs():
    """A function to verify command line argument input and report errors"""
    size = 3
    if len(sys.argv) < 3:
        print("Size not specified. Default size is 3.")
    elif len(sys.argv) > 3: # need to stop here
        print("Too many arguments")
        print("Usage: py tree -height [TREE_SIZE]")
    else:
        try:
            size = int(sys.argv[2])
        except ValueError as e:
            print(e, "TREE_SIZE must be a integer")
    return size

def printTree(size: int):
    """A function to print a correctly formatted Christmas tree of the specified size"""
    print("*".center(((size * 2) + 1)))
    midSpace = 1
    for sect in reversed(range(size)):
        print("/".rjust(sect + 1), "\\".rjust(midSpace))
        midSpace += 2
    print("-".center(((size * 2) + 1), "-"))
    print("#".center(((size * 2) + 1)))

def main():
    """A driver function for using checkArgs and printTree to build a Christmas tree"""
    size = 0
    size = checkArgs()
    printTree(size)

if __name__ == "__main__":
    main()