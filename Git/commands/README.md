- Show commit log history

```bash
git log
```

- Show which file git is tracking and are stagged/unstagged

```bash
git status
```

- Stage a file

```bash
git add [file-name]
```
- Make the repo chnages according to that commit

```bash
git reset [commit-hash]
```
- Change the date and time of last commit

```bash
git commit --amend --date="YYYY-MM-DD HH:MM:SS
```

- Remove/reset all the commits

```bash
git update-ref -d HEAD
```

- Check the git config

```bash
git config --list
```
- Create a new brach

```bash
git branch <branch-name>
```

- Checkout a branch

```bash
git checkout <branch-name>
```

- Create and Checkout branch with a single command

```bash
git checkout -b <branch-name>
```
