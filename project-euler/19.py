n30=(4,6,9,11)
weekday=('sunday','monday','tuesday','wednesday','thursday','friday','saturday')
months=[]

#1900 january 1 monday


n=0
for i in range(1,13):
    if i == 2:
        n+=28
    elif i in n30: 
        n+=30
    else: 
        n+=31
    
    try:print(i,n,weekday[n%7])
    except:print(n)
    
#print(months)
