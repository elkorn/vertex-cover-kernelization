// VC(G, T , k)
// 	Input: a graph G, a set T of tuples, and a positive integer k.
// 	Output: the size of a minimum vertex cover of G if the size is bounded by k;
// 	report failure otherwise.
// 	0. if |G| > 0 and k = 0 then reject;
// 	1. apply Reducing;
// 	2. pick a structure Γ of highest priority;
// 	3. if (Γ is a 2-tuple ({u, z}, 1)) or (Γ is a good pair (u, z) where z is
// 		almost-dominated by a vertex v ∈ N (u)) or (Γ is a vertex z with d(z) ≥ 7)
// 	then return
// 		min{1+VC(G − z, T ∪ (N (z), 2), k − 1), d(z)+ VC(G − N [z], T , k − d(z))};
// 	else /* Γ is a good pair (u, z) where z is not almost-dominated by by any
// 		vertex in N (u) */
// 		return
// 		min{1+VC(G − z, T , k − 1), d(z)+ VC(G − N [z], T ∪ (N (u), 2), k − d(z))};
//
// Reducing
// 	a. for each tuple (S, q) ∈ T do
// 		a.1. if |S| < q then reject;
// 		a.2. for every vertex u ∈ S do T = T ∪ {(S − {u}, q − 1)};
// 		a.3. if S is not an S independent set then
// 			T = T ∪ ( (u,v)∈E,u,v∈S {(S − {u, v}, q − 1)});
// 		a.4. if there exists v ∈ G such that |N (v) ∩ S| ≥ |S| − q + 1 then
// 			return (1+VC(G − v, T , k − 1)); exit;
// 	b. if Conditional General Fold(G) or Conditional Struction(G) in the
// 		given order is applicable then apply it; exit;
// 	c. if there are vertices u and v in G such that v dominates u then
// 		return (1+ VC(G − v, T , k − 1)); exit;
//
// Conditional General Fold
// 	if there exists a strong 2-tuple ({u, z}, 1) in T then
// 	if the repeated application of General Fold reduces the parameter by at
// 		least 2 then apply it repeatedly;
// 	else if the application of General-Fold reduces the parameter by 1 and
// 		(d(u) < 4)
// 	then apply it until it is no longer applicable;
// 	else apply General-Fold until it is no longer applicable;
//
// Conditional Struction
// 	if there exists a strong 2-tuple {u, v} in T then
// 	if there exists w ∈ {u, v} such that d(w) = 3 and the Struction is
// 		applicable to w then apply it;
// 	else if there exists a vertex u ∈ G where d(u) = 3 or d(u) = 4 and such that
// 		the Struction is applicable to u then apply it;
