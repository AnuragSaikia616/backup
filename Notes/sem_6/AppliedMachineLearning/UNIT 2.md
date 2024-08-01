**Supervised learning:** rationale and basics, learning from observations, bias and variance, why learning works: computational learning theory, occam's razor principle and overfitting avoidance, heuristic search in inductive learning, estimating generalization errors, metrics for assessing regression (numeric prediction) accuracy, metrics for assessing classification (pattern recognition) accuracy, an overview of the design cycle and issues in machine learning. 
**Learning with support vector machines** (SVM) and Random Forests-introduction, linear discriminant functions for binary classification, perceptron algorithm, linear maximal margin classifier for linearly separable data, linear soft margin classifier for overlapping classes, nonlinear classifier, regression by support vector machines, , Decision tree learning, Building a decision tree, combining weak to strong learners via random forest, choosing a split with information gain.

# Supervised learning 
## Basics
## Bias and variance
In machine learning, an error is a measure of how accurately an algorithm can make predictions for the previously unknown dataset. On the basis of these errors, the machine learning model is selected that can perform best on the particular dataset. There are mainly two types of errors in machine learning, which are:

![Bias and Variance in Machine Learning](https://static.javatpoint.com/tutorial/machine-learning/images/bias-and-variance-in-machine-learning2.png)

- **Reducible errors:** These errors can be reduced to improve the model accuracy. Such errors can further be classified into bias and Variance.  
- **Irreducible errors:** These errors will always be present in the model regardless of which algorithm has been used. The cause of these errors is unknown variables whose value can't be reduced.
### Bias
**_While making predictions, a difference occurs between prediction values made by the model and actual values/expected values_**, **_and this difference is known as bias errors or Errors due to bias_**. 

It can be defined as an inability of machine learning algorithms such as Linear Regression to capture the true relationship between the data points. 

Each algorithm begins with some amount of bias because bias occurs from assumptions in the model, which makes the target function simple to learn. 

A model has either:
- **Low Bias:** A low bias model will make fewer assumptions about the form of the target function.
- **High Bias:** A model with a high bias makes more assumptions, and the model becomes unable to capture the important features of our dataset. **A high bias model also cannot perform well on new data.**

Generally, a linear algorithm has a high bias, as it makes them learn fast. The simpler the algorithm, the higher the bias it has likely to be introduced. Whereas a nonlinear algorithm often has low bias.

Some examples of machine learning algorithms with low bias **are Decision Trees, k-Nearest Neighbours and Support Vector Machines**. At the same time, an algorithm with high bias is **Linear Regression, Linear Discriminant Analysis and Logistic Regression.**

**To reduce high bias:**
- Increase the input features as the model is underfitted.
- Decrease the regularization term.
- Use more complex models, such as including some polynomial features.

### Variance
The variance would specify the amount of variation in the prediction if the different training data was used. Ideally, a model should not vary too much from one training dataset to another, which means the algorithm should be good in understanding the hidden mapping between inputs and output variables. 

Variance errors are either of **low variance or high variance.**
- **Low variance** means there is a small variation in the prediction of the target function with changes in the training data set. 
- **High variance** shows a large variation in the prediction of the target function with changes in the training dataset.

A model that shows high variance learns a lot and perform well with the training dataset, and does not generalize well with the unseen dataset. As a result, such a model gives good results with the training dataset but shows high error rates on the test dataset.

Since, with high variance, the model learns too much from the dataset, it leads to overfitting of the model. A model with high variance has the below problems:

- A high variance model leads to overfitting.
- Increase model complexities.

![Bias and Variance in Machine Learning](https://static.javatpoint.com/tutorial/machine-learning/images/bias-and-variance-in-machine-learning3.png)

**To reduce high variance:**
- Reduce the input features or number of parameters as a model is overfitted.
- Do not use a much complex model.
- Increase the training data.
- Increase the Regularization term.

### Combination of bias and variance
![[Pasted image 20240621173537.png]]
### Identifying bias and variance
![Bias and Variance in Machine Learning](https://static.javatpoint.com/tutorial/machine-learning/images/bias-and-variance-in-machine-learning5.png)

High variance can be identified if the model has:
- Low training error and high test error.

High Bias can be identified if the model has:
- High training error and the test error is almost similar to training error.

### Bias variance tradeoff
![Bias and Variance in Machine Learning](https://static.javatpoint.com/tutorial/machine-learning/images/bias-and-variance-in-machine-learning6.png)

For an accurate prediction of the model, algorithms need a low variance and low bias. But this is not possible because bias and variance are related to each other:

- If we decrease the variance, it will increase the bias.
- If we decrease the bias, it will increase the variance.

So, we need to find a sweet spot between bias and variance to make an optimal model.
Hence, the **_Bias-Variance trade-off is about finding the sweet spot to make a balance between bias and variance errors._**
## Computational learning theory
## Overfitting avoidance
[CLICK!! 8 simple techniques to avoid overfitting- Towards data science][https://towardsdatascience.com/8-simple-techniques-to-prevent-overfitting-4d443da2ef7d]
## Heuristic search in inductive learning
## Metrics for assessing accuracy
[CLICK!! ][https://towardsdatascience.com/metrics-to-evaluate-your-machine-learning-algorithm-f10ba6e38234]
## Design cycle and issues in machine learning

# Support vector Machines
- A support vector machine (SVM) is a [supervised machine learning](https://www.ibm.com/topics/supervised-learning) algorithm that classifies data by finding an optimal line or hyperplane that maximizes the distance between each class in an N-dimensional space.
- SVMs are commonly used within classification problems.
- The number of features in the input data determine if the hyperplane is a line of 2D plane in a n dimentional space.
- Since multiple hyperplanes can exist between classes, maximizing the margin between the points enables the algorithm to find the best possible boundary.
- SVMs can be used for both lienar and non linear classification. But if the classes are not linearly separable then kernel functions are used to transform the data into a higher dimentional space. This application of kernel functions is called the "kernel trick".
- The choice of kernel depends on the characteristics of data and the use case.
### Equation of the Hyperplane

In an nnn-dimensional feature space, the equation of the hyperplane is given by:

w⋅x+b=0w \cdot x + b = 0w⋅x+b=0

where:

- www is the weight vector perpendicular to the hyperplane.
- xxx is the feature vector.
- bbb is the bias term (or intercept).
![[Pasted image 20240622171921.png]]
### Types of SVMs:
1. Linear SVM: **Linear SVM** is a type of Support Vector Machine that finds a linear hyperplane to separate the data points of different classes. It works well when the data is linearly separable, meaning that there exists a straight line (or hyperplane in higher dimensions) that can perfectly divide the data points of different classes.
2. Non Linear SVM: **Non-Linear SVM** is used when the data is not linearly separable. It employs the kernel trick to transform the original feature space into a higher-dimensional space where a linear separator can be found.
### Support Vector Machine Terminology

1. **Hyperplane:** Hyperplane is the decision boundary that is used to separate the data points of different classes in a feature space. In the case of linear classifications, it will be a linear equation i.e. wx+b = 0.
2. **Support Vectors:** Support vectors are the closest data points to the hyperplane, which makes a critical role in deciding the hyperplane and margin. 
3. **Margin**: Margin is the distance between the support vector and hyperplane. The main objective of the support vector machine algorithm is to maximize the margin.  The wider margin indicates better classification performance.
4. **Kernel**: Kernel is the mathematical function, which is used in SVM to map the original input data points into high-dimensional feature spaces, so, that the hyperplane can be easily found out even if the data points are not linearly separable in the original input space. Some of the common kernel functions are linear, polynomial, radial basis function(RBF), and sigmoid.
5. **Hard Margin:** The maximum-margin hyperplane or the hard margin hyperplane is a hyperplane that properly separates the data points of different categories without any misclassifications.
6. **Soft Margin:** When the data is not perfectly separable or contains outliers, SVM permits a soft margin technique. Each data point has a slack variable introduced by the soft-margin SVM formulation, which softens the strict margin requirement and permits certain misclassifications or violations. It discovers a compromise between increasing the margin and reducing violations.
7. **C:** Margin maximisation and misclassification fines are balanced by the regularisation parameter C in SVM. The penalty for going over the margin or misclassifying data items is decided by it. A stricter penalty is imposed with a greater value of C, which results in a smaller margin and perhaps fewer misclassifications.
8. **Hinge Loss:** A typical loss function in SVMs is hinge loss. It punishes incorrect classifications or margin violations. The objective function in SVM is frequently formed by combining it with the regularisation term.
9. **Dual Problem:** A dual Problem of the optimisation problem that requires locating the Lagrange multipliers related to the support vectors can be used to solve SVM. The dual formulation enables the use of kernel tricks and more effective computing.

### Hard SVM
### Soft SVM
[YOUTUBE VIDEO LINK!!][https://www.youtube.com/watch?v=7vSGI9FCCaY]
- Soft margin SVM is an extension of the Support Vector Machine(SVM) that allows for some misclassifications of the training data. This approach is useful when the data is not linearly separable, which is common is real life problems.
- The goal is to maximize the margin while minimizing the classification error.
- Soft SVMs introduce the concept of *slack variables* which allows some flexibility. This allows SVM to handle overlapping data points and noise.

![[Pasted image 20240622193352.png]]
![[Pasted image 20240622193414.png]]

Mathematically, decision function of a soft margin SVM can defined as:

![[Pasted image 20240622193709.png]]

Where:

- ξiξi​ are slack variables representing the margin violations.
- yi are the target labels.
- The term (1−ξi ) represents the minimum required margin for each data point.

The objective function of a soft margin SVM combines the margin maximization with a penalty term for margin violations, minimizing:
![[Pasted image 20240622193753.png]]

# Random Forest
**What is random forest?**
Random forest is a machine learning algorithm that ensembles multiple decision trees to reach a singular, more accureate result or prediction.
It is trained using the bagging method.
The idea of bagging is to combine multiple learning models to get better results.
**BAGGING = (BOOTSTRAPING + AGGREGATION)**
![two tree random forest](http://builtin.com/sites/www.builtin.com/files/styles/ckeditor_optimize/public/inline-images/national/two-tree-random-forest.png)

The method combines the principles of "bagging" (Bootstrap Aggregating) and the "random subspace method" to improve predictive performance and control overfitting.

### Key Concepts

#### 1. Decision Trees

A **decision tree** is a predictive model that maps observations about an item to conclusions about its target value. In trees used for classification, leaves represent class labels, and branches represent conjunctions of features that lead to those class labels.

#### 2. Bagging (Bootstrap Aggregating)

**Bagging** is a technique to improve the stability and accuracy of machine learning algorithms. It reduces variance and helps prevent overfitting. In bagging:

- Multiple datasets are created by sampling with replacement from the original dataset.
- A model (e.g., a decision tree) is trained on each bootstrapped dataset.
- Predictions from each model are averaged (regression) or voted on (classification) to produce a final prediction.

#### 3. Random Subspace Method

The **random subspace method** involves training each decision tree on a random subset of the features, not just a random subset of the data. This helps ensure that the trees are diverse and reduces the correlation between them, which improves the overall performance of the forest.
### Advantages

- **Robustness**: Random forests are robust to noise and overfitting due to averaging (in classification) or averaging/majority voting (in regression).
- **Feature Importance**: They can be used to assess the importance of features.
- **Handling High Dimensionality**: They perform well with large datasets and high-dimensional spaces.
### Hyperparameters

- **Number of Trees (`n_estimators`)**: The number of trees in the forest. More trees generally improve performance but increase computational cost.
- **Maximum Depth (`max_depth`)**: The maximum depth of the trees. Limiting depth can help control overfitting.
- **Minimum Samples Split (`min_samples_split`)**: The minimum number of samples required to split an internal node.
- **Minimum Samples Leaf (`min_samples_leaf`)**: The minimum number of samples required to be at a leaf node.
- **Maximum Features (`max_features`)**: The number of features to consider when looking for the best split.
### Simple Step-by-Step Algorithm for Random Forest

Here's a simple, step-by-step algorithm for building a Random Forest classifier:

1. **Prepare the Dataset**:
    
    - Collect and preprocess the data.
    - Split the data into training and testing sets.
2. **Initialize Parameters**:
    
    - Set the number of trees TTT (e.g., `n_estimators`).
    - Set other hyperparameters (e.g., `max_depth`, `min_samples_split`, `max_features`).
3. **Build Each Tree**:
    
    - For each tree ttt in the forest (from 111 to TTT):
        1. **Bootstrap Sample**:
            - Draw a bootstrap sample (with replacement) from the training data.
        2. **Build Decision Tree**:
            - Grow a decision tree from the bootstrap sample:
                - At each node, randomly select a subset of features (`max_features`).
                - Split the node using the feature that provides the best split according to the chosen criterion (e.g., Gini impurity for classification).
                - Continue splitting until the stopping criteria are met (`max_depth`, `min_samples_split`, etc.).
4. **Aggregate Predictions**:
    
    - For classification: Use majority voting.
    - For regression: Use the average of predictions.
5. **Make Predictions**:
    
    - For a new data point:
        - Pass the data point through each of the TTT trees.
        - Aggregate the predictions from all trees.