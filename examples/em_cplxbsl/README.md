# How It Works
In this example, we solve the following SDE with the following approximation method.

## SDE: Complex-valued Bessel Process
$$
\begin{align*}
\xi(t,z) &= \Re(z) + \int_0^t\frac{\xi(s,z)}{\xi(s,z)^2+\eta(s,z)^2}\,ds+B(t),\\
\eta(t,z) &= \Im(z) - \int_0^t\frac{\eta(s,z)}{\xi(s,z)^2+\eta(s,z)^2}\,ds,
\end{align*}
$$
where $z\in\mathbb{C}$.

## Method: Euler-Maruyama
Letting $\Delta t>0$ and $n\in\mathbb{Z}_{>0}$ be specified, define $\widehat{\xi}(0,z):=\Re(z)$, $\widehat{\eta}(0,z):=\Im(z)$ and
$$
\begin{align*}
\widehat{\xi}(t_{k},z)&:=\widehat{\xi}(t_{k-1},z)+\frac{\widehat{\xi}(t_{k-1},z)}{\widehat{\xi}(t_{k-1},z)^{2}+\widehat{\eta}(t_{k-1},z)^{2}}\,\Delta t + \sqrt{\Delta t}\,Z_{k},\\
\widehat{\eta}(t_{k},z)&:=\widehat{\eta}(t_{k-1},z)-\frac{\widehat{\eta}(t_{k-1},z)}{\widehat{\xi}(t_{k-1},z)^{2}+\widehat{\eta}(t_{k-1},z)^{2}}\,\Delta t,
\end{align*}
$$
for $k=1,2,\cdots,n$. Here, $t_{k}=\frac{k}{n}t$ and $Z_{k}$, $k=1,\cdots,n$ are i.i.d. standard normal random variables.
