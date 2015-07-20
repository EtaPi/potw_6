This is a continuation of last week's Problem. Either continue with the new language you just picked up, or start with another language unknown to you, and implement a command line calendar. It should take as input a year and a width, and output a table of the months and days in that year, with the number of columns equal to the input width.

For example, if the input was 2015 and 3, the output should look like this (Hopefully Outlook won't screw up the fixed width font):

```
                                  2015

         January                February                March         
   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa              
                1  2  3    1  2  3  4  5  6  7    1  2  3  4  5  6  7  
    4  5  6  7  8  9 10    8  9 10 11 12 13 14    8  9 10 11 12 13 14  
   11 12 13 14 15 16 17   15 16 17 18 19 20 21   15 16 17 18 19 20 21  
   18 19 20 21 22 23 24   22 23 24 25 26 27 28   22 23 24 25 26 27 28  
   25 26 27 28 29 30 31                          29 30 31              

           April                  May                    June
   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa   Su Mo Tu We Th Fr Sa
             1  2  3  4                   1  2       1  2  3  4  5  6  
    5  6  7  8  9 10 11    3  4  5  6  7  8  9    7  8  9 10 11 12 13
   12 13 14 15 16 17 18   10 11 12 13 14 15 16   14 15 16 17 18 19 20
   19 20 21 22 23 24 25   17 18 19 20 21 22 23   21 22 23 24 25 26 27
   26 27 28 29 30         24 25 26 27 28 29 30   28 29 30
                          31
```

... and so on for the rest of the year. Hopefully after coding this one up you'll hate date time manipulation and realize its not so simple.
