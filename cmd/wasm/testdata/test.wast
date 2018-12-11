(module
 (type $0 (func (param i64)))
 (type $1 (func (param i32)))
 (type $2 (func (param i32) (result i32)))
 (type $3 (func (param i32 i32) (result i32)))
 (import "env" "printi" (func $import$0 (param i64)))
 (import "env" "prints" (func $import$1 (param i32)))
 (import "env" "read_param" (func $import$2 (param i32) (result i32)))
 (table 0 anyfunc)
 (memory $0 1)
 (data (i32.const 4) " P\00\00")
 (data (i32.const 16) "init\00")
 (export "memory" (memory $0))
 (export "apply" (func $1))
 (func $0 (type $3) (param $var$0 i32) (param $var$1 i32) (result i32)
  (block $label$0 i32
   (call $import$0
    (i64.extend_s/i32
     (get_local $var$0)
    )
   )
   (call $import$1
    (get_local $var$1)
   )
   (i32.const 0)
  )
 )
 (func $1 (type $2) (param $var$0 i32) (result i32)
  (local $var$1 i32)
  (block $label$0 i32
   (block $label$1
    (br_if $label$1
     (i32.eqz
      (call $2
       (get_local $var$0)
       (i32.const 16)
      )
     )
    )
    (return
     (i32.const 0)
    )
   )
   (set_local $var$0
    (call $import$2
     (i32.const 1)
    )
   )
   (set_local $var$1
    (call $import$2
     (i32.const 2)
    )
   )
   (call $import$0
    (i64.extend_s/i32
     (get_local $var$0)
    )
   )
   (call $import$1
    (get_local $var$1)
   )
   (i32.const 0)
  )
 )
 (func $2 (type $3) (param $var$0 i32) (param $var$1 i32) (result i32)
  (local $var$2 i32)
  (local $var$3 i32)
  (block $label$0 i32
   (set_local $var$3
    (i32.load8_u
     (get_local $var$1)
    )
   )
   (block $label$1
    (br_if $label$1
     (i32.eqz
      (tee_local $var$2
       (i32.load8_u
        (get_local $var$0)
       )
      )
     )
    )
    (br_if $label$1
     (i32.ne
      (get_local $var$2)
      (i32.and
       (get_local $var$3)
       (i32.const 255)
      )
     )
    )
    (set_local $var$0
     (i32.add
      (get_local $var$0)
      (i32.const 1)
     )
    )
    (set_local $var$1
     (i32.add
      (get_local $var$1)
      (i32.const 1)
     )
    )
    (loop $label$2
     (set_local $var$3
      (i32.load8_u
       (get_local $var$1)
      )
     )
     (br_if $label$1
      (i32.eqz
       (tee_local $var$2
        (i32.load8_u
         (get_local $var$0)
        )
       )
      )
     )
     (set_local $var$0
      (i32.add
       (get_local $var$0)
       (i32.const 1)
      )
     )
     (set_local $var$1
      (i32.add
       (get_local $var$1)
       (i32.const 1)
      )
     )
     (br_if $label$2
      (i32.eq
       (get_local $var$2)
       (i32.and
        (get_local $var$3)
        (i32.const 255)
       )
      )
     )
    )
   )
   (i32.sub
    (get_local $var$2)
    (i32.and
     (get_local $var$3)
     (i32.const 255)
    )
   )
  )
 )
)

