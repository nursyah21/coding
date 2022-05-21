res=1
for i in range(1001, 1,-2):
    k = i**2
    for j in range(4):
       res+= k-j*(i-1)

print(res)
