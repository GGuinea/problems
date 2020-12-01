def read_integeres(filename):
    with open(filename) as f:
        return [int(x) for x in f]


tab = read_integeres("input")
for x in range(len(tab)):
    for y in range(x+1,len(tab)):
        for z in range(y+1,len(tab)):
            if tab[x] + tab[y] + tab[z] == 2020:
                print(tab[x]*tab[y]*tab[z])
