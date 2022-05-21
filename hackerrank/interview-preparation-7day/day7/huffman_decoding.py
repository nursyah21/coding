def decodeHuff(root, s):
    # Enter Your Code Here
    # iterate s
    curr = root
    for c in s:
        if c == '0':
            curr = curr.left
        else:
            curr = curr.right
        if curr.left == curr.right:
            print(curr.data, end="")
            curr = root


if __name__ == "__main__":
    ip = open("input00_huffman.txt", "r")

    freq = {}

    for i in ip.read():
        if freq.get(i) == None:
            freq[i] = 1
        else:
            freq[i] += 1

    print(freq)
