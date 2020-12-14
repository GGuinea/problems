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


lines = read_lines('input')
lines = remove_end_of_line(lines)

splitted = lines[1].split(",")

buses = []

maxBusId = 1
for bus in splitted:
    if bus == "x":
        buses.append(bus)
    else:
        buses.append(int(bus))

f = 1
nr = 1
time = 1
for bus in buses:
    if bus == "x":
        nr += 1
        continue
    while(time + nr) % int(bus) > 0:
        time += f
    f *= bus
    nr += 1

print(time+1)

#print(busMultiplication)
#print(len(buses))
