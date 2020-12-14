import re
import fileinput
import copy

def read_lines(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    file.close()
    return lines

def remove_end_of_line(lines):
    for item in range(0, len(lines)):
        lines[item] = lines[item].rstrip()
    return lines

value_data = []
mem = {}
lines = read_lines('input')
lines = remove_end_of_line(lines)

currentMask = ""
value = ""
valuePlace = 0
for line in lines:
    if line.startswith("mask"):
        splitted = line.split(" = ")
        mask = splitted[1]
    else:
        splitted = line.split(" = ")
        valuePlace = int(splitted[0][4:-1])
        value = int(splitted[1])
        binaryValue = bin(value)[2:]
        newValue = copy.deepcopy(binaryValue)
        newValue = newValue.zfill(36)
        newValue = list(newValue)
        for i in range(len(mask)):
            if mask[i] == "X":
                continue
            elif mask[i] == "1":
                newValue[i] = "1"
            elif mask[i] == "0":
                newValue[i] = "0"
        newValue = "".join(newValue)
        value = int(newValue,2)
        mem[valuePlace] = value

sumOfAll = 0
for k in mem:
    sumOfAll += mem[k] 
print(sumOfAll)






