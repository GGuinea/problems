import re
from collections import Counter

def read_lines(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    file.close()
    return lines

def remove_end_of_line(lines):
    for item in range(0, len(lines)):
        lines[item] = lines[item].rstrip()
    return lines

lines = read_lines('input')
lines = remove_end_of_line(lines)
answer = 0
groupAnswers = set() 
allAnswers = []
nuberOfPeopleInGroup = 0
for line in lines:
    if line == '':
        for char in groupAnswers:
            cnt = allAnswers.count(char)
            if cnt == nuberOfPeopleInGroup:
                answer += 1
        groupAnswers = set()
        allAnswers = []
        nuberOfPeopleInGroup = 0
        continue
    nuberOfPeopleInGroup += 1
    listOfChars = list(line)
    groupAnswers.update(listOfChars)
    allAnswers += listOfChars

print(answer)


