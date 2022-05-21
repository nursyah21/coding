def res():
  for i in range(1,999):
    for j in range(i,999):
      k= 1000 - i - j
      if k>0:
        if k**2 == i**2 + j**2:
         return i*j*k
         
print(res())
      

       
