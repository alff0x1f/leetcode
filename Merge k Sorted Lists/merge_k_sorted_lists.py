from typing import List, Optional, Tuple


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    @staticmethod
    def extract_min(lists: list) -> Tuple[int, list]:
        min_value_pos = 0
        for i, node in enumerate(lists):
            if node.val < lists[min_value_pos].val:
                min_value_pos = i
        if lists[min_value_pos].next is None:
            lists[-1], lists[min_value_pos] = lists[min_value_pos], lists[-1]
            return lists[-1].val, lists[:-1]
        min_value = lists[min_value_pos].val
        lists[min_value_pos] = lists[min_value_pos].next
        return min_value, lists

    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        lists = [node for node in lists if node]
        if not lists:
            return None
        min_value, lists = self.extract_min(lists)
        first_node = ListNode(min_value, None)
        current_node = first_node
        while lists:
            min_value, lists = self.extract_min(lists)
            current_node.next = ListNode(min_value, None)
            current_node = current_node.next
        return first_node
