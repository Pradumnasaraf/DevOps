---
sidebar_position: 1
title: Linux Introduction
description: An introduction to Linux and its architecture components.
tags: ["Linux", "Operating System", "Open Source"]
keywords: ["Linux", "Operating System", "Open Source"]
slug: "/linux"
---

### Overview of Linux and Why It Is So Popular

Linux is an open-source operating system that manages hardware and provides the foundation for software to run. It is heavily influenced by Unix ideas and is used across servers, laptops, cloud platforms, networking devices, and embedded systems.

Linux is widely used because it is open source, stable, flexible, and well suited for automation. It is also known for strong community support, broad tooling, and the ability to tailor systems for different use cases.

### Linux Architecture Components

The main architectural components that work together to form a complete Linux OS include Kernel, System Libraries, User Space, Shell, System Utilities, Daemons, and Configuration files.

- Kernel: The kernel is the central part of the Linux OS. It manages the computer's hardware resources and communicates with the software to provide a stable and consistent interface for us.

- System Libraries: System libraries are collections of pre-written code that provides common functions for other programs to use.

- User Space: User space mainly refers to the programs and applications that run on top of the operating system. This includes a wide range of applications such as text editors, web browsers, games, etc.

- Shell: Shell is a command-line interface that allows users to interact with the Linux operating system.

- System Utilities: System utilities are small programs that perform specific tasks such as file management, process management, network management, etc.

- Daemons: Daemons are the background processes that run continuously and perform specific tasks such as network services, log management, etc.

- Configuration files: Configuration files are used to configure various aspects of the Linux operating system such as the network settings, system settings, security settings, etc.

![architecture](https://user-images.githubusercontent.com/37767537/225990738-9e505c6d-bad0-4820-a2b2-4ce84ef286c9.jpg)

### Linux Distributions

Linux distributions, or distros, are packaged versions of Linux that combine the kernel with package managers, system tools, and often a desktop environment or server-focused defaults.

Examples of Linux distros include Ubuntu, Fedora, Debian, CentOS, Mint, Red Hat Enterprise Linux (RHEL), etc.

### Linux File Systems

The Linux file system is the way that the operating system organizes and stores files and directories on a computer's hard drive or other storage devices. It consists of a hierarchical structure where files and directories are arranged in a tree-like format starting from the root directory (represented by "/"). Each file and directory has a unique location within the file system hierarchy, and these locations can be referred to using a path similar to how file paths work in other operating systems.

There are several file systems commonly used with Linux, including:

- Ext4: This is the most widely used file system in Linux and it is known for its high performance, reliability, and scalability. It is a journaling file system meaning that it keeps track of changes made to the file system in a log that helps to ensure that data is not lost in the event of a system failure.

- XFS: This is a high-performance file system that is optimized for large-scale data storage and is commonly used in enterprise and high-performance computing environments. It supports large file sizes and high-bandwidth I/O which makes it well-suited for use with high-speed storage devices such as solid-state drives.

- Btrfs: This is a newer file system that is designed to be flexible, scalable, and easy to manage. It offers advanced features such as snapshots which allow administrators to easily backup and restore data and it also supports data and metadata checksums which help to prevent data corruption.

- NTFS: This is a file system commonly used in Microsoft Windows and it is supported by most modern Linux distributions. It is designed to be compatible with a wide range of devices including hard drives and removable storage devices and it supports advanced features such as compression and encryption.

- FAT32: This is an older file system that is commonly used in older Microsoft Windows systems and is also widely used for removable storage devices such as USB drives and has limitations such as a maximum file size of 4GB.

(Note: The machine from which the screenshot has been taken is having CentOS 8 Stream)

![Linux-1](https://user-images.githubusercontent.com/37767537/226102852-2ca35206-f833-44e0-994f-a003a88c6b84.png)

### Organization of Linux File System

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

These subdirectories may vary slightly between Linux distributions, but the overall structure stays fairly consistent. Exploring them one by one is one of the best ways to get comfortable with Linux.

![Linux-2](https://user-images.githubusercontent.com/37767537/226103260-c51190cf-3e9a-47e9-abe8-6b131227572d.png)

### What's next?

- [Linux Commands](./commands.md) - Learn about the commands that you can use with Linux.
- [Learning Resources](./learning-resources.md) - Learn more about Linux with these resources.
