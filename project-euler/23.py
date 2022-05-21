from itertools import combinations_with_replacement as cwr
def prop(n,num=1):
    return 1+sum([i for i in range(2,n//2+1) if n%i==0])

def prop2(n):
    n=[i if i == n/i else i+n/i for i in range(2,1+int(n**.5)) if n%i == 0]
    return sum(n)+1

maxp = 28123

abundant = [n for n in range(2,maxp) if prop2(n) > n]
sumAbundant = set()

for i in range(len(abundant)):
    for j in range(i+1):
        s=abundant[i]+abundant[j]
        sumAbundant.add(s)
        if s>28123:break

res=sum(set([n for n in range(28123)]) - sumAbundant)

print(res)


