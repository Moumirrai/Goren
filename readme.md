# Renaming Utility for DWG/PDF Files

This utility simplifies the process of renaming DWG/PDF files. It is designed to run on Windows operating systems.

## Getting Started

1. Download the latest release from the [Releases](https://github.com/Moumirrai/Goren/releases/latest) page.

2. When you run the utility for the first time, the Windows system may display a warning that "Windows has protected your PC". Click "More info" and then select "Run anyway" to continue.

3. **Simply opening the utility's interface by clicking on it will not trigger any action.**

4. To begin renaming files, drag and drop the original DWG/PDF files onto the executable. A folder named "RenamedFiles" will be created in the original file directory to store the renamed copies.

5. In some cases where the file naming rules are not met, a prefix "_ERR_" will be added to the new filename, indicating an error.

## Example

`_- SO 02-D-1-1-103C - PŮDORYS 1PP - ČÁST C.dwg` → `SO 02.D.1.1-103C_PŮDORYS 1PP - ČÁST C.dwg`

`_-Výkres - SO 02-D-1-1-110A - PŮDORYS STŘECHY NAD 6NP.dwg` → `SO 02.D.1.1.110A_PŮDORYS STŘECHY NAD 6NP.dwg`

`_-Výkres - ŘEZ A-A - ČÁST A + C.dwg` → `_ERR__-Výkres - ŘEZ A-A - ČÁST A + C.dwg` - error

## Configuration

The utility generates a configuration file named "renconfig.json" in the same directory as the executable file. This file allows for some customization, although modifying it is not necessary for the utility's basic functionality. If the configuration file is deleted, a new one will be created automatically upon the next use.

### Configuration File Example:

```json
{
  "marker": "SO ",
  "makeCopy": true,
  "outputDir": "RenamedFiles"
}
```

- The **marker** field specifies the text that the new filenames should start with. You can leave this unchanged.

- The **makeCopy** field can be set to "false" if you want the utility to directly rename files without creating copies in the "RenamedFiles" folder. This can be useful to save disk space, but it's recommended to keep backups of original files.

- The **outputDir** field specifies the name of the folder into which the renamed files will be copied."

## License
This utility is provided under the [MIT License](https://github.com/Moumirrai/Goren/blob/master/LICENSE).

## Support
For any questions or issues, please open an [issue](https://github.com/Moumirrai/Goren/issues/new) on the GitHub repository.
