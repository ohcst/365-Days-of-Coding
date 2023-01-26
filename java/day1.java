import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class day1 {
    public static void main(String[] arg) {
        ArrayList<Integer> arr1 = new ArrayList<Integer>(Arrays.asList(3, 2, 3));
        ArrayList<Integer> arr2 = new ArrayList<Integer>(Arrays.asList(2, 2, 1, 1, 1, 2, 2));
        majority(arr1);
        majority(arr2);
    }

    public static void majority(ArrayList<Integer> arr) {
        Map<Integer, Integer> count = new HashMap<Integer, Integer>();
        System.out.printf("Given array: %s%n", arr);
        int goal = (int) Math.floor(arr.size()/2)+1;
        System.out.printf("Goal: %s appearances", goal);

        for (int i : arr) {
            if (count.containsKey(i)) {
                count.put(i, count.get(i) + 1);
            } else {
                count.put(i, 1);
            }
        }

        count.forEach(
            (key, value)
            -> {if (value >= goal) {
                System.out.printf("%n%s is the majority element, appearing %s times%n", key, value);
            }}
        );  

    }
}