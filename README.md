Picsort
=================

This started as an effort to sort images into individual folders, I had gone to a place where I had taken 300 photos and it was boring to segregate them into distinct folders based on the name of the person the photo was in so I wrote a webapp to assist me with that task.

How does it work?
====================

Put the photos in the public folder and run the script insert.py, it'll generate insert statements, create a sqlite3 database with the DDL given below.

Then run the binary

I have packaged it as a zip, you can just download it from the releases folder on here, all you have to do is run the insert.py file from the command line and execute the insert statements into the sqlite3 database.

Then go to `localhost:8080`, you'll see a form which'll show the photo and a text box for the photo's tag, there is another input box stating the name of the photo, do not change it. After all photos are done the page will give the link to sorting the photos, click on the link and sit back and relax as it'll create folders for the names you have given in tags, for multiple names use a comma separated list for instance if the photo has three people, suraj kiran and rohit then put the tag as suraj,kiran,rohit. It'll create three folders and copy the photo in each folder, please make sure you give the right tag, else it'll give incorrect output! This doesn't use AI or face detection to find out who is in the photo.

Released under the MIT license, enjoy!
