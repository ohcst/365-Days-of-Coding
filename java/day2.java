import java.util.ArrayList;
import java.util.Arrays;

public class day2 {
    public static void main(String[] arg) {
        int num = 11;
        collatz(num);
    }

    public static void collatz(int num) {
        ArrayList<Integer> arr = new ArrayList<Integer>();
        arr.add(num);

        while (num != 1) {
            if (num % 2 == 0) {
                num = (int) Math.floor(num / 2);
                
            } else {
                num = (num * 3) + 1;
            }
            arr.add(num);
        }

        System.out.println(arr);
    }
}