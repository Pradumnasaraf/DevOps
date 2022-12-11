- `whereis` - Find the path of that executable file.

### List operation
- `ls` - Shows list.
    - `-a` - Hidden file.
    - `-l` - Persmission.
    - `-R` - Show sub dir.

### Changing dir operation
- `cd <folder-name>` - Change Dir.
- `cd ..` -   Go one Directory back.
- `cd` -    Go to home.
- `cd ../<foldername>` - Open a previous dir folder.
- `cd <path>` - Open a dir with the path.

### File/Folder Ope.
- `mkdir <new-dir-name>` - Create a new folder.
- `mkdir -p test/test1/test2` - Craete a dir between two directories.
- `touch <new-file-name>` - create a blank file.
- `pwd` - Present working directory.
- `cat <filename>` - Display file content.
- `cat > <new-file-name>` - Create a file.
- `dd if=/dev/zero of=bos_dosya bs=4G count=1`- create empty file with zeros
- `cat >> <filename>` - Append the file.
- `cat <filename> <filename2> ` - Display 2 files at a time.
- `cat <filename> <filename2> > <newfile-name>` - Merge both of file content in a single one.
- `cat <file-name> | tr > <new-file-name>` - Translate the file.
- `cut -c  1-2 <filename>` - cut the file with colum wise
- `echo "Hello" >> <file-name>`




+
### 
- `man <commad name>` - Know about the command usages and options.
- `man <commad name>` - know about the command.

### File/Folder operation
- `cp <file-name> <new-fie-name>` - Make a copy of a file in the current location.
- `mv <file-name> <dir-path>` - Move a file from a one dir to a another.
- `mv <file-name> <new-fie-name>` - Rename a file.
- `mv -R <dir-name> <dir-path>` - Move Dir
- `rm <file-name>` - Remove a file
- `rm -R <file-name>` - Delete a folder with dir included.
- `head <file-name>` - Will display first 10 lines of a file.
- `tail <file-name>` - Will display last 10 lines of a file.
    -`-n 2` - will display last 2 lines.
- `diff <file-1> <file-2>` - Show diff between the two files.
- `locate <file>` - To find out the file.  
- `find <file/folder-name>` - Find a file/folder.
- `find <dir-name>` - Find files inside the dir
- `find .-type d` - Show only dir.
    - `.-type f` -Sshow only files.
    - `.-type f -name "*.txt"` - Show only files with that specfic name.
    - `.-type f -iname "*.txt"` - Show only files with that specfic name - not case sentive (i)
    - `.-type f -mmin -20` - Show files which modify less than 20 min ago.
    - `.-type f -mmin +20` - show files which modify more than 20 min ago.
    - `.-type f -maxdepth 2` - Will only show 1 folder deep.
    - `.-size +1k` - will only show file/folder with size of 1kb


### System commands
- `ps aux` - processes which are running
- `df` - Check the capacity and storage details.
    - `m` - In megabyte)  or 
    - `hg` - (gigabyte).
- `du` - Disk usages capcity 
    - `-h` (human readable)
- `echo` - Get a output of a string
- `echo $PATH` - Check the path varibale
- `sudo` - Admin command
- `sudo chown root text.txt` - chnage owner
- `!<command-name>` - Run the pevious command
- `git add .; git commit -m "message"` - Run mutiple commands at a time
- `sort <file-name>"` - sort the file
- `job`
- `wget <url>` - download the file from the URL
- `top` - what processes are running
- `kill <process-id>` -stop that process
- `Uname` -
- `zip <file-1> <file-2>` - Zip Two or more files
- `Unzip <file-name>` - Unzip files
- `useradd <name>`
- `passwd <name>`
- `uname -<flag>` -o -m -r
- `lscpu` - get cpu details
- `free` - free memory 
- `vmstat` - virtual memory
- `lsof` - list all the open file
- `xdg-open <file-fath>` - open the folder (graphical window) of a file/folder with path.
- `xdg-open .` - open the folder of the current directory.
- `vi ~/.bashrc` - set your Alias
- `echo -n 'username' | base64` 

### Networking
- `nslookup google.com` 
- `netstat` - 
- `hostname` - 
- `whoami` - 
- `ping google.com`


### Permissions
- `chmod u=rwx,g=rxw,o=rwx <file-name>` READ, WRITE AND EXECUTE
- `chmod 777 <file-name>` - 4- Read, 2- Write, 1 - Execute
- `find . -perm 777 ` - will only show file/folder with size of 1kb
- `grep <keyword> <file-name>` - To search if the keyword is presnt in the file or not
- `grep -w <keyword> <file-name>` - To search if the keyword is presnt in the file or not (complete word)
- `grep -i <keyword> <file-name>` - To search if the keyword is presnt in the file or not (not case sens)
- `grep -n <keyword> <file-name>` - To search if the keyword is presnt in the file or not (Line number)
- `grep -B <keyword> <file-name>` - Show Line before that keyword
- `grep -win <keyword> ./*.txt` - To search if the keyword is presnt in the file in current dir
- `grep -win -r <keyword> .*` 
- `history | grep "ls -l"` - Pipping, we filter out the things

```
histroy
!<number-from-history>
```

# Operators

- `ping google.com & pingfacebook.com`
- `echo "google.com" && echo "facebook.com"` - second will oly run if first is sucessful
- `echo "google.com" && {echo "facebook.com"; eco "pradumnasaraf.co"}`
- `echo "google.com" || echo "pingfacebook.com"` 
- `echo "google.com" || echo "pingfacebook.com"` 
- `rm -r !(file.txt)`

- `printevnv` - to print all th env.
