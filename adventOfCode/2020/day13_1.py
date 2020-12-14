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

def calculate_deparute(userTime, busId):
    lastDepartue = 0
    while lastDepartue < userTime:
        lastDepartue += busId
    return lastDepartue

def find_closes(closestDeparture, userTime):
    nextAfterArriving = userTime + 1
    bestBusId = -1
    for item in closestDeparture:
        tmp  = closestDeparture[item] - userTime
        if tmp < nextAfterArriving:
            nextAfterArriving = tmp
            bestBusId = item


    print(bestBusId, nextAfterArriving)
    print(bestBusId*nextAfterArriving)

lines = read_lines('input')
lines = remove_end_of_line(lines)
userTime = int(lines[0])
closestDeparture = {}
print(userTime)
buses = []
splitted = lines[1].split(",")
for c in splitted:
    if c != "x":
        lastDepartue = calculate_deparute(userTime, int(c))
        closestDeparture[int(c)] = lastDepartue

find_closes(closestDeparture, userTime)
