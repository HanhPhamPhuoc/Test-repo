# Git-squash tutorial :
To squash commits into one, we use git rebase :
Syntax : 

```
$ git rebase -i <branch-name>
```

Two commit which is current commit and the commit selected will be displayed.
You can choose which to keep and which to squash, or rename commmit message
To `pick`, type `pick` following with the commit ID, or else, to `squash`, type `squash` or `s`, following with the commit id. There are many other option, you can find the help for more information.
