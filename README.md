# nil-is-not-nil

Run the code:
```sh
go run main.go
```

It turns out:
```text
(*A)(nil)
        is not nil
        is (*A)(nil)
nil
        is nil
        is not *A
(*T)(nil)
        is not nil
        is I, but not nil
        is (*T)(nil)
(I)(nil)
        is nil
        is not I
        is not *T
nil
        is nil
        is not I
        is not *T
```
