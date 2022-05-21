def prime(n):
  sieve = [True] * n

  def mark(m):
     for i in range(m*2,n,m):
       sieve[i]=False

  for i in range(2, int(n**.5)+1):
     if(sieve[i]):mark(i)
  
  res=0
  for i in range(2,n):
     if sieve[i]:res += i

  return res


print(prime(2000000))
