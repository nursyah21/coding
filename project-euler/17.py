def num2words(n):
    num = { 1:'one',
            2:'two',
            3:'three',
            4:'four',
            5:'five',
            6:'six',
            7:'seven',
            8:'eight',
            9:'nine',
            10:'ten',
            11:'eleven',
            12:'twelve',
            13:'thirteen',
            14:'fourteen',
            15:'fifteen',
            16:'sixteen',
            17:'seventeen',
            18:'eighteen',
            19:'nineteen',
            20:'twenty',
            30:'thirty',
            40:'fourty',
            50:'fifty',
            60:'sixty',
            70:'seventy',
            80:'eighty',
            90:'ninety',
            1000:'one-thousand'}

    if n in num:return num[n]

    if len(str(n)) == 2:
        a = n%10
        b = n-a
        return num[b] +'-'+ num[a]

    if len(str(n)) == 3:
        a = n//100
        b = n%100
        hundred = num[a] +'-hundred'
        if b==0: return hundred
        c = n-100*a
        if c in num: return hundred +' and ' + num[c]
        d = n%10
        c -= d
        return hundred +' and '+num[c] +'-'+num[d]
    
res = 0
for i in range(1,1001):
    words = num2words(i)
    res += len([i for i in words if i != ' ' and i != '-'])

print(res)
