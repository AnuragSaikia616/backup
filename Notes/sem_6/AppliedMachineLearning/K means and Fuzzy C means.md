## [K means clustering + Elbow method + Python implementation][https://www.javatpoint.com/k-means-clustering-algorithm-in-machine-learning]

## [Fuzzy C means clustering][https://medium.com/geekculture/fuzzy-c-means-clustering-fcm-algorithm-in-machine-learning-c2e51e586fff]
Fuzzy K-means clustering, also known as Fuzzy C-means (FCM) clustering, is an extension of the traditional K-means clustering algorithm. Unlike K-means, where each data point belongs to exactly one cluster, FCM allows each data point to belong to multiple clusters with varying degrees of membership. This makes FCM a more flexible and robust clustering method, particularly for datasets where cluster boundaries are not well-defined.

1. **Membership Matrix**:
    
    - Each data point has a membership value associated with each cluster, indicating the degree of belonging to that cluster.
    - These membership values are constrained to be between 0 and 1, and the sum of membership values for each data point across all clusters is 1.

2. **Fuzziness Coefficient (m)**:
    
    - This parameter determines the level of cluster fuzziness. Higher values of mmm result in fuzzier clusters, while m=1m = 1m=1 makes the algorithm equivalent to standard K-means.
3. **Cluster Centers Update**:
    
    - The centers of the clusters are updated using the weighted average of all data points, where the weights are the membership values.

### Steps in Fuzzy K-means Clustering:

![[Pasted image 20240605214301.png]]

The process flow of fuzzy c-means is enumerated below:

1. Assume a fixed number of clusters k.
2. Initialization: Randomly initialize the k-means μk associated with the clusters and compute the probability that each data point xi is a member of a given cluster k, P(point xi has label k|xi, k).
		example: ![[Pasted image 20240605214935.png]]
1. Iteration: Recalculate the centroid of the cluster as the weighted centroid given the probabilities of membership of all data points xi:
		![[Pasted image 20240605214626.png]]
4. Termination: Iterate until convergence or until a user-specified number of iterations has been reached (the iteration may be trapped at some local maxima or minima).	

### ADVANTAGES
1. Flexibility: Fuzzy clustering allows for overlapping clusters, which can be useful when the data has a complex structure or when there are ambiguous or overlapping class boundaries.
2. Robustness: Fuzzy clustering can be more robust to outliers and noise in the data, as it allows for a more gradual transition from one cluster to another.
3. Interpretability: Fuzzy clustering provides a more nuanced understanding of the structure of the data, as it allows for a more detailed representation of the relationships between data points and clusters.
### DISADVANTAGES
1. Complexity: Fuzzy clustering algorithms can be computationally more expensive than traditional clustering algorithms, as they require optimization over multiple membership degrees.
2. Model selection: Choosing the right number of clusters and membership functions can be challenging, and may require expert knowledge or trial and error.
3. If you’re interested in learning more about fuzzy clustering, you might consider reading “Fuzzy Clustering and Its Applications” by James C. Bezdek or “An Introduction to Fuzzy Clustering” by Witold Pedrycz and Fernando Gomide.