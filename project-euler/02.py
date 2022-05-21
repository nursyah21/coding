def fib(n):
  if n <= 1: return 1
  return fib(n-1) + fib(n-2)

def fib2(n,cache={}):
  if n <= 1: return 1
  if n not in cache: 
    cache[n] = fib2(n-1,cache) + fib2(n-2,cache)

  return cache[n]


sum=0
for i in range(1,100):
  if fib2(i) >= 4000000: break
  if fib2(i) % 2 == 0: sum+= fib2(i)

print(sum)
