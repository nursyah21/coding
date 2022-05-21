from functools import lru_cache

@lru_cache(None)
def div(n):
    k=1
    for i in range(2,n//2+1):
        if n%i == 0:k+=i
    return k

@lru_cache(None)
def div2(n):
    total=1
    for x in range(2, int(n**.5)+1):
        if n%x==0:
            total += x
            y = n // x
            if y > x: total += y
    return total

amicable = set()
n=10000

for i in range(1,n):
    j = div2(i)
    if i != j and div2(j) == i:
        amicable.add(i)

print(sum(amicable))


