(*main.HiroAst)(0xc00000d260)({
 Functions: ([]*main.Function) (len=2 cap=2) {
  (*main.Function)(0xc00008a280)({
   Name: (string) (len=3) "add",
   Args: ([]*main.Arg) (len=2 cap=2) {
    (*main.Arg)(0xc00008a2d0)({
     Pos: (lexer.Position) 1:8,
     VarName: (string) (len=1) "a",
     VarType: (string) (len=3) "int"
    }),
    (*main.Arg)(0xc00008a320)({
     Pos: (lexer.Position) 1:15,
     VarName: (string) (len=1) "b",
     VarType: (string) (len=3) "int"
    })
   },
   Return: (string) (len=3) "int",
   Body: ([]*main.Command) (len=4 cap=4) {
    (*main.Command)(0xc00008a460)({
     Pos: (lexer.Position) 2:2,
     End: (bool) false,
     Print: (*main.Print)(0xc00020e800)({
      Pos: (lexer.Position) 2:2,
      Expression: (string) (len=7) "\"hello\""
     }),
     Call: (*main.Call)(<nil>),
     Expression: (*main.Expression)(<nil>)
    }),
    (*main.Command)(0xc00008a4b0)({
     Pos: (lexer.Position) 3:2,
     End: (bool) false,
     Print: (*main.Print)(0xc00020e880)({
      Pos: (lexer.Position) 3:2,
      Expression: (string) (len=7) "\"hello\""
     }),
     Call: (*main.Call)(<nil>),
     Expression: (*main.Expression)(<nil>)
    }),
    (*main.Command)(0xc00008a500)({
     Pos: (lexer.Position) 4:2,
     End: (bool) false,
     Print: (*main.Print)(<nil>),
     Call: (*main.Call)(<nil>),
     Expression: (*main.Expression)(0xc00000e158)({
      Equality: (*main.Equality)(0xc00022ab60)({
       Comparison: (*main.Comparison)(0xc00022ab40)({
        Addition: (*main.Addition)(0xc00022ab20)({
         Multiplication: (*main.Multiplication)(0xc00022ab00)({
          Unary: (*main.Unary)(0xc00022aae0)({
           Op: (string) "",
           Unary: (*main.Unary)(<nil>),
           Primary: (*main.Primary)(0xc000214f90)({
            Float: (*float64)(<nil>),
            Int: (*int)(0xc000012cf0)(3),
            String: (*string)(<nil>),
            Bool: (*bool)(<nil>),
            Nil: (bool) false,
            SubExpression: (*main.Expression)(<nil>)
           })
          }),
          Op: (string) "",
          Next: (*main.Multiplication)(<nil>)
         }),
         Op: (string) "",
         Next: (*main.Addition)(<nil>)
        }),
        Op: (string) "",
        Next: (*main.Comparison)(<nil>)
       }),
       Op: (string) "",
       Next: (*main.Equality)(<nil>)
      })
     })
    }),
    (*main.Command)(0xc00008a5a0)({
     Pos: (lexer.Position) 5:1,
     End: (bool) true,
     Print: (*main.Print)(<nil>),
     Call: (*main.Call)(<nil>),
     Expression: (*main.Expression)(<nil>)
    })
   }
  }),
  (*main.Function)(0xc00008a6e0)({
   Name: (string) (len=4) "main",
   Args: ([]*main.Arg) <nil>,
   Return: (string) "",
   Body: ([]*main.Command) (len=3 cap=4) {
    (*main.Command)(0xc00008a780)({
     Pos: (lexer.Position) 8:5,
     End: (bool) false,
     Print: (*main.Print)(<nil>),
     Call: (*main.Call)(0xc00008a910)({
      Pos: (lexer.Position) 8:5,
      Name: (string) (len=3) "add",
      Args: ([]*main.Expression) (len=2 cap=2) {
       (*main.Expression)(0xc00000e210)({
        Equality: (*main.Equality)(0xc00022ae20)({
         Comparison: (*main.Comparison)(0xc00022ae00)({
          Addition: (*main.Addition)(0xc00022ade0)({
           Multiplication: (*main.Multiplication)(0xc00022ada0)({
            Unary: (*main.Unary)(0xc00022ace0)({
             Op: (string) "",
             Unary: (*main.Unary)(<nil>),
             Primary: (*main.Primary)(0xc0002152f0)({
              Float: (*float64)(<nil>),
              Int: (*int)(0xc000012d30)(3),
              String: (*string)(<nil>),
              Bool: (*bool)(<nil>),
              Nil: (bool) false,
              SubExpression: (*main.Expression)(<nil>)
             })
            }),
            Op: (string) "",
            Next: (*main.Multiplication)(<nil>)
           }),
           Op: (string) (len=1) "+",
           Next: (*main.Addition)(0xc00022adc0)({
            Multiplication: (*main.Multiplication)(0xc00022ad80)({
             Unary: (*main.Unary)(0xc00022ad60)({
              Op: (string) "",
              Unary: (*main.Unary)(<nil>),
              Primary: (*main.Primary)(0xc000215350)({
               Float: (*float64)(<nil>),
               Int: (*int)(0xc000012d40)(2),
               String: (*string)(<nil>),
               Bool: (*bool)(<nil>),
               Nil: (bool) false,
               SubExpression: (*main.Expression)(<nil>)
              })
             }),
             Op: (string) "",
             Next: (*main.Multiplication)(<nil>)
            }),
            Op: (string) "",
            Next: (*main.Addition)(<nil>)
           })
          }),
          Op: (string) "",
          Next: (*main.Comparison)(<nil>)
         }),
         Op: (string) "",
         Next: (*main.Equality)(<nil>)
        })
       }),
       (*main.Expression)(0xc00000e2a0)({
        Equality: (*main.Equality)(0xc00022af60)({
         Comparison: (*main.Comparison)(0xc00022af40)({
          Addition: (*main.Addition)(0xc00022af20)({
           Multiplication: (*main.Multiplication)(0xc00022af00)({
            Unary: (*main.Unary)(0xc00022aee0)({
             Op: (string) "",
             Unary: (*main.Unary)(<nil>),
             Primary: (*main.Primary)(0xc000215410)({
              Float: (*float64)(<nil>),
              Int: (*int)(0xc000012d50)(2),
              String: (*string)(<nil>),
              Bool: (*bool)(<nil>),
              Nil: (bool) false,
              SubExpression: (*main.Expression)(<nil>)
             })
            }),
            Op: (string) "",
            Next: (*main.Multiplication)(<nil>)
           }),
           Op: (string) "",
           Next: (*main.Addition)(<nil>)
          }),
          Op: (string) "",
          Next: (*main.Comparison)(<nil>)
         }),
         Op: (string) "",
         Next: (*main.Equality)(<nil>)
        })
       })
      }
     }),
     Expression: (*main.Expression)(<nil>)
    }),
    (*main.Command)(0xc00008a960)({
     Pos: (lexer.Position) 9:2,
     End: (bool) false,
     Print: (*main.Print)(0xc00020eb40)({
      Pos: (lexer.Position) 9:2,
      Expression: (string) (len=7) "\"hello\""
     }),
     Call: (*main.Call)(<nil>),
     Expression: (*main.Expression)(<nil>)
    }),
    (*main.Command)(0xc00008a9b0)({
     Pos: (lexer.Position) 10:1,
     End: (bool) true,
     Print: (*main.Print)(<nil>),
     Call: (*main.Call)(<nil>),
     Expression: (*main.Expression)(<nil>)
    })
   }
  })
 }
})
