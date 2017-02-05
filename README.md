# Delta
delta utility for time operations

delta.scala - script
delta.jar - compiled version
build.sh - build script
delta.sh - runscript for linux

Howto use
Purpouse of the script is to calculate work time on floating schedule
For example You got to work at 9:20 and spent 40min on lunch
Then you want to calculate how much u are on the work 
so if now is 18:30 you can calculate work time
$ scala delta.jar 18-30 10-20 -40
Total=7:30
So you are working for 7:30
