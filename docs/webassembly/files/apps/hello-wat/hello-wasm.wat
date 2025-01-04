(module
    (import "console" "log" (func $log (param i32)))
    (memory 1)
    (export "memory" (memory 0))
    (data (i32.const 0) "Hello, WebAssembly!")
    (func $main 
        i32.const 0
        call $log)
    (export "main" (func $main))
)