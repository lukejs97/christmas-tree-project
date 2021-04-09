import sys

size = 0
if len(sys.argv) == 3: # polish this up
    try:
        size = int(sys.argv[2])
    except ValueError as e:
        print(e, "TREE_SIZE must be a integer")
else:
    print("Incorrect amount of arguments")
    print("Usage: py tree -height [TREE_SIZE]")

print("*".center(((size * 2) + 1)))
midSpace = 1
for sect in reversed(range(size)):
    print("/".rjust(sect + 1), "\\".rjust(midSpace))
    midSpace += 2
print("-".center(((size * 2) + 1), "-"))
print("#".center(((size * 2) + 1)))

# comment stuff