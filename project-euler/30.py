res=0
for i in range(10, 6*(9**5)):
    a=sum(x**5 for x in map(int, str(i)))
    if a == i:
        res += a

print(res)
        
