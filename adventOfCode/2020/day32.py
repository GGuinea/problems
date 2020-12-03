def read_lines(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    file.close()
    return lines

def remove_end_of_line(lines):
    for item in range(0, len(lines)):
        lines[item] = lines[item].rstrip()
    return lines

def calculate_trees(right, down, lines):
    rightIncrement = right
    downIncrement = down
    lastRight = 0
    rightToCheck = 0
    numberOfTrees = 0
    tmpDown = down
    for singleLine in lines:
        if tmpDown != downIncrement:
            tmpDown += 1
            continue
        else:
            tmpDown = 1
        if singleLine[rightToCheck] == '#':
            numberOfTrees += 1
        rightToCheck = lastRight + rightIncrement
        if rightToCheck >= len(singleLine):
            rightToCheck = rightToCheck % len(singleLine)
        lastRight = rightToCheck
    return numberOfTrees

lines = read_lines('input')
lines = remove_end_of_line(lines)

output1 = calculate_trees(3,1,lines)
output2 = calculate_trees(1,1,lines)
output3 = calculate_trees(5,1,lines)
output4 = calculate_trees(7,1,lines)
output5 = calculate_trees(1,2,lines)

print(output1 * output2 * output3 * output4 *output5)
