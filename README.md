# SlcStrToCSV
SlcStrToCSV is very simple library to wrtie CSV files from slice of structs without setting a type of struct and borring adding each field. 

It uses reflect package, gets names of struct's fields and write them to the title of bytes.Buffer, values of fields are written to body. Also you can add to each field tag "title" to replace title of csv to these tags. The only thing you should do is to append all structs to the slice of interface and use func GetCSV(). Maybe in future I'll figure out how to do it simpler.

All values are written as strings. For this purpose I used package "github.com/spf13/cast".

## Instal link

```
go get github.com/metalCodeSolid/sliceofStructToCSV
```

## Usage

Your struct:
```
Foo struct{
...
Field type `title:"fieldTitle"` // title is optional
}
```
Just use this:
```
SlcStrToCSV.GetCSV(yourSlice []interface{}) (*bytes.Buffer, string) // string is an error, if string is nill -> OK
```
