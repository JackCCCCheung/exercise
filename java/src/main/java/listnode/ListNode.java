package listnode;

/**
 * @author JackCheung
 */
public class ListNode {

    protected int val;

    protected ListNode next;

    public ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }

    public ListNode(int val) {
        this.val = val;
    }

    public ListNode() {
    }

    public static ListNode build() {
        ListNode result = new ListNode(1);
        result.next = new ListNode(2);
        result.next.next = new ListNode(3);
        result.next.next.next = new ListNode(4);
        result.next.next.next.next = new ListNode(5);
        return result;
    }
}
