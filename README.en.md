# Flus (⚠️ Experimental)

**Flus** is an experimental program that functions to organize or move files based on their extensions to their respective directories.

## Installation

To use this program, you can follow the steps below:

1. Download this repository.
2. Navigate to the terminal.
3. Run the command `build.bat` or `build.sh`.

## Explanation -target

The `-target` argument is an option in the Flus program that is used to specify the target directory that you want to organize. When you run the program with this option, Flus will scan the files in the target directory and move them to their respective directories based on their file types.

For example, if you have a target directory with files of different types such as `.jpg`, `.pdf`, and `.zip`, Flus will move the `.jpg` files to a directory named `Images`, the `.pdf` files to a directory named `Documents`, and the `.zip` files to a directory named `Archives`.

You can specify the target directory by using the `-target` option followed by the directory path. For instance, if your target directory is located at `/home/user/documents`, you can run the program with the following command:

```powershell
flus -target /home/user/documents
```
