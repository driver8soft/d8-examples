       IDENTIFICATION DIVISION.
       PROGRAM-ID. test3.

       ENVIRONMENT DIVISION.

       CONFIGURATION SECTION.
       
       DATA DIVISION.

       FILE SECTION.

       WORKING-STORAGE SECTION.
      * Declare variables in the WORKING-STORAGE section
       01 PROG-NAME PIC X(8) VALUE "loop".
       01 COMMLEN PIC 9(9) COMP.
       01 COMMAREA.
           05 N-TIMES PIC S9(9) COMP.

       PROCEDURE DIVISION.
      * code goes here!
           INITIALIZE COMMAREA.

           DISPLAY "test3 example program"
           DISPLAY "Number of loops: " WITH NO ADVANCING.
           ACCEPT N-TIMES.

           COMPUTE COMMLEN = LENGTH OF COMMAREA.
           CALL "D8link" USING PROG-NAME COMMAREA COMMLEN.

           GOBACK.


