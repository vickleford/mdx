mdx
===

Markdown Executor (mdx) is intended to help test guides written in markdown.

Usage
---

Run mdx and pass it a filename to what you want to execute.

```shell
mdx demo/sample.md
```

It scans for blocks of code in the markdown labeled shell and tries to execute them. If it encounters any errors in the sh, it exits with a return code 1 indicating failure.

Motivation
---

Don't maintain a separate test runner for each guide.

Avoid the drift between the tester and the guide itself.

Make somebody (something?) run your actual guide every time; validate the guide itself instead of only validating some arbitrary assertions we are making about the files the guide uses.

Wishlist
---

A way to tell mdx about a guid's pre-requisites. For example, run the bootstrap before running the Ambarli guide.