# Go Type System: Key Implications & Examples

Go enforces strict typing and explicit conversions. Here are some important implications with examples:

---

## 1. Strict Function Argument Types

Many standard library functions require specific types (often `float64` for math).

**Example:**
```go
import "math"

a := 2
b := 3
// math.Pow requires float64, so convert explicitly:
result := math.Pow(float64(a), float64(b)) // 8.0
```

---

## 2. No Implicit Type Conversion

Go does **not** automatically convert between types. You must convert explicitly.

**Example:**
```go
var x int32 = 10
var y int = int(x) // explicit conversion
```

---

## 3. Integer Division Truncates

Dividing two integers results in an integer (fractional part is discarded).

**Example:**
```go
fmt.Println(5 / 2)   // Output: 2
fmt.Println(5.0 / 2) // Output: 2.5
```

---

## 4. Rune and Byte Types

- `rune` is an alias for `int32` (used for Unicode code points).
- `byte` is an alias for `uint8` (used for raw data).

**Example:**
```go
var r rune = 'A'   // Unicode code point
var b byte = 'A'   // Raw byte value
```

---

## 5. Strings Are Immutable

To modify a string, convert it to a slice of bytes or runes.

**Example:**
```go
s := "hello"
b := []byte(s)
b[0] = 'H'
s2 := string(b) // "Hello"
```

---

## 6. Short Variable Declaration

The `:=` shorthand can only be used inside functions, not at the package level.

**Example:**
```go
func main() {
    x := 10 // OK
}
// y := 20 // NOT OK at package level
```

---

## 7. Arrays vs Slices

- Arrays have a fixed size: `[5]int`
- Slices are dynamic: `[]int`

**Example:**
```go
var arr [3]int = [3]int{1, 2, 3}
var slice []int = []int{1, 2, 3}
slice = append(slice, 4)
```

---

## 8. Map Key Types

Only comparable types can be used as map keys (not slices, maps, or functions).

**Example:**
```go
m := map[string]int{"one": 1, "two": 2} // OK
// m2 := map[[]int]int{} // NOT OK, slices are not comparable
```

---

## Summary

- Always check function signatures for required types.
- Use explicit type conversions.
- Understand Go’s strict type system to avoid common pitfalls.