Assertions
===

Running the guides are good, right? But how do you actually make assertions about the results of a guide when verifying it with Markdown Executor?

Let's demo it by testing our own guide. Assertions will be marked by an html comment with shell commands but the comment must be prefixed with mdx.

Let's assume this guide teaches how to use grep.

First, you have to create a file to grep through.

```shell
touch no_excuses
```

Sweet, I like that song too. Now we need to add some lyrics.

```shell
cat > no_excuses <<EOF
every day it's something
hits me all so cold
find me sitting by myself
no excuses that i know
EOF
```

After that and a tiny guitar solo, we can begin using grep. Since we have no excuses, let's search for some excuses in the lyrics.

```shell
grep excuses no_excuses
```

You should see the correct line in the output

```
$ grep excuses no_excuses
no excuses that i know
```

Notice how you didn't have to pipe output from `cat` to `grep` since it knows how to take file names as part of its arguments.

That was fun but let's make sure we have the right number of excuses

```shell
grep excuses no_excuses
```

You should see the correct output

```
$ grep excuses no_excuses
no excuses that i know
$
```

<!--mdx
num_excuses=$(grep -c excuses no_excuses)
if [ 1 -ne $num_excuses ]; then
    exit 1
fi
# oh, and be sure comments are ignored...
-->

And that's how we make assertions that the guide is correct.

Don't forget to clean up.

```shell
rm no_excuses
```