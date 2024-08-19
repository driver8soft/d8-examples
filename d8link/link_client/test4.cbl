       IDENTIFICATION DIVISION.
       PROGRAM-ID. test4.

       ENVIRONMENT DIVISION.

       CONFIGURATION SECTION.
       
       DATA DIVISION.

       FILE SECTION.

       WORKING-STORAGE SECTION.
      * Declare variables in the WORKING-STORAGE section
       01 PROG-NAME PIC X(8) VALUE "loop".
       01 COMMLEN PIC 9(9) COMP.
       01 COMMAREA.
           05 N-TIMES PIC S9(9) COMP VALUE 1.

       01 N-CALLS PIC S9(9) COMP.

       PROCEDURE DIVISION.
      * code goes here!
           DISPLAY "test4 example program"
           DISPLAY "Number of gRPC calls: " WITH NO ADVANCING.
           ACCEPT N-CALLS.

           COMPUTE COMMLEN = LENGTH OF COMMAREA.
           PERFORM PROCESS N-CALLS TIMES.
           GOBACK.

           PROCESS.

           CALL "D8link" USING PROG-NAME COMMAREA COMMLEN.
           DISPLAY "Loop: " N-CALLS.
           SUBTRACT 1 FROM N-CALLS.

           


