from functools import reduce
from itertools import permutations

d='0123456789'

def fact(n):
    if n < 2: return 1
    return reduce(lambda x,y: x*y, [i for i in range(1,n+1)])

perm = permutations([i for i in range(0,10)])

for idx,i in enumerate(perm):
    if idx == 1000000-1:
        print(''.join(map(lambda x:str(x),i)))
        break
