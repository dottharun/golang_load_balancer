# Load balancers
- sits above web servers, between user requests
- and sends those requests to any of the servers in a collection of web servers

# Many types of algorithms
1. Round Robin - for every new user send the request to the next server sequentially
    - simple algorithm
2. Random select

# IMPLEMENTATION
Implement the L4 load balancer as a package.

- The type of load balancing is round robin,
- For health check both active check and passive check are supported
- Persistence is not supported (3rd party tools can be used)

