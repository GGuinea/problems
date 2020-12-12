import re
import fileinput

def read_lines(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    file.close()
    return lines

def remove_end_of_line(lines):
    for item in range(0, len(lines)):
        lines[item] = lines[item].rstrip()
    return lines

def change_direction(way, value, currentFacing):
    ways = ["N", "E", "S", "W"]
    chaningSteps = value / 90
    for i in range(0, len(ways)):
        if currentFacing == ways[i]:

            if way == "R":
                tmp = (i + chaningSteps) % len(ways)
                currentFacing = ways[int(tmp)]
                return currentFacing
            else: 
                tmp = (i + 4 - chaningSteps) % len(ways)
                currentFacing = ways[int(tmp)]
                return currentFacing



lines = read_lines('input')
lines = remove_end_of_line(lines) 
directions = []
values = []
for line in lines:
    directions.append(line[0])
    values.append(int(line[1:]))

X = 0
Y = 0
facingDirection = 'E'
size = len(directions)
i = 0
while True:
    if i == size:
        break
    if directions[i] == "F":
        directions[i] = facingDirection
        continue
    elif directions[i] == "E":
        X += values[i]
    elif directions[i] == "W":
        X -= values[i]
    elif directions[i] == "N":
        Y += values[i]
    elif directions[i] == "S":
        Y -= values[i]
    elif directions[i] == "R":
        facingDirection = change_direction(directions[i], values[i], facingDirection)
    elif directions[i] == "L":
        facingDirection = change_direction(directions[i], values[i], facingDirection)
    else:
        print("PROBLEM", directions[i])
    i += 1


print(abs(X)+abs(Y))

