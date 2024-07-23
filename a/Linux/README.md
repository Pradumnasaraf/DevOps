## Linux

### Learning resources

- [Linux kernel source tree](https://github.com/torvalds/linux)
- [Linux & Docker Fundamentals - Kubesimplify](https://www.youtube.com/live/EUu1E_YKGTw?feature=share)

### ⚫ Overview of Linux and Why it is so popular?

Linux is an open-source operating system (OS) that runs on a computer and controls all of the other software and hardware on the computer. Initially, Linux was built on the Unix operating system, which was created many decades ago as a platform for scientific and academic computing. It is used on a wide range of devices, including servers, desktop computers, laptops, cloud computing, and Mobile devices.

As Linux is open-source software so anyone can access its source code and make changes to it. Other than this Linux is known for its Security features. It has a large community of users and developers who are constantly working to identify and fix security vulnerabilities. Also, Linux is well known for its Customizability allowing users to modify and configure it in order to meet their specific needs.

### ⚫ Linux Architecture Components

The main architectural components that work together to form a complete Linux OS include Kernel, System Libraries, User Space, Shell, System Utilities, Daemons, and Configuration files.

- Kernel: The kernel is the central part of the Linux OS. It manages the computer's hardware resources and communicates with the software to provide a stable and consistent interface for us.

- System Libraries: System libraries are collections of pre-written code that provides common functions for other programs to use.

- User Space: User space mainly refers to the programs and applications that run on top of the operating system. This includes a wide range of applications such as text editors, web browsers, games, etc.

- Shell: Shell is a command-line interface that allows users to interact with the Linux operating system.

- System Utilities: System utilities are small programs that perform specific tasks such as file management, process management, network management, etc.

- Daemons: Daemons are the background processes that run continuously and perform specific tasks such as network services, log management, etc.

- Configuration files: Configuration files are used to configure various aspects of the Linux operating system such as the network settings, system settings, security settings, etc.

![architecture](https://user-images.githubusercontent.com/37767537/225990738-9e505c6d-bad0-4820-a2b2-4ce84ef286c9.jpg)

### ⚫ Linux Distributions

Linux distributions, also known as "distros," are different versions of the Linux operating system that are built 
and packaged with specific versions of the architectural components, the package management system used, the graphical user interface, and the tools and applications included in it.

Examples of Linux distros include Ubuntu, Fedora, Debian, CentOS, Mint, Red Hat Enterprise Linux (RHEL), etc.

### ⚫ Linux File Systems

The Linux file system is the way that the operating system organizes and stores files and directories on a computer's hard drive or other storage devices. It consists of a hierarchical structure where files and directories are arranged in a tree-like format starting from the root directory (represented by "/"). Each file and directory has a unique location within the file system hierarchy, and these locations can be referred to using a path similar to how file paths work in other operating systems.

There are several types of file systems available for use in Linux including below:

- Ext4: This is the most widely used file system in Linux and it is known for its high performance, reliability, and scalability. It is a journaling file system meaning that it keeps track of changes made to the file system in a log that helps to ensure that data is not lost in the event of a system failure.

- XFS: This is a high-performance file system that is optimized for large-scale data storage and is commonly used in enterprise and high-performance computing environments. It supports large file sizes and high-bandwidth I/O which makes it well-suited for use with high-speed storage devices such as solid-state drives.

- Btrfs: This is a newer file system that is designed to be flexible, scalable, and easy to manage. It offers advanced features such as snapshots which allow administrators to easily backup and restore data and it also supports data and metadata checksums which help to prevent data corruption.

- NTFS: This is a file system commonly used in Microsoft Windows and it is supported by most modern Linux distributions. It is designed to be compatible with a wide range of devices including hard drives and removable storage devices and it supports advanced features such as compression and encryption.

- FAT32: This is an older file system that is commonly used in older Microsoft Windows systems and is also widely used for removable storage devices such as USB drives and has limitations such as a maximum file size of 4GB.

(Note: The machine from which the screenshot has been taken is having CentOS 8 Stream)

![Linux-1](https://user-images.githubusercontent.com/37767537/226102852-2ca35206-f833-44e0-994f-a003a88c6b84.png)

### ⚫ Organization of Linux File System

In Linux, the file system is organized into a hierarchical structure, with the root directory (/) at the top of the hierarchy. The root directory contains several subdirectories, each with its own specific purpose that form the foundation of the Linux file system hierarchy. The subdirectories include:

- /bin - contains essential command-line utilities that are required for the system to function
- /boot - contains files needed for booting the system including the Linux kernel and boot loader configuration files
- /dev - contains device files that represent various devices and hardware components on the system
- /etc - contains configuration files for the system and various applications
- /home - contains the home directories for individual users
- /lib - contains libraries needed by the utilities in /bin and /sbin
- /media - a mount point for removable media such as CDs, DVDs, and USB drives
- /mnt - a mount point for temporarily mounting file systems
- /opt - contains third-party software packages
- /proc - A virtual file system that contains information about the system's process and kernel
- /root - the home directory for the root user
- /sbin - contains system executables, similar to /bin
- /sys - A virtual file system that provides an interface to the kernel's device drivers
- /tmp - Contains temporary files that can be deleted by the system when space is needed
- /usr - contains user utilities and libraries as well as data shared by multiple users
- /var - contains variable data such as log files, mail spools, and database files

These subdirectories may vary between different Linux distributions but the basic structure will remain the same. It is strongly recommended to explore each subdirectory separately to understand more about it.

![Linux-2](https://user-images.githubusercontent.com/37767537/226103260-c51190cf-3e9a-47e9-abe8-6b131227572d.png)
