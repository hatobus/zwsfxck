# zwsfxck
‍‍‍‍‍‍‍‍‍‫​‍‍‍‍‍‍‍‍‌‎
zwsfxck is the extends of brain fxck with zero width spaces.

zero width space have no space and couldn't see from human.
‬​‏‌‍‍‍‍‍‍‍‍‍‫​‍‍‌‎‬​‍‍‍‍‍‍‍
## Operand

```‏‌‍‍‍‍‍‍‍‍‍‫​‍‌‎‬​‍‍‍‍‍‍‏‏‌‍‍‍‍‍‍‍‍‍‫​‍‌‎‬​‏‌‍‍‍‍‍‍‍‍‍‫​‎‎‎‎‎
Increment the data pointer (>) ... U+200B
Decrement th‎‎‎‎‌‎‬e data pointer (<) ... U+200C

Icr​‎‎‎‎‎‎‎‎‏‌‍‍‍‍‍‍‍‍‍‫​‍‍‍‍‍‍‍‍ement the byte at the data pointer (+) ... U+200D
Decrement the byte at the data pointer (-) ... U+200E

O‌‎‬​‏‎‎‎‎‎‎utput the byte at the data po‎‏‍‍‏‍‍‍‍‍‍‍‍‏‎‎‏‍‍‍‍‍‏‎‎‎‎inter (.) ... U+200F
Accept one byte of input, string its value in the byte at the data pointeter (,) ... U202A

If the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next command ([) ... U+202B‎‎‎‏‌‍‍‍‍‍‍‍
If the byte at the data pointer is nonzero, then instead of moving the instruction pointer forward to the next command (]) ... U202C
```
‍‍‫​‎‎‎‎‎‎‎‌‎
## translate
‬​‎‎‎‎‎‎‎‎‏‌‍‍‍‍‍‍‍‍‍‫​‍‍‌‎‬​‍‍‍‍
```
sed -e "s/+/$(echo -ne '\U200d')/g" -e "s/-/$(echo -ne '\U200E')/g" -e "s/>/$(echo -ne '\U200B')/g" -e "s/</$(echo -ne '\U200C')/g" -e "s/\./$(echo -ne '\U200F')/g" -e "s/,/$(echo -ne '\U202A')/g" -e "s/\[/$(echo -ne '\U202B')/g" -e "s/\]/$(echo -ne '\U202C')/g" -i zwsfxck.bf
```
‍‍‍‍‏‌‍‍‍‍‍‍‍‍‍‫​‎‌‎
translated data

```
$ cat zwsample.bf | xxd 

00000000: e280 8de2 808d e280 8de2 808d e280 8de2  ................
00000010: 808d e280 8de2 808d e280 8de2 80ab e280  ................
00000020: 8be2 808d e280 8de2 808d e280 8de2 808d  ................
00000030: e280 8de2 808d e280 8de2 808b e280 8de2  ................
00000040: 808d e280 8de2 808d e280 8de2 808d e280  ................
00000050: 8de2 808d e280 8de2 808d e280 8de2 808b  ................
00000060: e280 8de2 808d e280 8de2 808d e280 8de2  ................
00000070: 808c e280 8ce2 808c e280 8ee2 80ac e280  ................
00000080: 8be2 808f e280 8be2 808d e280 8de2 808f  ................
00000090: e280 8de2 808d e280 8de2 808d e280 8de2  ................
000000a0: 808d e280 8de2 808f e280 8fe2 808d e280  ................
000000b0: 8de2 808d e280 8fe2 808b e280 8ee2 808f  ................
000000c0: e280 8ee2 808e e280 8ee2 808e e280 8ee2  ................
000000d0: 808e e280 8ee2 808e e280 8ee2 808e e280  ................
000000e0: 8ee2 808e e280 8fe2 808c e280 8de2 808d  ................
000000f0: e280 8de2 808d e280 8de2 808d e280 8de2  ................
00000100: 808d e280 8fe2 808e e280 8ee2 808e e280  ................
00000110: 8ee2 808e e280 8ee2 808e e280 8ee2 808f  ................
00000120: e280 8de2 808d e280 8de2 808f e280 8ee2  ................
00000130: 808e e280 8ee2 808e e280 8ee2 808e e280  ................
00000140: 8fe2 808e e280 8ee2 808e e280 8ee2 808e  ................
00000150: e280 8ee2 808e e280 8ee2 808f e280 8be2  ................
00000160: 808d e280 8f0a                           ......
```

## Build
```
$ go build -o ./zwsfxck
```

```
$ ./zwsfxck zwsample.bf

Hello world
```



## Example of use

- encrypt a message
    Zero width space are many varie‬​‎‎‎‎‏‎‎‎‎‏ties. So, if you are override some operations, You can use secret message.

Enjoy Happy Hacking!
