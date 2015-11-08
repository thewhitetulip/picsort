Picsort
=================

This started as an effort to sort images into individual folders, I had gone to a place where I had taken 300 photos and it was boring to segregate them into distinct folders based on the name of the person the photo was in so I wrote a webapp to assist me with that task.

How does it work?
====================

Put the photos in the public folder and run the script insert.py, it'll generate insert statements, create a sqlite3 database with the DDL given below.

Then run the binary

I have packaged it as a zip, you can just download it from the releases folder on here, all you have to do is run the insert.py file from the command line and execute the insert statements into the sqlite3 database.


