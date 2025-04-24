# pkg/functions/grammars/bnf_rules.go  
## Package: grammars  
  
Imports:  
- "encoding/json"  
- "regexp"  
  
External data and input sources:  
- The code defines a map called `PRIMITIVE_RULES` containing string representations of grammar rules for various primitive data types.  
- It also defines a regular expression `INVALID_RULE_CHARS_RE` to match invalid characters in grammar rules.  
- The `GRAMMAR_LITERAL_ESCAPE_RE` regular expression is used to match characters that need to be escaped in grammar literals.  
- The `GRAMMAR_LITERAL_ESCAPES` map provides escape sequences for specific characters.  
  
TODOs:  
- In the description of the `freestring` rule, there's a TODO comment mentioning that the grammar shouldn't forbid certain characters.  
  
### Grammar Rules  
  
This section defines grammar rules for various primitive data types, such as boolean, number, integer, string, and null. It also includes rules for arrays with and without newlines.  
  
### Utility Functions  
  
The code includes a utility function `jsonString` that takes an interface{} value and returns its JSON string representation.  
  
### Constants  
  
The code defines constants for space rule and array rules with and without newlines.  
  
### Regular Expressions  
  
The code defines regular expressions for matching invalid characters in grammar rules and escaping special characters in grammar literals.  
  
### Data Structures  
  
The code defines a map called `PRIMITIVE_RULES` containing string representations of grammar rules for various primitive data types.  
  
  
  
# pkg/functions/grammars/json_schema.go  
Package: grammars  
  
Imports:  
	"encoding/json"  
	"fmt"  
	"sort"  
	"strings"  
  
External data, input sources:  
	- The code takes a JSON schema as input, which is expected to be a map[string]interface{} type.  
  
TODOs:  
	- There are no TODO comments in the code.  
  
Summary:  
	- The JSONSchemaConverter struct is responsible for converting a JSON schema into a grammar. It takes a string representing the order of properties as input and uses it to create a map of property names to their order.  
	- The addRule method adds a new rule to the grammar with a given name and rule string. If a rule with the same name already exists, it appends a number to the name to create a unique identifier.  
	- The visit method recursively traverses the JSON schema and generates grammar rules for each part of the schema. It handles various schema types, such as objects, arrays, and primitive types.  
	- The resolveReference method handles references to definitions in the schema.  
	- The Grammar method takes a JSON schema and returns a grammar string, which can be used to generate text samples based on the schema.  
	- The GrammarFromBytes method takes a byte slice containing a JSON schema and returns a grammar string.  
  
	- The code defines a set of primitive rules for common data types, such as strings, numbers, and booleans.  
	- The code also includes a set of rules for handling arrays and objects, which are common data structures in JSON schemas.  
	- The code is designed to be flexible and can handle a wide range of JSON schemas.  
  
	- The code is well-documented and easy to understand.  
	- The code is well-structured and organized, making it easy to maintain and extend.  
	- The code is efficient and performs well, even when handling large JSON schemas.  
  
# pkg/functions/grammars/llama31_schema.go  
Package: grammars  
  
Imports:  
	"encoding/json"  
	"fmt"  
	"regexp"  
	"sort"  
	"strings"  
  
External data and input sources:  
	- The code uses JSON data as input, which is expected to be provided in the form of a map[string]interface{}.  
  
TODOs:  
	- There are no TODO comments in the provided code.  
  
Summary:  
	- The package contains a struct called LLama31SchemaConverter, which is used to convert a JSON schema into a grammar.  
	- The converter has a method called Grammar, which takes a JSON schema as input and returns a grammar string.  
	- The Grammar method first adds a rule called "freestring" to the converter's rules map.  
	- Then, it calls the visit method to traverse the schema and generate rules for each part of the schema.  
	- The visit method handles different types of schema elements, such as objects, arrays, and primitive types.  
	- For each element, it generates a rule that matches the structure of the element.  
	- Once all rules have been generated, the Grammar method returns the rules as a grammar string.  
	- The package also contains a method called GrammarFromBytes, which takes a byte slice containing JSON data as input and returns a grammar string.  
	- This method first unmarshals the JSON data into a map[string]interface{}, then calls the Grammar method to convert the schema into a grammar string.  
  
	- The code defines a number of constants and variables that are used in the conversion process.  
	- For example, there are constants for the different types of primitive values, such as strings, numbers, and booleans.  
	- There are also variables that store the rules for each type of schema element.  
	- The code uses regular expressions to match patterns in the schema, such as the names of properties and the types of values.  
	- It also uses string manipulation functions to format the rules and the grammar string.  
  
	- The code is well-organized and easy to understand.  
	- The use of comments and clear variable names makes it easy to follow the logic of the code.  
	- The code is also well-documented, with a number of comments explaining the purpose of the different functions and methods.  
  
# pkg/functions/grammars/options.go  
Package: grammars  
  
Imports:  
- SchemaConverterType  
  
External data, input sources:  
- None  
  
TODOs:  
- None  
  
Summary:  
The grammars package provides a GrammarOption struct and related functions for customizing the behavior of a grammar. The GrammarOption struct contains various options such as property order, prefix, and schema type. The package offers functions to enable or disable specific options, set prefixes and property orders, and specify schema types and function names. These functions can be used to configure the grammar according to the user's needs.  
  
- GrammarOption struct: This struct holds the configuration options for the grammar.  
- Apply method: This method allows users to apply multiple options to the GrammarOption struct.  
- Option functions: These functions enable or disable specific options, such as EnableMaybeArray, DisableParallelNewLines, EnableMaybeString, NoMixedFreeString, and ExpectStringsAfterJSON.  
- SetPrefix and SetPropOrder functions: These functions allow users to set the prefix and property order for the grammar.  
- WithSchemaType and WithFunctionName functions: These functions allow users to specify the schema type and function name for the grammar.  
  
  
  
# pkg/functions/grammars/rules.go  
Package name: grammars  
  
Imports:  
- "fmt"  
- "strings"  
- "github.com/mudler/LocalAI/pkg/utils"  
  
External data, input sources:  
- None  
  
TODOs:  
- None  
  
Summary:  
- The `Rules` type is a map of strings to strings, representing the rules of a grammar.  
- The `ToGrammar` function takes a set of rules and a list of options, and returns a string representation of the grammar.  
- The function first applies the options to a `GrammarOption` object.  
- It then iterates through the rules, and for each rule, it checks if the root rule needs to be modified based on the options.  
- If the root rule needs to be modified, it appends the modified root rule to the list of lines.  
- If the root rule does not need to be modified, it appends the original root rule to the list of lines.  
- Finally, the function joins the lines together with newline characters and returns the resulting string.  
  
  
  
# pkg/functions/grammars/types.go  
Package: grammars  
  
Imports:  
  
External data, input sources:  
  
TODOs:  
  
Summary:  
  
The grammars package provides functionality for converting between different schema types. It defines two schema types: JSONSchema and LLama31Schema. The package includes functions for converting between these schema types and for creating new schema types.  
  
The SchemaConverterType type is an enumeration that represents the different schema types. The JSONSchema and LLama31Schema constants represent the JSON and LLama 3.1 schema types, respectively.  
  
The String() method of the SchemaConverterType type returns a string representation of the schema type.  
  
The NewType() function takes a string argument representing the schema type and returns a SchemaConverterType value.  
  
The package is designed to be used in conjunction with other packages that work with schema data, such as the schema package.  
  
