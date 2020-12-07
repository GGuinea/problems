import re
import fileinput
from collections import deque, defaultdict
def read_lines(filename):
    file = open(filename, 'r')
    lines = file.readlines()
    file.close()
    return lines

target = 'shinygoldbag'
PARENT = defaultdict(list)
CONTENTS = defaultdict(list)
lines = list(fileinput.input())
lines.append('')
for line in lines:
    line = line.strip()
    if line:
        words = line.split()
        container = words[0]+words[1]+words[2]
        container = container[0:-1]
        if words[-3] == 'no':
            continue
        else:
            idx = 4
            while idx < len(words):
                bag = words[idx]+words[idx+1]+words[idx+2]+words[idx+3]
                if bag.endswith('.'):
                    bag = bag[:-1]
                if bag.endswith(','):
                    bag = bag[:-1]
                if bag.endswith('s'):
                    bag = bag[:-1]
                n = int(bag[0])
                while any([bag.startswith(d) for d in '0123456789']):
                    bag = bag[1:]

                PARENT[bag].append(container)
                CONTENTS[container].append((n, bag))
                idx += 4

SEEN = set()
Q = deque([target])
while Q:
    x = Q.popleft()
    if x in SEEN:
        continue
    SEEN.add(x)
    for y in PARENT[x]:
        Q.append(y)

print(len(SEEN)-1)

def count_all(bag_color):
    answer = 1
    for (n,y) in CONTENTS[bag_color]:
        answer += n*count_all(y)
    return answer 

print(count_all(target)-1)


