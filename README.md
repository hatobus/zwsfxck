# zwsfxck

zwsfxck is the extends of brain fxck with zero width spaces.

zero width space have no space and couldn't see from human.

## translate

```
sed -e "s/+/$(echo -ne '\U200d')/g" -e "s/-/$(echo -ne '\U200E')/g" -e "s/>/$(echo -ne '\U200B')/g" -e "s/</$(echo -ne '\U200C')/g" -e "s/\./$(echo -ne '\U200F')/g" -e "s/,/$(echo -ne '\U202A')/g" -e "s/\[/$(echo -ne '\U202B')/g" -e "s/\]/$(echo -ne '\U202C')/g" -i zwsfxck.bf
```
