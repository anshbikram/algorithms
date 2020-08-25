package com.github.anshbikram.algorithms;

import java.util.Arrays;

public class SpanningTree_Kruskal {

    private SpanningTree_Kruskal() {
    }

    public static void main(String[] args) {
        Graph graph = new Graph(4, 5);
        graph.addEgde(0, 1, 10);
        graph.addEgde(0, 2, 6);
        graph.addEgde(0, 3, 5);
        graph.addEgde(1, 3, 15);
        graph.addEgde(2, 3, 4);

        Edge[] mst = graph.runKruskalMst();
        int mstSum = 0;
        for (int e = 0; e < mst.length; e++) {
            System.out.printf("%s -> %s (%s)\n", mst[e].getSource(), mst[e].getDestination(), mst[e].getWeight());
            mstSum += mst[e].getWeight();
        }
        System.out.println("MST sum: " + mstSum);
    }

    public static class Graph {
        private int numOfVertices;
        private int numOfEdges;
        private Edge[] edges;
        private int edgeCount;

        public Graph(int numOfVertices, int numOfEdges) {
            this.numOfVertices = numOfVertices;
            this.numOfEdges = numOfEdges;
            this.edges = new Edge[numOfEdges];
        }

        public void addEgde(int src, int dest, int weight) {
            this.edges[edgeCount++] = new Edge(src, dest, weight);
        }

        public Edge[] runKruskalMst() {
            Arrays.sort(this.edges);
            Edge[] sortedEdges = this.edges;
            UnionSet[] unionSets = new UnionSet[numOfVertices];
            for (int v = 0; v < numOfVertices; v++) {
                unionSets[v] = new UnionSet(v, 1);
            }

            Edge[] selectedEdges = new Edge[numOfVertices - 1];

            int e = 0, v = 0;
            while (v < numOfVertices - 1) {
                Edge edge = sortedEdges[e++];
                int parentOfSrc = UnionSet.find(unionSets, edge.getSource());
                int parentOfDest = UnionSet.find(unionSets, edge.getDestination());
                if (parentOfSrc != parentOfDest) {
                    UnionSet.union(unionSets, edge.getSource(), edge.getDestination());
                    selectedEdges[v++] = edge;
                }

            }

            return selectedEdges;
        }
    }

    public static class UnionSet {
        private int parent;
        private int rank;

        public UnionSet(int parent, int rank) {
            this.parent = parent;
            this.rank = rank;
        }

        public int getParent() {
            return this.parent;
        }

        public int getRank() {
            return this.rank;
        }

        public void setParent(int parent) {
            this.parent = parent;
        }

        public static int find(UnionSet[] unionSets, int target) {
            if (unionSets[target].getParent() != target) {
                unionSets[target].setParent(find(unionSets, unionSets[target].getParent()));
            }

            return unionSets[target].getParent();
        }

        public static void union (UnionSet[] unionSets, int src, int dest) {
            int rootOfSrc = find(unionSets, src);
            int rootOfDest = find(unionSets, dest);

            if (unionSets[rootOfSrc].rank < unionSets[rootOfDest].rank) {
                unionSets[rootOfSrc].parent = unionSets[rootOfDest].parent;
            } else if (unionSets[rootOfSrc].rank > unionSets[rootOfDest].rank) {
                unionSets[rootOfDest].parent = unionSets[rootOfSrc].parent;
            } else {
                unionSets[rootOfDest].parent = unionSets[rootOfSrc].parent;
                unionSets[rootOfSrc].rank++;
            }
        }
    }

    public static class Edge implements Comparable<Edge> {
        private final int source;
        private final int destination;
        private final int weight;

        public Edge(int source, int destination, int weight) {
            this.source = source;
            this.destination = destination;
            this.weight = weight;
        }

        private int getSource() {
            return this.source;
        }

        private int getDestination() {
            return this.destination;
        }

        private int getWeight() {
            return this.weight;
        }

        @Override
        public int compareTo(Edge o) {
            return this.weight - o.weight;
        }
    }
}
