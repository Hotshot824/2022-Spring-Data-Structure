from collections import defaultdict
 
class Graph:

    def __init__(self):
        self.graph = defaultdict(list)
 
    def addEdge(self,u,v):
        self.graph[u].append(v)

    def BFS(self, s):
        visited = [False] * (max(self.graph) + 1)
        queue = []
        queue.append(s)
        visited[s] = True
        while queue:
            s = queue.pop(0)
            print(s)
 
            for i in self.graph[s]:
                if visited[i] == False:
                    queue.append(i)
                    visited[i] = True

if __name__ == "__main__":
    g = Graph()
    vlist = [
        [1, 2], [1, 3], [2, 1], [2, 4], [2, 5], 
        [3, 1], [3, 6], [3, 7], [4, 2], [4, 8],
        [5, 2], [5, 8], [6, 3], [6, 9], [7, 3], 
        [7, 9], [8, 4], [8, 5], [8, 10], [9, 6], 
        [9, 7], [9, 10], [10, 8], [10, 9]]

    for v in vlist:
        g.addEdge(v[0], v[1])

    g.BFS(1)
    
    