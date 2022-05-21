num=open("p067_triangle.txt").read()
num = [[int(j) for j in i.split()] for i in num.split('\n')]


for i in range(len(num)-1,0,-1):
    for j in range(len(num[i])-1):
        num[i-1][j] += max(num[i][j], num[i][j+1])

print(num[0][0])
