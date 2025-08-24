# Emi compiler

Backend for front-end code generation library based on Emi definition files. This will be the next compiler 
for Fireback projects.


## Emi definition

Emi definition is handling common contracts used across the software, such as entities, dtos, actions, web requests,
sockets, enums, errors, tasks and other parts.

## Target languages

Emi compiler is targeting to build backend code using Golang, and clients using different languages.
The type system is based on Go, and will be converted to different languages as clients. In theory,
we should be able to create backend codes out of the definition, but internally only go is needed,
therefor it's the only language which will be invested on.

In case of JavaScript, "Nest.js" framework has additional support, and can be directly used to generate controllers
and requests signature both for client and backend.