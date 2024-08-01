https://www.tutorialspoint.com/compiler_design/compiler_design_quick_guide.htm
syllabus:
Introduction to Language Processing System, Compiling Analysis of the Source Program, Phases of a Compiler, Compiler Construction Tools. Lexical Analysis –Regular Expression, Introduction to Finite Automata and Regular Expression, Conversion of Regular Expression to NFA, Role ofLexical Analyzer, Specification of Tokens

## [OVERVIEW][https://www.naukri.com/code360/library/compiler-design-notes]
### Q. What are the different phases of a compiler? 
A compiler is a program that translates source code written in a high-level programming language into machine code that can be executed by a computer's processor. The process of compilation involves several distinct phases, each responsible for a specific part of the translation process. Here are the main phases of a compiler:

1. **Lexical Analysis (Scanning)**:
   - **Purpose**: Converts the sequence of characters in the source code into a sequence of tokens.
   - **Output**: Tokens (e.g., keywords, identifiers, literals, operators, punctuation).
   - **Tools/Concepts**: Lexical analyzer (lexer), regular expressions.

2. **Syntax Analysis (Parsing)**:
   - **Purpose**: Analyzes the sequence of tokens to determine its grammatical structure according to the language's syntax rules.
   - **Output**: Parse tree (syntax tree).
   - **Tools/Concepts**: Parser, context-free grammars, BNF (Backus-Naur Form), parse tree.

3. **Semantic Analysis**:
   - **Purpose**: Checks for semantic errors and ensures that the syntax tree follows the rules of the language (e.g., type checking, scope resolution).
   - **Output**: Annotated syntax tree.
   - **Tools/Concepts**: Symbol table, type checker.

4. **Intermediate Code Generation**:
   - **Purpose**: Translates the syntax tree into an intermediate representation (IR) that is easier to optimize and translate into machine code.
   - **Output**: Intermediate code (e.g., three-address code, abstract syntax tree).
   - **Tools/Concepts**: Intermediate code generator.

5. **Code Optimization**:
   - **Purpose**: Improves the intermediate code to make it more efficient (e.g., reduce execution time, reduce memory usage).
   - **Output**: Optimized intermediate code.
   - **Tools/Concepts**: Optimizer, control flow analysis, data flow analysis.

6. **Code Generation**:
   - **Purpose**: Translates the optimized intermediate code into target machine code.
   - **Output**: Machine code or assembly code.
   - **Tools/Concepts**: Code generator, instruction selection, register allocation.

7. **Code Linking and Assembly**:
   - **Purpose**: Converts the machine code into an executable file, resolves addresses, and links with libraries.
   - **Output**: Executable file.
   - **Tools/Concepts**: Linker, assembler, loader.

Each of these phases plays a critical role in the overall compilation process, ensuring that the source code is accurately and efficiently transformed into executable code.

## Specification of tokens
