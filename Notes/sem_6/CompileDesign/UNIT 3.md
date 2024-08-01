# SYLLABUS: 
Semantic Analysis: Semantic Errors, Attribute Grammar, Synthesized attributes, Static Allocation, Stack Allocation, Heap Allocation, Activation Trees, Symbol Table, Intermediate Code Generation and Code Intermediate languages, Declarations, Assignment Statements, Boolean Expressions, Case Statements, DAG representation of Basic Blocks, A simple Code generator from DAG, Issues in the Design of Code Generator.


## Symantic Errors

We have mentioned some of the semantics errors that the semantic analyzer is expected to recognize:

- Type mismatch
- Undeclared variable
- Invalid pointer operations
- Unreachable code
- Invalid return types
- Multiple declaration of variable in a scope.
- Accessing an out of scope variable.
- Actual and formal parameter mismatch.

## STACK VS HEAP ALLOCATION
| Parameter                    | STACK                                                      | HEAP                                                    |
| ---------------------------- | ---------------------------------------------------------- | ------------------------------------------------------- |
| Basic                        | Memory is allocated in a contiguous block.                 | Memory is allocated in any random order.                |
| Allocation and De-allocation | Automatic by compiler instructions.                        | Manual by the programmer.                               |
| Cost                         | Less                                                       | More                                                    |
| Implementation               | Easy                                                       | Hard                                                    |
| Access time                  | Faster                                                     | Slower                                                  |
| Main Issue                   | Shortage of memory                                         | Memory fragmentation                                    |
| Locality of reference        | Excellent                                                  | Adequate                                                |
| Safety                       | Thread safe, data stored can only be accessed by the owner | Not Thread safe, data stored visible to all threads     |
| Flexibility                  | Fixed-size                                                 | Resizing is possible                                    |
| Data type structure          | Linear                                                     | Hierarchical                                            |
| Preferred                    | Static memory allocation is preferred in an array.         | Heap memory allocation is preferred in the linked list. |
| Size                         | Small than heap memory.                                    | Larger than stack memory.                               |
|                              |                                                            |                                                         |
## Symbol table
- Symbol table is a very important data structure in compiler design.
- Symbol table is used to store various information about entitires such as variables, constants, procedures, classes, objects, interfaces etc.
- It is used by both synthesis and analysis phases.
- The symbol table has the following purposes:
	1. It is used to store the names of entities in a structured format.
	2. it is used to determine the scope of a var.
	3. it is used to verify if a var has been declared.
	4. it is used for type checking...
#### Implementation 
Symbol tables can be implemented with **unordered list** if the compiler is being used to handle small amount of data.
A symbol tables can be implemented as:
1. Linear(sorted or unsorted) list
2. Hash table
3. Binary search tree
Symbol tables are mostly implemented as hash tables.
#### Operations in symbol table
1. insert() :
2. lookup() :
## Activation Record
An activation record holds all the crutial information required to call a procedure. It stores immediate values as well as temporary values. An *activation tree* is a structure used to represent the function calls.
- The execution of a precedure is called an activation.
- Each procecure's activation record is pushed into the stack in a predictatble order.
- The activations are pushed in a haierachical manner consisting of parents and children.
- Each procedure can branch into multiple children procudures to form a tree like structure called the activation tree.
#### What does activation record store?
1. local vars: holds the data that is local to the execution of the function.
2. temp vars: to store the values of evaluted expressions.
3. Machine status: to store the status of the machine just before the fn call.
4. Access links: to store the local var data of other activation records.
5. Control links: to store the address of other activation records.
6. Actual parameters: values of arguments passed into fns.
7. Return values: to store the result of a function call.

![[Pasted image 20240610165546.png]]

## Intermediate code
- Intermediate code is used to translate source code into machine code.
- It acts like a bridge between high level language and machine leavel language.
- If a compiler directly converts source code into machine code without converting it into intermediate code first, then a specific compiler would be required for each new type of machine.
- Intermediate code helps by making a machine independent code.
- *Portability*: because of this it can be used by multiple computer architectures. Compiler can generate code for multiple platforms with minimal changes.

Examples of intermediate code are:
1. 3 address code
2. Abstract syntax tree
3. DAG
4. Postfix notation
#### Postfix notation
- It is a useful form of intermediate code representation for expressions.
- Here the the operator follows the operand.
- Example: 
		INFIX: (a + b ) * c
		POSTFIX: ab+c*
- It is also called 'suffix' or 'reverse polish' notation.
- It is a linear representation of syntax tree.

#### Abstract Syntax tree (AST)
- In parse tree most of the child nodes have no siblings.
- In syntax tree (AST) we can eliminate this extra information as it is actually not needed.
- In syntax tree interior nodes are operator and the leaf nodes are operands.
- Example: id + id * id
	![[Pasted image 20240610171239.png]]

#### Three address codes (TAC)
- 3 address code is an intermediate code. It is used by optimizing compilers.
- It uses at most 3 operands.
- It can be easily converted into assembly code.
- *Simplicity* it can represent complex expressions in straightforward way.
- USES in compiler design:
	1. **Syntax analysis**: Conversion into TAC.
	2. **Optimization**: Performing transformations such as constant folding, step reduction and common subexpression elimination.
	3. **Code generation**: Converting TAC into machine code or into lower level intermediate code.
	![[Pasted image 20240610172400.png]]