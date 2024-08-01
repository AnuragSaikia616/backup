#### Self organizing neural networks
- SONN or self organizing feature maps or kohenn maps are a type of unsuprevised learning model in ANN.
- The model generates a 2d feature map of the input space.
- **Competitive learning** method is useed for learning process.
- It is used to show higher dimentional data in lower dimention. Such projection is called **Topology conserving**.
- SOM has two layers- (a) fully connected input layer (b) Output (map) layer.
- [GFG SONNs](https://www.geeksforgeeks.org/ann-self-organizing-neural-network-sonn/)
##### ALGORITHM: [video](https://www.youtube.com/watch?v=9ZhwKv_bUx8)
1. Initialization:
	Initialize a grid of neurons, each associated with a weight vector of the same dimension as the input data.
	Randomly initialize the weights of the neurons.
	
2. Neighborhood Function:
	Define a neighborhood function that quantifies the influence each neuron has on its neighbors during training. This function typically decreases with distance from the winning neuron.
	
3. Training:
	For each input vector in the dataset:
	Find the Best Matching Unit (BMU): the neuron with the weight vector closest to the input vector.
	Update the weights of the BMU and its neighbors based on the input and the neighborhood function.
	Adjust the learning rate and neighborhood size over time to gradually decrease them.
	
4. Repeat:
	Repeat the training process for a specified number of epochs.
	
5. Clustering and Visualization:
	After training, each neuron represents a cluster, and similar input vectors are mapped to nearby neurons in the grid.
	The SOM can be visualized to reveal the underlying structure of the input data.

##### Implementation in python
```python
import numpy as np

def initialize_som(input_size, grid_size):
    return np.random.rand(grid_size[0], grid_size[1], input_size)

def find_best_matching_unit(weights, input_vector):
    distances = np.linalg.norm(weights - input_vector, axis=2)
    return np.unravel_index(np.argmin(distances), distances.shape)

def update_weights(weights, input_vector, bmu_index, learning_rate=0.1):
    weights[bmu_index] += learning_rate * (input_vector - weights[bmu_index])

def train_som(weights, data, epochs=100, learning_rate=0.1):
    for epoch in range(epochs):
        for input_vector in data:
            bmu_index = find_best_matching_unit(weights, input_vector)
            update_weights(weights, input_vector, bmu_index, learning_rate)

# Generate random 2D data for demonstration
np.random.seed(42)
data = np.random.rand(100, 2)

# Create and train a Self-Organizing Map
input_size = 2
grid_size = (5, 5)
som_weights = initialize_som(input_size, grid_size)
train_som(som_weights, data, epochs=100, learning_rate=0.1)

```

#### Perceptron network
##### Algorithm:
1.  Initialize weights and bias to small random values
2. Activation Function: 
		A Step function is used in perceptron.
3. Training: 
	iterate throught the training data.
	for each input predict the outupt
	calculate the error
	update the weigts based of the error and learning rate
4. Repeat the process untill convergence or a specific number iof epochs 

##### Implementation in python
```python
import numpy as np

class Perceptron:
    def __init__(self, input_size, learning_rate=0.01, epochs=100):
        self.weights = np.random.rand(input_size)
        self.bias = np.random.rand()
        self.learning_rate = learning_rate
        self.epochs = epochs

    def step_function(self, z):
        return 1 if z >= 0 else 0

    def predict(self, inputs):
        z = np.dot(self.weights, inputs) + self.bias
        return self.step_function(z)

    def train(self, training_data):
        for epoch in range(self.epochs):
            for inputs, true_label in training_data:
                prediction = self.predict(inputs)
                error = true_label - prediction

                # Update weights and bias
                self.weights += self.learning_rate * error * inputs
                self.bias += self.learning_rate * error

```