names = open('p022_names.txt').read().replace('"','').split(',')
names.sort()
hasil=0

for idx, i in enumerate(names):
    res = sum([ord(j)-64 for j in i]) * (idx+1)
    hasil += res


print(hasil)
    
#print()
