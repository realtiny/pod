// Copyright 2021 shuaimingzhou(at)gmail(dot)com.
// All rights reserved. License that can be found in the LICENSE file.

/*
Package pod is tiny web framwork. The idea mainly come from
Expressjs and Negroni. It is much more smaller and simpler also
very easier to use and extend. Any Struct/Type can be made to
a pod.Handler by adding a function ServeHTTP(w, r, next).
Then you can push the implement of pop.Handler to the statck.

Example:

  m1 := ...
  m2 := ...
  m3 := ...
  app := pod.New()
  app.Push(m1)
  app.Push(m2)
  app.Push(m3)
  app.Run(":8888")

pod.handle is the entry of pod.Handler stack.
The execution order of pod.Handlers is:

  m1      // Before next(w, r)
    m2    // Before next(w, r)
      m3  // ServeHTTP(w, r, nil)
    m2    // After next(w, r)
  m1      // After next(w, r)

Any functionality should be made to a reusable Handler such as
Staics Serving, Router, Sessions Manager, Logging, Database Accessing,
Authentication/Authorization and so on.
*/
package pod
