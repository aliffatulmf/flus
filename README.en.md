# Flus

[ID-id](README.md)

The **Flus** command-line tool offers a streamlined solution for efficiently managing directory operations, such as copying and moving files between different locations. This tool provides a range of customizable options to cater to diverse requirements across various operating systems.

## Usage

To leverage the capabilities of the **Flus** tool, you need to provide specific command-line arguments as follows:

```powershell
PS C:\> flus.exe [options]
```

## Available Options

- `-target`: Specify the target directory for processing. Files will be copied or moved to this directory.

- `-move`: Activate move mode to relocate files from the source directory to the target directory, instead of copying.

- `-buffer`: Define the buffer size for copying files. The buffer size significantly affects copying efficiency. The default buffer size is set to 64 KB.

## Example

Here's an example illustrating the usage of the **Flus** tool:

```powershell
PS C:\> flus.exe -move -buffer 128000 -target C:\path\to\target\directory
```

In this example, the tool will move files from the source directory to the specified target directory using a buffer size of 128 KB.

## Important Notes

- Replace `C:\path\to\target\directory` with the actual path of your target directory.

- Exercise caution when using move mode, as files will be deleted from the source directory upon successful transfer to the target directory.

- Adjusting the buffer size can significantly impact the speed of copying or moving files, as well as memory usage. Experiment with various buffer sizes to find the optimal configuration for your system.

## Running the Tool

To run the **Flus** tool on your operating system, follow these steps:

1. Download the executable file (`flus`) from the official repository or source.

2. Open a terminal or command prompt.

3. Navigate to the directory containing the `flus` executable using the `cd` command.

4. Execute the tool with the desired options:

   ```powershell
   PS C:\> flus.exe [options]
   ```

## Note

While this documentation provides a general overview of the **Flus** tool's functionalities, its compatibility and performance on specific operating systems other than Windows have not been extensively tested.
