# Cracking the Code Inteview - Section VI - Big O

This is just some notes and examples from the book to make me more familiar with the problems involving Big O

### Example 1:

What's the runtime of the below code?

```java
void foo(int[] array) {
  int sum = 0;
  int product = 1;
  for (int i = 0; i < array.length; i++) {
    sum += array[i];
  }

  for (int i = 0; i < array.length; i++) {
    product *= array[i];
  }

  System.out.println(sum + ", " + product);
}
```
This code will take O(N) time. Doesn't matter if there are two iterations through the array.


### Example 2:

What's the runtime of code below?
```java
void printPairs(int[] array) {
  for (int i = 0; i < array.length; i++) {
    for (int j = 0; j < array.length; j++) {
      System.out.println(array[i] + ", " + array[j]);
    }
  }
}
```
The inner loop has O(N) iterations and it is called N times. Therefore, the runtime is O(N^2).

Another way we can see this is by inspecting what the "meaning" of the code is. It is printing all pairs(two-element sequences). There are O(N^2) pairs; therefore, the runtime is O(N^2)


### Example 3:

This is very similar code to the above example, but now the inner for loop starts at i + 1
```java
void printUnorderedPairs(int[] array) {
  for (int i = 0; i < array.length; i++) {
    for (int j = i + 1; j < array.length; j++) {
      System.out.println(array[i] + ", " + array[j]);
    }
  }
}
```
Although the inner loop does N/2 work on average it runs in N times so the total work is N^2/2 which is O(N^2).

### Example 4:
This is similar to the above, but now we have two different arrays.
```java
void  printUnorderedPairs(int[] arrayA, int[] arrayB) {
  for (int i = 0; i < arrayA.length; i++) {
    for (int j = 0; j < arrayB.length; j++) {
      if (arrayA[i] < arrayB[j]) {
        System.out.println(arrayA[i] + ", " + arrayB[j]);
      }
    }
  }
}
```
If we analize the code, we see that the if-statement within j's for loop is O(1) time since it's just a sequence of constant-time statements.

So for each element of arrayA the inner for loop goes through b iterations, where b = arrayB.length.
If a = arrayA.length, then the runtime is O(ab). Don't confuse with O(N^2) because there are two loops. There are two loops but with two different inputs. Both matter. This is an extremely common mistake.


### Example 5:
What about this strange bit of code?
```java
void  printUnorderedPairs(int[] arrayA, int[] arrayB) {
  for (int i = 0; i < arrayA.length; i++) {
    for (int j = 0; j < arrayB.length; j++) {      
      for (int k = 0; k < 1000000; k++>) {
        System.out.println(arrayA[i] + ", " + arrayB[j]);
      }
    }
  }
}
```
Notice that nothing has really changed here. 100,000 units of work is still constant, so the runtime is O(ab)