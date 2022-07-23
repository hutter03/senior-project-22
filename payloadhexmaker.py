import re
file = open("filedump.txt", "r")
raw = file.read()
no_space = raw.replace('\n', '')
splitting = re.split("(.{2})", no_space)

final = ""
for x in range(len(splitting)):
	if (x % 2) == 0 and x != (len(splitting)-1) and x != 0:
		final = final + ","
	else:
		final = final + splitting[x]

print(final)
