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


## Example of use

- encrypt a message
    Zero width space are many varie‬​‎‎‎‎‏‎‎‎‎‏ties. So, if you are override some operations, You can use secret message.

Enjoy Happy Hacking!
