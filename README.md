# Description
ThisChord is a discord bot which offers a wide range of loosely related audio capabilities.

This is a project for honing my Golang skills. Any feedback would be appreciated.

# Implemented functionalities

hue

# To be implemented:
## Ear training
Inspired by Duolingo, keep your streak of ear training alive.
Choose between intervals and chord progression in whichever configuration you would like.
### Commands:
```
I suspect some interactive messages might be used here.
```
## Song recognition
Invite ThisChord to your voice channel and recognize songs which is currently played by a User.
### Commands:
```
- !listen - will order ThisChord to join the channel you are currently occupying
            and listen in on any user that is currently in it
- !listen \<user\> - allows ThisChord to listen in on a specific \<user\>. 
                     If no user is found, ThisChord won't join.
```
## Word counting
Ever wondered how many times your friend said "NICE" during a game? Allow ThisChord to count it for you!
### Commands:
```
- !count <word> - count the number of occurances of <word>.
                  First version will have only one counter available.
- !stopcount - finishes the count.
```
## History
You may want to check what was the song that ThisChord caught two weeks ago. No problem! ThisChord will remember your searches!
### Commands:
```
- !history - shows history of all commands and their outputs.
```
# Docs
As it is a training project I've prepared a [document](https://docs.google.com/document/d/1W3S16ZyBtvxUayCGRXPnlIU-Qi-OG7pPScvwyb1n6JM/edit?usp=sharing) documenting my journey.
