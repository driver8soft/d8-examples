       IDENTIFICATION DIVISION.
       PROGRAM-ID. test1.

       ENVIRONMENT DIVISION.

       CONFIGURATION SECTION.
       
       DATA DIVISION.

       FILE SECTION.

       WORKING-STORAGE SECTION.
      * In COBOL, you declare variables in the WORKING-STORAGE section
       01 PROG-NAME PIC X(8) VALUE "hello".
       01 COMMLEN PIC 9(9) COMP.
       01 COMMAREA.
           05 INPUT-NAME PIC X(20).

       PROCEDURE DIVISION.
      * code goes here!
           INITIALIZE COMMAREA.

           DISPLAY "test1 example program"
           DISPLAY "Name: " WITH NO ADVANCING.
           ACCEPT INPUT-NAME.

           COMPUTE COMMLEN = LENGTH OF COMMAREA.
           CALL "D8link" USING PROG-NAME COMMAREA COMMLEN.

           GOBACK.


