package listnode;

/**
 * @author zhangxiaohe@zhihu.com
 */
public class RemoveNthFromEnd {

    public static void main(String[] args) {
        ListNode listNode = ListNode.build();
        removeNthFromEnd(listNode,2);
    }

    public static ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode reulst=new ListNode(-1);
        reulst.next=head;
        ListNode fast = head, slow = reulst;
        for (int i = 0; i < n; i++) {
            fast = fast.next;
        }
        while (fast != null) {
            slow = slow.next;
            fast = fast.next;
        }
        slow.next=slow.next.next;
        return reulst.next;

    }
}
