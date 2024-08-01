There are many problems which can be solved using genetic algorithm.
Encoding of the chromosome is the first step to solving these problems.
The goal to represent the solution in the form of a string just as in a chromosome.
Each gene controls particular characteristics of the individual, similarly each bit/char in the string represents characteristics of the solution.

The 3 main methods that we are going to study here are-
1 . Binary Encoding 2. Permutation Encoding 3. Value Encoding

#### Binary Encoding

Binary Encoding is the most common.

- Every bit in the string represents a characteristic of the solution.
- Example of a problem: Knapsack problem.
- The problem: Given a list of items with weights and values, find the maximum value that can be put in a knapsack of capacity W.
- Encoding: Each bit says, if the corresponding item is in the knapsack or not.

#### Permutation Encoding

- Used in ordering problems.
- The string represents a sequence.

Example: Travelling salesman problem.

#### Value Encoding

- Values such as real numbers, characters, etc are used.

Example: Finding the weights for neural network.
