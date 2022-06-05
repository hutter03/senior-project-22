# senior-project-22
a polymorphic engine that i coded in about 45 hrs for my highschool senior project

### the main go file
-reads from the two files ending in a .yeppers file extenstion </br>
-creates a random encryption key and encrypts the contents of payloadhex.yeppers using a xor cipher </br>
-the encrypted payload and encryption key are put into decrypt-exec.yeppers as an array of integers </br>
-junk code is then added </br>
-a file named output.c is saved which contains the code to decrypt and execute the </br>

### payloadhex.yeppers
-this is the file of hex values that convert-dump-to-hex-string-for-c.py has made </br>

### decrypt-exec.yeppers
-a c template for decrypting and executing the encrypted binary </br>

### convert-dump-to-hex-string-for-c.py (creative name ik)
-reads a hex dump generated from xxd and formats it for the main go file by inserting a comma between every hex value </br>
