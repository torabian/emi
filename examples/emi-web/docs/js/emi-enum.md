---
sidebar_position: 4
---

# Enums in modules

Besides defining inline enums, emi definition allows for standalone enums which is really useful,
for sharing enums between multiple files.

 
```yaml
name: emiEnums
enums:
  - name: enum1
    fields:
      - k: Key1
        value: Value1
      - k: Key2
        value: Value2

```


```ts
/**
 * Enum Enum1
 */
export enum Enum1 {
  /*
   *
   **/
  Key1 = "Key1",
  /*
   *
   **/
  Key2 = "Key2",
}

```