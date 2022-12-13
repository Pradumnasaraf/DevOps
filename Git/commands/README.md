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

## **Push**

- Push chnages to remote repo
  
  ```bash
  git push <remote> <branch>
  ```
  
- force the push even if it results in a non-fast-forward merge
  
  ```bash
  git push <remote> --force # Use the flag in case you know what you’re doing.
  ```

- Push all of your local branches to the specified remote.
  
  ```bash
  git push <remote> --all
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

## **Configure**

- name and email address
  
  ```bash
  git config --global user.name "username" 

  git config --global user.email "email address"
  # Your username and email address should be the same as the one used with your git hosting provider i.e. github, bitbucket etc
  ```

- default editor
  
  ```bash
  git config --global core.editor "vim"
  # Use "code --wait" to set VS Code as default editor
  ```

- external diff tool
  
  ```bash
  git config --global diff.external "meld"
  # You can change "meld" to "emerge" or "kompare"
  ```

- default merge tool
  
  ```bash
  git config --global merge.tool "meld"
  # You can change "meld" to "emerge", "gvimdiff", "kdiff3", "vimdiff", and "tortoisemerge"
  ```

- color
  
  ```bash
  git config --global color.ui auto # Enables colorization of CLI output
  ```

- add the GPG key
  
  ```bash
  git config --global user.signingkey <your-secret-gpg-key>
  # If you’re taking work from others on the internet and want to verify that commits are actually from a trusted source.
  ```

## **Delete**

- Branch
  
  ```bash
  git branch -D <branch name>
  ```

- Tag
  
  ```bash
  git tag -d v<tag version>
  ```

- Remote
  
  ```bash
  git remote rm <remote>
  ```

- Untracked files
  
  ```bash
  git clean -<flag>
  # replace -<flag> with:
  # -i for interactive command
  # -n to preview what will be removed
  # -f to remove forcefully
  # -d to remove directories
  # -X to remove ignored files
  ```

- Files from index
  
  ```bash
  git rm --cached <file or dir>
  ```

- Local branches that don't exist at remote
  
  ```bash
  git remote prune <remote-name>

  ```

## **Merge**

- Another branch to current branch
  
  ```bash
  git merge <branch-name>
  ```

- Merge a single file from one branch to another.
  
  ```bash
  git checkout <branch name> <path to file> --patch
  ```

## **Modify**

- last/latest commit message
  
  ```bash
  git commit --amend
  ```

- Repo's remote url
  
  ```bash
  git remote set-url <alias> <url> # <alias> is your remote name e.g origin
  ```

- Change date and time of the commit

  ```bash
  git commit --amend --date="YYYY-MM-DD HH:MM:SS
  ```

## **Pull**

- Pull the specified remote’s copy of the current branch and merge it into local
  
  ```bash
  git pull <remote>
  ```

- Gives output during a pull (displays the pulled content and the merge details)
  
  ```bash
  git pull --verbose
  ```

- Pull changes and prevent merge conflicts

  ```bash
  git pull --ff-only # applies the remote changes only if they can be fast-forwarded
  ```

## **Rebase**

- An origin branch into working branch
  
  ```bash
  git pull --rebase origin <branch name>
  ```

- Local branch into my working branch
  
  ```bash
  git rebase <branch name>
  ```

- And skip commits
  
  ```bash
  git rebase --skip
  # In case of conflicts use this command to discard of your own changes in the current commit
  # and apply the changes from an incoming branch
  ```

- And continue after resolving conflicts
  
  ```bash
  git rebase --continue
  # Use it whenever conflicts detected therefore you can resolve these conflicts manually and use this command to continue your rebase operation
  ```

## **Rename**

- Branch
  
  ```bash
  git branch -m <new name> # while working in the branch
  git branch -m <old name> <new name> # from outside the branch
  ```

- Remote
  
  ```bash
  git remote rename <oldname> <newname>
  ```

## **Reset**

- a specific commit

  ```bash
  git revert <commit-hash> # Get a commit hash by using `git log`
  ```

- a specific file

  ```bash
  git checkout <repo>/<branch> <filename>
  ```

- To last commit

  ```bash
  git reset --hard
  ```

- To last commit on remote branch

  ```bash
  git reset --hard <repo>/<branch>
  ```

- Remove/reset all commits

  ```bash
  git update-ref -d HEAD
  ```

## **Squash**

- commits in pull request into single commit

  ```bash
  git rebase -i <branch name>
  ```

- last n number of commit into one

  ```bash
  git reset --soft HEAD~N # N for number of commits you want to squash
  git add .
  git commit -m <message>
  ```

## **Stash**

- Create stash (Tracked and Untracked files)

  ```bash
  git stash
  ```

- Create a new branch and apply stash

  ```bash
  git stash branch <branch name> <stash id>
  ```

- Delete

  ```bash
  git stash clear # all stashed changes
  git stash drop <stash id> # specific stash
  ```

- View the contents of a stash

  ```bash
  git stash show -p <stash id> #Leave stash ID to see the latest stash
  ```

- Apply

  ```bash
  git stash apply
  git stash apply <stash id> # stash id can be gotten when you run git stash list
  git stash pop <stash id> # Stash id optional. Add it if you want to apply and delete a specific stash otherwise leave to pop the latest stash
  ```

- View list of stashed changes

  ```bash
  git stash list
  ```

## **View**

- Status of project

  ```bash
  git status
  ```

- Commit(s) log

  ```bash
  git log    # View all logs
  git log -n # for last n number of commits
  # to exit you have to press (q)
  ```

- uncommitted changes

  ```bash
  git diff
  ```

- Committed changes

  ```bash
  git diff
  ```

- repo's remote url

  ```bash
  git remote -v
  ```

- repo's remote url

  ```bash
  git branch # The active branch is prefixed with *
  ```

- repo's remote url

  ```bash
  git tag
  ```
