def gcd(a,b): 
  if b == 0:return a
  return gcd(b, a%b)

def lcm(a,b): return int(a*b / gcd(a,b))

num=1
for i in range(1,200):
  num=lcm(num,i)

print(num)
    
    
