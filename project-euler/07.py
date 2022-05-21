def isPrime(n):
  if n <= 1: return False
  for i in range(2,int(n**.5)+1):
     if n % i == 0: return False
  return True

def findPrime(n):
  k=1
  for i in range(n):
    if isPrime(i):
      if(k>=10001):return i
      k+=1
    
  


print(findPrime(300000))
