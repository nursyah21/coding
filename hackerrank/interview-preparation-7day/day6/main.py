def legoBlocks(n, m):
    M = 1000000007   # most likely used to prevent integer overflow

    a = [0,1,2,4,8] # a[i] is the number of all wall layouts with width i 
    
    for j in range(5,m+1):  # this formula executes only when we have width 5 or more to generate number of wall layouts with width > 4
        # the number of row arrangements (aka layouts) for width j is the sum of the previous 4 number of row arrangements
        # Examples: 
            # possible arrangements of | | | | | | ->
            # arr(| | | | |x|) + arr(| | | |x|x|) + arr(| | |x|x|x|) + arr(| |x|x|x|x|)
            # possible arrangements of | | | | | | | ->
            # arr(| | | | | |x|) + arr(| | | | |x|x|) + arr(| | | |x|x|x|) + arr(| | |x|x|x|x|)
        # Explanation:
        # | | | | |  <- visual representation of a row with width 4
        # arr(| | | |) is the number of possible layouts for row with width 3
        # | | | |x| is number of possible layouts for row with width 3
            # imagine that the box marked 'x' is taken: already occupied with a single block and cannot change
            # similarly:
                # | | |x|x| means the 2 boxes to the right is fixed (occupied by a single block with width 2)
        # arr(| | | |x|) = arr(| | | |) since the box with 'x' is marked as fixed
        a.append((a[j-1]+a[j-2]+a[j-3]+a[j-4])%M)
    #print("i", [i for i in range(len(a))])
    #print("a - ex", a)

    for i in range(m+1): # this will give us number of possible wall layouts with height n 
        # when expanding from 1 row to n rows, number of possible combinations is x**n
        # see cartesian product: permutations with repetition
        a[i] = pow(a[i],n,M)
    #print("a - n", a, "  -> number of possible layouts that have height n and width i")

    # let r[i] be the number of good layouts that have height n, and width i
    r = [a[i] for i in range(m+1)] # start with all of them
    # r=a if m >= 4, len(r)<len(a) if m < 4
    #print("r", r, "  -> number of good layouts that have height n and width i")
    #print("r[i] -= (r[j]*a[i-j])")
    # define break_point: the leftmost column that has a vertical line all the way through
    # define bad layouts:
    #  - layouts that has a break_point
    #  - layouts that can be split into 2 smaller layouts
    for i in range(1,m+1):
        # for each width, remove bad layouts
        res =0
        for j in range(1,i):
            # moving j from left to right, starting with having the break in column1 and moving it 1 step to the right with every iteration
            # given that the break_point is in column j, a[j-i] is the number of possible layouts for the right side
            # r[j] is the number of good layouts for the left side of column j
            #   ie. the number of ways to get the break point in column j ignoring the right side of column j
            # therefore, r[j] * a[j-1] is the number of bad layouts with the break_point at column j
            #print("i", i, "  j", j, "  r[i]", r[i], "  r[j]", r[j], "  a[i-j]", a[i-j])
            res = (r[j]*a[i-j])
            r[i] -=  res# subtract the number of bad layouts, when the FIRST vertical break in the wall appears at index j 
        #print("i", i, "  r[i]", r[i])
        print(res, r[i],end=" ")
        r[i] %= M # make the computations easier: likely prevents overflow error as well
    # pick out the width needed and return the number there
    print(r)
    return r[m]

#print(legoBlocks(4, 4))
#print(legoBlocks(2, 2))
res = legoBlocks(8, 6)
if res == 402844528: print("success")
else: print("failed")