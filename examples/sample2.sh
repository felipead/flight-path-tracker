#!/bin/bash

curl -0 -v http://localhost:8080/flight_paths \
-H "Expect:" \
-H 'Content-Type: application/json' \
--data-binary @- << EOF
{
    "flight_legs": [
        ["GRU", "MIA"],
        ["MIA", "ORD"],
        ["YUL", "JFK"],
        ["SFO", "YUL"],
        ["ORD", "SFO"],
        ["CNF", "GRU"],
        ["JFK", "LHR"]
    ]
}
EOF
