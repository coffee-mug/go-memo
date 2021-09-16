#!/bin/bash

# build last version of executable. Binary name must not be changed
go build -o app

# sync useful files with the server
rsync -avz -e "ssh -p 3500" --progress app assets templates lucas@159.89.228.197:/home/lucas/notes.scrapiste.com

# restart the process
ssh -t lucas@159.89.228.197 -p 3500 'sudo systemctl restart notes-scrapiste'

echo 'Deployment done'
