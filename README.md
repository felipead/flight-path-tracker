# Flight Path Tracker

## Instructions

Install Go 1.22, and the Make tool. Then run:

```shell
$ make build
$ make run
```

It will start the HTTP server on port 8080 on localhost.

### Examples

The folder `examples/` contains a list of sample HTTP requests using cURL. After starting the server, feel free to execute those samples.

### Running tests

```shell
$ make test
```

## API Specification

### Definitions

From [here](https://aviation.stackexchange.com/questions/14567/what-is-the-difference-between-slice-segment-and-leg):
> A flight is defined by the IATA as the operation of one or more flight legs with the same flight designator. Unlike a flight segment, a flight may involve one or more aircraft. The IATA defines a leg as the operation of an aircraft from one scheduled departure station to its next scheduled arrival station. A flight segment can include one or more legs operated by a single aircraft with the same flight designator.

### The `/calculate` endpoint

This endpoint expects a JSON payload containing the list of flight legs that are part of a given flight itinerary.

```
POST /calculate

Content-Type: application/json

{
    "flight_legs": [
        ["IND", "EWR"], 
        ["SFO", "ATL"], 
        ["GSO", "IND"], 
        ["ATL", "GSO"]
    ]
}

200 OK

{
    "flight_start_end": ["SFO", "EWR"]
}
```

#### Constraints and validations

- At least one flight leg must be provided.
- A flight leg must be declared as a list of two strings.
- A flight leg cannot point to itself. The following will throw an error: `["JFK", "JFK"]`
- No loops shall be present in the path. The implementation will raise an error if it detects any loops. We detect loops by checking the presence of multiple inbound or outbound flight legs for any given airport code.

#### Security considerations

For security reasons, we are not accepting or returning a JSON array in the HTTP request or response bodies. This is to avoid [JSON Hijacking](https://stackoverflow.com/questions/43717574/javascript-why-shouldnt-the-server-respond-with-a-json-array), a common exploit.

Instead, we are using a JSON object, and the array is embedded as a property.

### Errors

The API will obey to the [HTTP response status code convention](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status). More specifically, it will return:

- `400 Bad Request` for malformed JSON payloads and invalid inputs

Errors should be returned using the following JSON structure:

```json
{
    "error": true,
    "retryable": false,
    "message": "unable to unmarshal flight leg ..."
}
```

---

## Context

### Story

There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

### Goal

To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

### Required JSON structure

```
• [["SFO", "EWR"]]                                                 => ["SFO", "EWR"]
• [["ATL", "EWR"], ["SFO", "ATL"]]                                 => ["SFO", "EWR"]
• [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]
```

### Specifications

- Your microservice must listen on port 8080 and expose the flight path tracker under the `/calculate` endpoint.
- Create a private GitHub repo and add https://github.com/taariq as a collaborator to the project. Please only add the collaborators when you are sure you are finished.
- Define and document the format of the API endpoint in the README.
- Use Golang and/or any tools that you think will help you best accomplish the task at hand.
- When you are done with the assignment, follow up and reply-all to the email that directed you to this document. Include your private github link and an estimate of how long you spent on the task and any interesting ideas you wish to share.
