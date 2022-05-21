from num2words import num2words #num2words not builts in

res=0
for i in range(1,1001):
  words = num2words(i)
  num =  [i for i in words if i != '-' and i != ' ']
  res += len(num)

print(res)

