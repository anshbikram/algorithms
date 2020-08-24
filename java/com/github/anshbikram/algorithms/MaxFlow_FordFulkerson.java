package com.github.anshbikram.algorithms;

import java.util.LinkedList;

public class MaxFlow_FordFulkerson {

    public static void main(String[] args) {
        int[][] adjacencyMatrix = new int[][] {
            {0, 16, 13, 0, 0, 0}, 
            {0, 0, 10, 12, 0, 0}, 
            {0, 4, 0, 0, 14, 0}, 
            {0, 0, 9, 0, 0, 20}, 
            {0, 0, 0, 7, 0, 4}, 
            {0, 0, 0, 0, 0, 0}
        };

        System.out.println("max flow is " + fordFulkersonMaxFlow(adjacencyMatrix, 0, 5));
    }

    private static int fordFulkersonMaxFlow(int[][] graph, int source, int sink) {
        int[][] residualGraph = new int[graph.length][graph.length];
        for (int i = 0; i < graph.length; i++) {
            for (int j = 0; j < graph.length; j++) {
                residualGraph[i][j] = graph[i][j];
            }
        }

        int maxFlow = 0;

        int[] parentPath = new int[residualGraph.length];
        while (bfs(residualGraph, source, sink, parentPath)) {

            int minFlowInPath = Integer.MAX_VALUE;
            for (int vertex = sink; vertex != source; vertex = parentPath[vertex]) {
                minFlowInPath = Math.min(minFlowInPath, residualGraph[parentPath[vertex]][vertex]);
            }

            for (int vertex = sink; vertex != source; vertex = parentPath[vertex]) {
                residualGraph[parentPath[vertex]][vertex] -= minFlowInPath;
                residualGraph[vertex][parentPath[vertex]] += minFlowInPath;
            }

            maxFlow += minFlowInPath;
        }

        return maxFlow;
    }

    private static boolean bfs(int[][] residualGraph, int source, int sink, int[] parentPath) {
        boolean[] visited = new boolean[residualGraph.length];
        LinkedList<Integer> queue = new LinkedList<>();
        queue.add(source);

        while (!queue.isEmpty()) {
            int srcVertex = queue.poll();
            
            for (int targetVertex = 0; targetVertex < residualGraph.length; targetVertex++) {
                if (!visited[targetVertex] && residualGraph[srcVertex][targetVertex] > 0) {
                    visited[targetVertex] = true;
                    parentPath[targetVertex] = srcVertex;
                    queue.add(targetVertex);
                }
            }
        }

        return visited[sink];
    }
}
