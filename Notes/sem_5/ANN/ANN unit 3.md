---
file link: "[[ANN unit 3]]"
---
- Linear association:
	Linear association in the context of associative memory refers to the relationship between *input and* 
	*output patterns that can be represented by a linear transformation*. In other words, the *output pattern* 
	*is a linear combination of the input pattern's components*. This linear association is often established 
	during the training phase of an associative memory network.
	
	Let's consider a simple linear association scenario with a set of input patterns (X) and their corresponding output patterns (Y).
	The linear association can be represented as:
	
	[ Y = W.X + B ]
	
	- \( Y \): Output pattern
	- \( X \): Input pattern
	- \( W \): Weight matrix
	- \( B \): Bias vector
	
	In this linear association model, the weight matrix \( W \) captures the linear relationships between the input and output patterns, and the bias vector \( B \) represents any additional constant terms.
	
	### Training Linear Associations:
	
	1. **Hebbian Learning (Autoassociative):**
	   - If it's an autoassociative memory (associating input patterns with themselves), the weight matrix can be trained using Hebbian learning. The Hebbian learning rule strengthens the connections between neurons that fire together.
	
	2. **Delta Rule (Heteroassociative):**
	   - If it's a heteroassociative memory (associating different input and output patterns), the Delta rule or other supervised learning methods may be used. The Delta rule adjusts weights based on the difference between the predicted and desired outputs.
	
- Pattern association: 
	Pattern association in an associative network refers to the ability of the network to recall or associate specific input patterns with corresponding output patterns. Associative networks are designed to store relationships between different patterns and retrieve the associated patterns when given partial or degraded input.
	
	There are two main types of associative memories: autoassociative memory and heteroassociative memory.
	
	1. **Autoassociative Memory:**
	   - **Objective:** Recall the same pattern that was stored.
	   - **Architecture:** The network is trained to associate an input pattern with itself.
	   - **Example:** If a specific input pattern is presented, the network should recall the same pattern as the output.
	   - **Application:** Autoassociative memories are often used for tasks like pattern completion or noise reduction.
	
	2. **Heteroassociative Memory:**
	   - **Objective:** Associate one set of patterns with another set.
	   - **Architecture:** The network is trained to associate different input patterns with specific output patterns.
	   - **Example:** Given a set of input patterns, the network should recall the corresponding associated output patterns.
	   - **Application:** Heteroassociative memories are used in tasks where different sets of data need to be related or mapped to each other.
	
	### How Pattern Association Works:
	
	1. **Training Phase:**
	   - The connections (weights) between neurons are adjusted to store the associations.
	2. **Testing or Recall Phase:**
	   - When presented with a partial or degraded version of an input pattern, the network should be able to recall the complete or correct associated output pattern.
	3. **Hebbian Learning or Delta Rule:**
	   - The learning rules used during training are often Hebbian learning or the Delta rule, depending on whether the memory is autoassociative or heteroassociative.
	4. **Activation and Recall:**
	   - The activation of certain neurons in the network retrieves associated patterns.
	   - If part of a familiar pattern is presented as input, the network should be able to reconstruct the complete pattern.
	
	- **Network Architecture:** Fully connected network where neurons are connected to all other neurons (except themselves).
	- **Learning Rule:** Hebbian learning.
	- **Operation:** Given a partial or noisy input pattern, the network iteratively updates its state until it converges to a stable state, which should be the closest stored pattern.
	
- Associative memory networks: [tutorials point](https://www.tutorialspoint.com/artificial_neural_network/artificial_neural_network_associate_memory.htm)
	It is like a simplified version of human brain.
	It is used to associate relationship between input and  output patterns.
	It is a content addressabel memory. It enables recollection of data based on the similarity between input patterns and output patterns.
	**content addressable memory (CAM)** refers to a mechanism that allows the network to retrieve information based on the content or features of the input data rather than relying on explicit memory addresses		
	TYPES: 
	1. Auto associative memory:
		- Capable of retrieval of piece of data based on only partial input information presentation.
		- Recovers stored data that most closely remembles the input pattern.
		- Trained using HEBBAIN LEARNING.
		- example: *hopfield network*
		- The training set is equal to the target set. [S] = [T]
		- *Goal is to retrieve the same pattern even if partial distorted pattern is presented.*
		 Distorted pattern(input) --> Associative memory --> Perfect pattern(output)
	1. Hetero associative memroy:
		- Capable of retrieval of data of one category when data of another category is presented.
		- *Goal is the map distinct input pattern to corresponding output pattern.*
		- Training set is not equal to the target set. [S] != [T]
		- Often trained using supervised learning algorithms like DELTA RULE
		- example: *Bidirectional associative Memory* (BAM)

- Hebbs Rule [video](https://www.youtube.com/watch?v=Vl57hK0iFoU) [notes](https://medium.com/analytics-vidhya/hebb-network-c38598e1a7a1)
	Hebbian learning is a synaptic learning rule that states that the connection between two neurons should be strengthened if they are both active simultaneously.
	
	Implementation in python:
	```python
		import numpy as np

		def hebb(input_pattern, target_pattern):
			n_patterns, pattern_size = input_pattern.shape
			# Init weights and biases as 0
			weights_bias = np.zeroes(pattern_size)
			# For earch pattern update the weights using hebbs rule
			for i, pattern in enumerate(input_pattern):
				weights_bias += weights_bias + pattern*target_pattern[i]
				
			return weights_bias


		# We are using hebbs rule to implement a network for a AND gate
		train_set = np.array([[1,1,1],[1,-1,1],[-1,1,1],[-1,-1,1]])
		target_set = np.array([1,-1,-1,-1])
	```
	 ^ba500f
- Delta Rule and gradient decent [notes](https://medium.com/@neuralnets/delta-learning-rule-gradient-descent-neural-networks-f880c168a804) [video](https://www.youtube.com/watch?v=ktGm0WCoQOg)
	The Delta rule, also known as the Widrow-Hoff rule or the Least Mean Squares (LMS) algorithm, is a supervised learning algorithm used for adjusting the weights of a neural network to minimize the difference between the predicted output and the target output. 
	
	Implementation in python:
	```python
		import numpy as np
	
		# Using delta rule to implement gradient decent
		def delta_rule(training_imputs, training_outputs, lr=0.1, epochs=100):
			num_patterns, pattern_size = training_inputs.shape
			_, output_size = training_outputs.shape
	
			weights = np.random.rand(input_size, ouput_size)
			for epoch in range(epochs):
				for i in range(num_patterns):
					input_pattern = training_inputs[i]
					desired_output = training_output[i]
					output = np.dot(input_pattern, weights)
					error = desired_output - output
					# derivative(gradient) of mean squared error wrt weights is equal to -err*input
					gradient = -np.dot(error, input_pattern)
					# update weights by w = w + (-lr * de/dw)
					weights += -learning_rate * gradient
					
			return weights
	```
	
- [Bidirectional Associative Memory](https://studyglance.in/nn/display.php?tno=8&topic=Bidirectional-Associative-Memory):
		
	Bidirectional Associative Memory, or BAM for short, is a type of recurrent neural network designed for associative memory recall. Unlike traditional feed-forward networks, BAM operates in two layers, with each layer serving as both input and output. This bidirectional nature allows BAM to recall patterns in two sets of neurons, enabling associative memory recall in both directions.
	
	In simpler terms, given a pattern in one set, BAM can recall the corresponding pattern in the other set and vice versa, making it "bidirectional".
	
	### How does BAM Work?
	
	1. **Training:** BAM is trained on paired patterns. Each pair consists of patterns from two sets (let's call them X and Y). The weight matrix is adjusted based on the outer product of these patterns.
		    
	1. **Recall:** Given an incomplete or noisy pattern from set X (or Y), BAM iteratively updates the neurons in both sets until the network stabilizes to a known pattern in set Y (or X).
	    
	3. **Stability:** One of the unique features of BAM is its guaranteed stability. It ensures that the recalled patterns won't oscillate indefinitely.
	
- Asociation encoding and decoding:
	Encoding(training): 
		**Objective**: To store associations between input patterns and corresponding outupt patterns in the synaptic weights of the network.
		**Process**:
		1. Input patterns (X) and output patterns (Y) are provided as **Trainin Data**.
		2.  Weights are updated using a **Hebbian learning** rule during training process.
			   [INSERT HEBBIAN RULE]
	Decoding(recall):
		**Objective**: To retrieve an output pattern froma  given input pattern.
		**Process**:
		1. Input pattern: A partial or complete input pattern is provided to get the corresponding output.
		2. Recall operation: The dot product of input pattern X and weights W is calculated.
			The operation can be expressed as: Y = sign(X .WT) where sign is the trigo sign function.
		3. Output pattern: The output pattern Y is the result of the recall operation.

- Stability: 
	last page of [[Unit- 3 Notes_Associate Memory Network.pdf#page=11&selection=82,0,82,9|Unit- 3 Notes_Associate Memory Network, page 11]] 