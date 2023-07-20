# ISBN Verifier

**Exercise:**
Create a Golang program which checks whether the number is a ISBN or not
Here are the requirements that are needed:

* The ISBN-10 verification process is used to validate book identification numbers. These normally contain dashes and look like: ``3-598-21508-8``.
* The ISBN-10 format is 9 digits (0 to 9) plus one check character (either a digit or an X only). In the case the check character is an X, this represents the value '10'.
* The program should also able to avoid empty inputs.
* Should able to avoid letters other than 'X'.
* The number of digits in ISBN should be 10.


**Example:**

Let's take the ISBN-10 ``3-598-21508-8``. We plug it in to the formula, and get:

*(3 * 10 + 5 * 9 + 9 * 8 + 8 * 7 + 2 * 6 + 1 * 5 + 5 * 4 + 0 * 3 + 8 * 2 + 8 * 1) mod 11 == 0*

Since the result is 0, this proves that our ISBN is valid.

**Test Cases:**
```
3-598-21507-A
3-598-2X507-9
00
3598P215088
3598215088
359821507X
98245726788
```