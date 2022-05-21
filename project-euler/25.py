from functools import *

@lru_cache(maxsize=None)
def fib(n):
  if n < 2: return 1
  return fib(n-2) + fib(n-1)

for i in range(100000):
  res=fib(i)
  lres=len(str(res)) 
  if lres == 1000:
    print(i+1,res,lres)
    break
