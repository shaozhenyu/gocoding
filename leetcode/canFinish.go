func canFinish(numCourses int, prerequisites [][]int) bool {
    adj := Digraph(numCourses, prerequisites)
    marked := make([]bool, numCourses)
    onStack := make([]bool, numCourses)
    hasCycle := false
    for i := 0; i < numCourses; i++ {
        if !marked[i] {
            dfs(i, adj, marked, onStack, &hasCycle)
        }
        if hasCycle {
            return !hasCycle
        }
    }
    return !hasCycle
}

func dfs(v int, adj [][]int, marked, onStack []bool, hasCycle *bool) {
    if len(adj[v]) == 0 {
        return
    }
    marked[v] = true
    onStack[v] = true
    for _, w := range adj[v] {
        if *hasCycle {
            return
        }
        if !marked[w] {
            dfs(w, adj, marked, onStack, hasCycle)
        }
        if onStack[w] {
            *hasCycle = true
            return
        }
    }
    onStack[v] = false
}

func Digraph(n int, G [][]int) [][]int {
    adj := make([][]int, n)
    for i := 0; i < n; i++ {
        adj[i] = make([]int, 0)
    }
    for _, g := range G {
        adj[g[1]] = append(adj[g[1]], g[0])
    }
    return adj
}
