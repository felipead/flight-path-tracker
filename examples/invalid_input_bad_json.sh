#!/bin/bash

curl -0 -v -X POST http://localhost:8080/calculate \
-H "Expect:" \
-H 'Content-Type: application/json' \
--data-binary @- << EOF
{
    "flight_legs": [
        ["IND", "EWR"],
        ["SFO", "ATL"],
        ["GSO", "IND"],
EOF
