# 📚 Type Casting and Inference in Go

---

## 🧠 Type Inference
(**Go automatically figures out the type**)

When you declare a variable with `:=`, Go **infers** the type.

Example:

```go
name := "ChatGPT"     // inferred as string
age := 5              // inferred as int
isCool := true        // inferred as bool
height := 5.9         // inferred as float64
```

✅ No need to manually write `var name string = "ChatGPT"` — Go figures it out.

> **Note:**  
> You can only use `:=` **inside functions**.  
> Outside functions, you must use `var`.

---

## 🎯 Type Casting
(**Manually converting types**)

Go does **NOT automatically convert** types.  
You must **explicitly cast** them.

Example:

```go
var x int = 10
var y float64 = float64(x) // cast int to float64
```

Or:

```go
var f float64 = 3.5
var i int = int(f) // cast float64 to int (truncates to 3)
```

✅ Always use the format: `TargetType(value)`

---

## 🚀 Common Type Casting Examples

| From | To | Example |
|:----|:---|:--------|
| int → float64 | `float64(x)` |
| float64 → int | `int(f)` |
| string → int | `strconv.Atoi("123")` |
| int → string | `strconv.Itoa(123)` |
| []byte → string | `string(bytes)` |
| string → []byte | `[]byte(str)` |

> **Important:**  
> For `string ↔ int`, use the `strconv` package, not direct casting.

---

## ⚡ Example: String to Int

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Conversion error:", err)
        return
    }
    fmt.Println(num + 1) // Output: 124
}
```

---

## 🛑 Common Mistakes to Avoid
- Trying `string(123)` — ❌ This gives weird ASCII output, not `"123"`.
  - Correct way: `strconv.Itoa(123)`
- Forgetting to check for errors when using `strconv.Atoi()`.
- Expecting Go to **automatically** convert between types — it **never** does.

---

## ✅ Quick Summary

| Concept | Description |
|:-------|:-------------|
| Type Inference | Go guesses the type (with `:=`) |
| Type Casting | You manually convert types (`Type(value)`) |

---

# 🔥 Pro Tip
> "**Go is strict about types, but that’s why Go code is super safe.**"
> 
> "**Always cast explicitly — don't expect auto magic.**"

---
