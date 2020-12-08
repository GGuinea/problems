import re
import fileinput
from collections import defaultdict
def read_lines(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    file.close()
    return lines

def remove_end_of_line(lines):
    for item in range(0, len(lines)):
        lines[item] = lines[item].rstrip()
    return lines

lines = list(fileinput.input())
lines = remove_end_of_line(lines)
operations = []
values = []
ACCUMULATOR = 0
visited = set()
for line in lines:
    splitted = line.split(" ")
    if splitted[1].startswith('+'):
        operations.append(splitted[0])
        values.append(int(splitted[1][1:]))
    else:
        operations.append(splitted[0])
        values.append(int(splitted[1]))

loop = True
i = 0
print(operations)
while(loop):
    if operations[i] == "nop":
        if i in visited:
            print(ACCUMULATOR)
            loop = False
            break
        visited.add(i)
        i+=1
    elif operations[i] == "acc" :
        if i in visited:
            print(ACCUMULATOR)
            loop = False
            break
        visited.add(i)
        ACCUMULATOR += values[i]
        i+=1
    elif operations[i] == "jmp":
        if i in visited:
            print(ACCUMULATOR)
            loop = False
            break
        visited.add(i)
        i += values[i]
    else:
        print("PROBLEM")
        break

