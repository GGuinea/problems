import re
import fileinput
from itertools import combinations
import copy

data = [list(l.strip()) for l in list(fileinput.input())]

ROW = len(data)
COLUMN = len(data[0])

while True:
    newList = copy.deepcopy(data)
    change = False
    for r in range(ROW):
        for c in range(COLUMN):
            occupied = 0 # counter for occupied sits around you
            for dr in [-1,0,1]: # three values in loop for row, actual(0) up and down
                for dc in [-1,0,1]: # same as above but for coulumn
                    if not (dr == 0 and dc ==0):  # we are not intreseted in your current possiton
                        rr = r + dr # rr = row to check, r = the main iterator, dr = inner iterator for closest neigborhood
                        cc = c + dc # same as above but for column
                        if 0 <= rr < ROW and 0 <=cc <COLUMN and data[rr][cc] == '#': # cheching if we are in boudries, and if closest sit is occupied
                            occupied += 1 # if yes, then we are counting neigbour
            if data[r][c] == 'L': # checking if current postiion is free
                if occupied == 0:  # checking if there is no neigbourhoods
                    newList[r][c] = "#" # if no neigbourhoods then we are sitting
                    change = True # flag to not break the loop
            elif data[r][c] == "#" and occupied >= 4: # checking if current postiion is occupied, and if has many neigbourhoods
                    newList[r][c] = "L" #if yes then we are going away
                    change = True # not to break the loop
    if not change: # checking if in iteration was any change
        break # if not then we are ending
    data = copy.deepcopy(newList) # if program continue, we have to make deep copy of list

ans = 0
for i in range(ROW):
    for j in range(COLUMN):
        if data[i][j] == "#":
            ans+= 1
print(ans)

