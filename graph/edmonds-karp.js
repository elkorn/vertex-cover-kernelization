
function ek(E, C, s, t) {
    var n = c.length;
    console.log("n", n);
    var F = [];
    for (var i = 0; i < n; i++) {
        F[i] = [];
        for (var j = 0; j < n; j++) {
            F[i][j] = 0;
        }
    }

    console.log("F", F);

    while (true) {
        var P = [];
        for (var i = 0; i < n; i++) {
            P[i] = -1;
        }

        P[s] = s;
        var M = [];
        M[s] = 100000;
        var Q = [];
        Q.unshift(s); // queue
        var shouldContinue = true;
        while (Q.length && shouldContinue) {
            var u = Q.pop();
            var broken = false;
            E[u].forEach(function(v) {
                if (broken) return;
                if (C[u][v] - F[u][v] > 0 && P[v] == -1) {
                    P[v] = u;
                    M[v] = Math.min(M[u], C[u][v] - F[u][v]);
                    if (v !== t) {
                        Q.unshift(v);
                    } else {
                        while (P[v] !== u) {
                            u = P[v];
                            F[u][v] += M[t];
                            F[v][u] -= M[t];
                            v = u;
                        }

                        shouldContinue = false;
                        broken = true;
                    }
                }
            });
        }

        if (P[t] == -1) {
            var sum = 0;
            F[s].forEach(function(x) {
                sum += x;
            });

            return sum;
        }
    }
}

var e = [
    [1],
    [2],
    []
]; // adjacency list
var c = [
    [1, 1, 1],
    [1, 1, 1],
    [1, 1, 1]
]; // capacities

console.log(ek(e, c, 0, 2));
console.log("c", c);
