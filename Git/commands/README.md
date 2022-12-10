# Git commands

## **Add**

- New changes
  
  ```bash
  git add <file.ext> # To add a specific file
  
  git add . # To add all the files in the current directory
  ```

- New branch
  
  ```bash
  git branch <new name> # and remain in the current branch 
  
  git checkout -b <new name> # and switch to the new branch

  git checkout -b <new name> <another branch> # From another branch

  ```

- New remote repository
  
  ```bash
  git remote add <shortname> <url>
  ```

- Annotated tag
  
  ```bash
  git tag -a v1.4 -m "my version 1.4"
  git push --tags
  ```

## **Cherry-pick**

- An commit from the origin branch into my working branch
  
  ```bash
  git cherry-pick <commit-hash> <commit-hash>
  ```

## **Clone**

- Existing repo into a new directory
  
  ```bash
  git clone <repo-url> <directory> # Replace "directory" with the directory you want
  ```

- Existing repo into the current directory
  
 ```bash
  git clone <repo-url> . # The current directory is represented with a "."
  ```

- existing repo along with submodules into the current directory
  
 ```bash
  git clone --recurse-submodules <repo-url> . 
  ```

- submodules after cloning the existing repo
  
 ```bash
  git submodule update --init --recursive 
  ```

## **Commit**

- commit all local changes in tracked files
  
  ```bash
  git commit -a
  ```

- commit all staged changes
  
  ```bash
  git commit -m <message> # Replace <message> with your commit message.
  ```

## **Compare two commits**

- and output results in the terminal
  
  ```bash
  git diff <sha1> <sha2> # the sha hash of the commits you want to compare.

  ```

- and output result to a file
  
  ```bash
  git diff <sha1> <sha2> > diff.txt
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
