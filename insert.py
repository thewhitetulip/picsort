import os

files = os.listdir("./public")
print(files)
print(os.curdir)

stmt = 'insert into picture(name) values("'

for file in files:
    if os.path.isfile("./public/"+file):
        stmt += file + '");'
        print(stmt)
