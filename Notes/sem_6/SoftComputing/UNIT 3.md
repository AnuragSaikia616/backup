[FUZZY NOTES!!!][https://codecrucks.com/what-is-crisp-set/]



# Fuzzy Relation
Fuzzy relation defines the mapping of variables from one [fuzzy set](https://codecrucks.com/what-and-why-fuzzy-set/?relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=2&relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=2&relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=2&relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=2&relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=2) to another. Like [crisp relation](https://codecrucks.com/crisp-relation-definition-types-and-operations/?relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=0&relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=0&relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=0&relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=0&relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=0&relatedposts_hit=1&relatedposts_origin=460&relatedposts_position=0), we can also define the relation over fuzzy sets.

Let A be a fuzzy set on universe X and B be a fuzzy set on universe Y, then the Cartesian product between fuzzy sets A and B will result in a fuzzy relation R which is contained with the full Cartesian product space or it is a subset of the cartesian product of fuzzy subsets. Formally, we can define fuzzy relation as,

R = A x B

and

R ⊂ (X x Y)

where the relation R has a membership function,

μR(x, y) = μA x B(x, y) = min( μA(x), μB(y) )

A binary fuzzy relation R(X, Y) is called a **bipartite graph** if X ≠ Y.

A binary fuzzy relation R(X, Y) is called **directed graph** or **digraph** if X = Y. , which is denoted as R(X, X) = R(X2)

#### Operations on fuzzy relations
1. Union: max(A,B)
2. Intersection: min(A,b)
3. Complement: 1-R
4. Projection:  sup(R) for X,Y where sup=supermum of the set

# Fuzzyfication V/S Defuzzification
|S.No.|Comparison|Fuzzification|Defuzzification|
|---|---|---|---|
|1.|Basic|Precise data is converted into imprecise data.|Imprecise data is converted into precise data.|
|2.|Definition|Fuzzification is the method of converting a crisp quantity into a fuzzy quantity.|Defuzzification is the inverse process of fuzzification where the mapping is done to convert the fuzzy results into crisp results.|
|3.|Example|Like, Voltmeter|Like, Stepper motor and D/A converter|
|4.|Methods|Intuition, inference, rank ordering, angular fuzzy sets, neural network, etcetera.|Maximum membership principle, centroid method, weighted average method, center of sums, etcetera.|
|5.|Complexity|It is quite simple.|It is quite complicated.|
|6.|Use|It can use IF-THEN rules for fuzzifying the crisp value.|It uses the center of gravity methods to find the centroid of the sets.|
# Fuzzy composition (minmax)
### Fuzzy composition
**Fuzzy composition** can be defined just as it is for crisp (binary) relations. Suppose R is a fuzzy relation on X × Y, S is a fuzzy relation on Y × Z, and T is a fuzzy relation on X × Z; then,

**Fuzzy Max–Min composition** is defined as:
![[Pasted image 20240629033734.png]]


# Fuzzy membership functions
>Fuzzy membership functions are fundamental components in fuzzy logic systems, defining how each point in the input space is mapped to a membership value between 0 and 1.
>These functions describe the degree to which a particular element belongs to a fuzzy set.
>Different types of membership functions can be used depending on the specific application and characteristics of the data.
>Fuzzy membership function is used to convert the crisp input provided to the fuzzy inference system.
>Fuzzy logic itself is not fuzzy, rather it deals with the fuzziness in the data.

some membership functions-
1. Singleton
2. Triangular
3. Trapezoidal
4. Gaussian

