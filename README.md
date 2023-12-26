# Heft
## A simple HTTP load tester


Heft is a simple tool for testing the performance of web servers under load.

### Key features:

- HTTP load testing: Simulates multiple concurrent users accessing a web server to assess its performance under stress.
- GET requests only: Currently supports testing 
GET requests, the most common type of HTTP request.
- Simple design: Likely offers a user-friendly CLI and straightforward setup for ease of use.

### Potential use cases:

- Measuring website performance: Evaluating how a website handles various traffic loads.
- Identifying bottlenecks: Pinpointing areas of a website that might slow down under heavy traffic.
- Testing server capacity: Assessing how many concurrent users a server can handle before performance degrades.
Optimizing website performance: Using results to make informed decisions about website optimization strategies.

## Build Instructions

### Prerequisites:

Ensure you have Go installed on your system. Check by running `go version` in your terminal.

Get the code:

- Clone the repository using 
`git clone https://github.com/ankitsridhar16/heft.git`

Build the executable:

Run `go build cmd/heft/heft.go` in the terminal to create the Heft executable file.

## Usage

### Running Heft

Navigate to the build directory: 
- In your terminal, go to the directory where the Heft executable was built (usually the project's root directory).

  - Execute the command:
  Run the following command to start Heft:
  `./heft -u <URL> [-n <num_requests>] [-c <concurrent_requests>]`
  - Replace `<URL>` with the actual URL you want to test.
    Optionally, specify:
    - -n: Number of requests to perform (default: 1).
    - -c: Number of concurrent requests to run (default: 1).
    
  #### Example
    `./heft -u https://www.example.com -n 100 -c 10`

  Output:

Heft will perform the load test and print the results to the console, including:

- Total successful requests (2XX status codes)
- Failed requests (5XX status codes)
- Requests per second
- Total request time (minimum, maximum, and mean)
- Time to first byte (minimum, maximum, and mean)
- Time to last byte (minimum, maximum, and mean)