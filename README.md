Picsort
=================

This started as an effort to sort images into individual folders, I had gone to a place where I had taken 300 photos 
and it was boring to segregate them into distinct folders based on the name of the person the photo was in so 
I wrote a webapp to assist me with that task.

How to make it work?
===================

1. either clone the repo or do a go get, or download the release from the link above
2. create a public folder if it doesn't exist and paste your photos in it
3. run the binary or do a `go run main.go`
4. then open localhost:8080 in the browser, it'll show you pages with the image and a input box for the tag
   use a comma seperated tag list to indicate multiple tags like c,java,python
5. when all the photos are tagged, the page will give you the link to sort the images, it'll generate one folder for each tag in the result folder

Note: start the webapp only after you have pasted the photos in the public folder this app is programmed to read the public folder and insert the files into
the database as one entry, so make sure you have only images pasted, it won't work if you paste any random file. When you are done pasting photos, start the server,
it'll read the folder and populate the database.

This webapp uses a sqlite3 database which is a part of the repo. In case you want to make your own sqlite3 file, use this schema
`CREATE TABLE picture( name varchar(100), tags varchar(100));`

Released under the MIT license, enjoy!
