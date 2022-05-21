from numpy import arange


def arraySum(numbers):
    if len(numbers) == 0:
        return 0
    return numbers[0] + numbers[1:]
    res = 0
    for i in numbers:
        res += i
    return res

def reverseList(arr):
    n = len(arr)
    for i in range(n//2):
        arr[i],arr[n-1-i] = arr[n-i-1], arr[i]
    
    return arr

#23 // 2 = 11, 1
#12 // 2 = 6, 0
#5 // 2 = 2, 1
#2 // 2 = 0, 0
#2 // 2 = 1, 0
#3 // 2 = 1, 1
#1 // 2 = 0, 0

def fourBit(num):
    # Olog(n)
    res = ""
    while num > 0:
        remainder = num %2
        num //= 2
        res += str(remainder)

    return res[3]


""" def binary(number):
    res = ""
    for 
    return res
 """


def dnaStrant(s):
    n = len(s)
    #reverse string
    res = list(s)

    for i in range(n//2): 
        res[i], res[n-i-1] = res[n-i-1], res[i]
    

    for idx,i in enumerate(res):
        if i == "G":
            res[idx] = "C"
        elif i == "C":
            res[idx] = "G"
        elif i == "T":
            res[idx] = "A"
        elif i == "A":
            res[idx] = "T"
    
    
    return ''.join(res)


#3,6
#3,4,5,6
#3^4, 3^5, 3^6
#4^5 4^6
#5^6

def bitLogic(lo, hi, k):
    #On^2
    res = -9999

    for i in range(lo,hi):
        for j in range(i+1, hi+1):
            temp = i ^ j
            if temp <= k and temp > res:
                res = temp

    return res


if __name__ == "__main__":
    #print(arraySum([3,13,4,11,9]))
    #print(fourBit(23))
    """ print(reverseList([1,3,2,4,5]))
    print(reverseList([1,3,2,5])) """
    #print(dnaStrant("GTACAG"))
    print(bitLogic(3,5,6))