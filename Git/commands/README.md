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
- To change the date and time of last commit

```bash
git commit --amend --date="YYYY-MM-DD HH:MM:SS
```

- To remove/reset all the commits

```bash
git update-ref -d HEAD
```

- `git config --list`
- `git branch <branch-name>` - Create a new branch
- `git checkout <branch-name>` - Navigate among the created branches
