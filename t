(*compiler.HiroAst)(0xc00000d668)({
 Functions: ([]*compiler.Function) (len=2 cap=2) {
  (*compiler.Function)(0xc0000827e0)({
   Name: (string) (len=3) "add",
   Args: ([]*compiler.Arg) (len=2 cap=2) {
    (*compiler.Arg)(0xc000010230)({
     Pos: (lexer.Position) 1:8,
     VarName: (string) (len=1) "x",
     VarType: (string) (len=3) "int"
    }),
    (*compiler.Arg)(0xc000010280)({
     Pos: (lexer.Position) 1:15,
     VarName: (string) (len=1) "y",
     VarType: (string) (len=3) "int"
    })
   },
   Return: (string) (len=3) "int",
   Body: ([]*compiler.Command) (len=1 cap=1) {
    (*compiler.Command)(0xc0000103c0)({
     Pos: (lexer.Position) 2:2,
     Let: (*compiler.Let)(<nil>),
     Print: (*compiler.Print)(<nil>),
     If: (*compiler.If)(<nil>),
     Call: (*compiler.Call)(<nil>),
     Expression: (*compiler.Expression)(0xc000075280)({
      IsAsync: (bool) false,
      CheckedForAsync: (bool) false,
      Equality: (*compiler.Equality)(0xc000233760)({
       Comparison: (*compiler.Comparison)(0xc000233740)({
        Addition: (*compiler.Addition)(0xc000233720)({
         Multiplication: (*compiler.Multiplication)(0xc0002336e0)({
          Unary: (*compiler.Unary)(0xc000233620)({
           Op: (string) "",
           Unary: (*compiler.Unary)(<nil>),
           Primary: (*compiler.Primary)(0xc0002050c0)({
            Call: (*compiler.Call)(<nil>),
            Float: (*float64)(<nil>),
            Int: (*int)(<nil>),
            String: (*string)(<nil>),
            Bool: (*bool)(<nil>),
            Nil: (bool) false,
            Variable: (*string)(0xc0000751d0)((len=1) "x"),
            SubExpression: (*compiler.Expression)(<nil>)
           })
          }),
          Op: (string) "",
          Next: (*compiler.Multiplication)(<nil>)
         }),
         Op: (string) (len=1) "+",
         Next: (*compiler.Addition)(0xc000233700)({
          Multiplication: (*compiler.Multiplication)(0xc0002336c0)({
           Unary: (*compiler.Unary)(0xc0002336a0)({
            Op: (string) "",
            Unary: (*compiler.Unary)(<nil>),
            Primary: (*compiler.Primary)(0xc000205140)({
             Call: (*compiler.Call)(<nil>),
             Float: (*float64)(<nil>),
             Int: (*int)(<nil>),
             String: (*string)(<nil>),
             Bool: (*bool)(<nil>),
             Nil: (bool) false,
             Variable: (*string)(0xc000075230)((len=1) "y"),
             SubExpression: (*compiler.Expression)(<nil>)
            })
           }),
           Op: (string) "",
           Next: (*compiler.Multiplication)(<nil>)
          }),
          Op: (string) "",
          Next: (*compiler.Addition)(<nil>)
         })
        }),
        Op: (string) "",
        Next: (*compiler.Comparison)(<nil>)
       }),
       Op: (string) "",
       Next: (*compiler.Equality)(<nil>)
      })
     })
    })
   },
   End: (bool) true
  }),
  (*compiler.Function)(0xc0000829c0)({
   Name: (string) (len=4) "main",
   Args: ([]*compiler.Arg) <nil>,
   Return: (string) "",
   Body: ([]*compiler.Command) (len=5 cap=8) {
    (*compiler.Command)(0xc0000106e0)({
     Pos: (lexer.Position) 6:5,
     Let: (*compiler.Let)(0xc000205580)({
      Pos: (lexer.Position) 6:5,
      Var: (string) (len=1) "a",
      Expr: (*compiler.Expression)(0xc000075570)({
       IsAsync: (bool) false,
       CheckedForAsync: (bool) false,
       Equality: (*compiler.Equality)(0xc000233be0)({
        Comparison: (*compiler.Comparison)(0xc000233bc0)({
         Addition: (*compiler.Addition)(0xc000233ba0)({
          Multiplication: (*compiler.Multiplication)(0xc000233b60)({
           Unary: (*compiler.Unary)(0xc000233aa0)({
            Op: (string) "",
            Unary: (*compiler.Unary)(<nil>),
            Primary: (*compiler.Primary)(0xc0002054c0)({
             Call: (*compiler.Call)(0xc0000108c0)({
              Pos: (lexer.Position) 6:13,
              Name: (string) (len=3) "add",
              Args: ([]*compiler.Expression) (len=2 cap=2) {
               (*compiler.Expression)(0xc000075430)({
                IsAsync: (bool) false,
                CheckedForAsync: (bool) false,
                Equality: (*compiler.Equality)(0xc000233940)({
                 Comparison: (*compiler.Comparison)(0xc000233920)({
                  Addition: (*compiler.Addition)(0xc000233900)({
                   Multiplication: (*compiler.Multiplication)(0xc0002338e0)({
                    Unary: (*compiler.Unary)(0xc0002338c0)({
                     Op: (string) "",
                     Unary: (*compiler.Unary)(<nil>),
                     Primary: (*compiler.Primary)(0xc000205380)({
                      Call: (*compiler.Call)(<nil>),
                      Float: (*float64)(<nil>),
                      Int: (*int)(0xc000012ff0)(2),
                      String: (*string)(<nil>),
                      Bool: (*bool)(<nil>),
                      Nil: (bool) false,
                      Variable: (*string)(<nil>),
                      SubExpression: (*compiler.Expression)(<nil>)
                     })
                    }),
                    Op: (string) "",
                    Next: (*compiler.Multiplication)(<nil>)
                   }),
                   Op: (string) "",
                   Next: (*compiler.Addition)(<nil>)
                  }),
                  Op: (string) "",
                  Next: (*compiler.Comparison)(<nil>)
                 }),
                 Op: (string) "",
                 Next: (*compiler.Equality)(<nil>)
                })
               }),
               (*compiler.Expression)(0xc000075460)({
                IsAsync: (bool) false,
                CheckedForAsync: (bool) false,
                Equality: (*compiler.Equality)(0xc000233a80)({
                 Comparison: (*compiler.Comparison)(0xc000233a60)({
                  Addition: (*compiler.Addition)(0xc000233a40)({
                   Multiplication: (*compiler.Multiplication)(0xc000233a20)({
                    Unary: (*compiler.Unary)(0xc000233a00)({
                     Op: (string) "",
                     Unary: (*compiler.Unary)(<nil>),
                     Primary: (*compiler.Primary)(0xc000205400)({
                      Call: (*compiler.Call)(<nil>),
                      Float: (*float64)(<nil>),
                      Int: (*int)(0xc000013000)(3),
                      String: (*string)(<nil>),
                      Bool: (*bool)(<nil>),
                      Nil: (bool) false,
                      Variable: (*string)(<nil>),
                      SubExpression: (*compiler.Expression)(<nil>)
                     })
                    }),
                    Op: (string) "",
                    Next: (*compiler.Multiplication)(<nil>)
                   }),
                   Op: (string) "",
                   Next: (*compiler.Addition)(<nil>)
                  }),
                  Op: (string) "",
                  Next: (*compiler.Comparison)(<nil>)
                 }),
                 Op: (string) "",
                 Next: (*compiler.Equality)(<nil>)
                })
               })
              }
             }),
             Float: (*float64)(<nil>),
             Int: (*int)(<nil>),
             String: (*string)(<nil>),
             Bool: (*bool)(<nil>),
             Nil: (bool) false,
             Variable: (*string)(<nil>),
             SubExpression: (*compiler.Expression)(<nil>)
            })
           }),
           Op: (string) "",
           Next: (*compiler.Multiplication)(<nil>)
          }),
          Op: (string) (len=1) "+",
          Next: (*compiler.Addition)(0xc000233b80)({
           Multiplication: (*compiler.Multiplication)(0xc000233b40)({
            Unary: (*compiler.Unary)(0xc000233b20)({
             Op: (string) "",
             Unary: (*compiler.Unary)(<nil>),
             Primary: (*compiler.Primary)(0xc000205540)({
              Call: (*compiler.Call)(<nil>),
              Float: (*float64)(<nil>),
              Int: (*int)(0xc000013010)(2),
              String: (*string)(<nil>),
              Bool: (*bool)(<nil>),
              Nil: (bool) false,
              Variable: (*string)(<nil>),
              SubExpression: (*compiler.Expression)(<nil>)
             })
            }),
            Op: (string) "",
            Next: (*compiler.Multiplication)(<nil>)
           }),
           Op: (string) "",
           Next: (*compiler.Addition)(<nil>)
          })
         }),
         Op: (string) "",
         Next: (*compiler.Comparison)(<nil>)
        }),
        Op: (string) "",
        Next: (*compiler.Equality)(<nil>)
       })
      })
     }),
     Print: (*compiler.Print)(<nil>),
     If: (*compiler.If)(<nil>),
     Call: (*compiler.Call)(<nil>),
     Expression: (*compiler.Expression)(<nil>)
    }),
    (*compiler.Command)(0xc0000109b0)({
     Pos: (lexer.Position) 7:5,
     Let: (*compiler.Let)(0xc000205780)({
      Pos: (lexer.Position) 7:5,
      Var: (string) (len=1) "b",
      Expr: (*compiler.Expression)(0xc0000756a0)({
       IsAsync: (bool) false,
       CheckedForAsync: (bool) false,
       Equality: (*compiler.Equality)(0xc000233fa0)({
        Comparison: (*compiler.Comparison)(0xc000233f80)({
         Addition: (*compiler.Addition)(0xc000233f60)({
          Multiplication: (*compiler.Multiplication)(0xc000233f40)({
           Unary: (*compiler.Unary)(0xc000233f20)({
            Op: (string) "",
            Unary: (*compiler.Unary)(<nil>),
            Primary: (*compiler.Primary)(0xc000205740)({
             Call: (*compiler.Call)(0xc000010b90)({
              Pos: (lexer.Position) 7:13,
              Name: (string) (len=3) "add",
              Args: ([]*compiler.Expression) (len=2 cap=2) {
               (*compiler.Expression)(0xc0000755f0)({
                IsAsync: (bool) false,
                CheckedForAsync: (bool) false,
                Equality: (*compiler.Equality)(0xc000233dc0)({
                 Comparison: (*compiler.Comparison)(0xc000233da0)({
                  Addition: (*compiler.Addition)(0xc000233d80)({
                   Multiplication: (*compiler.Multiplication)(0xc000233d60)({
                    Unary: (*compiler.Unary)(0xc000233d40)({
                     Op: (string) "",
                     Unary: (*compiler.Unary)(<nil>),
                     Primary: (*compiler.Primary)(0xc000205680)({
                      Call: (*compiler.Call)(<nil>),
                      Float: (*float64)(<nil>),
                      Int: (*int)(0xc000013020)(2),
                      String: (*string)(<nil>),
                      Bool: (*bool)(<nil>),
                      Nil: (bool) false,
                      Variable: (*string)(<nil>),
                      SubExpression: (*compiler.Expression)(<nil>)
                     })
                    }),
                    Op: (string) "",
                    Next: (*compiler.Multiplication)(<nil>)
                   }),
                   Op: (string) "",
                   Next: (*compiler.Addition)(<nil>)
                  }),
                  Op: (string) "",
                  Next: (*compiler.Comparison)(<nil>)
                 }),
                 Op: (string) "",
                 Next: (*compiler.Equality)(<nil>)
                })
               }),
               (*compiler.Expression)(0xc000075620)({
                IsAsync: (bool) false,
                CheckedForAsync: (bool) false,
                Equality: (*compiler.Equality)(0xc000233f00)({
                 Comparison: (*compiler.Comparison)(0xc000233ee0)({
                  Addition: (*compiler.Addition)(0xc000233ec0)({
                   Multiplication: (*compiler.Multiplication)(0xc000233ea0)({
                    Unary: (*compiler.Unary)(0xc000233e80)({
                     Op: (string) "",
                     Unary: (*compiler.Unary)(<nil>),
                     Primary: (*compiler.Primary)(0xc000205700)({
                      Call: (*compiler.Call)(<nil>),
                      Float: (*float64)(<nil>),
                      Int: (*int)(0xc000013030)(3),
                      String: (*string)(<nil>),
                      Bool: (*bool)(<nil>),
                      Nil: (bool) false,
                      Variable: (*string)(<nil>),
                      SubExpression: (*compiler.Expression)(<nil>)
                     })
                    }),
                    Op: (string) "",
                    Next: (*compiler.Multiplication)(<nil>)
                   }),
                   Op: (string) "",
                   Next: (*compiler.Addition)(<nil>)
                  }),
                  Op: (string) "",
                  Next: (*compiler.Comparison)(<nil>)
                 }),
                 Op: (string) "",
                 Next: (*compiler.Equality)(<nil>)
                })
               })
              }
             }),
             Float: (*float64)(<nil>),
             Int: (*int)(<nil>),
             String: (*string)(<nil>),
             Bool: (*bool)(<nil>),
             Nil: (bool) false,
             Variable: (*string)(<nil>),
             SubExpression: (*compiler.Expression)(<nil>)
            })
           }),
           Op: (string) "",
           Next: (*compiler.Multiplication)(<nil>)
          }),
          Op: (string) "",
          Next: (*compiler.Addition)(<nil>)
         }),
         Op: (string) "",
         Next: (*compiler.Comparison)(<nil>)
        }),
        Op: (string) "",
        Next: (*compiler.Equality)(<nil>)
       })
      })
     }),
     Print: (*compiler.Print)(<nil>),
     If: (*compiler.If)(<nil>),
     Call: (*compiler.Call)(<nil>),
     Expression: (*compiler.Expression)(<nil>)
    }),
    (*compiler.Command)(0xc000010be0)({
     Pos: (lexer.Position) 8:5,
     Let: (*compiler.Let)(<nil>),
     Print: (*compiler.Print)(0xc00025e300)({
      Pos: (lexer.Position) 8:5,
      Expression: (*compiler.Expression)(0xc000075780)({
       IsAsync: (bool) false,
       CheckedForAsync: (bool) false,
       Equality: (*compiler.Equality)(0xc000262380)({
        Comparison: (*compiler.Comparison)(0xc000262360)({
         Addition: (*compiler.Addition)(0xc000262340)({
          Multiplication: (*compiler.Multiplication)(0xc000262320)({
           Unary: (*compiler.Unary)(0xc000262300)({
            Op: (string) "",
            Unary: (*compiler.Unary)(<nil>),
            Primary: (*compiler.Primary)(0xc000205940)({
             Call: (*compiler.Call)(0xc000010dc0)({
              Pos: (lexer.Position) 8:11,
              Name: (string) (len=3) "add",
              Args: ([]*compiler.Expression) (len=2 cap=2) {
               (*compiler.Expression)(0xc0000756f0)({
                IsAsync: (bool) false,
                CheckedForAsync: (bool) false,
                Equality: (*compiler.Equality)(0xc0002621a0)({
                 Comparison: (*compiler.Comparison)(0xc000262180)({
                  Addition: (*compiler.Addition)(0xc000262160)({
                   Multiplication: (*compiler.Multiplication)(0xc000262140)({
                    Unary: (*compiler.Unary)(0xc000262120)({
                     Op: (string) "",
                     Unary: (*compiler.Unary)(<nil>),
                     Primary: (*compiler.Primary)(0xc000205880)({
                      Call: (*compiler.Call)(<nil>),
                      Float: (*float64)(<nil>),
                      Int: (*int)(0xc000013040)(2),
                      String: (*string)(<nil>),
                      Bool: (*bool)(<nil>),
                      Nil: (bool) false,
                      Variable: (*string)(<nil>),
                      SubExpression: (*compiler.Expression)(<nil>)
                     })
                    }),
                    Op: (string) "",
                    Next: (*compiler.Multiplication)(<nil>)
                   }),
                   Op: (string) "",
                   Next: (*compiler.Addition)(<nil>)
                  }),
                  Op: (string) "",
                  Next: (*compiler.Comparison)(<nil>)
                 }),
                 Op: (string) "",
                 Next: (*compiler.Equality)(<nil>)
                })
               }),
               (*compiler.Expression)(0xc000075720)({
                IsAsync: (bool) false,
                CheckedForAsync: (bool) false,
                Equality: (*compiler.Equality)(0xc0002622e0)({
                 Comparison: (*compiler.Comparison)(0xc0002622c0)({
                  Addition: (*compiler.Addition)(0xc0002622a0)({
                   Multiplication: (*compiler.Multiplication)(0xc000262280)({
                    Unary: (*compiler.Unary)(0xc000262260)({
                     Op: (string) "",
                     Unary: (*compiler.Unary)(<nil>),
                     Primary: (*compiler.Primary)(0xc000205900)({
                      Call: (*compiler.Call)(<nil>),
                      Float: (*float64)(<nil>),
                      Int: (*int)(0xc000013050)(3),
                      String: (*string)(<nil>),
                      Bool: (*bool)(<nil>),
                      Nil: (bool) false,
                      Variable: (*string)(<nil>),
                      SubExpression: (*compiler.Expression)(<nil>)
                     })
                    }),
                    Op: (string) "",
                    Next: (*compiler.Multiplication)(<nil>)
                   }),
                   Op: (string) "",
                   Next: (*compiler.Addition)(<nil>)
                  }),
                  Op: (string) "",
                  Next: (*compiler.Comparison)(<nil>)
                 }),
                 Op: (string) "",
                 Next: (*compiler.Equality)(<nil>)
                })
               })
              }
             }),
             Float: (*float64)(<nil>),
             Int: (*int)(<nil>),
             String: (*string)(<nil>),
             Bool: (*bool)(<nil>),
             Nil: (bool) false,
             Variable: (*string)(<nil>),
             SubExpression: (*compiler.Expression)(<nil>)
            })
           }),
           Op: (string) "",
           Next: (*compiler.Multiplication)(<nil>)
          }),
          Op: (string) "",
          Next: (*compiler.Addition)(<nil>)
         }),
         Op: (string) "",
         Next: (*compiler.Comparison)(<nil>)
        }),
        Op: (string) "",
        Next: (*compiler.Equality)(<nil>)
       })
      })
     }),
     If: (*compiler.If)(<nil>),
     Call: (*compiler.Call)(<nil>),
     Expression: (*compiler.Expression)(<nil>)
    }),
    (*compiler.Command)(0xc000010e10)({
     Pos: (lexer.Position) 9:5,
     Let: (*compiler.Let)(<nil>),
     Print: (*compiler.Print)(0xc00025e450)({
      Pos: (lexer.Position) 9:5,
      Expression: (*compiler.Expression)(0xc000075800)({
       IsAsync: (bool) false,
       CheckedForAsync: (bool) false,
       Equality: (*compiler.Equality)(0xc0002624c0)({
        Comparison: (*compiler.Comparison)(0xc0002624a0)({
         Addition: (*compiler.Addition)(0xc000262480)({
          Multiplication: (*compiler.Multiplication)(0xc000262460)({
           Unary: (*compiler.Unary)(0xc000262440)({
            Op: (string) "",
            Unary: (*compiler.Unary)(<nil>),
            Primary: (*compiler.Primary)(0xc000205a00)({
             Call: (*compiler.Call)(<nil>),
             Float: (*float64)(<nil>),
             Int: (*int)(<nil>),
             String: (*string)(<nil>),
             Bool: (*bool)(<nil>),
             Nil: (bool) false,
             Variable: (*string)(0xc0000757e0)((len=1) "b"),
             SubExpression: (*compiler.Expression)(<nil>)
            })
           }),
           Op: (string) "",
           Next: (*compiler.Multiplication)(<nil>)
          }),
          Op: (string) "",
          Next: (*compiler.Addition)(<nil>)
         }),
         Op: (string) "",
         Next: (*compiler.Comparison)(<nil>)
        }),
        Op: (string) "",
        Next: (*compiler.Equality)(<nil>)
       })
      })
     }),
     If: (*compiler.If)(<nil>),
     Call: (*compiler.Call)(<nil>),
     Expression: (*compiler.Expression)(<nil>)
    }),
    (*compiler.Command)(0xc000010f00)({
     Pos: (lexer.Position) 10:5,
     Let: (*compiler.Let)(<nil>),
     Print: (*compiler.Print)(0xc00025e5a0)({
      Pos: (lexer.Position) 10:5,
      Expression: (*compiler.Expression)(0xc000075880)({
       IsAsync: (bool) false,
       CheckedForAsync: (bool) false,
       Equality: (*compiler.Equality)(0xc000262600)({
        Comparison: (*compiler.Comparison)(0xc0002625e0)({
         Addition: (*compiler.Addition)(0xc0002625c0)({
          Multiplication: (*compiler.Multiplication)(0xc0002625a0)({
           Unary: (*compiler.Unary)(0xc000262580)({
            Op: (string) "",
            Unary: (*compiler.Unary)(<nil>),
            Primary: (*compiler.Primary)(0xc000205b00)({
             Call: (*compiler.Call)(<nil>),
             Float: (*float64)(<nil>),
             Int: (*int)(<nil>),
             String: (*string)(<nil>),
             Bool: (*bool)(<nil>),
             Nil: (bool) false,
             Variable: (*string)(0xc000075860)((len=1) "a"),
             SubExpression: (*compiler.Expression)(<nil>)
            })
           }),
           Op: (string) "",
           Next: (*compiler.Multiplication)(<nil>)
          }),
          Op: (string) "",
          Next: (*compiler.Addition)(<nil>)
         }),
         Op: (string) "",
         Next: (*compiler.Comparison)(<nil>)
        }),
        Op: (string) "",
        Next: (*compiler.Equality)(<nil>)
       })
      })
     }),
     If: (*compiler.If)(<nil>),
     Call: (*compiler.Call)(<nil>),
     Expression: (*compiler.Expression)(<nil>)
    })
   },
   End: (bool) true
  })
 }
})
([]string) <nil>
([]string) (len=1 cap=1) {
 (string) (len=1) "b"
}
([]string) (len=1 cap=1) {
 (string) (len=1) "a"
}
