class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverse_linked_list(self, head: ListNode) -> ListNode:
        prev = None
        current = head
        while current:
            nxt = current.next
            current.next = prev
            prev = current
            current = nxt
        return prev


if __name__ == "__main__":
    head = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
    solution = Solution()
    reversed_head = solution.reverse_linked_list(head)

    result = []
    while reversed_head:
        result.append(reversed_head.val)
        reversed_head = reversed_head.next
    print(result)
