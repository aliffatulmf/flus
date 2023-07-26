# Flus (⚠️experimental)

[English](README.en.md)

**Flus** is an experimental program that functions to organize or move files based on their extensions to corresponding directories.

## Installation

To use this program, you can follow these steps:

1. Download this repository.
2. Navigate to the terminal.
3. Run the `build.bat` or `build.sh` command.

## Explanation

### target

The `-target` argument is an option in the Flus program used to specify the target directory you want to organize. When you run the program with this option, Flus will scan the files in the target directory and move them to appropriate directories based on their file types.

For example, if you have a target directory with files of various types such as `.jpg`, `.pdf`, and `.zip`, Flus will move `.jpg` files to a directory named `Images`, `.pdf` files to a directory named `Documents`, and `.zip` files to a directory named `Archives`.

You can specify the target directory by using the `-target` option followed by the directory path. For instance, if your target directory is located at `/home/user/documents`, you can run the program with the following command:

```powershell
flus -target /home/user/documents
```

### unsafe

The `-unsafe` argument is an option that allows users to enable the `unsafe` mode or skip verification steps after the file copying process is completed. Using this option aims to speed up the copying process, but it should be noted that this method also carries the potential risk of data corruption due to the lack of verification after copying is finished.

In this program, the usual verification method is to compare the Hash results of the original file and the copied file. However, when the `unsafe` mode is enabled, this verification process is bypassed to expedite execution.

If you want to enable the `unsafe` mode, you can run the program with the following command:

```powershell
flus -target /home/user/documents -unsafe
```

However, it is recommended to be cautious when using this option. Make sure you only enable the `unsafe` mode when you are confident that the source files and the copying process can be fully trusted, and the risk of data loss or damage due to negligence in verification is acceptable.

### move

By default, this program will copy files to the appropriate directories based on their file types. However, if you want to delete the original files after the copying process is complete, you can use the `-move` option.

⚠️ CAUTION! Using the `-move` option with `-unsafe` will delete the original files without verification. Make sure you understand the potential risks before using this option.

If you want to enable the `-move` option, you can run the program with the following command:

```powershell
flus -target /home/user/documents -move
```

It is advisable not to use the `-move` option if you are unsure about the copying process. Delete the copied files manually and ensure that before deletion, the copied data has been verified and is not damaged.