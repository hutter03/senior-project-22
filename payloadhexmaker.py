import re
file = open("filedump.txt", "r")
raw = file.read()
no_space = raw.replace('\n', '')
splitting = re.split("(.{2})", no_space)
final = splitting[0]
for x in range(len(splitting)-1):
        x += 1
        if (x % 2) == 0 and x != (len(splitting)-1):
                final += ","
        else:
                final += splitting[x]
print(final)
