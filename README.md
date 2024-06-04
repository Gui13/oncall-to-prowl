
Oncall to prowl
======

This is a very simple Grafana Oncall to Prowl software.

In Oncall, just create a new outgoing webhook an use this program's `/event` URL as target.

Using the PROWL_API_KEY passed in args or through an environment variable, this program will receive webhook calls and 
translate it to Prowl notifications on your phone.



To test it:

    curl localhost:8080/event -X POST  -v --data @event.json


To build it:

    make all


To clean it:

    make clean