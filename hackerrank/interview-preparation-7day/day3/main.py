def caesarCipher(s, k):
    # Write your code here
    res = list(s)
    for idx,i in enumerate(s):
        v = ord(i)
        if v >= 65 and v <= 90:
            v += k%26
            if v > 90:v -= 26
        if v >= 97 and v <= 122:
            v += k %26
            if v > 122:v -= 26

        res[idx] = chr(v)

    return ''.join(res)

if __name__ == "__main__":
    print(caesarCipher("nursyah",43))
    print(caesarCipher("nursyah",26))
    # print(caesarCipher("Hello_World!",4))
    print(caesarCipher("6DWV95HzxTCHP85dvv3NY2crzt1aO8j6g2zSDvFUiJj6XWDlZvNNr", 87))
    # print(caesarCipher("6DWV95HzxT", 87))