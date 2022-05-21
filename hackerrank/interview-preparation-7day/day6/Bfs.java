package day6;

// public class Bfs {
//     public static void main(String[] args) {
//         int EDGE_DISTANCE = 6;

//         // fill all distances with -1 (unreachable)
//         int[] distances = new int[n];
//         Arrays.fill(distances, -1);

//         Queue<Integer> nodes = new LinkedList<>();

//         // add start node into queue, and its distance (to itself) is 0;
//         nodes.add(s);
//         distances[s - 1] = 0; // index of s in the distance is `s-1`.

//         while (!nodes.isEmpty()) {
//             int node = nodes.poll();
//             /*
//              * Loop all edges.
//              * Edges means two nodes have connection. so the connected node's distance is my
//              * (node) distance + EDGE_DISTANCE.
//              *
//              * Tricky point:
//              * Edge just represents node 1 has connection with node 2, not the sequence.
//              * So FROM could be edge[0] or edge[1].
//              * We need to take care of both ends.
//              */
//             for (int[] edge : edges) {
//                 if (edge[0] == node && distances[edge[1] - 1] == -1) {
//                     distances[edge[1] - 1] = distances[node - 1] + EDGE_DISTANCE;
//                     nodes.add(edge[1]);
//                 } else if (edge[1] == node && distances[edge[0] - 1] == -1) {
//                     distances[edge[0] - 1] = distances[node - 1] + EDGE_DISTANCE;
//                     nodes.add(edge[0]);
//                 }
//             }
//         }

//         // Exclude start node from distances.
//         int[] result = new int[n - 1];
//         int j = 0;
//         for (int i = 0; i < distances.length; i++) {
//             if (s - 1 == i) {
//                 continue;
//             }
//             result[j] = distances[i];
//             j++;
//         }

//         return result;
//     }
// }