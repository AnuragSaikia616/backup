# SYLLABUS
**Learning with neural networks** - towards cognitive machine, neuron models, network architectures, perceptron, linear neuron and the widrow-hoff learning rule, the error-correction delta rule, multi-layer perceptron (MLP) networks and the error-backpropagation algorithm, multi-class discrimination with MLP networks, radial basis functions (RBF) networks, genetic-neural systems 
**Fuzzy inference systems**-introduction, cognitive uncertainty and fuzzy rule-base, fuzzy quantification of knowledge, fuzzy rule-base and approximate reasoning, Mamdani model for fuzzy inference systems, takagi-sugeno fuzzy model, neuro-fuzzy inference systems, genetic-fuzzy systems

## Convolutional neural network 

[CLICK!!][https://www.simplilearn.com/tutorials/deep-learning-tutorial/convolutional-neural-network]
Convolutional Neural Networks (CNNs) are a class of deep learning algorithms primarily used for processing structured grid data, such as images. CNNs leverage spatial hierarchies in data by employing convolution operations to detect features at various levels of abstraction. Here's an explanation of CNNs and the different steps involved in their operation:
### Components and Steps in CNNs

1. **Convolution Layer**
   - **Purpose**: To detect features in the input data using filters (kernels).
   - **Operation**: A filter slides (convolves) over the input image, computing the dot product between the filter weights and the input. This produces a feature map.
   - **Formula**: For an input \( X \) and filter \( W \),
     \[
     (X * W)[i, j] = \sum_m \sum_n X[i+m, j+n] \cdot W[m, n]
     \]

2. **Activation Function**
   - **Purpose**: To introduce non-linearity into the model, allowing it to learn more complex patterns.
   - **Common Functions**: ReLU (Rectified Linear Unit), Sigmoid, Tanh.
   - **ReLU Formula**:
     \[
     \text{ReLU}(x) = \max(0, x)
     \]

3. **Pooling Layer**
   - **Purpose**: To reduce the spatial dimensions of the feature maps, lowering the computational load and controlling overfitting.
   - **Types**: Max Pooling, Average Pooling.
   - **Operation**: Pooling operations take the maximum or average value in a specified window (usually 2x2) across the feature map.
   - **Max Pooling Formula**:
     \[
     \text{MaxPool}(X)[i, j] = \max(X[i:i+f, j:j+f])
     \]
     where \( f \) is the size of the pooling window.

4. **Flattening**
   - **Purpose**: To convert the pooled feature map into a one-dimensional vector that can be fed into fully connected layers.
   - **Operation**: The multidimensional feature maps are flattened into a single vector.

5. **Fully Connected Layer (Dense Layer)**
   - **Purpose**: To combine features learned by the convolutional and pooling layers to make final predictions.
   - **Operation**: Each neuron in the fully connected layer is connected to every neuron in the previous layer, similar to a traditional neural network.
   - **Formula**: For weights \( W \), bias \( b \), and input \( x \),
     \[
     y = W \cdot x + b
     \]

6. **Output Layer**
   - **Purpose**: To produce the final output of the network, such as class probabilities in classification tasks.
   - **Activation Function**: Typically, softmax for multi-class classification or sigmoid for binary classification.
   - **Softmax Formula**:
     \[
     \text{Softmax}(x_i) = \frac{e^{x_i}}{\sum_{j} e^{x_j}}
     \]

### Workflow of a CNN

1. **Input**: An image or any grid-like data.
2. **Convolution**: Apply multiple convolution layers to extract features from the input.
3. **Activation**: Apply non-linear activation functions to the feature maps.
4. **Pooling**: Reduce the spatial dimensions of the feature maps.
5. **Flattening**: Convert the pooled feature maps into a one-dimensional vector.
6. **Fully Connected Layers**: Combine the features to make predictions.
7. **Output**: Generate the final prediction (e.g., class labels).

### Example Architecture
![[Pasted image 20240623022722.png]]
CNNs are powerful for tasks involving image data, providing robust feature extraction capabilities through their hierarchical structure and convolutional operations.