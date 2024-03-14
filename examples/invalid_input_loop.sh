#!/bin/bash

curl -0 -v http://localhost:8080/calculate \
-H "Expect:" \
-H 'Content-Type: application/json' \
--data-binary @- << EOF
{
    "flight_legs": [
        ["IND", "EWR"],
        ["SFO", "ATL"],
        ["GSO", "IND"],
        ["ATL", "GSO"],
        ["EWR", "SFO"]
    ]
}
EOF
