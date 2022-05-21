def div(n,k=1): #complexity n/2
  for i in range(1,n//2+1):
    if (n%i): continue
    k+=1
  return k

def div2(n,k=1): #complexicty log(n)
  for i in range(1,int(n**.5)):
    if (n%i): continue
    k+=1

  if k%2 == 0: return k*2
  return (k-1)*2

def anum(n,a=0):
  for i in range(1,n+1):a+=i
  return a

def solve():
  n=1
  a=anum(n)
  t=384

  for i in range(n+1,100000+1):
    a+=i
    res=div2(a)
    if(res >= 500):
      print(f"i: {i} a:{a} res:{res}")
      break

if __name__ == "__main__":
  solve()
