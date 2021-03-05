"""
import time
    start = time.process_time()
    stop = time.process_time()
    result = stop - start
    print(result)
"""

# BINARY TO DECIMAL
def bintodec(x):
    if isinstance(x, str):
        print(str(x) + " is the original string.")
        y = []
        z = []
        value = 0
        power = 0
        for i in range(len(x)):
            y.append(x[i])
        for j in range(len(y)):
            z.append(y.pop())
        for k in range(len(z)):
            if z[k] == "1":
                value += (int(z[k]) * (2**power))
            power += 1
        return value
    else:
        raise TypeError("You must use string values for this converter.")

# DECIMAL TO BINARY
def dectobin(x):
    print(str(x) + " is the original string.")
    y = int(x)
    z = bin(y)
    return z

# DECIMAL TO HEXADECIMAL
# Tested vs. function that utilizes list(reverse()) & this is faster
def dectohex(x):
    print(str(x) + " is the original string.")
    if x > 9:
        value = []
        value2 = []
        for i in range(x):
            if x >= 16:
                y = int(x % 16)
                value.append(y)
                x = int(x // 16)
            else:
                y = int(x % 16)
                value.append(y)
                value.append("0x")
                break
        for j in range(len(value)):
            value2.append(value.pop())
        for k in range(len(value2)):
            if value2[k] == 10:
                value2[k] = 'A'
            if value2[k] == 11:
                value2[k] = 'B'
            if value2[k] == 12:
                value2[k] = 'C'
            if value2[k] == 13:
                value2[k] = 'D'
            if value2[k] == 14:
                value2[k] = 'E'
            if value2[k] == 15:
                value2[k] = 'F'
    else:
        value2 = []
        value2.append(x)
    for l in range(len(value2)):
        value2[l] = str(value2[l])
    value2 = ''.join(value2)
    return value2

# CONVERT HEXADECIMAL TO DECIMAL
def hextodec(x):
    print(str(x) + " is the original string.")
    if isinstance(x, str):
        value = []
        value2 = 0
        for i in range(len(x)):
            value.append(x[i])
        if value[0] == '0' and value[1] == 'x' or value[0] == '0' and value[1] == 'X':
            value.remove(value[0])
            value.remove(value[0])
        for j in range(len(value)):
            if value[j] == '0' or value[j] == '1' or value[j] == '2' or value[j] == '3' or value[j] == '4' or value[j] == '5' or value[j] == '6' or value[j] == '7' or value[j] == '8' or value[j] == '9':
                value[j] = int(value[j])
            if value[j] == 'A':
                value[j] = 10
            if value[j] == 'B':
                value[j] = 11
            if value[j] == 'C':
                value[j] = 12
            if value[j] == 'D':
                value[j] = 13
            if value[j] == 'E':
                value[j] = 14
            if value[j] == 'F':
                value[j] = 15
        for k in range(len(value)):
            try:
                value[k] = (int(value[k]) * int(16**((len(value)-k)-1)))
            except:
                raise TypeError("All HEX characters, ie ABCDEF, must be capitalized in string values.")
        for l in range(len(value)):
            value2 += value[l]
        return value2
    else:
        raise TypeError("String values must be used for this HEX converter.")

def bintohex(x):
    y = bintodec(x)
    z = dectohex(y)
    return z

def hextobin(x):
    y = hextodec(x)
    z = dectobin(y)
    return z

# EXECUTE
hw62231 = hextodec("0x55866AE76080")
hw62232 = hextodec("0x55866AE76088")
hw62233 = hextodec("0x55866AE77FC0")

hw62 = hw62232 - hw62231
hw63 = hw62233 - hw62231

print(str(hw62231))
print(str(hw62232))
print(str(hw62233))
print(str(hw62))
print(str(hw63))

"""
test1 = bintodec("1111")
print(test1)
test2 = dectohex(test1)
print(test2)
test3 = hextodec(test2)
print(test3)
test4 = dectobin(test3)
print(test4)
test5 = hextodec("0xAAABF112149978973987DE999930281")
print(test5)
test6 = dectobin(test5)
print(test6)
test7 = bintodec(test6)
print(test7)
test8 = dectohex(test7)
print(test8)
test9 = hextobin(test8)
print(test9)
test10 = bintohex(test9)
print(test10)
print(hextodec("CC"))
print(hextodec("157"))
print(hextodec("D8"))
print(hextodec("BB29"))
print(hextodec("7A"))
print(hextodec("8177"))
print(hextodec("A011"))
print(hextodec("99"))
print(hextodec("2B36"))
print(hextodec("FACE"))
print(hextodec("8DB3"))
print(hextodec("FF"))
print(bintodec("0010"))

"""
