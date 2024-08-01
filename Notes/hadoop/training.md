#### Ensemble Methods

Ensemble is learning by aggregating multiple models.

- Types:

1. Bagging(Random Forest):
   Bagging is also called bootstrap aggregation.
   Can reduce the variance of those algorithms which have high variance.
   Some algorithms which have high variance are decision trees, support vector machines, KNN.

   Classification- Majority
   Regression- Average

   Out of bag(OOB) sampling: refers to those data points that are not included in the bootstrap sample for training the model.

2. Boosting
3. Stacking
4. Voting

#### Sampling types:

1. Sampling with replacement: Element can be selected multiple times.
2. Sampling without replacement: Element can be selected only one at a time.

#### Convolutional neural network

1D 2D 3D

- Activation Functions

1. ReLU: Rectified Linear Unit or Leaky ReLU. It is a piecewise linear function that will output 0 for negative input values and the input value for positive input values.
2. Sigmoid
3. Tanh
4. Softmax

# Java libraries

In Java, certain classes and interfaces are used to handle a variety of functionalities, such as file I/O and data management. Let's go through the specifics of the `BufferedReader`, `IOException`, `InputStreamReader`, `ArrayList`, and the `List` interface, including their common imports and usage:

### 1. `BufferedReader`

**Import Statement:**

```java
import java.io.BufferedReader;
```

**Description:**
`BufferedReader` reads text from a character-input stream, buffering characters to provide for the efficient reading of characters, arrays, and lines. The buffer size may be specified, or the default size may be used. The default is large enough for most purposes.

**Usage Example:**

```java
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class ReadFile {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new FileReader("example.txt"));
        String line;
        while ((line = reader.readLine()) != null) {
            System.out.println(line);
        }
        reader.close();
    }
}
```

This example reads a file line by line using `BufferedReader`.

### 2. `IOException`

**Import Statement:**

```java
import java.io.IOException;
```

**Description:**
`IOException` is a checked exception that is thrown to indicate that an I/O operation failed or was interrupted. Any code that deals with I/O in Java is required to handle or throw `IOException`.

**Usage Example:**
Already included in the `BufferedReader` example above, where `IOException` is handled in the method signature.

### 3. `InputStreamReader`

**Import Statement:**

```java
import java.io.InputStreamReader;
```

**Description:**
`InputStreamReader` is a bridge from byte streams to character streams. It reads bytes and decodes them into characters using a specified charset. It forms the basis for wrapping around byte-oriented streams to be used in character-based processing.

**Usage Example:**

```java
import java.io.InputStreamReader;
import java.io.BufferedReader;

public class ReadFromStandardInput {
    public static void main(String[] args) {
        InputStreamReader isr = new InputStreamReader(System.in);
        BufferedReader br = new BufferedReader(isr);
        System.out.println("Enter your name:");
        try {
            String name = br.readLine();
            System.out.println("Hello, " + name);
            br.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```

This code reads from standard input (keyboard) using `InputStreamReader`.

### 4. `ArrayList`

**Import Statement:**

```java
import java.util.ArrayList;
```

**Description:**
`ArrayList` is an implementation of the `List` interface that uses a dynamic array to store its elements. It provides methods to manipulate the size of the array that stores the elements.

**Usage Example:**

```java
import java.util.ArrayList;

public class Example {
    public static void main(String[] args) {
        ArrayList<String> list = new ArrayList<>();
        list.add("Java");
        list.add("Python");
        list.add("C++");
        System.out.println(list);
    }
}
```

This example initializes an `ArrayList`, adds several elements to it, and prints the list.

### 5. `List` Interface

**Import Statement:**

```java
import java.util.List;
```

**Description:**
`List` is an interface in Java's collection framework that represents an ordered collection (also known as a sequence). This interface allows precise control over where each element is inserted and can access elements by their integer index.

**Usage Example:**

```java
import java.util.List;
import java.util.ArrayList;

public class ListExample {
    public static void main(String[] args) {
        List<String> list = new ArrayList<>();
        list.add("HTML");
        list.add("CSS");
        list.add("JavaScript");
        System.out.println(list);
    }
}
```

This example uses `List` to declare a variable and `ArrayList` for the actual implementation.

These classes and interfaces provide a robust framework for performing a variety of essential tasks in Java, such as reading from and writing to files, handling exceptions, and managing collections of objects.

#  Hadoop libraries
The import statements you've provided are specific to Apache Hadoop's MapReduce framework, which is used for processing large data sets with a distributed algorithm on a Hadoop cluster. Let's go through each of these import statements to explain what classes they bring into a Java program and how they are typically used in a Hadoop application:

### 1. `org.apache.hadoop.conf.Configuration`
- **Usage:** The `Configuration` class is part of Hadoop’s core libraries, providing access to configuration parameters. Configurations for Hadoop are typically specified in XML files, and this class provides methods to read these configuration settings. It is used to configure various aspects of the job such as input paths, output paths, mapper classes, reducer classes, and any other specific job settings.

### 2. `org.apache.hadoop.fs.Path`
- **Usage:** The `Path` class represents a file path in Hadoop's filesystem (HDFS). It is used to specify the input and output directories for MapReduce jobs. When setting up a job, you specify paths as instances of this class.

### 3. `org.apache.hadoop.io.IntWritable` and `org.apache.hadoop.io.Text`
- **Usage:** These classes are part of Hadoop's custom writable types that serve as the Java primitives for Hadoop's serialization framework. `IntWritable` is a wrapper class for the `int` primitive data type, making it a serializable object. `Text` is Hadoop's implementation for a serializable and comparable string. These classes are commonly used as input and output types for keys and values in MapReduce tasks.

### 4. `org.apache.hadoop.mapreduce.Job`
- **Usage:** The `Job` class represents a MapReduce job configuration. It allows the user to configure the job, submit it, control its execution, and query the state. The job specifies settings like which input/output formats to use, which mapper and reducer classes to run, and other job tuning parameters.

### 5. `org.apache.hadoop.mapreduce.Mapper`
- **Usage:** The `Mapper` class is used to define a map job in MapReduce. Each mapper processes a split of data and produces key-value pairs as output. The output from the mapper becomes the input to the reducer. This class is extended by any job that requires map capabilities.

### 6. `org.apache.hadoop.mapreduce.Reducer`
- **Usage:** The `Reducer` class is used to define a reduce job in MapReduce. It processes the key-value pairs generated by mappers and typically performs a summary operation, emitting a smaller, aggregated result. Like `Mapper`, this class must be extended to define specific reduce operations for a job.

### 7. `org.apache.hadoop.mapreduce.lib.input.FileInputFormat`
- **Usage:** This class is an input format for reading data from files in HDFS. It defines how input files are read and split into records for processing by the Mapper. The class is typically extended to handle custom input formats if the default behaviors are not suitable.

### 8. `org.apache.hadoop.mapreduce.lib.output.FileOutputFormat`
- **Usage:** This class handles formatting and writing output to files in HDFS. It defines the output specifications for MapReduce jobs, such as setting the output path, managing output files, and formatting output keys and values. This class can also be extended to define custom output formats.

These classes are crucial for setting up, configuring, running, and managing the data flow of MapReduce jobs on a Hadoop cluster, facilitating the distributed processing of large data sets.


# Clustering using mapReduce
### 1. Initializing 


# 