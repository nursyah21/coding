n=1000000
def chain(n): #loop
  k=1
  while n != 1:
    if n%2 == 0: n //=2
    else: n = 3*n + 1
    k+=1
  return k

def chain2(n,cache={}): #memoization
  if n not in cache:
     if n == 1: cache[n] = 1
     if n % 2: cache[n] = 1 + chain2(3*n + 1)
     else: cache[n] = 1 + chain2(n/2)

  return cache[n]


max_=0
i_=0
for i in range(2,n):
  res=chain2(i)
  if res>max_:
    i_=i
    max_=res

print(f"chain:{max_} result:{i_}")
