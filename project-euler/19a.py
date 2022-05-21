from calendar import weekday

res = sum(1 for y in range(1901,2001) for m in range(1,13) if weekday(y,m,1) == 6)
print(res)
