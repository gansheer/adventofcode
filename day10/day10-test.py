#! /usr/bin/env python3
# coding: utf-8
from collections import deque

def build_graph(adapters, device_joltage):
  inversed_ordered_adapter =  (list(reversed(sorted(adapters)))) 
  print(inversed_ordered_adapter)
  graph = [[] for i in range(device_joltage+1)] 
  print(len(graph))
  for adapter in inversed_ordered_adapter:
    graph[adapter] = []
    if (adapter - 1) in inversed_ordered_adapter:
      graph[adapter].append(adapter-1)
    if (adapter - 2) in inversed_ordered_adapter:
      graph[adapter].append(adapter-2)
    if (adapter - 3) in inversed_ordered_adapter:
      graph[adapter].append(adapter-3)
  print(graph)
  return graph

def bfs(G, s):
    shortest = [float('+Inf')]*len(G)
    count = [0]*len(G)

    shortest[s] = 0
    count[s] = 1

    Q = deque([s])

    while Q:
        u = Q.popleft()
        for v in G[u]:
            if not count[v]: 
                Q.append(v)

            if shortest[u]+1 <= shortest[v]:
                shortest[v] = shortest[u]+1
                count[v] += count[u]
    return count


def main():
    input_file = open('input_test', 'r') 
    lines = input_file.readlines()
    clean_lines = map(lambda x: int(x.rstrip("\n")),lines)
    print(clean_lines)


    device_joltage = max(clean_lines) + 3

    graph = build_graph(clean_lines, device_joltage)
    print(graph)

    G = [
        [1, 2, 3],
        [4],
        [4],
        [4],
        []
    ]
    print(bfs(G, 0))
    print(bfs(graph, 0))
    print(sum(bfs(graph, 19)))

if __name__ == '__main__':
    main()