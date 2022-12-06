# Today I Learned (Go)

## Maps are iterable but not in a predictable order

Doing this will give a different order each time:

```go
m := map[string]string{...}
for key, value := range m {
    // ...
}
```

Solution: sort keys first

```go
  var keys []string
  for key := range m {
      keys = append(keys, key)
  }
  sort.Strings(keys)
  for _, key := range keys {
      value := m[key]
      // ...
  }
```

## Zero Value for generic type

When working with generics, we can't return `nil` when `T any` as we are not working with the pointer or interface.

```go
func foo[T any]() T {
	return nil
}
```

This will show error: `cannot use nil as T value in return statement`

Instead, many solutions:

- Returns the [zero value](https://go.dev/ref/spec#The_zero_value) of the generic type: `*new(T)`

- Declaring an empty object then return it:

```go
func foo[T any]() T {
	var empty T
    return empty
}
```

- If returning the pointer of the generic type, then nil is correct
