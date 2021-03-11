# cache
Last used last deleted [string]interface{}.

Simple cache implementation for general purpose, especially intended for use inside the server.

## Why I need this.
I needed simple and reliable cache feature to use server internal cache which shoud be robust and reliable.
There are a lot of clever cache technic but these seems to be capable in the specific situation. So, just simple "last used last delete" is versatile and less error prone and in case if my application will have a specific tendency to use variable, at this time it might be good time to add special hash technic to this.
