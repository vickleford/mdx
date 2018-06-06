Linked Code Blocks
===

Sometimes you need code blocks to be "linked" because they are doing related things.
For example, you may do some related things but add explanations in between.
Alternatively, there could be a cleanup step at the end.

Let's try it out. This sets up an environment variable to link to other code blocks.

```shell
export WHATEVER="some_name.txt"
```
Now let's use it.

```shell
echo "heehaw" > $WHATEVER
cat $WHATEVER
```

You should hear a donkey in the output.

```
heehaw
```

Good. Time to clean up.

```shell
rm $WHATEVER
test ! -e $WHATEVER
```

That was fun.
