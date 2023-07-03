# How It Works
In this example, we solve the following SDE with the following approximation method.

## SDE: Black-Scholes
```math
X(t,x) = x + \int_0^t\mu X(s,x)\,ds+\int_0^t\sigma X(s,x)\,dB(s),
```
where $`x, \mu, \sigma\in\mathbb{R}`$.

## Method: Euler-Maruyama
Letting $\Delta t>0$ and $`n\in\mathbb{Z}_{>0}`$ be specified, define $`\widehat{X}(0,x):=x`$ and
```math
\widehat{X}(t_{k}, x) := \widehat{X}(t_{k-1}, x)
+ \mu \widehat{X}(t_{k-1}, x)\,\Delta t + \sigma \widehat{X}(t_{k-1}, x)\, \sqrt{\Delta t}\,Z_{k},
```
for $`k=1,2,\cdots,n`$. Here, $`t_{k}=\frac{k}{n}t`$ and $`Z_{k}$, $k=1,\cdots,n`$ are i.i.d. standard normal random variables.
