import os
import sys

for name, value in os.environ.items():
    print("{0}: {1}".format(name, value))

print(sys.argv[1:])
