// Compiler targets of the playground: `value` is the wasm function that compiles
// a whole module for that language.
export const targets = [
  { value: "goGen", label: "Golang" },
  { value: "kotlinGen", label: "Kotlin" },
  { value: "jsGenModule", label: "JavaScript" },
  { value: "pythonGenModule", label: "Python" },
  { value: "dartGenModule", label: "Dart" },
  { value: "csharpGenModule", label: "C#" },
  { value: "javaGenModule", label: "Java" },
  { value: "phpGenModule", label: "PHP" },
  { value: "cGenModule", label: "C" },
  { value: "cppGenModule", label: "C++" },
  { value: "sqlQueryPredict", label: "QueryPredict(SQL)" },
  { value: "entitySqlGen", label: "Entity(SQL)" },
  { value: "preprocessorGen", label: "Preprocessor" },
  { value: "postmanGen", label: "Postman" },
  { value: "openapiGen", label: "OpenApi" },
  { value: "mdGen", label: "Markdown" },
  { value: "swiftGen", label: "Swift(All)" },
];

export const targetLabel = (value: string) =>
  targets.find((t) => t.value === value)?.label ?? value;
