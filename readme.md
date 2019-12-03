# Bulb
I recently set up an office in my home with the hope that I would be working
from home more often.  My office is down in the basement of my home in a 
finished room.  I had a concern that my kids would come down and rush into my
office when I was in the middle of a call or busy doing heads down work.

I started to look at various Busy Light indicators.  Most were expensive and 
tied into software that I wouldn't be using.  I had a couple of basic 
requirements and 1 very strong desire.

1) The light had to be outside the room, thus probably wireless
2) There isn't a well placed outlet outside the room.  So needed to be battery
powered or plug in some other way
3) Needs to be able to show multiple colors
4) Needs to be very obvious to those entering the hallway
5) Desire to be able to be set via the command line

I thought about building this with a Raspberry PI, some LEDs, and an SNS topic
or something.  That would have worked, but would have required a lot of work 
to get up and going.  Then I thought about Smart Bulbs.  I could control them 
with an app on my phone and covers all of the requirements except the command
line request.  After looking I found that a lot of them integrate with IFTTT
and that allowed me to write this tool.

Now I have several IFTTT actions that will change the colors of the bulbs.  All
I have to do is run
```
bulb set red
```  
and an HTTP request is fired off to IFTTT and a couple of seconds later the bulb
color is updated and my family knows if I am in the middle of something or not.
