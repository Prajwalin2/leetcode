package main

import "fmt"

func main() {
	testcases := [][]int{
		{1, 1, 3},
		{2, 1, 3},
		{3, 4, 5},
		{3, 3, 5},
		{4, 1, 2},
		{10, 1, 100},
		{100, 1, 1000000000},
	}
	for _, t := range testcases {
		fmt.Println(zigZagArrays(t[0], t[1], t[2]))
	}
}

const MOD = 1_000_000_007

var fact, invFact []int64

func modPow(a, b int64) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return res
}

func precomputeFact(n int) {
	fact = make([]int64, n+2)
	invFact = make([]int64, n+2)
	fact[0] = 1
	for i := 1; i <= n+1; i++ {
		fact[i] = fact[i-1] * int64(i) % MOD
	}
	invFact[n+1] = modPow(fact[n+1], MOD-2)
	for i := n; i >= 0; i-- {
		invFact[i] = invFact[i+1] * int64(i+1) % MOD
	}
}

func lagrange(y []int64, t int64) int64 {
	d := len(y) - 1
	if t >= 0 && t <= int64(d) {
		return y[t]
	}
	pref := make([]int64, d+1)
	suff := make([]int64, d+1)
	pref[0] = 1
	for i := 0; i <= d; i++ {
		pref[i] = (t - int64(i) + MOD) % MOD
		if i > 0 {
			pref[i] = pref[i] * pref[i-1] % MOD
		}
	}
	suff[d] = 1
	for i := d; i >= 0; i-- {
		suff[i] = (t - int64(i) + MOD) % MOD
		if i < d {
			suff[i] = suff[i] * suff[i+1] % MOD
		}
	}
	res := int64(0)
	for i := 0; i <= d; i++ {
		num := int64(1)
		if i > 0 {
			num = num * pref[i-1] % MOD
		}
		if i < d {
			num = num * suff[i+1] % MOD
		}
		den := invFact[i] * invFact[d-i] % MOD
		term := y[i] * num % MOD * den % MOD
		if (d-i)&1 == 1 {
			term = (MOD - term) % MOD
		}
		res = (res + term) % MOD
	}
	return res
}

func polyPrefixSum(y []int64, mVal int64) ([]int64, int64) {
	d := len(y) - 1
	pref := make([]int64, d+2)
	pref[0] = y[0]
	for t := 1; t <= d; t++ {
		pref[t] = (pref[t-1] + y[t]) % MOD
	}
	yNext := lagrange(y, int64(d+1))
	pref[d+1] = (pref[d] + yNext) % MOD
	total := int64(0)
	if mVal-1 <= int64(d+1) {
		if mVal-1 >= 0 {
			total = pref[mVal-1]
		}
	} else {
		total = lagrange(pref, mVal-1)
	}
	return pref, total
}

func zigZagArrays(n int, l int, r int) int {
	if n == 1 {
		return (r - l + 1) % MOD
	}
	m := int64(r - l + 1)
	precomputeFact(n)

	P := []int64{0, 1}
	Q := []int64{(m - 1 + MOD) % MOD, (m - 2 + MOD) % MOD}

	if n == 2 {
		return int(m * (m - 1) % MOD)
	}

	for i := 1; i < n-1; i++ {
		prefP, totalP := polyPrefixSum(P, m)

		pNext := lagrange(P, int64(i+1))

		newP := make([]int64, i+2)
		newP[0] = 0
		for t := 1; t <= i+1; t++ {
			newP[t] = (newP[t-1] + Q[t-1]) % MOD
		}

		newQ := make([]int64, i+2)
		for t := 0; t <= i; t++ {
			newQ[t] = (totalP - prefP[t] + MOD) % MOD
		}
		newQ[i+1] = (newQ[i] - pNext + MOD) % MOD

		P = newP
		Q = newQ
	}

	_, totalP := polyPrefixSum(P, m)
	_, totalQ := polyPrefixSum(Q, m)

	ans := (totalP + totalQ) % MOD
	return int(ans)
}
