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
print(operations)
operationsCopy = list(operations)
for change in range(len(operations)):
    operations = list(operationsCopy)
    if operations[change] == 'nop':
        operations[change] = 'jmp'
    elif operations[change] == 'jmp':
        operations[change] = 'nop'
    else:
        continue
    i = 0
    ACCUMULATOR = 0
    t = 0
    while 0 <= i < len(operations) and t < 1000:
        t+=1
        if operations[i] == "nop":
            i+=1
        elif operations[i] == "acc" :
            ACCUMULATOR += values[i]
            i+=1
        elif operations[i] == "jmp":
            i += values[i]
        else:
            print("PROBLEM")
            break
    if i == len(operations):
        print(ACCUMULATOR)

print(ACCUMULATOR)

