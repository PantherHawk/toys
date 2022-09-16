import java.lang.IllegalArgumentException;
import java.lang.IndexOutOfBoundsException;
import java.util.Set;
import java.util.HashSet;

final class UndirectedGraph {
	private final boolean[][] adjacencyMatrix;

	UndirectedGraph(int numberOfNodes) {
		if (numberOfNodes < 0) {
			throw new IllegalArgumentException();
		}
		adjacencyMatrix = new boolean[numberOfNodes][numberOfNodes];
	}

	int numberOfNodes() {
		return adjacencyMatrix.length;
	}

	private void validateIndex(int index) {
		if (index < 0 || index >= numberOfNodes()) {
			throw new IndexOutOfBoundsException();
		}
	}

	boolean areAdjacent(int firstNode, int secondNode) {
		validateIndex(firstNode);
		validateIndex(secondNode);

		return adjacencyMatrix[firstNode][secondNode];
	}

	void setAdjacent(int firstNode, int secondNode, boolean areAdjacent) {
		validateIndex(firstNode);
		validateIndex(secondNode);

		adjacencyMatrix[firstNode][secondNode] = areAdjacent;
		adjacencyMatrix[secondNode][firstNode] = areAdjacent;
	}

	int findDegreesOfSeparation(int firstNode, int secondNode) {
		Set<Integer> visited = new HashSet<>();
		int degrees = 0;
		return findDegreesOfSeparation(firstNode, secondNode, visited, degrees);
	}

	int findDegreesOfSeparation(int firstNode, int secondNode, Set<Integer> visited, int degrees) {
		int ret;
		int min = Integer.MAX_VALUE;

		visited.add(firstNode);

		if (firstNode == secondNode || adjacencyMatrix[firstNode][secondNode]) {
			return degrees;
		}

		degrees++;
		for (int i = 0; i < adjacencyMatrix[firstNode].length; i++) {
			if (!visited.contains(i)) {
				ret = findDegreesOfSeparation(i, secondNode, visited, degrees);
				min = Math.min(min, ret);
			}
		}
		return min;
	}
	public static void main(String[] args) {
		UndirectedGraph friends = new UndirectedGraph(8);
		friends.setAdjacent(0, 1, true);
		friends.setAdjacent(1, 2, true);
		System.out.println(friends.findDegreesOfSeparation(0, 2));
	}
}
