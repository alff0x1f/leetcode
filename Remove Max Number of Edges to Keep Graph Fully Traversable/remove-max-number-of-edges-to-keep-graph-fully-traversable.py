from typing import List


class EdgeGroups:
    def __init__(self, parent=None):
        if parent:
            self.parent = parent
        else:
            self.parent = {}

    def find(self, node):
        if node not in self.parent:
            return node
        self.parent[node] = self.find(self.parent[node])
        return self.parent[node]

    def connect(self, node1, node2) -> bool:
        head1 = self.find(node1)
        head2 = self.find(node2)
        if head1 == head2:
            return False
        self.parent[head1] = head2
        return True

    def is_all_connected(self, n: int) -> bool:
        heads_count = 0
        for i in range(1, n + 1):
            if i == self.parent.get(i):
                heads_count += 1
                if heads_count > 1:
                    return False
        return True


class Solution:
    def maxNumEdgesToRemove(self, n: int, edges: List[List[int]]) -> int:
        removed_count = 0
        groups_both = EdgeGroups()
        for edge in edges:
            if edge[0] == 3:  # both
                if not groups_both.connect(edge[1], edge[2]):
                    removed_count += 1

        groups_alice = groups_both
        groups_bob = EdgeGroups(groups_both.parent.copy())

        for edge in edges:
            if edge[0] == 1:  # alice
                if not groups_alice.connect(edge[1], edge[2]):
                    removed_count += 1
            elif edge[0] == 2:  # bob
                if not groups_bob.connect(edge[1], edge[2]):
                    removed_count += 1

        if not groups_alice.is_all_connected(n):
            return -1
        if not groups_bob.is_all_connected(n):
            return -1

        return removed_count

