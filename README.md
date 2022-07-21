# SOLID
Golang's implementation of SOLID principles

## What is SOLID?
The SOLID principles tell us how to arrange our functions and data, 
and the goal of the principles is the creation of mind-level software structures that:  
- Tolerate changes
- Are easy to understand
- Are the basis of components that could be used in many software systems  

## SOLID Principles
### Single responsability:
“A module should have one, and only one reason to change”

### Open-Close principle:
“A software artifact should be open for extension but closed for modifications”

### Liskov substitution principle
“What is wanted here is something like the following substitution property: If for each object o1 of type S there is an object o2 of type T such that for all programs P defined in terms of T, the behavior of P is unchanged when o1 is substituted for o2 then S is a subtype of T”

### Interface segregation
“Clients should not be forced to depend on methods they don't use”

### Dependency inversion
“High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should not depend on abstractions”
