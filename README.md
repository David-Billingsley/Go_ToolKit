# Go_ToolKit
## List of functions to help the end user with their projects
  To use the tool add toolkit "github.com/David-Billingsley/Go_ToolKit" to your code base

  And in your function put var tools toolkit.Tools
  
## RandomString
  This function takes an input ( whole number ), and outputs a random alpha numeric string ( abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 ) for the user to use

## FixJson
  This function takes a string represntation of JSON and removes unnneded info from the string. It takes to parmeters one is the JSON string and the other is what you want removed. This returns the JSON as a string.
  This lets the user get the inner part of the JSON for the reutrned result form some API calls.
    Example:
      items:{[

## EpochConverMil
  This function takes a int64 number and returns time.  This is used to help where you have an [Epoc time]( https://www.epochconverter.com/ ) vaule and need to see its time in the tradional date format.

## DateStrParse
  This function takes a string representation of a date and converts it into a proper date format.

## DownloadFile
  This function takes a URL and filename.  It will download the file form the URL and save it with the filename provided.
