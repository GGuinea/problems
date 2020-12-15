import re
import fileinput
import copy
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

lines = read_lines("input")
lines = remove_end_of_line(lines)
line = lines[0]
numbers = defaultdict(list) 

line = line.split(",")

lastNumber = -1
iterator = 1
for i in range(0, len(line)):
    numbers[int(line[i])] = [i+1]
    lastNumber = int(line[i])
    iterator+=1

while True:
    occurances = numbers[lastNumber]
    nextNumber = 0
    if len(occurances) == 1:
        if not numbers[nextNumber]:
            numbers[nextNumber] = [iterator]
        else:
            newOccurences = numbers[nextNumber]
            newOccurences.append(iterator)
            newOccurences.sort(reverse = True)
            numbers[nextNumber] = newOccurences[0:3]
    else:
        nextNumber = occurances[0] - occurances[1]
        if not numbers[nextNumber]:
            numbers[nextNumber] = [iterator]
        else:
            newOccurences = numbers[nextNumber]
            newOccurences.append(iterator)
            newOccurences.sort(reverse = True)
            numbers[nextNumber] = newOccurences[0:3]
    if nextNumber == 201:
        print(iterator)

    if iterator == 30000000:
        print(nextNumber)
        exit()
    lastNumber = nextNumber 
    iterator+=1


