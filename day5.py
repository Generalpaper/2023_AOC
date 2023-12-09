D = open("inputs/5.txt","r+").read()
# L = D.split('\n')

parts = D.split('\n\n')

seed, *others = parts
seed = seed.split(':')[1].split()
for s in seed:
    for map in others:
        lines = map[1:]
    print(lines)


# abandoning this btw