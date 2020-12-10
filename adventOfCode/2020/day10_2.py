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
lines.append(0)
lines.sort()
lines.append(max(lines)+3)
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

DP = {}
def dp(i):
    if i == len(lines)-1:
        return 1
    if i in DP:
        return DP[i]
    answer = 0
    for j in range(i+1, len(lines)):
        if lines[j] - lines[i] <= 3:
            answer += dp(j)
    DP[i] = answer
    return answer

print(numberOfOne*numberOfThree)
print(dp(0))

