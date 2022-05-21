def fact(n):
  k=1
  for i in range(2,n+1):k*=i
  return k

n=20
print(fact(n*2)/fact(n)/fact(n))
