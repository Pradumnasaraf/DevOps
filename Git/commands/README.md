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
