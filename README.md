# ballclock

The source just needs to be installed in a go environment and then can be run by doing:

    ~/go/src/ballclock/bclock/ballclock$ go install

    $ ballclock 30
    30 balls cycle after 15 days.
    Completed in 5.601976 milliseconds ( 0.005601976 seconds )

    $ ballclock 30 325
    {"Min":[],"FiveMin":[22,13,25,3,7],"Hour":[6,12,17,4,15],"Main":[11,5,26,18,2,30,19,8,24,10,29,20,16,21,28,1,23,14,27,9]}
    Completed in 0.2123 milliseconds ( 0.00021229999999999998 seconds )
