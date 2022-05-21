maxPal=0
for i in range(999,900,-1):
  for j in range(999,99,-1):
    num=str(i*j)
    if(num == num[::-1]):
       if(i*j > maxPal):maxPal=i*j
       break

print(maxPal)
