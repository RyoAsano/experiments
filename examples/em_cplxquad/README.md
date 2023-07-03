# How It Works
In this example, we solve the following SDE with the following approximation method.

## SDE with Injection
Letting $x=\Re(z)$ and $y=\Im(z)$, we solve
$$
    Z(t, z) := \left(X(t, x),\, y\,e^{Y(t, x)}\right),
$$
where
$$
\begin{align*}
X(t, x) &= x + \int_0^t\left(X(s, x)^{2}-Y(s, x)^{2}\right)\,dB(s),\\
Y(t, x) &= \int_0^t2X(s, x)Y(s, x)\,dB(s).
\end{align*}
$$

## Method: Euler-Maruyama
Letting $\Delta t>0$ and $n\in\mathbb{Z}_{>0}$ be specified, define $\widehat{X}(0,x):=x$, $\widehat{Y}(0,x):=0$ and
$$
\begin{align*}
\widehat{X}(t_{k},x)&:=\widehat{X}(t_{k-1},x)+\left(\widehat{X}(t_{k-1}, x)^{2}-\widehat{Y}(t_{k-1},x)^{2}\right)\,\sqrt{\Delta t}\,Z_{k},\\
\widehat{Y}(t_{k},x)&:=\widehat{Y}(t_{k-1},x)+2\widehat{X}(t_{k-1},x)\widehat{Y}(t_{k-1},x)\,\sqrt{\Delta t}\,Z_{k},
\end{align*}
$$
for $k=1,2,\cdots,n$. Here, $t_{k}=\frac{k}{n}t$ and $Z_{k}$, $k=1,\cdots,n$ are i.i.d. standard normal random variables.
