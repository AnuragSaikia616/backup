# Syllabus
**Genetic Algorithm:** History of Genetic Algorithms (GA), Working Principle, Various Encoding methods, Fitness function, GA Operators- Reproduction, Crossover, Mutation, Convergence of GA, Bit wise operation in GA, Multi-level Optimization.GA based Backpropagation Networks: GAbased Weight Determination, K - factor determination in Columns.

# History of genetic algorithm
# What is genetic algorithm
Genetic Algorithms (GAs) are search based algorithms based on the concepts of natural selection and genetics. GAs are a subset of a much larger branch of computation known as **Evolutionary Computation**.

Genetic Algorithms are being widely used in different real-world applications, for example, **Designing electronic circuits, code-breaking, image processing, and artificial creativity.**

Key terms-
- **Population:** Population is the subset of all possible or probable solutions, which can solve the given problem.
- **Chromosomes:** A chromosome is one of the solutions in the population for the given problem, and the collection of gene generate a chromosome.
- **Gene:** A chromosome is divided into a different gene, or it is an element of the chromosome.
- **Allele:** Allele is the value provided to the gene within a particular chromosome.
- **Fitness Function:** The fitness function is used to determine the individual's fitness level in the population. It means the ability of an individual to compete with other individuals. In every iteration, individuals are evaluated based on their fitness function.
- **Genetic Operators:** In a genetic algorithm, the best individual mate to regenerate offspring better than parents. Here genetic operators play a role in changing the genetic composition of the next generation.

## How Genetic Algorithm Work?

The genetic algorithm works on the evolutionary generational cycle to generate high-quality solutions. These algorithms use different operations that either enhance or replace the population to give an improved fit solution.

It basically involves five phases to solve the complex optimization problems, which are given as below:

- **Initialization**
- **Fitness Assignment**
- **Selection**
- **Reproduction**
- **Termination**

### 1. Initialization

The process of a genetic algorithm starts by generating the set of individuals, which is called population. Here each individual is the solution for the given problem. An individual contains or is characterized by a set of parameters called Genes. Genes are combined into a string and generate chromosomes, which is the solution to the problem. One of the most popular techniques for initialization is the use of random binary strings.

ADVERTISEMENT

ADVERTISEMENT

![Genetic Algorithm in Machine Learning](https://static.javatpoint.com/tutorial/machine-learning/images/genetic-algorithm-in-machine-learning2.png)

### 2. Fitness Assignment

Fitness function is used to determine how fit an individual is? It means the ability of an individual to compete with other individuals. In every iteration, individuals are evaluated based on their fitness function. The fitness function provides a fitness score to each individual. This score further determines the probability of being selected for reproduction. The high the fitness score, the more chances of getting selected for reproduction.

### 3. Selection

The selection phase involves the selection of individuals for the reproduction of offspring. All the selected individuals are then arranged in a pair of two to increase reproduction. Then these individuals transfer their genes to the next generation.

There are three types of Selection methods available, which are:

- Roulette wheel selection
- Tournament selection
- Rank-based selection

### 4. Reproduction

After the selection process, the creation of a child occurs in the reproduction step. In this step, the genetic algorithm uses two variation operators that are applied to the parent population. The two operators involved in the reproduction phase are given below:

- **Crossover:** The crossover plays a most significant role in the reproduction phase of the genetic algorithm. In this process, a crossover point is selected at random within the genes. Then the crossover operator swaps genetic information of two parents from the current generation to produce a new individual representing the offspring.  
    ![Genetic Algorithm in Machine Learning](https://static.javatpoint.com/tutorial/machine-learning/images/genetic-algorithm-in-machine-learning3.png)  
    The genes of parents are exchanged among themselves until the crossover point is met. These newly generated offspring are added to the population. This process is also called or crossover. Types of crossover styles available:
    - One point crossover
    - Two-point crossover
    - Livery crossover
    - Inheritable Algorithms crossover
- **Mutation**  
    The mutation operator inserts random genes in the offspring (new child) to maintain the diversity in the population. It can be done by flipping some bits in the chromosomes.  
    Mutation helps in solving the issue of premature convergence and enhances diversification. The below image shows the mutation process:  
    Types of mutation styles available,
    
    - **Flip bit mutation**
    - **Gaussian mutation**
    - **Exchange/Swap mutation**
    
      
    ![Genetic Algorithm in Machine Learning](https://static.javatpoint.com/tutorial/machine-learning/images/genetic-algorithm-in-machine-learning4.png)

### 5. Termination

After the reproduction phase, a stopping criterion is applied as a base for termination. The algorithm terminates after the threshold fitness solution is reached. It will identify the final solution as the best solution in the population.

## General Workflow of a Simple Genetic Algorithm

![Genetic Algorithm in Machine Learning](https://static.javatpoint.com/tutorial/machine-learning/images/genetic-algorithm-in-machine-learning5.png)

## Advantages of Genetic Algorithm

- The parallel capabilities of genetic algorithms are best.
- It helps in optimizing various problems such as discrete functions, multi-objective problems, and continuous functions.
- It provides a solution for a problem that improves over time.
- A genetic algorithm does not need derivative information.

ADVERTISEMENT

## Limitations of Genetic Algorithms

- Genetic algorithms are not efficient algorithms for solving simple problems.
- It does not guarantee the quality of the final solution to a problem.
- Repetitive calculation of fitness values may generate some computational challenges.

## Difference between Genetic Algorithms and Traditional Algorithms

- A search space is the set of all possible solutions to the problem. In the traditional algorithm, only one set of solutions is maintained, whereas, in a genetic algorithm, several sets of solutions in search space can be used.
- Traditional algorithms need more information in order to perform a search, whereas genetic algorithms need only one objective function to calculate the fitness of an individual.
- Traditional Algorithms cannot work parallelly, whereas genetic Algorithms can work parallelly (calculating the fitness of the individualities are independent).
- One big difference in genetic Algorithms is that rather of operating directly on seeker results, inheritable algorithms operate on their representations (or rendering), frequently appertained to as chromosomes.
- One of the big differences between traditional algorithm and genetic algorithm is that it does not directly operate on candidate solutions.
- Traditional Algorithms can only generate one result in the end, whereas Genetic Algorithms can generate multiple optimal results from different generations.
- The traditional algorithm is not more likely to generate optimal results, whereas Genetic algorithms do not guarantee to generate optimal global results, but also there is a great possibility of getting the optimal result for a problem as it uses genetic operators such as Crossover and Mutation.
- Traditional algorithms are deterministic in nature, whereas Genetic algorithms are probabilistic and stochastic in nature.

# Encoding methods in genetic algorithm
LINKS:
>[medium.com- encoding techniques][https://medium.com/geekculture/encoding-techniques-in-genetic-algorithm-371bccbe4bf7]
>[tutorials point- encoding][https://www.obitko.com/tutorials/genetic-algorithms/encoding.php]

### Binary Encoding

Binary encoding is the most common, mainly because first works about GA used this type of encoding.
In **binary encoding** every chromosome is a string of **bits**, **0** or **1**.

|              |                          |
| ------------ | ------------------------ |
| Chromosome A | 101100101100101011100101 |
| Chromosome B | 111111100000110000011111 |
_Example of chromosomes with binary encoding_

Binary encoding gives many possible chromosomes even with a small number of alleles. On the other hand, this encoding is often not natural for many problems and sometimes corrections must be made after crossover and/or mutation.

> **Example of Problem:** Knapsack problem  
> **The problem:** There are things with given value and size. The knapsack has given capacity. Select things to maximize the value of things in knapsack, but do not extend knapsack capacity.  
> **Encoding:** Each bit says, if the corresponding thing is in knapsack.

### Permutation Encoding
Permutation encoding can be used in ordering problems, such as travelling salesman problem or task ordering problem.
In **permutation encoding**, every chromosome is a string of numbers, which represents number in a **sequence**.

|              |                           |
| ------------ | ------------------------- |
| Chromosome A | 1  5  3  2  6  4  7  9  8 |
| Chromosome B | 8  5  6  7  2  3  1  4  9 |
_Example of chromosomes with permutation encoding_

Permutation encoding is only useful for ordering problems. Even for this problems for some types of crossover and mutation corrections must be made to leave the chromosome consistent (i.e. have real sequence in it).

> **Example of Problem:** Travelling salesman problem (TSP)  
> **The problem:** There are cities and given distances between them.Travelling salesman has to visit all of them, but he does not to travel very much. Find a sequence of cities to minimize travelled distance.  
> **Encoding:** Chromosome says order of cities, in which salesman will visit them.  

### Value Encoding

Direct value encoding can be used in problems, where some complicated value, such as real numbers, are used. Use of binary encoding for this type of problems would be very difficult.

In **value encoding**, every chromosome is a string of some values. Values can be anything connected to problem, form numbers, real numbers or chars to some complicated objects.

|              |                                            |
| ------------ | ------------------------------------------ |
| Chromosome A | 1.2324  5.3243  0.4556  2.3293  2.4545     |
| Chromosome B | ABDJEIFJDHDIERJFDLDFLFEGT                  |
| Chromosome C | (back), (back), (right), (forward), (left) |
_Example of chromosomes with value encoding_
Value encoding is very good for some special problems. On the other hand, for this encoding is often necessary to develop some new crossover and mutation specific for the problem.

> **Example of Problem:** Finding weights for neural network  
> **The problem:** There is some neural network with given architecture. Find weights for inputs of neurons to train the network for wanted output.  
> **Encoding:** Real values in chromosomes represent corresponding weights for inputs.  
###   Tree Encoding
Tree encoding is used mainly for evolving programs or expressions, for **genetic programming**.
In **tree encoding** every chromosome is a tree of some objects, such as functions or commands in programming language.

|                                                                                 |                                                                                 |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| Chromosome A                                                                    | Chromosome B                                                                    |
| ![image](https://www.obitko.com/tutorials/genetic-algorithms/images/treec1.gif) | ![image](https://www.obitko.com/tutorials/genetic-algorithms/images/treec2.gif) |
| ( +  x  ( /  5  y ) )                                                           | ( do_until  step  wall )                                                        |
_Example of chromosomes with tree encoding_
Tree encoding is good for evolving programs. Programing language LISP is often used to this, because programs in it are represented in this form and can be easily parsed as a tree, so the crossover and mutation can be done relatively easily.

> **Example of Problem:** Finding a function from given values  
> **The problem:** Some input and output values are given. Task is to find a function, which will give the best (closest to wanted) output to all inputs.  
> **Encoding:** Chromosome are functions represented in a tree.  

# Selection methods
### Roulette Wheel Selection

In a roulette wheel selection, the circular wheel is divided as described before. A fixed point is chosen on the wheel circumference as shown and the wheel is rotated. The region of the wheel which comes in front of the fixed point is chosen as the parent. For the second parent, the same process is repeated.

![Roulette Wheel Selection](https://www.tutorialspoint.com/genetic_algorithms/images/roulette_wheel_selection.jpg)

It is clear that a fitter individual has a greater pie on the wheel and therefore a greater chance of landing in front of the fixed point when the wheel is rotated. Therefore, the probability of choosing an individual depends directly on its fitness.

Implementation wise, we use the following steps −

- Calculate S = the sum of a finesses.
    
- Generate a random number between 0 and S.
    
- Starting from the top of the population, keep adding the finesses to the partial sum P, till P<S.
    
- The individual for which P exceeds S is the chosen individual.
    

### Stochastic Universal Sampling (SUS)

Stochastic Universal Sampling is quite similar to Roulette wheel selection, however instead of having just one fixed point, we have multiple fixed points as shown in the following image. Therefore, all the parents are chosen in just one spin of the wheel. Also, such a setup encourages the highly fit individuals to be chosen at least once.

![SUS](https://www.tutorialspoint.com/genetic_algorithms/images/sus.jpg)

It is to be noted that fitness proportionate selection methods don’t work for cases where the fitness can take a negative value.

## Tournament Selection

In K-Way tournament selection, we select K individuals from the population at random and select the best out of these to become a parent. The same process is repeated for selecting the next parent. Tournament Selection is also extremely popular in literature as it can even work with negative fitness values.

![Tournament Selection](https://www.tutorialspoint.com/genetic_algorithms/images/tournament_selection.jpg)

## Rank Selection

Rank Selection also works with negative fitness values and is mostly used when the individuals in the population have very close fitness values (this happens usually at the end of the run). This leads to each individual having an almost equal share of the pie (like in case of fitness proportionate selection) as shown in the following image and hence each individual no matter how fit relative to each other has an approximately same probability of getting selected as a parent. This in turn leads to a loss in the selection pressure towards fitter individuals, making the GA to make poor parent selections in such situations.

![Rank Selection](https://www.tutorialspoint.com/genetic_algorithms/images/rank_selection.jpg)

In this, we remove the concept of a fitness value while selecting a parent. However, every individual in the population is ranked according to their fitness. The selection of the parents depends on the rank of each individual and not the fitness. The higher ranked individuals are preferred more than the lower ranked ones.

|Chromosome|Fitness Value|Rank|
|---|---|---|
|A|8.1|1|
|B|8.0|4|
|C|8.05|2|
|D|7.95|6|
|E|8.02|3|
|F|7.99|5|

## Random Selection

In this strategy we randomly select parents from the existing population. There is no selection pressure towards fitter individuals and therefore this strategy is usually avoided.

# Cross over
1. one point cross over
2. Multipoint cross over
3. Uniform crossover
4. Artithmetic Recombination
5. Davis Order Crossover 
# Mutation
1. Bit flip
2. Random Resetting
3. Scramble 
4. Swap 
5. Inversion