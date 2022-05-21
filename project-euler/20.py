a=1
for i in range(2,101):a*=i
num=[int(i) for i in str(a).strip()]

print(sum(num))
