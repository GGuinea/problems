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

lastForBusId={}
def calculate_deparute(timestamp, busId):
    if busId == "x":
        return 1
    lastDepartue = lastForBusId[busId]
    while lastDepartue < timestamp:
        lastForBusId[busId] = lastDepartue
        lastDepartue += busId
        if timestamp == lastDepartue:
            return 1 
    return 0 

def find_timestamp(buses):
    timestamp = 0
    howManyTrue = 0
    while True:
        howManyTrue = 0
        tmp = timestamp
        for bus in buses:
            result = calculate_deparute(tmp, bus)
            if result == 0:
                break
            howManyTrue += result
            tmp+=1
        if howManyTrue == len(buses):
            print("LEGIT:",  timestamp)
            break
        timestamp += buses[0]

lines = read_lines('input')
lines = remove_end_of_line(lines)

splitted = lines[1].split(",")

buses = []

for bus in splitted:
    if bus == "x":
        buses.append(bus)
    else:
        buses.append(int(bus))
        lastForBusId[int(bus)] = 0
print(len(buses))
find_timestamp(buses)
