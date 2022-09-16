final class DirectedAcyclicGraph {
	
	Node root;
	DirectedAcyclicGraph() {
	}
	class Node {
		int cost;
		Node[] children;
		Node parent;
		
		Node(int cost) {
			this.cost = cost;	
		}

		void addChild(Node child) {
			Node[] newArray = new Node[this.children.length + 1];
			for (int i = 0; i < this.children.length; i++) {
				newArray[i] = this.children[i];
			}
			newArray[newArray.length - 1] = child;
			this.children = newArray;
			child.parent = this;
		}
		int getCheapestCost(Node rootNode) {
                	int n = this.children.length;

                	if (n == 0) {
                        	return rootNode.cost;
                	} else {
                        	int minCost = Integer.MAX_VALUE;
                        	for (int i = 0; i < n - 1; i++) {
                                	int tempCost = getCheapestCost(rootNode.children[i]);
                                	if (tempCost < minCost) {
                                        	minCost = tempCost;
                               		}
                        	}
                		return minCost + rootNode.cost;
			}
		}
	}

	public static void main(String[] args) {
		DirectedAcyclicGraph dg = new DirectedAcyclicGraph();
		Node root = dg.new Node(0);
		root.addChild(dg.new Node(5));
		root.addChild(dg.new Node(3));
		root.addChild(dg.new Node(6));
		int result = root.getCheapestCost(root);
		System.out.println(result);
	}

}
