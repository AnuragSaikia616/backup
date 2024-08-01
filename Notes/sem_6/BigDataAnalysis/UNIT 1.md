syllabus:
Introduction To Big Data - Distributed file system, Big Data and its importance, Four Vs, Drivers for Big data, big data analytics, big data applications. Algorithms using map reduce, Matrix-Vector Multiplication by Map Reduce.


#### What is big data
Big Data is the acquisition and extraction of a large amount of data from numerous sources, including social media platforms, Web Logs, sensors, IoT devices, etc. The data may be structured, semi-structured or even unstructured. Additionally, Big Data help in the generation of insightful information from massive data. It helps refine marketing techniques, campaigns, machine learning projects, and various analytics operations. 

#### Why is big data
The rapid expansion and development of social media platforms and mobile apps have resulted in large volumes of data flooding the market. Billions of users have shown interest and are attracted to social media platforms, due to which data is rising. As a result, the need to handle, process and store these large volumes of data requires Big Data. Furthermore, the business organisations in the market are at an additional advantage considering that Big Data Analytics has been revolutionising the IT sector. Big Data uses new-age analytics, mining, statistics and machine learning. Moreover, with the help of Big Data, organisations can now perform multiple operations, store TBs of data, and pre-process, analyse and visualise it effectively. 

#### Big data vs traditional data

| **Aspect**         | **Big Data**                                                | **Traditional Data**                                |
| ------------------ | ----------------------------------------------------------- | --------------------------------------------------- |
| **Volume**         | Extremely large datasets (petabytes, exabytes, or more)     | Smaller datasets (gigabytes, terabytes)             |
| **Variety**        | Structured, semi-structured, and unstructured data          | Primarily structured data                           |
| **Velocity**       | High-speed data generation and processing                   | Slower data generation and batch processing         |
| **Veracity**       | Includes uncertain or imprecise data, requiring validation  | Generally more accurate and well-defined            |
| **Value**          | Extracting valuable insights from vast and varied datasets  | Traditional analytics for business intelligence     |
| **Storage**        | Distributed storage systems (Hadoop, NoSQL databases)       | Relational databases (SQL, traditional DBMS)        |
| **Processing**     | Parallel processing, real-time analytics (MapReduce, Spark) | Batch processing, transactional systems             |
| **Scalability**    | Highly scalable across distributed systems                  | Limited scalability with vertical scaling           |
| **Tools**          | Hadoop, Spark, NoSQL databases, data lakes                  | RDBMS (MySQL, Oracle, SQL Server), data warehouses  |
| **Data Sources**   | Diverse sources (social media, IoT, sensors, logs)          | Structured sources (ERP, CRM systems)               |
| **Complexity**     | Complex data integration and analysis                       | Less complex, predefined schemas                    |
| **Infrastructure** | Requires high-performance, distributed infrastructure       | Can be managed with traditional servers and storage |

#### Importance of big data
The importance of Big Data revolves around more than just the amount of data a company has. Big Data’s importance revolves around how the company can use the collected data. Effectively, the importance of Big Data for companies has various reasons, which include the following: 

- Cost Savings- storing the collection of large volumes of data occurs efficiently with enough security and confidentiality. Accordingly, Big Data tools like Hadoop, Apache, Spark, etc., helps keep the data. 

- Time-saving- Big Data uses real-time in-memory analytics that helps collect data from multiple sources. Using Hadoop further helps analyse the data and make efficient decisions faster. 

- Understanding Customer Behaviour: the market conditions change with changes in consumer behaviour. Effectively, Big Data helps identify these changes and enables organisations to produce products with higher demand. Further, it helps in achieving a competitive advantage for the company. 

- Social Media Listening: Social media platforms help companies to reach their target customers. By undertaking sentiment analysis using Big Data tools, companies can gain feedback from customers on what they think about the brand. Furthermore, Big Data can help in improving the online presence of companies. 

- Boost Customer Acquisition and Retention: Businesses can identify and understand their customers, demands, and needs. Significantly, Big Data tools helps analyse customer demand trends and patterns. It helps in acquiring new customers and retaining customers. 

- Innovation and Product Development: Companies can use Big Data tools to enable product development and foster innovation. Effectively, it ensures that companies can grow and enhance profitability in the market. 


#### Types of big data

- Structured Data- the data collected, analysed and stored in a fixed format using spreadsheets and databases are structured. With the help of advancing techniques, it has been possible to extract meaningful value from these data. 
- Unstructured Data: any form of data that has no proper structure or an unknown form is unstructured data. This type of data is challenging to derive valuable insights from because of the raw nature of the data.
- Semi-structured Data: this type of data consists of both structured and unstructured data. A semi-structured data has a structured form, but the data cannot be defined. 

#### The V's of big data (chatacteristics of big data)
Big data is a collection of data from many different sources and is often describe by five characteristics: ==volume, value, variety, velocity, and veracity==.

- Volume: the size and amounts of big data that companies manage and analyze

- Value: the most important “V” from the perspective of the business, the value of big data usually comes from insight discovery and pattern recognition that lead to more effective operations, stronger customer relationships and other clear and quantifiable business benefits

- Variety: the diversity and range of different data types, including unstructured data, semi-structured data and raw data

 - Velocity: the speed at which companies receive, store and manage data – e.g., the specific number of social media posts or search queries received within a day, hour or other unit of time
  
- Veracity: the “truth” or accuracy of data and information assets, which often determines executive-level confidence


#### Applications of big data (benefits of big data)
The use of Big Data analytics in today’s market allows businesses to have immense advantages. The explanation of the benefits of Big Data is as follows: 

- Customer Acquisition and Retention 

	Big Data helps analyse customer behaviour in the market and collects data on customer feedback regarding their purchase of different products. Accordingly, customers in the current market demand to be treated respectfully and acknowledged for their investment. Especially in the case of online purchases that customers indulge in, they want to receive gratitude for the same. Significantly, Big Data helps in this aspect and thanks the customers for their investment, increasing engagement. Additionally, there are times when customers complain about specific products and require a brand to take action. Big Data helps take real-time steps by checking the customer profile and enabling reputation management. 

- Product Development

	Big Data helps companies to enable product development within their business. Therefore, Understanding customer demands and feedback on existing products and engaging with them through social media helps collect data. Moreover, it allows companies to make innovations within their products for redevelopment to gain higher customer satisfaction.

- Enhance Manufacturing Processes

	With the help of Big Data, you can make minor changes within a product’s images and test different variations of Computer-Aided Design. Furthermore, it helps understand the impact of minor changes and becomes a crucial step in manufacturing. 

- Competitive Advantage 

	Businesses use Predictive Analysis to analyse future trends and patterns in the market. Big Data facilitates the analysis to analyse the data and provides valuable insights. For instance, identifying trends and patterns from social media feed and news reports and analysing them using Big Data helps you understand your competitor’s strategies. As a result, it helps develop strategies that might take you ahead of your competitors. 

- Risk Management
	
	Business organisations typically operate in high-risk environments and require practical solutions and plans to eliminate the risks. Big Data helps eradicate chances by planning risk management processes and strategies. 

- Market Trends and Patterns

	Big Data analytics help identify the customer’s trends and patterns in terms of the type of products and services they demand. By focusing on customer feedback, companies can understand in-depth requirements. It further helps the company to induce customisations within their products for higher customer engagement. 

#### Algorithms using map reduce
MapReduce is a programming model and associated implementation for processing and generating large datasets that is amenable to parallelization across multiple nodes in a cluster (a distributed computing environment). It consists of two main phases: the Map phase and the Reduce phase. MapReduce algorithms are designed to leverage this framework to perform distributed computations efficiently. Here are some common algorithms implemented using MapReduce:

- Word Count: This is one of the simplest examples of a MapReduce algorithm. It involves counting the frequency of each word in a large corpus of text documents. The Map phase emits key-value pairs where the key is the word and the value is 1. The Reduce phase aggregates the counts for each word by summing up the values associated with each key.

- PageRank: PageRank is a link analysis algorithm used by search engines to rank web pages in search results. In the Map phase, each page emits its page rank divided by the number of outbound links to each linked page. The Reduce phase aggregates these contributions and computes the new page ranks based on the damping factor and the contributions from incoming links.

- K-Means Clustering: K-Means is a popular clustering algorithm used for partitioning a dataset into K clusters. In the Map phase, each data point is assigned to the nearest cluster centroid. The Reduce phase updates the cluster centroids based on the mean of the data points assigned to each cluster. This process iterates until convergence.

- Matrix Multiplication: Matrix multiplication is a fundamental operation in linear algebra. In a distributed environment, it can be efficiently implemented using MapReduce. The Map phase partitions the input matrices into blocks and emits intermediate key-value pairs representing partial products. The Reduce phase aggregates these partial products to compute the final result.

- Naive Bayes Classifier: Naive Bayes is a probabilistic classification algorithm commonly used in machine learning. In the Map phase, each training instance emits key-value pairs representing the conditional probabilities of each feature given each class label. The Reduce phase aggregates these probabilities to compute the final likelihoods and class priors.

- Inverted Index: An inverted index is a data structure used in information retrieval systems to map terms to the documents in which they occur. In the Map phase, each document emits key-value pairs representing the terms and their corresponding document IDs. The Reduce phase aggregates these pairs to construct the inverted index.