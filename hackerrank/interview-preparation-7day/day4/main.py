def digsum(x):
    return str(sum((int(i) for i in list(x))))

if __name__ == "__main__":
    n = "861568688"
    print(digsum(n))