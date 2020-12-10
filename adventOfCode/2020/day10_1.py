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

lines.sort()
jolts = 0
numberOfOne = 0
numberOfThree = 0
numberOfTwo = 0
for number in lines:
    if (number - jolts) == 1:
        numberOfOne += 1
        jolts = number 
    elif (number - jolts) == 2:
        jolts = number 
    elif (number - jolts) == 3:
        numberOfThree += 1
        jolts = number 
    else:
        print("Problem!", numberOfOne, numberOfThree, number, jolts)
        break

numberOfThree+=1
print(numberOfOne, numberOfThree)
print(numberOfOne*numberOfThree)


