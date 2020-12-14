import re
import fileinput
import copy

def read_lines(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    file.close()
    return lines

def remove_end_of_line(lines):
    for item in range(0, len(lines)):
        lines[item] = lines[item].rstrip()
    return lines

value_data = []
mem = {}
lines = read_lines('input')
lines = remove_end_of_line(lines)

mask = ""
value = ""
valuePlace = 0
for line in lines:
    if line.startswith("mask"):
        splitted = line.split(" = ")
        mask = splitted[1]
        continue
    else:
        splitted = line.split(" = ")
        valuePlace = int(splitted[0][4:-1])
        value = int(splitted[1])

    valuePlaceBinary = bin(valuePlace)[2:]
    newPlace = valuePlaceBinary
    newPlace = newPlace.zfill(36)
    newPlace = list(newPlace)
    places = []
    places.append(newPlace)
    for i in range(len(mask)):
        if mask[i] == "X":
            size = len(places)
            for place in range(0, size):
                tmp = copy.deepcopy(places[place])
                tmp1 = copy.deepcopy(places[place])
                tmp[i] = '0'
                tmp1[i] = '1'
                places.append(tmp)
                places.append(tmp1)
        elif mask[i] == "1":
            size = len(places)
            for place in range(0, size):
                places[place][i] = '1'

    newPlaces = []
    for place in places:
        newPlace = "".join(place)
        newPlaces.append(newPlace)

    for place in newPlaces:
        position = int(place, 2)
        mem[position] = value

sumOfAll = 0
for k in mem:
    sumOfAll += mem[k] 
print(sumOfAll)






