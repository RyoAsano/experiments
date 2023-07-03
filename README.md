# Stochastic Calculus in Go
This library is intended to provide miscellaneous tools for stochastic calculus. Currently, this library supports to:

- define a custom *stochastic differential equation (SDE)*
- define a custom *strong approximation method*
- *solve* a given SDE with a given method

# Getting Started
Maybe the most trivial way to get yourself accustomed with this library is to check the [examples](examples) folder, pick up whatever subfolder you want and build it to execute `main.go`. 

Currently, there are 3 examples available:
- [Black-Scholes SDE & Euler-Maruyama method](examples/em_bs/)
- [Complex-Valued Bessel Process & Euler-Maruyama method](examples/em_cplxbsl/)
- [Complex-Valued Custom SDE & Euler-Maruyama method](examples/em_cplxquad/)

Each of them has its own expanatory text in README.md, so if you are not sure where to start, refer yourself to whichever of them you like.

Executing the binary yields a sample path of the SDE in the `db` directory (which is created only after the execution has been done for the first time).