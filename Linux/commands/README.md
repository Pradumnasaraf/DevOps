- `whereis` - Find the path of that executable file.

### List operation
- `ls` - Shows list.
    - `-a` - Hidden file.
    - `-l` - Permission.
    - `-R` - Show sub dir.

### Changing dir operation
- `cd <folder-name>` - Change Dir.
- `cd ..` -   Go one Directory back.
- `cd` -    Go to home.
- `cd ../<foldername>` - Open a previous dir folder.
- `cd <path>` - Open a dir with the path.

### File/Folder Ope.
- `mkdir <new-dir-name>` - Create a new folder.
- `mkdir -p test/test1/test2` - Create a dir between two directories.
- `touch <new-file-name>` - create a blank file.
- `pwd` - Present working directory.
- `cat <filename>` - Display file content.
- `cat > <new-file-name>` - Create a file.
- `dd if=/dev/zero of=bos_dosya bs=4G count=1`- create empty file with zeros
- `cat >> <filename>` - Append the file
- `cat  <filename1> >> <filename2>` - Append the content of the file filename1 at the end of the file <filename2>.
- `cat <filename> <filename2> ` - Display 2 files at a time.
- `cat <filename> <filename2> > <newfile-name>` - Merge both of file content in a single one.
- `cat <file-name> | tr > <new-file-name>` - Translate the file.
- `cut -c  1-2 <filename>` - cut the file column wise
- `echo "Hello" >> <file-name>`
- `man <commad name>` - Know about the command usages and options.
- `man <commad name>` - know about the command.

### File/Folder operation
- `cp <file-name> <new-fie-name>` - Make a copy of a file in the current location.
- `mv <file-name> <dir-path>` - Move a file from one dir to another.
- `mv <file-name> <new-fie-name>` - Rename a file.
- `mv -R <dir-name> <dir-path>` - Move Dir
- `rm <file-name>` - Remove a file permanently.
- `rm -R <file-name>` - Delete a folder with dir included.
- `head <file-name>` - Will display first 10 lines of a file.
- `tail <file-name>` - Will display last 10 lines of a file.
    -`-n 2` - will display last 2 lines.
- `diff <file-1> <file-2>` - Show diff between the two files.
- `locate <file>` - To find out the file.  
- `find <file/folder-name>` - Find a file/folder.
- `find <dir-name>` - Find files inside the dir
- `find .-type d` - Show only dir.
    - `.-type f` - show only files.
    - `.-type f -name "*.txt"` - Show only files with that specific name.
    - `.-type f -iname "*.txt"` - Show only files with that specific name - not case sensitive (i)
    - `.-type f -mmin -20` - Show files which modify less than 20 min ago.
    - `.-type f -mmin +20` - show files which modify more than 20 min ago.
    - `.-type f -maxdepth 2` - Will only show 1 folder deep.
    - `.-size +1k` - will only show file/folder with size of 1kb


### System commands
- `ps aux` - processes which are running
- `systemctl [option] [service]` - interact with a process
    - We can do 4 `option` with `systemctl`
        - start
        - stop
        - enable
        - disable
    - Example, `systemctl start apache2`     
- `df` - Check the capacity and storage details.
    - `m` - In megabyte  or 
    - `hg` - In gigabyte.
- `du` - Disk usages capcity 
    - `-h` (human readable)
- `echo` - Get a output of a string
- `echo $PATH` - Check the path variable
- `sudo` - Admin command
- `sudo chown root text.txt` - change owner
- `!<command-name>` - Run the previous command
- `git add .; git commit -m "message"` - Run multiple commands at a time
- `sort <file-name>"` - sort the file
- `job` - show the jobs
- `wget <url>` - download the file from the URL
- `top` - what processes are running
- `kill <process-id>` -stop that process
- `Uname` - show the system info
- `zip <file-1> <file-2>` - Zip Two or more files
- `Unzip <file-name>` - Unzip files
- `useradd <name>` - add a user
- `passwd <name>` - set a password for the user
- `uname -<flag>` -o -m -r
- `lscpu` - get cpu details
- `free` - free memory 
- `vmstat` - virtual memory
- `lsof` - list all the open file
- `xdg-open <file-fath>` - open the folder (graphical window) of a file/folder with path.
- `xdg-open .` - open the folder of the current directory.
- `vi ~/.bashrc` - set your Alias
- `echo -n 'username' | base64` - encode the username to base64
- `echo -n 'encoded' | base64 -d` - decode the username to base64

### Networking
- `nslookup google.com` - To check the IP address of the domain.
- `netstat` - To check the network status.
- `hostname` - To check the hostname.
- `whoami` -  To check the current user.
- `ping google.com` - To check the connectivity.


### Permissions
- `chmod u=rwx,g=rxw,o=rwx <file-name>` READ, WRITE AND EXECUTE
- `chmod 777 <file-name>` - 4- Read, 2- Write, 1 - Execute
- `find . -perm 777 ` - shows files with all permissions(rwx)
- `grep <keyword> <file-name>` - To search if the keyword is presnt in the file or not
- `grep -w <keyword> <file-name>` - To search if the keyword is present in the file or not (complete word)
- `grep -i <keyword> <file-name>` - To search if the keyword is present in the file or not (not case sens)
- `grep -n <keyword> <file-name>` - To search if the keyword is present in the file or not (Line number)
- `grep -B <keyword> <file-name>` - Show Line before that keyword
- `grep -win <keyword> ./*.txt` - To search if the keyword is present in the file in current dir
- `grep -win -r <keyword> .*` 
- `history | grep "ls -l"` - Piping, we filter out the things

```
history
!<number-from-history>
```

# Operators

- `ping google.com & ping facebook.com` - run both the commands at the same time
- `echo "google.com" && echo "facebook.com"` - second will only run if first is successful
- `echo "google.com" && {echo "facebook.com"; eco "pradumnasaraf.co"}` 
- `echo "google.com" || echo "pingfacebook.com"` - second will only run if first is not successful
- `rm -r !(file.txt)` - delete all files except file.txt
- `printevnv` - to print all th env.
