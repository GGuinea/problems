import re
import fileinput
from itertools import combinations
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
lines = list(map(int, lines))

pointer = 0
preamble = 25
badNumber = 0

def findPair(lst, value):
    res = []
    for var in combinations(lst, 2):
        if var[0] + var[1] == value and var[0] != var[1]:
            res.append((var[0], var[1]))
    return res

def find_contigous_sum(lst, value):
    for i in range(0, len(lst)):
        for j in range(i, len(lst)):
            sumOfIntegeres = sum(lst[i:j])
            if sumOfIntegeres == value:
                return min(lst[i:j]) + max(lst[i:j])

for i in range(preamble, len(lines)):
    if (i - preamble) < 0:
        pointer = 0
    else:
        pointer = i - preamble

    allPairs = findPair(lines[pointer:i], lines[i])
    if len(allPairs) == 0:
        print(lines[i])
        badNumber = lines[i]
        print(find_contigous_sum(lines[0:i], badNumber))
        break



