# Issues
Working with a really simple program using huge Fyne List, we can see that there is two problems.
At the top of the list everything is working, but as we go closer to the bottom two leaks appears:
1. It's start to be more and more slow (lag)
2. An offset appear when we select an item, that make the list go up with no reason (the closer to the end, the bigger the offset is)

As you can see the code is very simple and apparently didn't have any performance issues.
I'm on Mac M1 (no possibility for me to test on another device).

# Video
https://www.loom.com/share/0c2d5393fad9451cad7e8b0bb5d26676