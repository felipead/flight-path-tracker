#!/bin/bash

curl -0 -v http://localhost:8080/flight_paths \
-H "Expect:" \
-H 'Content-Type: application/json' \
--data-binary @- << EOF
{
    "flight_legs": [
        ["IND", "EWR"],
        ["SFO", "ATL"],
        ["EWR"],
        ["GSO", "IND"],
        ["ATL", "GSO"]
    ]
}
EOF
