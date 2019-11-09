# game-builder
a simple command line for building dungeon crawler games (under construction!!!)

<b>Commands:</b><br/>
buildobject: create an object.<br/>
allobject: view all objects by name and index<br/>
viewobject: allows you to view object (type in name)<br/>
placeobject: place object on the map (type in co-ordinates)<br/>
buildmap: allows you to create an X by X map by an index<br/>
viewworld: prints out all maps by index.<br/>
viewmap: prints out map by index. First map would be 0, second 1 etc.<br/>
play: initiates the game<br/>
q: exit the game<br/>

Objects:<br/>
ID - generated automatically<br/>
name - string<br/>
description - string<br/>
health - integer<br/>
attack - integer<br/>

Maps are a 3D slice generated by x, y width/length:<br/>
i.e a 3,3 map will be:<br/>
000<br/>
000<br/>
000<br/>

Placing objects on the map:<br/>
First you type in name of object you wish to place.<br/>
then the map index i.e first map is 0, second map will be 1 and so on.<br/>
Then you choose the x,y coordinates to place the object.<br/>
The first object you generate will be indexed at 0 and will present itself as "1" on the map. This object will be the hero character.<br/>
If objects have attack, then they can be fought against.<br/>
If objects have 0 attack, then no combat will be initiated and you can do other things like provide descriptive text.<br/>

Play:<br/>
When you press play, you will be able to control the hero character and navigate the world.<br/>

TODO:<br/>
Error handling...<br/>
Implement save functionality.<br/>
implement load functionality.<br/>

Long-term TODO:<br/>
Create a GUI on top of this interface, maybe a web app, so that it can be used to create dnd-lite scenarios.<br/>

:)<br/>




