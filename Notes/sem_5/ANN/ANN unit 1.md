---
Semester: 5
Syllabus: "[[sem_5&6_syllabus.pdf#page=8]]"
---
#### Biological neurons
##### Convolutional neural networks(CNN)
- [IBM](https://www.ibm.com/topics/convolutional-neural-networks)
- Superior performance with image, video, audio signals. 
- Used often in Compuiter Vision.
- Consists of three main types of layers- 
	1. Convolutional layer 
	2. Pooling layer
	3. Fullt connected (FC) layer
##### Recurrent neural network(RNN)
- uses sequential or time series data.
- The output of the layer is added to the next input and is then fed back to the same layer.
- They share the same paramters with each layer of the network.
- Typically consists of a single layer
- Aplication- NLPs, speech recognition, language translation, image captioning etc
- [Recurrent neural networks](https://www.youtube.com/watch?v=_aCuOwF1ZjU) [IBM](https://www.ibm.com/topics/recurrent-neural-networks)
- We can STACK RNNS to get more complex output.
- PROBLEMS: vanishing gradient(exponentially worse for RNN :( )
- SOLUTION: gating(to decide when to forget and remember given input)

#### Error based learning and Memory based learning
**Error-Based Learning:**
1. **Objective:** Minimize the difference between predicted and actual output.
2. **Approach:** Adjust model parameters to reduce prediction errors.
3. **Common Technique:** Backpropagation in neural networks.
4. **Learning Type:** Often associated with supervised learning using labeled datasets.
5. **Feedback:** Utilizes explicit error signals to update the model during training.

**Memory-Based Learning:**
1. **Objective:** Make predictions based on stored experiences or instances.
2. **Approach:** Relies on similarity between new input and stored instances.
3. **Common Technique:** k-Nearest Neighbors (k-NN) algorithm.
4. **Learning Type:** Instance-based learning; model memorizes training instances.
5. **Feedback:** No explicit adjustment of parameters; uses stored instances directly for prediction.
#### Hebbian learnining [[ANN unit 3#^ba500f]]
#### Competitive learning
The output neurons compete with themselves to become active.

Unlike hebbian learning where multiple neurons can become active(fired), in competitive learning only one neuron can be fired at a time.

This feature helps in classification.

Three basic elements of competitive learning:
1. All the neurons are structural similar(expect for weights).
2. Limit imposed on the strength of each neuron.
3. A mechanism the permits neurons to compete.(a winner takes all)

[Youtube video on Competitive learning](https://www.youtube.com/watch?v=lvJ8YcYTyx8)
#### Implementation of a ANN
```python
import tensorflow as tf
from tensorflow.keras import layers, models

inputs = tf.random.normal(1000, 20)
targets = tf.random.normal(1000, 1)

# Build the neural network
model = models.Sequential([
		   layers.Dense(64, activation='relu', input_shape=(20,)),
		   layers.Dense(32, activation='relu'),
		   layers.Dense(1)
])

# Compile the model
model.compile(optimizer='adam', loss='mse', metrics=['mae'])

# Print the model summary
model.summary()

# Train the model on the dataset
model.fit(inputs, targets, epochs=10, batch_size=32, validation_split=0.2)
```