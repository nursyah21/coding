def isPrime(n):
  if n <= 1: return False
  for i in range(2,int(n**.5)+1):
     if n % i == 0: return False
  return True

n=600851475143
i=2

while i*i < n:
  break
  while n%i == 0: n /= i
  i += 1


for i in range(int(n**.5),2,-1):
  if(n %i == 0 and isPrime(i)):
     print(i)
     break
