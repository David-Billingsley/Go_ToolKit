# Go_ToolKit
  This libary is free of use, and is a helpful list of functions I have used over and over in code.

## List of functions to help the end user with their projects
  To use the tool add toolkit "github.com/David-Billingsley/Go_ToolKit" to your code base

  And in your function put var tools toolkit.Tools

# Strings:
## RandomString
  This function takes an input ( whole number ), and outputs a random alpha numeric string ( abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 ) for the user to use

# JSON:
## FixJson
  This function takes a string represntation of JSON and removes unnneded info from the string. It takes to parmeters one is the JSON string and the other is what you want removed. This returns the JSON as a string.
  This lets the user get the inner part of the JSON for the reutrned result form some API calls.
    Example:
      items:{[

# Date:
## EpochConver
  This function takes a string type int64 number and returns time.  This is used to help where you have an [Epoc time]( https://www.epochconverter.com/ ) vaule and need to see its time in the tradional date format. If you enter type of micro it does micro epoc time and milli does milli epoc time.

## DateStrParse
  This function takes a string representation of a date and converts it into a proper date format.

# File:
## DownloadFile
  This function takes a URL and filename.  It will download the file form the URL and save it with the filename provided.

## FilePathInSameDir
  This function takes a filename.  It gets the path of the bundled exe file, and returns the path with the filename attached. 

## CreateDirIfNotExist
  This function takes a filepath, and will create it if it does not exist.

## DeleteDir
  This function takes a filepath, and will delete it

## CopyDir
  This function takes a Orginating file path, and destination path.  It will move the files to this new path.

## BlankFileGen
  This function takes a full filepath and generates a blank file.

# URL
  
## Slugify
  This function takes a string and outputs a string. It creats a string with " - " so that you can use the string in the URLs