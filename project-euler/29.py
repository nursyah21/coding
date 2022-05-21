num = set()

for i in range(2,101):
    for j in range(2,101):
        num.add(i**j)

print(len(num))
