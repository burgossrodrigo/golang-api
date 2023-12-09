# golang-restapi
 
Water Jug Problem Solver API
Overview
This API solves the classic Water Jug problem, given three parameters: the capacities of two jugs (X and Y) and a target volume (Z). The solution finds a way to measure exactly Z gallons using only the two jugs.

Technologies
Language: Go
Framework: Gin
Cross-Origin Resource Sharing (CORS): Handled by custom middleware
API Endpoints
POST /api/calculate
Solves the Water Jug problem based on the provided jug capacities and target volume.

Request
Path: /api/calculate
Method: POST
Content-Type: application/json
Body:
```
{
    "x": int, // Capacity of jug X
    "y": int, // Capacity of jug Y
    "z": int  // Target volume in jug Z
}

```

Response
Content-Type: application/json
Body:
Success: Array of Step objects detailing each step to reach the target volume.
Failure: JSON object with {"result": "No Solution"}
Step Object

```
{
    "description": string,   // Description of the step
    "state": {               // Current state of the jugs
        "x": int,            // Current volume in jug X
        "y": int,            // Current volume in jug Y
        "z": int             // Current volume in jug Z
    }
}

```

TypeScript Usage
Here is an example of how to call the /api/calculate endpoint using TypeScript:

```
type CalculationRequest = {
    x: number;
    y: number;
    z: number;
};

type JugState = {
    x: number;
    y: number;
    z: number;
};

type Step = {
    description: string;
    state: JugState;
};

async function calculateWaterJugSolution(request: CalculationRequest): Promise<Step[]> {
    const response = await fetch('http://localhost:8080/api/calculate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(request),
    });

    if (!response.ok) {
        throw new Error('Server responded with an error!');
    }

    return response.json();
}

// Example usage
const request: CalculationRequest = { x: 3, y: 5, z: 4 };
calculateWaterJugSolution(request)
    .then(steps => console.log(steps))
    .catch(error => console.error(error));

```

Running the Server
To run the server:

Navigate to the directory containing the Go code.
Run `go run .`` to start the server.
The server will listen on http://localhost:8080.