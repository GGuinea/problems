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

def change_direction(way, value, x, y):
    chaningSteps = value / 90
    if way == "R":
        for j in range(0, int(chaningSteps)):
            newX = y
            newY = x*(-1)
            x = newX
            y = newY
        return x, y
    else: 
        for j in range(0, int(chaningSteps)):
            newX = y * (-1)
            newY = x
            x = newX
            y = newY
        return x, y



lines = read_lines('input')
lines = remove_end_of_line(lines) 
directions = []
values = []
for line in lines:
    directions.append(line[0])
    values.append(int(line[1:]))

X = 0
Y = 0
waypintX = 10
waypintY = 1
facingDirection = 'E'
size = len(directions)
i = 0
while True:
    if i == size:
        break
    if directions[i] == "F":
        X += waypintX * values[i]
        Y += waypintY * values[i]
    elif directions[i] == "E":
        waypintX += values[i]
    elif directions[i] == "W":
        waypintX -= values[i]
    elif directions[i] == "N":
        waypintY += values[i]
    elif directions[i] == "S":
        waypintY -= values[i]
    elif directions[i] == "R":
        waypintX, waypintY = change_direction(directions[i], values[i], waypintX, waypintY)
    elif directions[i] == "L":
        waypintX, waypintY = change_direction(directions[i], values[i], waypintX, waypintY)
    else:
        print("PROBLEM", directions[i])
    i += 1


print(X, Y)
print(abs(X)+abs(Y))

