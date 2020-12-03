def read_lines(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    file.close()
    return lines

def remove_end_of_line(lines):
    for item in range(0, len(lines)):
        lines[item] = lines[item].rstrip()
    return lines

#def calculate_treess(right, down, lines):


lines = read_lines('input')
lines = remove_end_of_line(lines)
print(lines)
rightIncrement = 3
downIncrement = 1
lastRight = 0
rightToCheck = 0
numberOfTrees = 0

for singleLine in lines:
    if singleLine[rightToCheck] == '#':
        numberOfTrees+=1
    rightToCheck = lastRight + rightIncrement
    if rightToCheck >= len(singleLine):
        rightToCheck = rightToCheck % len(singleLine)
    lastRight = rightToCheck

print(numberOfTrees)


