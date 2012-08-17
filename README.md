Make some monte carlo pie calculations

Currently just does a simple (concurrent) PI approximation, as a proof of concept.

# General approach
1. Define domain of possible inputs
2. Generate input randomy with some probability distribution
3. Perform deterministic computation on inputs
4. Aggregate results

# For a circle
1. Draw a square with a circle in i
2. Uniformly scatter points over square
3. Count number of points inside circle and total number of points
4. Ratio of the two counts leads to PI
