Example
===

This demonstrates how mdx executes stuff.

```shell
ls -a
```

It should be able to tolerate multiple steps.

```shell
echo "hello"
```

It should not run steps like this.

```
# this should fail if ran manually but should not
# be executed at all by mdx.
ls asdflkjsdflkjsdffsd
```

Yay!
