Memo serves to store simple memos, much more like a second sapce for my brain.
I would love to blog but I don't feel like typing 5000 words articles long, I 
just want a simple way to remind common stuff I have seen at work, personal life
and keep them at hand. 

We'll build this application in go, trying to keep it as simple as possible yet 
maintainable. We'l try to go fast too to avoid losing motivation. 

# DATA STRUCTURES
The main data structure will be the Memo, which should be comprised of: 

- A question / Title
- A Body
- A creation Date
- A last updated date. 

# NEEDS
I need to be able to save a memo
I need to be able to edit a memo's title and body; update time shoudl be computed automatically.
I need to be able to browse the full list of memos, presented by Title and paginated 20 per 20. 
I need to be able to search for individual memos by keywords contained in the title.

We'll see user ownership and authentication later

# STORAGE
Those memo will be persisted to a boltdb database. 

# PRESENTATION
The memo will be accessible through a simple website. We'll make use of go native
templating engine. 

