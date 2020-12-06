import re

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
for line in lines:
    if line == '':
        answer += len(groupAnswers)
        groupAnswers = set()
        continue
    listOfChars = list(line)
    groupAnswers.update(listOfChars)

print(answer)


